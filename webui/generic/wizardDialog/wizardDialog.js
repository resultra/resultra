function updateDialogProgress(dialog,progressVal) {
	var dlgParams = $(dialog).data("dialogParams")
	
	console.log("Update progress: " + dlgParams.progressDivID + " val:" + progressVal)
	
	$(dlgParams.progressDivID).css('width', progressVal+'%').attr('aria-valuenow', progressVal);
}

function updateDialogToPanelConfig(dialog, panelConfig) {
	updateDialogProgress(dialog,panelConfig.progressPerc)
	$(dialog).data("currPanelConfig",panelConfig)
	$(dialog).dialog("option","buttons",panelConfig.dlgButtons)	
}

function transitionToNextWizardDlgPanel(dialog, nextPanelConfig) {
	function showNextPanel() {
		$(nextPanelConfig.divID).show("slide",{direction:"right"},200);
	}
	
	var currPanelConfig = $(dialog).data("currPanelConfig")
	$(currPanelConfig.divID).hide("slide",{direction:"left"},200,showNextPanel);
	
	updateDialogToPanelConfig(dialog,nextPanelConfig)
}

function transitionToPrevWizardDlgPanel(dialog, prevPanelConfig) {
	function showPrevPanel() {
		$(prevPanelConfig.divID).show("slide",{direction:"left"},200);
	}

	var currPanelConfig = $(dialog).data("currPanelConfig")
	$(currPanelConfig.divID).hide("slide",{direction:"right"},200,showPrevPanel);

	updateDialogToPanelConfig(dialog,prevPanelConfig)
}


function setWizardDialogButtons(dialog,buttons) {
	$(dialog).dialog("option","buttons",buttons)
}

function transitionToNextWizardDlgPanelByID(dialog, nextPanelID) {
	
	var panelConfigByID = $(dialog).data("wizardDialogPanelsByID")
		
	transitionToNextWizardDlgPanel(dialog,panelConfigByID[nextPanelID].config)
}


function transitionToPrevWizardDlgPanelByPanelID(dialog, prevPanelID) {
	
	var panelConfigByID = $(dialog).data("wizardDialogPanelsByID")
		
	transitionToPrevWizardDlgPanel(dialog,panelConfigByID[prevPanelID].config)
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
	
	dialog.data("dialogParams",dlgParams)
			
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
	updateDialogToPanelConfig(dialog,firstPanelConfig)
	
	$(dlgParams.dialogDivID).dialog("option","buttons",firstPanelConfig.dlgButtons)
	
	// Clear any previous entries validation errors. The message blocks by 
	// default don't clear their values with 'clear', so any remaining error
	// messages need to be removed from the message blocks within the panels.
// TODO - Use Bootstrap form validation to clear any previous errors
//	$('.wizardPanel').form('clear') // clear any previous entries
	$('.wizardErrorMsgBlock').empty()
		
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

