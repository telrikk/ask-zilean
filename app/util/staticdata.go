package util

import "github.com/telrikk/lol-go-api/staticdata"

var itemsImagePrefix string
var uiImagePrefix string
var championImagePrefix string
var mapImagePrefix string
var mapData staticdata.MapData
var championData staticdata.ChampionData
var itemData staticdata.ItemData

func init() {
	realm, err := GetServiceFactory().StaticDataService().GetRealmData()
	if err != nil {
		panic(err)
	}
	championImagePrefix = realm.CDN + "/" + realm.LatestVersion["champion"] + "/img/"
	itemsImagePrefix = realm.CDN + "/" + realm.LatestVersion["items"] + "/img/"
	// HACK: not in Riot's "Realm" object
	uiImagePrefix = realm.CDN + "/5.5.1/img/"
	mapImagePrefix = realm.CDN + "/" + realm.LatestVersion["map"] + "/img/"

	dataRequest := staticdata.NewStaticDataRequest("en_US", realm.CurrentVersion)
	mapDataResponse, err := GetServiceFactory().StaticDataService().GetMapData(*dataRequest)
	if err != nil {
		panic(err)
	}
	mapData = *mapDataResponse

	filteredDataRequest := staticdata.NewFilterableStaticDataRequest("en_US", realm.CurrentVersion, []string{"image"})
	championDataResponse, err := GetServiceFactory().StaticDataService().GetChampionData(*filteredDataRequest, true)
	if err != nil {
		panic(err)
	}
	championData = *championDataResponse

	itemRequest := staticdata.NewFilterableStaticDataRequest("en_US", realm.CurrentVersion, []string{"tags"})
	itemDataResponse, err := GetServiceFactory().StaticDataService().GetItemData(*itemRequest)
	if err != nil {
		panic(err)
	}
	itemData = *itemDataResponse
}

// ChampionData returns all of the current static champion data, keyed by champion ID
func ChampionData() *staticdata.ChampionData {
	return &championData
}

// MapData returns all of the current static map data, keyed by map ID
func MapData() *staticdata.MapData {
	return &mapData
}

// ItemData returns all of the current static item data, keyed by item ID
func ItemData() *staticdata.ItemData {
	return &itemData
}

// ChampionImagePrefix returns the current prefix for LoL champion static images (up to '/img/')
func ChampionImagePrefix() string {
	return championImagePrefix
}

// ItemsImagePrefix returns the current prefix for LoL item static images (up to '/img/')
func ItemsImagePrefix() string {
	return championImagePrefix
}

// MapImagePrefix returns the current prefix for LoL map static images (up to '/img/')
func MapImagePrefix() string {
	return championImagePrefix
}

// UIImagePrefix returns the current prefix for LoL UI static images, e.g. gold, (up to '/img/')
func UIImagePrefix() string {
	return uiImagePrefix
}
