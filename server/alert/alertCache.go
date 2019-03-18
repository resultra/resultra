// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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
