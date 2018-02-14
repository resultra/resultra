package alert

import (
	"fmt"
	"github.com/hashicorp/golang-lru"
	"strings"
)

/* AlertsCache is a simple LRU cache for generated alerts */

var AlertsCache *lru.Cache

func init() {
	cache, err := lru.New(128)

	if err != nil {
		panic(fmt.Sprintf("Failure initializing alerts cache: %v", err))
	}
	AlertsCache = cache
}

func alertCacheKey(databaseID string, userID string, userIsAdmin bool) string {
	return fmt.Sprintf("%v-%v-%v", databaseID, userID, userIsAdmin)
}

func RemoveTrackerDatabaseCacheEntries(databaseID string) {
	keys := AlertsCache.Keys()
	for _, currKey := range keys {

		var validType bool
		keyStr, validType := currKey.(string)
		if validType {
			if strings.HasPrefix(keyStr, databaseID) {
				AlertsCache.Remove(keyStr)
			}
		}

	}
}
