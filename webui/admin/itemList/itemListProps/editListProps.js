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
			parentTableID: listInfo.parentTableID,
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
		
	initDatabaseTOC(itemListPropsContext.databaseID)
		
	initUserDropdownMenu()
		
		var listElemPrefix = "itemList_"
		
		var getItemListParams = { listID: itemListPropsContext.listID }
		jsonAPIRequest("itemList/get",getItemListParams,function(listInfo) {
			var filterPropertyPanelParams = {
				elemPrefix: listElemPrefix,
				tableID: listInfo.parentTableID,
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
				tableID: listInfo.parentTableID,
				resortFunc: function() {}, // no-op
				initDoneFunc: function() {}, // no-op
				saveUpdatedSortRulesFunc: saveDefaultListSortRules}
	
	
			initSortRecordsPane(sortPaneParams)
				
			initItemListNameProperties(listInfo)
				
			initItemListFormProperties(listInfo)
				

		}) // set record's number field value
	
})