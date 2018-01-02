

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
		
		var listItemsInfo = []
		this.$parentContainer.empty()
		for(var listIndex = 0; listIndex < currRecordSetWindowSize; listIndex++) {
			var listItemInfo = populateOneListItem(listIndex)
			listItemsInfo.push(listItemInfo)
		}
		loadMultipleFormViewContainers(viewListContext,listItemsInfo,populationDoneCallback)
	
	}
	
	
	function reloadRecordsIntoContainers() {
		
		var listItemIndex = 0
		
		listItemControllerSelf.$parentContainer.find(".listItemContainer").each(function() {
			
			var $listItemContainer = $(this)		
			var recordRef = currRecordSet.getRecordRefAtWindowIndex(listItemIndex)
			if(recordRef !== null)
			{
				$listItemContainer.show()
				loadRecordIntoFormLayout($listItemContainer,recordRef)
	
				// Update footer to reflect where the current record is in list of currently loaded records
				$('#recordNumLabel').text(currRecordSet.recPageLabel())
		
		
			} else {
				$listItemContainer.hide()
			}
			listItemIndex++
		})
		
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
		reloadRecordsIntoContainers()		
	
		// Enable the buttons to page through the records
		if(currRecordSet.numRecords() > 0) {
			enableRecordButtons(true)
		}
	
	}
	
	
	// Initially disabled the buttons for paging through the records. They'll be 
	// enabled once the records are loaded.
	enableRecordButtons(false)
	
	var $nextRecordButton = $('#nextRecordButton')
	
	// ListItemController is a singleton, so only 1 controller can be bound to the
	// the next or previous buttons. The ListItemController will be re-allocated when
	// navigating to the list, so only the latest ListItemController can respond
	// to the list navigation events.
	$nextRecordButton.unbind("click")
	$nextRecordButton.click(function(e){
	         e.preventDefault();
			 if(currRecordSet.advanceToNextPage()) {
			 	reloadRecordsIntoContainers()
			 }
	});
	
	var $prevRecordButton = $('#prevRecordButton')
	$prevRecordButton.unbind("click")
	$prevRecordButton.click(function(e){
	         e.preventDefault();
			 if(currRecordSet.advanceToPrevPage()) {
				 console.log("Advance to next record")
			 	reloadRecordsIntoContainers()
			 } 
	});
		
}