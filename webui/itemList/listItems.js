

function ListItemController($parentContainer) {
	
	this.listItemsInfo = []
	this.$parentContainer = $parentContainer
	
	// this isn't normally accessible from inner functions when this function
	// is instantiated as an object. Below is a conventional work-around.
	var listItemControllerSelf = this
	
	var currRecordSet = null
	var currRecordSetWindowSize = 1
	var currListContext = viewListContext
	
	this.populateListViewWithListItemContainers = function (viewListContext,populationDoneCallback) {
		
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
				reloadRecordsIntoContainers()
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
		
			listItemControllerSelf.$parentContainer.append($listItemContainer)
			
			return listItemInfo
		}
		
		this.listItemsInfo = []
		this.$parentContainer.empty()
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
	
	this.setPageSize = function(newPageSize) {
		currRecordSetWindowSize = Number(newPageSize)
		
		// Re-initialize the list view, then repopulate it with the records.
		listItemControllerSelf.populateListViewWithListItemContainers(currListContext,function() {
			currRecordSet.setWindowSize(currRecordSetWindowSize)
			reloadRecordsIntoContainers()
		})
		
	}
	
	this.setForm = function(formID) {
		var newViewContext = {
			databaseID: currListContext.databaseID,
			listID: currListContext.listID,
			formID: formID
		}
		// Re-initialize the list view, then repopulate it with the records.
		listItemControllerSelf.populateListViewWithListItemContainers(newViewContext,function() {
			if (currRecordSet !== null) {
				reloadRecordsIntoContainers()	
			}
		})
	}
	
	this.setFormAndPageSize = function(formID, pageSize) {
		currRecordSetWindowSize = Number(pageSize)
		if(currRecordSet != null) {
			currRecordSet.setWindowSize(currRecordSetWindowSize)
		}
		this.setForm(formID)
	}
	
	
	
	function enableRecordButtons(isEnabled)
	{
		var isDisabled = true;
		if(isEnabled) { isDisabled=false }
		$('#prevRecordButton').prop("disabled",isDisabled)
		$('#nextRecordButton').prop("disabled",isDisabled)
	}
		
	this.setRecordData =  function(recordData) {
		
		currRecordSet = new RecordSet(recordData,currRecordSetWindowSize);
		if(currRecordSet.numRecords() > 0) {
			reloadRecordsIntoContainers()		
		}
	
		// Enable the buttons to page through the records
		if(currRecordSet.numRecords() > 0) {
			enableRecordButtons(true)
		}
	
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
		
}