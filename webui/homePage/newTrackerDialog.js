function openNewTrackerDialog(pageContext) {
	
	var $newTrackerDialogForm = $('#newTrackerDialogForm')
	var $templateSelection = $('#newTrackerTemplateSelection')
	
	var validator = $newTrackerDialogForm.validate({
		rules: {
			newTrackerNameInput: {
				minlength: 3,
				required: true,
				remote: {
					url: '/api/database/validateNewTrackerName',
					data: {
						trackerName: function() { return $('#newTrackerNameInput').val(); }
					}
				} // remote
			}, // newFormNameInput
		},
		messages: {
			newTrackerNameInput: {
				required: "Tracker name is required"
			}
		}
	})

	validator.resetForm()
	
	
	function initNewTrackerDescriptionEditor(initialDescription) {
		var $descInput = $props.find(".adminGeneralTrackerDescriptionInput")
	
		function setTrackerDescription(description) {
			var setDescParams = {
				databaseID:trackerDatabaseInfo.databaseID,
				description:description
			}
			jsonAPIRequest("database/setDescription",setDescParams,function(dbInfo) {
			})
		
		}
	
		$descInput.html(trackerDatabaseInfo.description)
	
		$descInput.dblclick(function() {
			if (!inlineCKEditorEnabled($descInput)) {
			
				var editor = enableInlineCKEditor($descInput)
				$descInput.focus()
			
				editor.on('blur', function(event) {
					var popupMsg = editor.getData();
				
					setTrackerDescription(popupMsg)
							
					disableInlineCKEditor($descInput,editor)
				})
			
			}
		})
		
	}
	
	function getTemplateLists(templatesCallback) {
		
		var templateLists = {}
		var templateListsRemaining = 2
		function processOneTemplateList() {
			templateListsRemaining--
			if (templateListsRemaining <= 0) {
				templatesCallback(templateLists)
			}
		}
		
		var getDBListParams = {} // no parameters necessary - gets the tracker list for the currently signed in user
		jsonAPIRequest("database/getTemplateList",getDBListParams,function(accountTemplateList) {
			templateLists.accountTemplates = accountTemplateList
			processOneTemplateList()
		})
		
		jsonAPIRequest("database/getFactoryTemplateList",getDBListParams,function(factoryTemplateList) {
			templateLists.factoryTemplates = factoryTemplateList
			processOneTemplateList()
		})
		
		
	}
	
	var templateTrackerInfoByID = {}

	getTemplateLists(function(templateLists) {
		
		$templateSelection.empty()

		// Only include the "no template" option if there are no factory templates. The factory templates
		// should include a minimal template for setting up a starter tracker.
		if(templateLists.factoryTemplates.length <= 0) {
			$templateSelection.append(selectOptionHTML("","No template"))	
		}
		
		if(templateLists.accountTemplates.length > 0) {
			var $accountTemplateOptGroup = $('<optgroup label="Workspace Templates"></optgroup>')
			for (var trackerIndex=0; trackerIndex<templateLists.accountTemplates.length; trackerIndex++) {	
				var trackerInfo = templateLists.accountTemplates[trackerIndex]
				trackerInfo.templateSource = "account"
				$accountTemplateOptGroup.append(selectOptionHTML(trackerInfo.databaseID,trackerInfo.databaseName))
				templateTrackerInfoByID[trackerInfo.databaseID] = trackerInfo
			}
			$templateSelection.append($accountTemplateOptGroup)
		}
		if(templateLists.factoryTemplates.length > 0) {
			var $factoryTemplateOptGroup = $('<optgroup label="Factory Templates"></optgroup>')
			for (var trackerIndex=0; trackerIndex<templateLists.factoryTemplates.length; trackerIndex++) {	
				var trackerInfo = templateLists.factoryTemplates[trackerIndex]
				trackerInfo.templateSource = "factory"
				$factoryTemplateOptGroup.append(selectOptionHTML(trackerInfo.databaseID,trackerInfo.databaseName))
				templateTrackerInfoByID[trackerInfo.databaseID] = trackerInfo
			}
			$templateSelection.append($factoryTemplateOptGroup)
		}
		
		initSelectControlChangeHandler($templateSelection,function(selectedDatabaseID) {	
			
			var $descGroup = $('#newTrackerTemplateDescriptionGroup') 
				
			if (selectedDatabaseID.length > 0) {
				var trackerInfo = templateTrackerInfoByID[selectedDatabaseID]
				console.log("new template tracker database selected: " + JSON.stringify(trackerInfo))
			
				
				if(trackerInfo.description.length > 0) {
					var $templateDesc = $('#newTrackerTemplateDescription')
					$templateDesc.html(formatInlineContentHTMLDisplay(trackerInfo.description))
					$descGroup.show()
				} else {
					$descGroup.hide()
				}
				
			} else {
				$descGroup.hide()			
			}
			
		})
	})
		
	$('#newTrackerDialog').modal('show')
	
	initButtonClickHandler('#newTrackerSaveButton',function() {
		console.log("New Tracker save button clicked")
		if($newTrackerDialogForm.valid()) {	
			
			var selectedDatabaseID = $templateSelection.val()
			var templateSource = null
			if (selectedDatabaseID !== null && selectedDatabaseID.length > 0) {
				var trackerInfo = templateTrackerInfoByID[selectedDatabaseID]
				templateSource = trackerInfo.templateSource
			}
			
			
			var newTrackerParams = {  
				name: $('#newTrackerNameInput').val(),
				templateSource: templateSource,
				templateDatabaseID: selectedDatabaseID
			}
			jsonAPIRequest("database/new",newTrackerParams,function(newTrackerInfo) {
				console.log("Created new tracker: " + JSON.stringify(newTrackerInfo))
				addTrackerListItem(pageContext,newTrackerInfo)	
				$('#newTrackerDialog').modal('hide')
			})
			

		}
	})
	
}