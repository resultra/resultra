$(document).ready(function() {
	
	
	function initItemListNameProperties(listInfo) {
	
		var $nameInput = $('#itemListPropsNameInput')
	
		var $listNameForm = $('#itemListNamePropertyForm')
		
		$nameInput.val(listInfo.name)
		
		
		var remoteValidationParams = {
			url: '/api/itemList/validateListName',
			data: {
				listID: function() { return listInfo.listID },
				listName: function() { return $nameInput.val() }
			}	
		}
	
		var validationSettings = createInlineFormValidationSettings({
			rules: {
				itemListPropsNameInput: {
					minlength: 3,
					required: true,
					remote: remoteValidationParams
				} // newRoleNameInput
			}
		})	
	
	
		var validator = $listNameForm.validate(validationSettings)
	
		initInlineInputValidationOnBlur(validator,'#itemListPropsNameInput',
			remoteValidationParams, function(validatedName) {		
				var setNameParams = {
					listID:listInfo.listID,
					newListName:validatedName
				}
				jsonAPIRequest("itemList/setName",setNameParams,function(listInfo) {
					console.log("Done changing list name: " + validatedName)
				})
		})	

		validator.resetForm()
	
	} // initItemListNameProperties
	
	function initItemListFormProperties(listInfo) {
		var selectFormParams = {
			menuSelector: "#itemListDefaultFormSelection",
			parentDatabaseID: itemListPropsContext.databaseID,
			initialFormID: listInfo.formID
		}	
		populateFormSelectionMenu(selectFormParams)
		var $formSelection = $("#itemListDefaultFormSelection")
		initSelectControlChangeHandler($formSelection, function(selectedFormID) {

			var setFormParams = {
				listID: listInfo.listID,
				formID: selectedFormID
			}	
			jsonAPIRequest("itemList/setForm",setFormParams,function(saveReply) {
				console.log("Done setting form for list")
			})			
		})
		
	} // initItemListFormProperties
	
	
	var zeroPaddingInset = { top:0, bottom:0, left:0, right:0 }


	$('#editItemListPropsPage').layout({
			inset: zeroPaddingInset,
			north: fixedUILayoutPaneParams(40),
			west: {
				size: 250,
				resizable:false,
				slidable: false,
				spacing_open:4,
				spacing_closed:4,
				initClosed:false // panel is initially open	
			}
		})
		
	var tocConfig = {
		databaseID: itemListPropsContext.databaseID,
		newItemFormButtonFunc: openSubmitFormDialog
	}
	initDatabaseTOC(tocConfig)
		
	initUserDropdownMenu()
		
		var listElemPrefix = "itemList_"
		
		var getItemListParams = { listID: itemListPropsContext.listID }
		jsonAPIRequest("itemList/get",getItemListParams,function(listInfo) {
			var filterPropertyPanelParams = {
				elemPrefix: listElemPrefix,
				databaseID: itemListPropsContext.databaseID,
				defaultFilterRules: listInfo.properties.defaultFilterRules,
				initDone: function () {},
				updateFilterRules: function (updatedFilterRules) {
					var setDefaultFiltersParams = {
						listID: listInfo.listID,
						filterRules: updatedFilterRules
					}
					jsonAPIRequest("itemList/setDefaultFilterRules",setDefaultFiltersParams,function(updatedList) {
						console.log(" Default filters updated")
					}) // set record's number field value
				
				}
			
			}
			initFilterPropertyPanel(filterPropertyPanelParams)
			
			
			var preFilterElemPrefix = "itemListPreFilter_"
			var preFilterPropertyPanelParams = {
				elemPrefix: preFilterElemPrefix,
				databaseID: itemListPropsContext.databaseID,
				defaultFilterRules: listInfo.properties.preFilterRules,
				initDone: function () {},
				updateFilterRules: function (updatedFilterRules) {
					var setPreFiltersParams = {
						listID: listInfo.listID,
						filterRules: updatedFilterRules
					}
					jsonAPIRequest("itemList/setPreFilterRules",setPreFiltersParams,function(updatedList) {
						console.log(" Pre filters updated")
					}) // set record's number field value
				
				}
			
			}
			initFilterPropertyPanel(preFilterPropertyPanelParams)
			
			function saveDefaultListSortRules(sortRules) {
				console.log("Saving default sort rules for list: " + JSON.stringify(sortRules))
				var saveSortRulesParams = {
					listID:listInfo.listID,
					sortRules: sortRules
				}
				jsonAPIRequest("itemList/setDefaultSortRules",saveSortRulesParams,function(saveReply) {
					console.log("Done saving default sort rules")
				})			

			}
	
	
			var sortPaneParams = {
				defaultSortRules: listInfo.properties.defaultRecordSortRules,
				databaseID: itemListPropsContext.databaseID,
				resortFunc: function() {}, // no-op
				initDoneFunc: function() {}, // no-op
				saveUpdatedSortRulesFunc: saveDefaultListSortRules}
				
				
			var $pageSizeSelection = $('#adminItemListPageSizeSelection')
			$pageSizeSelection.val(listInfo.properties.defaultPageSize)
			initNumberSelectionChangeHandler($pageSizeSelection,function(newPageSize) {
				var savePageSizeParams = {
					listID:listInfo.listID,
					pageSize: newPageSize
				}
				
				jsonAPIRequest("itemList/setDefaultPageSize",savePageSizeParams,function(saveReply) {
					console.log("Done saving default page size")
				})			
				
			})
	
	
			initSortRecordsPane(sortPaneParams)
				
			initItemListNameProperties(listInfo)
				
			initItemListFormProperties(listInfo)
				

		}) // set record's number field value
	
})