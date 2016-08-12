function initDesignFormFilterProperties(formID) {
	
	function changeDefaultFilterSelection(filterID, isChecked) {
		
	}
	
	$('#designFormDefaultFilterList').empty()
	addFilterToFilterCheckboxList('#designFormDefaultFilterList', 'defaultFilter_',
		{filterID:"filter1",name:"Filter 1"},changeDefaultFilterSelection)
	addFilterToFilterCheckboxList('#designFormDefaultFilterList', 'defaultFilter_',
		{filterID:"filter2",name:"Filter 2"},changeDefaultFilterSelection)


	function changeAvailableFilterSelection(filterID, isChecked) {
		
	}
	
	$('#designFormAvailableFilterList').empty()
	addFilterToFilterCheckboxList('#designFormAvailableFilterList', 'availableFilter_',
		{filterID:"filter1",name:"Filter 1"},changeAvailableFilterSelection)
	addFilterToFilterCheckboxList('#designFormAvailableFilterList', 'availableFilter_',
		{filterID:"filter2",name:"Filter 2"},changeAvailableFilterSelection)
	addFilterToFilterCheckboxList('#designFormAvailableFilterList', 'availableFilter_',
		{filterID:"filter3",name:"Filter 3"},changeAvailableFilterSelection)
	addFilterToFilterCheckboxList('#designFormAvailableFilterList', 'availableFilter_',
		{filterID:"filter4",name:"Filter 4"},changeAvailableFilterSelection)

}