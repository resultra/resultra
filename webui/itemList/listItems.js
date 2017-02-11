

function ListItemController() {
	
	this.listItemsInfo = []
	
	// this isn't normally accessible from inner functions when this function
	// is instantiated as an object. Below is a conventional work-around.
	var listItemControllerSelf = this
	
	var currRecordSet = null
	var currRecordSetWindowSize = 1
	
	this.populateListViewWithListItemContainers = function (populationDoneCallback) {
		
		function populateOneListItem(itemWindowIndex) {
			var $listItemContainer = $('<div class="listItemContainer"></div>')
			
			if (itemWindowIndex % 2 === 0) {
				$listItemContainer.addClass("listItemEvenItem")
			} else {
				$listItemContainer.addClass("listItemOddItem")

			}
		
			initObjectCanvasContainerSelectionBehavior($listItemContainer, function() {
				hideSiblingsShowOne('#listViewProps')
			})
	
			function getCurrentRecord() {
				console.log("getCurrentRecord: item list: window index = " + itemWindowIndex)
				return  currRecordSet.getRecordRefAtWindowIndex(itemWindowIndex)
			}
			function updateCurrentRecord(updatedRecordRef) {
				currRecordSet.updateRecordRef(updatedRecordRef)
				loadCurrentRecordsIntoItemLayoutContainers()
			}
	
			var recordProxy = {
				changeSetID: MainLineFullyCommittedChangeSetID,
				getRecordFunc: getCurrentRecord,
				updateRecordFunc: updateCurrentRecord
			}
		
			var listItemInfo = {
				$listItemContainer: $listItemContainer,
				recordProxy: recordProxy
			}
		
			$('#layoutCanvas').append($listItemContainer)
			
			return listItemInfo
		}
		
		this.listItemsInfo = []
		$('#layoutCanvas').empty()
		for(var listIndex = 0; listIndex < currRecordSetWindowSize; listIndex++) {
			var listItemInfo = populateOneListItem(listIndex)
			this.listItemsInfo.push(listItemInfo)
		}
		loadMultipleFormViewContainers(viewListContext,this.listItemsInfo,populationDoneCallback)
	
	}
	
	
	function reloadRecordsIntoContainers() {
		for(var listItemIndex = 0; listItemIndex < listItemControllerSelf.listItemsInfo.length; listItemIndex++) {
			var currListItem = listItemControllerSelf.listItemsInfo[listItemIndex]
		
			var recordRef = currListItem.recordProxy.getRecordFunc()
			if(recordRef != null)
			{
				currListItem.$listItemContainer.show()
				loadRecordIntoFormLayout(currListItem.$listItemContainer,recordRef)
	
				// Update footer to reflect where the current record is in list of currently loaded records
				$('#recordNumLabel').text(currRecordSet.recPageLabel())
		
		
			} else {
				currListItem.$listItemContainer.hide()
			}
		
		}
		// If the record changed, and one of the form components is already loaded, it needs to be 
		// re-selected so the sidebar can be re-initialized with any settings specific to this 
		// record.
		reselectCurrentObjectSelection()
	}
	
	this.loadCurrentRecordsIntoItemLayoutContainers = function ()
	{
		reloadRecordsIntoContainers()
	}
	
	
	function loadFormData(reloadRecordParams, formDataCallback) {
		var numDataSetsRemainingToLoad = 2
	
		var formData =  {}
	
		function oneDataSetLoaded() {
			numDataSetsRemainingToLoad -= 1
			if(numDataSetsRemainingToLoad <= 0) {
				formDataCallback(formData)
			}
		}
	
		jsonAPIRequest("recordRead/getFilteredSortedRecordValues",reloadRecordParams,function(recordsData) {
			formData.recordData = recordsData
			oneDataSetLoaded()
		})
	
		var globalParams = { parentDatabaseID: viewListContext.databaseID }
		jsonAPIRequest("global/getValues",globalParams,function(globalVals) {
			formData.globalVals = globalVals
			oneDataSetLoaded()
		})
	
	}
	
	function enableRecordButtons(isEnabled)
	{
		var isDisabled = true;
		if(isEnabled) { isDisabled=false }
		$('#prevRecordButton').prop("disabled",isDisabled)
		$('#nextRecordButton').prop("disabled",isDisabled)
		$('#newRecordButton').prop("disabled",isDisabled)
	}
	
	function enableNewRecordButton()
	{
		$('#newRecordButton').prop("disabled",false)
	}
	
	this.reloadRecords =  function(reloadParams) {
		
		currReloadParams = reloadParams
	
		loadFormData(reloadParams,function(formData) {
			currGlobalVals = formData.globalVals	
			currRecordSet = new RecordSet(formData.recordData,currRecordSetWindowSize);
			if(currRecordSet.numRecords() > 0) {
				reloadRecordsIntoContainers()		
			}
		
			// Enable the buttons to page through the records
			if(currRecordSet.numRecords() > 0) {
				enableRecordButtons(true)
			}
			else {
				enableNewRecordButton() // just enable the "New Record" button
			}
		
		})	
	}
	
	
	// Initially disabled the buttons for paging through the records. They'll be 
	// enabled once the records are loaded.
	enableRecordButtons(false)
	
	$('#nextRecordButton').click(function(e){
	         e.preventDefault();
			 if(currRecordSet.advanceToNextPage()) {
			 	reloadRecordsIntoContainers()
			 }
	});
	
	$('#prevRecordButton').click(function(e){
	         e.preventDefault();
			 if(currRecordSet.advanceToPrevPage()) {
				 console.log("Advance to next record")
			 	reloadRecordsIntoContainers()
			 } 
	});
	
	function createNewRecord() {
		var newRecordsParams = {parentDatabaseID:viewListContext.databaseID}
		jsonAPIRequest("recordUpdate/newRecord",newRecordsParams,function(newRecordRef) {
			currRecordSet.appendNewRecord(newRecordRef);
			currRecordSet.jumpToRecord(newRecordRef.recordID)
			reloadRecordsIntoContainers()
		}) // getRecord
	
	}
	
	$('#newRecordButton').click(function(e){ createNewRecord() });
	
	var $pageSizeSelection = $('#viewListPageSizeSelection')
	$pageSizeSelection.val(currRecordSetWindowSize)
	initSelectControlChangeHandler($pageSizeSelection, function(newPageSize){
		currRecordSetWindowSize = Number(newPageSize)
		// Re-initialize the list view, then repopulate it with the records.
		listItemControllerSelf.populateListViewWithListItemContainers(function() {
			currRecordSet.setWindowSize(currRecordSetWindowSize)
			reloadRecordsIntoContainers()
		})
	})
	
	
	
}