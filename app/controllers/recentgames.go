package controllers

import (
	"strconv"

	"github.com/revel/revel"
	model "github.com/telrikk/ask-zilean/app/recentgames"
	"github.com/telrikk/ask-zilean/app/util"
	"github.com/telrikk/lol-go-api/game"
	"github.com/telrikk/lol-go-api/match"
)

var queueDescriptions = map[string]string{
	"ARAM_5x5":                        "(ARAM)",
	"BILGEWATER_5x5":                  "(Bilgewater)",
	"BILGEWATER_ARAM_5x5":             "(ARAM)",
	"ASCENSION_5x5":                   "(Ascension)",
	"NORMAL_5x5_BLIND":                "(Blind)",
	"BOT_5x5":                         "(Bots)",
	"BOT_5x5_BEGINNER":                "(Bots)",
	"BOT_5x5_INTERMEDIATE":            "(Bots)",
	"BOT_5x5_INTRO":                   "(Bots)",
	"BOT_TT_3x3":                      "(Bots)",
	"BOT_ODIN_5x5":                    "(Bots)",
	"BOT_URF_5x5":                     "(Bots)",
	"COUNTER_PICK":                    "(Counter Pick)",
	"CUSTOM":                          "(Custom)",
	"NORMAL_5x5_DRAFT":                "(Draft)",
	"FIRSTBLOOD_1x1":                  "(First Blood)",
	"FIRSTBLOOD_2x2":                  "(First Blood)",
	"HEXAKILL":                        "(Hexakill)",
	"KING_PORO_5x5":                   "(King Poro)",
	"SR_6x6":                          "(Hexakill)",
	"NIGHTMARE_BOT_5x5_RANK1":         "(Bots)",
	"NIGHTMARE_BOT_5x5_RANK5":         "(Bots)",
	"NIGHTMARE_BOT_5x5_RANK2":         "(Bots)",
	"ODIN_5x5_BLIND":                  "(Dominion)",
	"ODIN_5x5_DRAFT":                  "(Odin)",
	"ONEFORALL_MIRRORMODE_5x5":        "(One For All)",
	"ONEFORALL_5x5":                   "(One For All)",
	"RANKED_PREMADE_3x3":              "(Team Ranked)",
	"RANKED_PREMADE_5x5":              "(Team Ranked)",
	"RANKED_SOLO_5x5":                 "(Ranked)",
	"RANKED_TEAM_5x5":                 "(Team Ranked)",
	"RANKED_TEAM_3x3":                 "(Team Ranked)",
	"NORMAL_3x3":                      "(Normal)",
	"GROUP_FINDER_5x5":                "(Team Builder)",
	"TEAM_BUILDER_DRAFT_UNRANKED_5x5": "(Normal)",
	"TEAM_BUILDER_DRAFT_RANKED_5x5":   "(Ranked)",
	"URF_5x5":                         "(URF)",
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
	fullGameData, err := model.GetFullGameData(recentGames)
	if err != nil {
		return util.HandleError(*c.Controller, err)
	}
	var games []model.RecentGame
	for _, recentGame := range recentGames {
		translatedGame := translateRecentGame(recentGame, fullGameData, summonerID, name)
		if translatedGame.ID != 0 {
			games = append(games, translatedGame)
		}
	}
	response := new(model.Response)
	response.Results = games
	return c.RenderJson(response)
}

func translateRecentGame(game game.Game, gameData map[int]match.Detail, summonerID int, summonerName string) model.RecentGame {
	translatedGame := new(model.RecentGame)
	fullGame := gameData[game.GameID]
	nameData := make(map[int]string)
	for _, participant := range fullGame.ParticipantIdentities {
		// HACK: bad data on Riot's side, will need to use summoner API for non-ranked games
		nameData[participant.ParticipantID] = participant.Player.SummonerName
	}
	for _, player := range fullGame.Participants {
		translatedPlayer := translatePlayer(player, nameData)
		translatedGame.Players = append(translatedGame.Players, translatedPlayer)
	}
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
	description, ok := queueDescriptions[fullGame.QueueType]
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
	translatedPlayer.IsWinner = player.Stats.Winner
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
