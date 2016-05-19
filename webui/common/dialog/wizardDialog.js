function updateDialogProgress(progressSelector, progressVal) {
	console.log("Update progress: " + progressSelector + " val:" + progressVal)
	$(progressSelector).css('width', progressVal+'%').attr('aria-valuenow', progressVal);
}

function transitionToNextWizardDlgPanel(dialog, progressSelector, currPanelConfig, nextPanelConfig) {
	function showNextPanel() {
		$(nextPanelConfig.divID).show("slide",{direction:"right"},200);
	}
	$(currPanelConfig.divID).hide("slide",{direction:"left"},200,showNextPanel);
	
	updateDialogProgress(progressSelector,nextPanelConfig.progressPerc)
	
	$(dialog).dialog("option","buttons",nextPanelConfig.dlgButtons)
}

function setWizardDialogButtons(dialog,buttons) {
	$(dialog).dialog("option","buttons",buttons)
}

function transitionToNextWizardDlgPanelByID(dialog, progressSelector, currPanelID, nextPanelID) {
	
	var panelConfigByID = $(dialog).data("wizardDialogPanelsByID")
	assert(panelConfigByID !== undefined, "Can't call this transitionToNextWizardDlgPanelByID before openWizardDialog()")
	assert(panelConfigByID[currPanelID] !== undefined, 
		"transitionToNextWizardDlgPanelByID: panel id not configured for dialog:" + currPanelID)
	assert(panelConfigByID[nextPanelID] !== undefined, 
			"transitionToNextWizardDlgPanelByID: panel id not configured for dialog:" + nextPanelID)
	
	transitionToNextWizardDlgPanel(dialog,progressSelector,
			panelConfigByID[currPanelID].config,panelConfigByID[nextPanelID].config)
}


function transitionToPrevWizardDlgPanel(dialog, progressSelector, currPanelConfig, prevPanelConfig) {
	function showPrevPanel() {
		$(prevPanelConfig.divID).show("slide",{direction:"left"},200);
	}
	$(currPanelConfig.divID).hide("slide",{direction:"right"},200,showPrevPanel);
	
	updateDialogProgress(progressSelector,prevPanelConfig.progressPerc)
			
	$(dialog).dialog("option","buttons",prevPanelConfig.dlgButtons)
}

function transitionToPrevWizardDlgPanelByPanelID(dialog, progressSelector, currPanelID, prevPanelID) {
	
	var panelConfigByID = $(dialog).data("wizardDialogPanelsByID")
	assert(panelConfigByID !== undefined, "Can't call this transitionToPrevWizardDlgPanelByPanelID before openWizardDialog()")
	assert(panelConfigByID[currPanelID] !== undefined, 
		"transitionToPrevWizardDlgPanelByPanelID: panel id not configured for dialog:" + currPanelID)
	assert(panelConfigByID[prevPanelID] !== undefined, 
			"transitionToPrevWizardDlgPanelByPanelID: panel id not configured for dialog:" + prevPanelID)
	
	transitionToPrevWizardDlgPanel(dialog,progressSelector,
			panelConfigByID[currPanelID].config,panelConfigByID[prevPanelID].config)
}


function getFormFormInfoByPanelID(dialog, panelID) {
	var panelInfoByID = $(dialog).data("wizardDialogPanelsByID")
	
	var panelInfo = panelInfoByID[panelID]
	assert(panelInfo !== undefined, "Missing panel information for panel ID = " + panelID)
	return panelInfo.formInfo
}

function openWizardDialog(dlgParams) {
				
	var firstPanelConfig = dlgParams.panels[0]
	
	var dialog = $(dlgParams.dialogDivID)
			
    $(dlgParams.dialogDivID).dialog({
		autoOpen: false,
		width: dlgParams.width,
		height: dlgParams.height, 
		resizable: false,
		modal: true,
		buttons: firstPanelConfig.dlgButtons,
		close: dlgParams.closeFunc
    });
 
    $(dlgParams.dialogDivID).find( "form" ).on( "submit", function( event ) {
      	event.preventDefault();
		//saveNewTextBox()
		// TODO - reimplement save with enter key
		console.log("Save not implemented with enter key")
    });
	
	$( ".wizardPanel" ).hide() // hide all the panels
	$(firstPanelConfig.divID).show() // show the first panel
	$(dlgParams.dialogDivID).dialog("option","buttons",firstPanelConfig.dlgButtons)
	
	// Clear any previous entries validation errors. The message blocks by 
	// default don't clear their values with 'clear', so any remaining error
	// messages need to be removed from the message blocks within the panels.
// TODO - Use Bootstrap form validation to clear any previous errors
//	$('.wizardPanel').form('clear') // clear any previous entries
	$('.wizardErrorMsgBlock').empty()

	updateDialogProgress(dlgParams.progressDivID,10)
	
//	$(dlgParams.progressDivID).progress({percent:0});
	
	var panelsByID = {}
	for(var panelIndex = 0; panelIndex != dlgParams.panels.length; panelIndex++) {
		
		var panelConfig = dlgParams.panels[panelIndex]
		
		// Each panel is expected to return an object from initialization which can later
		// be used to reference the panel and form fields to gather up the 
		var panelFormInfo = dlgParams.panels[panelIndex].initPanel(dialog)
		assert(panelFormInfo !== undefined, "Panel form info not returned from init function")
		
		var panelInfo = {
			config: panelConfig,
			formInfo: panelFormInfo
		}
		
		assert(panelConfig.panelID !== undefined, "Missing panelID on dialog panel configuration")
		panelsByID[panelConfig.panelID] = panelInfo
	}

	// Store a map of panel IDs to the panel configurations. Different dialog panels can reference these unique panel IDs
	// when transitioning to the next panel or previous panel (see functions above).
	$(dlgParams.dialogDivID).data("wizardDialogPanelsByID",panelsByID)
	
	$(dlgParams.dialogDivID).dialog("open")
	
} // openWizardDialog

function initWizardDialog(dialogDivID) {
	// Initialize the newBarchart dialog with the minimum parameters. This is necessary
	// to hide the dialog from view when the document is initially loaded. The
	// dialog is fully re-initialized just prior to it being opened.
    $(dialogDivID).dialog({ autoOpen: false })

	
}

