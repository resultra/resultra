function initUILayoutPanes()
{
	var zeroPaddingInset = { top:0, bottom:0, left:0, right:0 }
	
	// Initialize the page layout
	var mainLayout = $('#layoutPage').layout({
		inset: zeroPaddingInset,
		north: fixedUILayoutPaneParams(40),
	})
				
}

$(document).ready(function() {	
	 
	initUILayoutPanes()
				
	initUserDropdownMenu()
	
	var $submitFormPageCanvas = $('#submitFormPageLayoutCanvas')
	
	
	var newRecordsParams = {
		parentDatabaseID:submitFormPageContext.databaseID,
		isDraftRecord:true
	}
	jsonAPIRequest("recordUpdate/newRecord",newRecordsParams,function(newRecordRef) {

		var newRecord = newRecordRef
		
		function getCurrentRecord() {
			return newRecord
		}
	
		function updateCurrentRecord(updatedRecordRef) {
			newRecord = updatedRecordRef
			loadRecordIntoFormLayout($submitFormPageCanvas,newRecord)
		}
		
		function finalizePageAfterFormComponentsLoaded() {
			console.log("Submit form: form components loaded")
		}
	
		var recordProxy = {
			changeSetID: MainLineFullyCommittedChangeSetID,
			getRecordFunc: getCurrentRecord,
			updateRecordFunc: updateCurrentRecord
		}
		
		
		var formContext = {
			databaseID: submitFormPageContext.databaseID,
			formID: submitFormPageContext.formID
		}
		
		
		loadFormViewComponents($submitFormPageCanvas,formContext,recordProxy,finalizePageAfterFormComponentsLoaded)
		
		initButtonClickHandler('#submitFormPageSubmitButton', function() {
			
			
			validateFormValues($submitFormPageCanvas,function(validationResult) {
				if(validationResult === true) {
					console.log("Saving form results")
			
					var recordDraftParams = {
						recordID: newRecord.recordID,
						isDraftRecord: false
					}
			
					jsonAPIRequest("record/setDraftStatus",recordDraftParams,function(response) {
						console.log("Submit form: form data submitted/saved")
						// Reset the form and re-load a different record.
						jsonAPIRequest("recordUpdate/newRecord",newRecordsParams,function(newRecordRef) {
							updateCurrentRecord(newRecordRef)
						})
					})
				} else {
					console.log("Form validation failed: not saving form results")
				}
			})	

		})
		
		var formLinkParams = { formLinkID: submitFormPageContext.formLinkID }
		jsonAPIRequest("formLink/get",formLinkParams,function(formLink) {
			var defaultVals = formLink.properties.defaultValues
			if (defaultVals.length > 0) {
				// Apply the default values before loading the form.
				var defaultValParams = {
					parentDatabaseID: submitFormPageContext.databaseID,
					recordID: newRecordRef.recordID,
					changeSetID: MainLineFullyCommittedChangeSetID,
					defaultVals: defaultVals }
				jsonAPIRequest("recordUpdate/setDefaultValues",defaultValParams,function(updatedRecordRef) {				
					updateCurrentRecord(updatedRecordRef)
				})
			}
		})
		
		
		

		
	}) // getRecord
	
	
					
}); // document ready