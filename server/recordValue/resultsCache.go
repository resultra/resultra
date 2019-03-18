// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package recordValue

import (
	"fmt"
	"github.com/hashicorp/golang-lru"
)

/* ResultsCache is a simple LRU cache for an unfiltered set of records with their calculated fields's values
   also computed. Unless a record's value is changed, most of the time the cache will have an up to date
   results cache entry. This is useful, for example, for dashboards, which load the same set or records, but
   filter them differently for different results. */

var ResultsCache *lru.Cache

func init() {
	cache, err := lru.New(512)

	if err != nil {
		panic(fmt.Sprintf("Failure initializing value results cache: %v", err))
	}
	ResultsCache = cache
}
