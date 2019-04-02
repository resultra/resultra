// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package dashboardController

import (
	"github.com/resultra/resultra/server/dashboard/values"
	"testing"
)

func testOneLabelGrouping(t *testing.T, valGrouping values.ValGrouping,
	inputVal float64, expectedLabel string) {
	labelInfo := bucketedNumberGroupLabelInfo(inputVal, valGrouping)
	if labelInfo == nil {
		t.Errorf("testOneLabelGrouping (fail): failed to generate label for %v with grouping %+v",
			inputVal, valGrouping)
	} else {
		if labelInfo.label != expectedLabel {
			t.Errorf("testOneLabelGrouping (fail): incorrect label for %v with grouping %+v: expecting %v, got %v",
				inputVal, valGrouping, expectedLabel, labelInfo.label)

		} else {
			t.Logf("testOneLabelGrouping (pass): %v", labelInfo.label)
		}
	}
}

func TestNumberGroupings(t *testing.T) {
	t.Logf("Testing number groupings")

	bucketWidth := 0.5
	start := -2.0
	end := 2.0

	groupByFieldID := "testFieldID"
	groupValsBy := values.ValGroupByBucket

	valGrouping := values.ValGrouping{
		GroupValsByFieldID:    &groupByFieldID,
		GroupValsBy:           &groupValsBy,
		GroupByValBucketWidth: &bucketWidth,
		BucketStart:           &start,
		BucketEnd:             &end}

	testOneLabelGrouping(t, valGrouping, -2.1, "< -2")
	testOneLabelGrouping(t, valGrouping, -2, "-2 to -1.5")
	testOneLabelGrouping(t, valGrouping, -0.6, "-1 to -0.5")
	testOneLabelGrouping(t, valGrouping, -0.5, "-0.5 to 0")
	testOneLabelGrouping(t, valGrouping, -0.1, "-0.5 to 0")
	testOneLabelGrouping(t, valGrouping, 0.0, "0 to 0.5")
	testOneLabelGrouping(t, valGrouping, 0.6, "0.5 to 1")
	testOneLabelGrouping(t, valGrouping, 2.0, "2 to 2.5")
	testOneLabelGrouping(t, valGrouping, 2.5, "> 2")

}
