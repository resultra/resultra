function createAlternateViewLookupTable(altViews) {
	var altViewLookup = {}
	$.each(altViews,function(index,altView) {
		if (altView.formID != null) {
			altViewLookup[altView.formID] = altView
		} else if (altView.tableID != null) {
			altViewLookup[altView.tableID] = altView
		}
	})
	return altViewLookup
}
