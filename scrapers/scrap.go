package scrapers

import (
	"encoding/json"

	"github.com/bpetetot/leafer/db"
	"github.com/coocood/freecache"
)

const cacheSize = 10 * 1024 * 1024
const expire = 60

var cache = freecache.NewCache(cacheSize)

// Scrap the search term to get info
func Scrap(search string) db.Media {
	media, err := getFromCache(search)
	if err != nil {
		media = scrapAnilist(search)
		setToCache(search, &media)
		return media
	}
	return media
}

func getFromCache(key string) (db.Media, error) {
	result, err := cache.Get([]byte(key))
	if err != nil {
		return db.Media{}, err
	}
	var media db.Media
	json.Unmarshal(result, &media)
	return media, nil
}

func setToCache(key string, media *db.Media) {
	content, err := json.Marshal(media)
	if err == nil {
		cache.Set([]byte(key), content, expire)
	}
}
