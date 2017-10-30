$(document).ready(function() {
	
	function initSharedLinkProperties(linkInfo) {
		
		function initSharedLinkPropertyControls(linkInfo) {
			var $linkDisplay = $('#adminFormLinkShareLinkLink')
			var $linkFormGroup = $('#adminFormLinkShareLinkLinkFormGroup')
			
			if(linkInfo.sharedLinkEnabled) {
				$linkFormGroup.show()
				var linkURL = formLinkPropsContext.siteBaseURL + "submitForm/" + linkInfo.sharedLinkID
				$linkDisplay.val(linkURL)
			} else {
				$linkFormGroup.hide()
				$linkDisplay.val("")
			}
		}
				
		initCheckboxChangeHandler('#adminFormLinkShareLink', 
					linkInfo.sharedLinkEnabled, function(sharedLinkEnabled) {
						
			console.log("Form link shared link enabled: " + sharedLinkEnabled)
						
			if (sharedLinkEnabled) {
				var enableSharedLinkParams = { formLinkID: linkInfo.linkID }
				jsonAPIRequest("formLink/enableSharedLink",enableSharedLinkParams,function(updatedLinkInfo) {
					console.log("Done setting form for formLink: " + JSON.stringify(updatedLinkInfo))
					initSharedLinkPropertyControls(updatedLinkInfo)
				})			
			
			} else {
				var disableSharedLinkParams = { formLinkID: linkInfo.linkID }
				jsonAPIRequest("formLink/disableSharedLink",disableSharedLinkParams,function(updatedLinkInfo) {
					console.log("Done setting form for formLink: " + JSON.stringify(updatedLinkInfo))
					initSharedLinkPropertyControls(updatedLinkInfo)
				})			
			
			}
	
		})
		initSharedLinkPropertyControls(linkInfo)
		
	}
	
	
	function initFormLinkNameProperties(linkInfo) {
	
		var $nameInput = $('#formLinkPropsNameInput')
	
		var $nameForm = $('#formLinkNamePropertyForm')
		
		$nameInput.val(linkInfo.name)
		
		
		var remoteValidationParams = {
			url: '/api/generic/stringValidation/validateItemLabel',
			data: {
				label: function() { return $nameInput.val(); }
			}
		}
	
		var validationSettings = createInlineFormValidationSettings({
			rules: {
				itemListPropsNameInput: {
					minlength: 3,
					required: true,
					remote: remoteValidationParams
				}
			}
		})	
	
	
		var validator = $nameForm.validate(validationSettings)
	
		initInlineInputValidationOnBlur(validator,'#formLinkPropsNameInput',
			remoteValidationParams, function(validatedName) {		
				var setNameParams = {
					formLinkID:linkInfo.linkID,
					newName:validatedName
				}
				jsonAPIRequest("formLink/setName",setNameParams,function(updatedLinkInfo) {
					console.log("Done changing form link name: " + validatedName)
				})
		})	

		validator.resetForm()
	
	
	} // initFormLinkNameProperties
	
	
	function initFormLinkFormProperties(linkInfo) {
		var selectFormParams = {
			menuSelector: "#formLinkPropFormSelection",
			parentDatabaseID: formLinkPropsContext.databaseID,
			initialFormID: linkInfo.formID
		}	
		populateFormSelectionMenu(selectFormParams)
		var $formSelection = $("#formLinkPropFormSelection")
		initSelectControlChangeHandler($formSelection, function(selectedFormID) {

			var setFormParams = {
				formLinkID: linkInfo.linkID,
				formID: selectedFormID
			}	
			jsonAPIRequest("formLink/setForm",setFormParams,function(updatedLinkInfo) {
				console.log("Done setting form for formLink")
			})			
		})
		
	} // initItemListFormProperties
	
	function initIncludeInSidebarProperty(linkInfo) {
		
		initCheckboxChangeHandler('#adminFormLinkIncludeInSidebar', 
					linkInfo.includeInSidebar, function(newVal) {
			var setIncludeSidebarParams = {
				formLinkID: linkInfo.linkID,
				includeInSidebar: newVal
			}
			jsonAPIRequest("formLink/setIncludeInSidebar",setIncludeSidebarParams,function(updatedLinkInfo) {
				console.log("Done setting form for formLink")
			})			
			
		})
		
	}
	
	var zeroPaddingInset = { top:0, bottom:0, left:0, right:0 }

	$('#editFormLinkPropsPage').layout({
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
		
	initAdminSettingsTOC(formLinkPropsContext.databaseID,"settingsTOCFormLinks")
		
		
	initUserDropdownMenu()
		
		var formLinkElemPrefix = "formLink_"
		
		var getFormLinkParams = { formLinkID: formLinkPropsContext.linkID }
		jsonAPIRequest("formLink/get",getFormLinkParams,function(linkInfo) {
			
			initFormLinkNameProperties(linkInfo)
			initFormLinkFormProperties(linkInfo)
			initIncludeInSidebarProperty(linkInfo)
			initSharedLinkProperties(linkInfo)
	
			var defaultValPropParams = {
				databaseID: formLinkPropsContext.databaseID,
				elemPrefix: "formLink_",
				defaultDefaultValues: linkInfo.properties.defaultValues,
				updateDefaultValues: function(updatedDefaultVals) {
					console.log("Updating default values for form button: " + JSON.stringify(updatedDefaultVals))
			
					var setDefaultValsParams = {
						formLinkID: linkInfo.linkID,
						defaultValues: updatedDefaultVals }
			
					jsonAPIRequest("formLink/setDefaultVals",setDefaultValsParams,function(updatedFormLink) {
					})
				}
			}
			initDefaultValuesPropertyPanel(defaultValPropParams)
	
	
	
		})
	
})