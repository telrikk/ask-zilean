package recentgames

import (
	"strings"

	"github.com/telrikk/ask-zilean/app/util"
	"github.com/telrikk/lol-go-api/game"
	"github.com/telrikk/lol-go-api/match"
)

// Item is an item which was purchased in a recent game
type Item struct {
	ImageURL  string `json:"imageURL"`
	IsTrinket bool   `json:"isTrinket"`
}

// Player is a summoner involved in a recent game
type Player struct {
	ChampionImageURL string `json:"championImageURL"`
	ChampionName     string `json:"championName"`
	SummonerName     string `json:"summonerName"`
	Kills            int    `json:"kills"`
	Deaths           int    `json:"deaths"`
	Assists          int    `json:"assists"`
	Items            []Item `json:"items"`
	CreepScore       int    `json:"creepScore"`
	Gold             int    `json:"gold"`
}

// RecentGame is a recently played game
type RecentGame struct {
	MapImageURL        string   `json:"mapImageURL"`
	MapName            string   `json:"mapName"`
	Players            []Player `json:"players"` // all players in the game
	ID                 int      `json:"id"`
	Summoner           Player   `json:"summoner"` // the summoner who is using the app
	CreepScoreImageURL string   `json:"creepScoreImageURL"`
	StatsImageURL      string   `json:"statsImageURL"`
	ItemsImageURL      string   `json:"itemsImageURL"`
	ChampionImageURL   string   `json:"championImageURL"`
	GoldImageURL       string   `json:"goldImageURL"`
}

// Response contains a list of recently played games for a summoner
type Response struct {
	Results []RecentGame `json:"results"`
}

// GetSummonerID returns the ID of a summoner given their name, or an error
func GetSummonerID(summonerName string) (int, error) {
	summonerData, err := util.GetServiceFactory().
		SummonerService().
		GetSummonerDataByName([]string{summonerName})
	if err != nil {
		return 0, err
	}
	summonerID := summonerData.SummonerData[strings.ToLower(summonerName)].ID
	return summonerID, nil
}

// GetFullGameData ,given a list of games, returns a data structure with additional information,
// keyed on a map by match ID
func GetFullGameData(games []game.Game) (map[int]match.Detail, error) {
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

// GetRecentGames gets the recent games for a summoner, given their ID
func GetRecentGames(summonerID int) ([]game.Game, error) {
	recentGames, err := util.GetServiceFactory().RecentGamesService().List(summonerID)
	if err != nil {
		return nil, err
	}
	return recentGames.Games, nil
}
