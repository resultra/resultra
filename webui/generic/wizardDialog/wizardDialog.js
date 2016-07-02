

function updateDialogProgress($dialog,progressVal) {
	var dlgParams = $dialog.data("dialogParams")
	
	console.log("Update progress: " + dlgParams.progressDivID + " val:" + progressVal)
	
	$(dlgParams.progressDivID).css('width', progressVal+'%').attr('aria-valuenow', progressVal);
}

function updateDialogToPanelConfig($dialog, panelConfig) {
		
	panelConfig.transitionIntoPanel($dialog)
	updateDialogProgress($dialog,panelConfig.progressPerc)
	$dialog.data("currPanelConfig",panelConfig)
}

function transitionToNextWizardDlgPanel($dialog, nextPanelConfig) {
	function showNextPanel() {
		$(nextPanelConfig.divID).show("slide",{direction:"right"},200);
	}
	
	var currPanelConfig = $dialog.data("currPanelConfig")
	$(currPanelConfig.divID).hide("slide",{direction:"left"},200,showNextPanel);
	
	updateDialogToPanelConfig($dialog,nextPanelConfig)
}

function transitionToPrevWizardDlgPanel($dialog, prevPanelConfig) {
	function showPrevPanel() {
		$(prevPanelConfig.divID).show("slide",{direction:"left"},200);
	}

	var currPanelConfig = $dialog.data("currPanelConfig")
	$(currPanelConfig.divID).hide("slide",{direction:"right"},200,showPrevPanel);

	updateDialogToPanelConfig($dialog,prevPanelConfig)
}


function setWizardDialogButtonSet(buttonClassName) {
	$('.wizardPanelButton').hide() // hide all buttons with the wizardPanelButton class
	$('.'+buttonClassName).show() // show specific buttons with the given class name
}


// getWizardDialogPanelVals should be called after a panel has fully validated its
// form contents and is ready to transition out of the panel. Other schemes for storing
// the panel's data for later retrieval are possible, but this method is considered a good
// separation of concerns for the dialog panel; in other words the dialog panel code has
// full knowledge of the HTML fields *and* value format from those fields (e.g., number vs text)
// is thus best positioned to "package up" a data object with the fully validated results.
function getWizardDialogPanelVals($dialog,panelID) {
	
	var panelsByID = $dialog.data("wizardDialogPanelsByID")
	
	var panelVals = panelsByID[panelID].getPanelVals()
	
	return panelVals
}

function transitionToNextWizardDlgPanelByID($dialog, nextPanelID) {
	
	var panelConfigByID = $dialog.data("wizardDialogPanelsByID")
		
	transitionToNextWizardDlgPanel($dialog,panelConfigByID[nextPanelID])
}


function transitionToPrevWizardDlgPanelByPanelID($dialog, prevPanelID) {
	
	var panelConfigByID = $dialog.data("wizardDialogPanelsByID")
		
	transitionToPrevWizardDlgPanel($dialog,panelConfigByID[prevPanelID])
}

function openWizardDialog(dlgParams) {
	var firstPanelConfig = dlgParams.panels[0]
	
	var $dialog = $(dlgParams.dialogDivID)
	
	$dialog.removeData() // remove all jQuery data from the dialog
	$dialog.data("dialogParams",dlgParams)
			 
    $dialog.find( "form" ).on( "submit", function( event ) {
      	event.preventDefault();
		//saveNewTextBox()
		// TODO - reimplement save with enter key
		console.log("Save not implemented with enter key")
    });
	
	$( ".wizardPanel" ).hide() // hide all the panels
	$(firstPanelConfig.divID).show() // show the first panel
	
	// Optional minBodyHeight parameter
	if(dlgParams.hasOwnProperty('minBodyHeight')) {
		$dialog.find('.modal-body').css('min-height',dlgParams.minBodyHeight)
	}
	
	updateDialogToPanelConfig($dialog,firstPanelConfig)
			
	var panelsByID = {}
	for(var panelIndex = 0; panelIndex != dlgParams.panels.length; panelIndex++) {
		
		var panelConfig = dlgParams.panels[panelIndex]
		
		dlgParams.panels[panelIndex].initPanel($dialog)
				
		assert(panelConfig.panelID !== undefined, "Missing panelID on dialog panel configuration")
		panelsByID[panelConfig.panelID] = panelConfig
	}

	// Store a map of panel IDs to the panel configurations. Different dialog panels can reference these unique panel IDs
	// when transitioning to the next panel or previous panel (see functions above).
	$dialog.data("wizardDialogPanelsByID",panelsByID)
	
	$dialog.modal('show')
	
}

