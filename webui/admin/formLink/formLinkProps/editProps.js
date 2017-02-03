$(document).ready(function() {
	
	
	function initFormLinkNameProperties(linkInfo) {
	
/*		var $nameInput = $('#itemListPropsNameInput')
	
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
	*/
	
	} // initFormLinkNameProperties
		
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
		
	var tocConfig = {
		databaseID: formLinkPropsContext.databaseID,
		newItemFormButtonFunc: openSubmitFormDialog
	}
	initDatabaseTOC(tocConfig)
		
	initUserDropdownMenu()
		
		var formLinkElemPrefix = "formLink_"
		
		var getFormLinkParams = { linkID: formLinkPropsContext.linkID }
//		jsonAPIRequest("formLink/get",getFormLinkParams,function(linkInfo) {
//			initFormLinkNameProperties(listInfo)
//		})
	
})