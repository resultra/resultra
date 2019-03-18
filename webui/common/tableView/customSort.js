// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function registerTableViewCustomSortFuncs() {
	
	function sortValCompare(a,b,sortValFunc) {
		var aSortVal = sortValFunc(a)
		var bSortVal = sortValFunc(b)
		
		if(aSortVal > bSortVal) {
			return 1
		} else if (aSortVal < bSortVal) {
			return -1
		} else {
			return 0
		}
		
	}
	
	function sortBoolAsc(a,b) {
		function boolSortVal(x) {
			if (x===true) {
				return 1
			} else if (x===false) {
				return 0
			} else {
				return -1 // order undefined values after true or false
			}
		}
		return sortValCompare(a,b,boolSortVal)		
	}
	
	function sortNumAsc(a,b) {
		function numSortVal(x) {
			if(x==="" || isNaN(x)) {
				return Number.NEGATIVE_INFINITY
			} else {
				return Number(x)
			}
		}
		return sortValCompare(a,b,numSortVal) 
	}
	
	jQuery.extend( jQuery.fn.dataTableExt.oSort, {
		"custom-bool-asc": function (a, b) {
			return sortBoolAsc(a,b);
		},
		"custom-bool-desc": function (a, b) {
			return sortBoolAsc(a,b) * -1;
		},
		"custom-num-asc": function (a, b) {
			return sortNumAsc(a,b);
		},
		"custom-num-desc": function (a, b) {
			return sortNumAsc(a,b) * -1;
		}
	});
	
}