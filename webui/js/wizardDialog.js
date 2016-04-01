function transitionToNextWizardDlgPanel(dialog, progressSelector, currPanelConfig, nextPanelConfig) {
	function showNextPanel() {
		$(nextPanelConfig.divID).show("slide",{direction:"right"},200);
	}
	$(currPanelConfig.divID).hide("slide",{direction:"left"},200,showNextPanel);
	
	$(progressSelector).progress({percent:nextPanelConfig.progressPerc});
	
	
	$(dialog).dialog("option","buttons",nextPanelConfig.dlgButtons)
}

function transitionToPrevWizardDlgPanel(dialog, progressSelector, currPanelConfig, prevPanelConfig) {
	function showPrevPanel() {
		$(prevPanelConfig.divID).show("slide",{direction:"left"},200);
	}
	$(currPanelConfig.divID).hide("slide",{direction:"right"},200,showPrevPanel);
	
	$(progressSelector).progress({percent:prevPanelConfig.progressPerc});
	
	$(dialog).dialog("option","buttons",prevPanelConfig.dlgButtons)
}


function openWizardDialog(dlgParams) {
				
	var firstPanelConfig = dlgParams.panels[0]
			
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
	$('.wizardPanel').form('clear') // clear any previous entries
	$('.wizardErrorMsgBlock').empty()
	
	
	$(dlgParams.progressDivID).progress({percent:0});
	
	for(var panelIndex = 0; panelIndex != dlgParams.panels.length; panelIndex++) {
		dlgParams.panels[panelIndex].initPanel()
	}

	$(dlgParams.dialogDivID).dialog("open")
	
} // openWizardDialog

function initWizardDialog(dialogDivID) {
	// Initialize the newBarchart dialog with the minimum parameters. This is necessary
	// to hide the dialog from view when the document is initially loaded. The
	// dialog is fully re-initialized just prior to it being opened.
    $(dialogDivID).dialog({ autoOpen: false })

	
}
