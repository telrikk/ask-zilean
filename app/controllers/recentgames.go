package controllers

import (
	"strconv"

	"github.com/revel/revel"
	model "github.com/telrikk/ask-zilean/app/recentgames"
	"github.com/telrikk/ask-zilean/app/util"
	"github.com/telrikk/lol-go-api/game"
)

var gameSubTypeMap = map[string]string{
	"BOT":                "(Bots)",
	"NORMAL":             "(Normal)",
	"ODIN_UNRANKED":      "(Dominion)",
	"RANKED_PREMADE_3x3": "(Team Ranked)",
	"RANKED_PREMADE_5x5": "(Team Ranked)",
	"NORMAL_3x3":         "(Twisted Treeline)",
	"BOT_3x3":            "(Bots)",
	"ARAM_UNRANKED_5x5":  "(ARAM)",
	"URF":                "(URF)",
	"URF_BOT":            "(Bots)",
	"ASCENSION":          "(Ascension)",
	"NIGHTMARE_BOT":      "(Bots)",
	"COUNTER_PICK":       "(Counter Pick)",
	"KING_PORO":          "(King Poro)",
	"BILGEWATER":         "(Bilgewater)",
	"HEXAKILL":           "(Hexakill)",
	"SR_6x6":             "(Hexakill)",
	"FIRSTBLOOD_1x1":     "(First Blood)",
	"FIRSTBLOOD_2x2":     "(First Blood)",
	"ONEFORALL_5x5":      "(One For All)",
	"RANKED_SOLO_5x5":    "(Ranked)",
	"RANKED_TEAM_5x5":    "(Team Ranked)",
	"RANKED_TEAM_3x3":    "(Team Ranked)",
}

var mapNames = map[string]string{
	"NewTwistedTreeline": "Twisted Treeline",
	"SummonersRift":      "Summoner's Rift",
	"CrystalScar":        "CrystalScar",
	"SummonersRiftNew":   "Summoner's Rift",
	"ProvingGroundsNew":  "Proving Grounds",
}

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
	var games []model.RecentGame
	for _, recentGame := range recentGames {
		translatedGame := translateRecentGame(recentGame, summonerID, name)
		if translatedGame.ID != 0 {
			games = append(games, translatedGame)
		}
	}
	response := new(model.Response)
	response.Results = games
	return c.RenderJson(response)
}

func translateRecentGame(game game.Game, summonerID int, summonerName string) model.RecentGame {
	translatedGame := new(model.RecentGame)
	translatedGame.Summoner = translateSummoner(game, summonerName)
	mapImageName := util.MapData().Data[strconv.Itoa(game.MapID)].Image.Full
	translatedGame.MapImageURL = util.MapImagePrefix() + "map/" + mapImageName
	translatedGame.ID = game.GameID
	translatedGame.GoldImageURL = "/public/img/gold.png"
	translatedGame.ChampionImageURL = util.UIImagePrefix() + "ui/champion.png"
	translatedGame.StatsImageURL = "/public/img/score.png"
	translatedGame.ItemsImageURL = "/public/img/items.png"
	translatedGame.CreepScoreImageURL = "/public/img/minion.png"
	rawMapName := util.MapData().Data[strconv.Itoa(game.MapID)].MapName
	readableMapName, ok := mapNames[rawMapName]
	if !ok {
		readableMapName = "Unknown Map"
	}
	translatedGame.MapName = readableMapName
	description, ok := gameSubTypeMap[game.SubType]
	if !ok {
		description = "(Other)"
	}
	translatedGame.QueueDescription = description
	return *translatedGame
}

func translateSummoner(game game.Game, summonerName string) model.Player {
	translatedPlayer := new(model.Player)
	translatedPlayer.SummonerName = summonerName
	translatedPlayer.Kills = game.Stats.ChampionsKilled
	translatedPlayer.Assists = game.Stats.Assists
	translatedPlayer.Deaths = game.Stats.NumDeaths

	championImageName := util.ChampionData().Data[strconv.Itoa(game.ChampionID)].Image.Full
	translatedPlayer.ChampionImageURL = util.ChampionImagePrefix() + "champion/" + championImageName
	translatedPlayer.ChampionName = util.ChampionData().Data[strconv.Itoa(game.ChampionID)].Name
	translatedPlayer.Items = []model.Item{}
	for _, ItemID := range []int{game.Stats.Item0, game.Stats.Item1, game.Stats.Item2,
		game.Stats.Item3, game.Stats.Item4, game.Stats.Item5, game.Stats.Item6} {
		if ItemID > 0 {
			translatedPlayer.Items = append(translatedPlayer.Items, translateItem(ItemID))
		}
	}
	translatedPlayer.CreepScore = game.Stats.MinionsKilled + game.Stats.NeutralMinionsKilled
	translatedPlayer.Gold = game.Stats.GoldEarned
	translatedPlayer.IsWinner = game.Stats.Win
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
