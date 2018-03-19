function initItemListPropsSettingsPageContent(databaseID, listInfo) {


	function initItemListNameProperties(listInfo) {

		var $nameInput = $('#itemListPropsNameInput')

		var $listNameForm = $('#itemListNamePropertyForm')

		$nameInput.blur() // prevent auto-focus

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
		
		var $formSelection = $("#itemListDefaultFormSelection")	
		
		function setDefaultView(viewParams) {
			var setDefaultViewParams = {
				listID: listInfo.listID,
				view: viewParams
			}
			jsonAPIRequest("itemList/setDefaultView",setDefaultViewParams,function(saveReply) {
				console.log("Done setting default view for list (form)")
			})
		}
		var defaultViewConfig = {
			setViewCallback: setDefaultView,
			databaseID: databaseID,
			initialView: listInfo.properties.defaultView
		}
		initItemListViewSelection(defaultViewConfig)
		
	} // initItemListFormProperties


		var listElemPrefix = "itemList_"

		var getItemListParams = { listID: listInfo.listID }
		jsonAPIRequest("itemList/get",getItemListParams,function(listInfo) {
			
			var filterPropertyPanelParams = {
				elemPrefix: listElemPrefix,
				databaseID: databaseID,
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
			
			var filterFieldParams = {
				elemPrefix: listElemPrefix,
				databaseID: databaseID,
				defaultFilterFields: listInfo.properties.defaultFilterFields,
				setFilterFieldsCallback: function(fields) {
					var setFieldsParams = {
						listID: listInfo.listID,
						defaultFilterFields: fields
					}
					jsonAPIRequest("itemList/setDefaultFilterFields",setFieldsParams,function(updatedList) {
						console.log(" Default filters updated")
					}) // set record's number field value
					
				}
			}
			initFilterFieldSelection(filterFieldParams)


			var preFilterElemPrefix = "itemListPreFilter_"
			var preFilterPropertyPanelParams = {
				elemPrefix: preFilterElemPrefix,
				databaseID: databaseID,
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
				databaseID: databaseID,
				resortFunc: function() {}, // no-op
				initDoneFunc: function() {}, // no-op
				saveUpdatedSortRulesFunc: saveDefaultListSortRules}
			initSortRecordsPane(sortPaneParams)
				
			var sortFieldParams = {
				elemPrefix: "itemListSortFields_",
				label: "Limit sorting to selected fields",
				databaseID: databaseID,
				defaultFields: listInfo.properties.defaultSortFields,
				setFieldsCallback: function(fields) {
					var setFieldsParams = {
						listID: listInfo.listID,
						defaultSortFields: fields
					}
					jsonAPIRequest("itemList/setDefaultSortFields",setFieldsParams,function(updatedList) {
						console.log(" Default filters updated")
					}) // set record's number field value
				}
			}
			initFieldSelectionChecklist(sortFieldParams)
			
			initListRolePrivProperties(listInfo.listID)

			initItemListNameProperties(listInfo)

			initItemListFormProperties(listInfo)

			initAlternateFormsProperties(listInfo)


		}) // set record's number field value
				

		initSettingsPageButtonLink('#listPropsBackToListListLink',"lists")
		
}



function navigateToItemListProps(databaseID,listInfo) {
	var editPropsContentURL = '/admin/itemList/' + listInfo.listID
	setSettingsPageContent(editPropsContentURL,function() {
		initItemListPropsSettingsPageContent(databaseID,listInfo)
	})
}

