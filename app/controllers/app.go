package controllers

import (
	"strconv"
	"strings"

	"github.com/revel/revel"
	"github.com/telrikk/ask-zilean/app/util"
	"github.com/telrikk/lol-go-api/game"
	"github.com/telrikk/lol-go-api/match"
	"github.com/telrikk/lol-go-api/staticdata"
	apiutil "github.com/telrikk/lol-go-api/util"
)

// App is the Ask Zilean app
type App struct {
	*revel.Controller
}

// Index renders the page for Ask Zilean
func (c App) Index() revel.Result {
	return c.Render()
}

type Item struct {
	ImageURL string `json:"imageURL"`
}

type Player struct {
	ChampionImageURL string `json:"championImageURL"`
	ChampionName     string `json:"championName"`
	SummonerName     string `json:"summonerName"`
	Kills            int    `json:"kills"`
	Deaths           int    `json:"deaths"`
	Assists          int    `json:"assists"`
	Items            []Item `json:"items"`
	Gold             int    `json:"gold"`
}

type RecentGame struct {
	MapImageURL string   `json:"mapImageURL"`
	MapName     string   `json:"mapName"`
	Players     []Player `json:"players"`
	ID          int      `json:"id"`
}

type RecentGamesResponse struct {
	Results []RecentGame `json:"results"`
}

var imagePrefix string
var mapData staticdata.MapData
var championData staticdata.ChampionData

func init() {
	realm, err := util.GetServiceFactory().StaticDataService().GetRealmData()
	if err != nil {
		panic(err)
	}
	imagePrefix = realm.CDN + "/" + realm.CurrentVersion + "/img/"

	dataRequest := staticdata.NewStaticDataRequest("en_US", realm.CurrentVersion)
	mapDataResponse, err := util.GetServiceFactory().StaticDataService().GetMapData(*dataRequest)
	if err != nil {
		panic(err)
	}
	mapData = *mapDataResponse

	filteredDataRequest := staticdata.NewFilterableStaticDataRequest("en_US", realm.CurrentVersion, []string{"image"})
	championDataResponse, err := util.GetServiceFactory().StaticDataService().GetChampionData(*filteredDataRequest, true)
	if err != nil {
		panic(err)
	}
	championData = *championDataResponse
}

// RecentGames returns json with recent game information
func (c App) RecentGames(name string) revel.Result {
	summonerData, err := util.GetServiceFactory().
		SummonerService().
		GetSummonerDataByName([]string{name})
	if err != nil {
		return handleError(c, err)
	}
	summonerID := summonerData.SummonerData[strings.ToLower(name)].ID
	recentGames, err := util.GetServiceFactory().RecentGamesService().List(summonerID)
	if err != nil {
		return handleError(c, err)
	}
	var games []RecentGame
	fullGameData, err := getFullGameData(recentGames.Games)
	if err != nil {
		return handleError(c, err)
	}
	for _, recentGame := range recentGames.Games {
		games = append(games, translateRecentGame(recentGame, fullGameData))
	}
	response := new(RecentGamesResponse)
	response.Results = games
	return c.RenderJson(response)
}

func getFullGameData(games []game.Game) (map[int]match.Detail, error) {
	gameData := make(map[int]match.Detail)
	for _, game := range games {
		fullGame, err := util.GetServiceFactory().MatchService().Get(game.GameID, false)
		if err != nil {
			return nil, err
		}
		gameData[fullGame.MatchID] = *fullGame
	}
	return gameData, nil
}

func translateRecentGame(game game.Game, gameData map[int]match.Detail) RecentGame {
	translatedGame := new(RecentGame)
	fullGame := gameData[game.GameID]
	nameData := make(map[int]string)
	for _, participant := range fullGame.ParticipantIdentities {
		nameData[participant.ParticipantID] = participant.Player.SummonerName
	}
	for _, player := range fullGame.Participants {
		translatedPlayer := new(Player)
		translatedPlayer.SummonerName = nameData[player.ParticipantID]
		translatedPlayer.Kills = player.Stats.Kills
		translatedPlayer.Assists = player.Stats.Assists
		translatedPlayer.Deaths = player.Stats.Deaths
		championImageName := championData.Data[strconv.Itoa(player.ChampionID)].Image.Full
		translatedPlayer.ChampionImageURL = imagePrefix + "champion/" + championImageName
		translatedPlayer.ChampionName = championData.Data[strconv.Itoa(player.ChampionID)].Name
		translatedPlayer.Items = []Item{}
		for _, ItemID := range []int{player.Stats.Item0, player.Stats.Item1, player.Stats.Item2,
			player.Stats.Item3, player.Stats.Item4, player.Stats.Item5, player.Stats.Item6} {
			if ItemID > 0 {
				translatedPlayer.Items = append(translatedPlayer.Items, translateItem(ItemID))
			}
		}
		translatedPlayer.Gold = player.Stats.GoldEarned
		translatedGame.Players = append(translatedGame.Players, *translatedPlayer)
	}
	mapImageName := mapData.Data[strconv.Itoa(game.MapID)].Image.Full
	translatedGame.MapImageURL = imagePrefix + "map/" + mapImageName
	translatedGame.MapName = mapData.Data[strconv.Itoa(game.MapID)].MapName
	translatedGame.ID = game.GameID
	return *translatedGame
}

func translateItem(itemID int) Item {
	item := new(Item)
	item.ImageURL = imagePrefix + "item/" + strconv.Itoa(itemID) + ".png"
	return *item
}

func handleError(c App, err error) revel.Result {
	riotError, isRiotError := err.(apiutil.APIError)
	if isRiotError && riotError.Status.StatusCode == 404 {
		return c.NotFound(riotError.Status.Message)
	}
	return c.RenderError(err)
}
