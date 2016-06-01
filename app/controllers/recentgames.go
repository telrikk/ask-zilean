package controllers

import (
	"strconv"

	"github.com/revel/revel"
	model "github.com/telrikk/ask-zilean/app/recentgames"
	"github.com/telrikk/ask-zilean/app/util"
	"github.com/telrikk/lol-go-api/game"
	"github.com/telrikk/lol-go-api/match"
)

// RecentGames returns json with recent game information
func (c App) RecentGames(name string) revel.Result {
	summonerID, err := model.GetSummonerID(name)
	if err != nil {
		return util.HandleError(*c.Controller, err)
	}
	recentGames, err := model.GetRecentGames(summonerID)
	if err != nil {
		return util.HandleError(*c.Controller, err)
	}
	fullGameData, err := model.GetFullGameData(recentGames)
	if err != nil {
		return util.HandleError(*c.Controller, err)
	}
	var games []model.RecentGame
	for _, recentGame := range recentGames {
		translatedGame := translateRecentGame(recentGame, fullGameData, summonerID)
		if translatedGame.ID != 0 {
			games = append(games, translatedGame)
		}
	}
	response := new(model.Response)
	response.Results = games
	return c.RenderJson(response)
}

func translateRecentGame(game game.Game, gameData map[int]match.Detail, summonerID int) model.RecentGame {
	translatedGame := new(model.RecentGame)
	fullGame := gameData[game.GameID]
	nameData := make(map[int]string)
	summonerParticipantID := 0

	for _, participant := range fullGame.ParticipantIdentities {
		// HACK: bad data on Riot's side
		if participant.Player.SummonerName == "" {
			return *translatedGame
		}
		nameData[participant.ParticipantID] = participant.Player.SummonerName
		if participant.Player.SummonerID == summonerID {
			summonerParticipantID = participant.ParticipantID
		}
	}
	for _, player := range fullGame.Participants {
		translatedPlayer := translatePlayer(player, nameData)
		translatedGame.Players = append(translatedGame.Players, translatedPlayer)
		if player.ParticipantID == summonerParticipantID {
			translatedGame.Summoner = translatedPlayer
		}
	}
	mapImageName := util.MapData().Data[strconv.Itoa(game.MapID)].Image.Full
	translatedGame.MapImageURL = util.MapImagePrefix() + "map/" + mapImageName
	translatedGame.MapName = util.MapData().Data[strconv.Itoa(game.MapID)].MapName
	translatedGame.ID = game.GameID
	translatedGame.GoldImageURL = "/public/img/gold.png"
	translatedGame.ChampionImageURL = util.UIImagePrefix() + "ui/champion.png"
	translatedGame.StatsImageURL = "/public/img/score.png"
	translatedGame.ItemsImageURL = "/public/img/items.png"
	translatedGame.CreepScoreImageURL = "/public/img/minion.png"
	return *translatedGame
}

func translatePlayer(player match.Participant, nameData map[int]string) model.Player {
	translatedPlayer := new(model.Player)
	translatedPlayer.SummonerName = nameData[player.ParticipantID]
	translatedPlayer.Kills = player.Stats.Kills
	translatedPlayer.Assists = player.Stats.Assists
	translatedPlayer.Deaths = player.Stats.Deaths
	championImageName := util.ChampionData().Data[strconv.Itoa(player.ChampionID)].Image.Full
	translatedPlayer.ChampionImageURL = util.ChampionImagePrefix() + "champion/" + championImageName
	translatedPlayer.ChampionName = util.ChampionData().Data[strconv.Itoa(player.ChampionID)].Name
	translatedPlayer.Items = []model.Item{}
	for _, ItemID := range []int{player.Stats.Item0, player.Stats.Item1, player.Stats.Item2,
		player.Stats.Item3, player.Stats.Item4, player.Stats.Item5, player.Stats.Item6} {
		if ItemID > 0 {
			translatedPlayer.Items = append(translatedPlayer.Items, translateItem(ItemID))
		}
	}
	translatedPlayer.CreepScore = player.Stats.MinionsKilled + player.Stats.NeutralMinionsKilled
	translatedPlayer.Gold = player.Stats.GoldEarned
	return *translatedPlayer
}

func translateItem(itemID int) model.Item {
	item := new(model.Item)
	item.ImageURL = util.ItemsImagePrefix() + "item/" + strconv.Itoa(itemID) + ".png"
	fullItem := util.ItemData().Data[strconv.Itoa(itemID)]
	for _, tag := range fullItem.Tags {
		if tag == "Trinket" {
			item.IsTrinket = true
		}
	}
	return *item
}
