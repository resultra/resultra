function transitionToNextDlgPanel(dialog, currPanelConfig, nextPanelConfig) {
	function showNextPanel() {
		$('#'+ nextPanelConfig.divID).show("slide",{direction:"right"},200);
	}
	$("#" + currPanelConfig.divID).hide("slide",{direction:"left"},200,showNextPanel);
	
	$('#newBarChartProgress').progress({percent:nextPanelConfig.progressPerc});
	
	
	$(dialog).dialog("option","buttons",nextPanelConfig.dlgButtons)
}

function transitionToPrevDlgPanel(dialog, currPanelConfig, prevPanelConfig) {
	function showPrevPanel() {
		$('#'+ prevPanelConfig.divID).show("slide",{direction:"left"},200);
	}
	$("#" + currPanelConfig.divID).hide("slide",{direction:"right"},200,showPrevPanel);
	
	$('#newBarChartProgress').progress({percent:prevPanelConfig.progressPerc});
	
	$(dialog).dialog("option","buttons",prevPanelConfig.dlgButtons)
}


var barChartXAxisPanelConfig = {
	divID: "newBarChartDlgXAxisPanel",
	progressPerc:80,
	dlgButtons: { 
		"Next" : function() { 
			if($( "#newBarChartDlgXAxisPanel" ).form('validate form')) {
				transitionToNextDlgPanel(this,barChartXAxisPanelConfig,barChartYAxisPanelConfig)
			} // if validate panel's form
		},
		"Cancel" : function() { $(this).dialog('close'); },
 	}, // dialog buttons
	initPanel: function() {
		$('#xAxisFieldSelection').dropdown()
		$('#xAxisGroupBySelection').dropdown()
		$('#xAxisSortSelection').dropdown()
	},	
}


var barChartYAxisPanelConfig = {
	divID: "newBarChartDlgYAxisPanel",
	progressPerc:20,
	dlgButtons: { 
		"Previous": function() {
			transitionToPrevDlgPanel(this,barChartYAxisPanelConfig,barChartXAxisPanelConfig)	
		 },
		"Done" : function() { 
			if($( "#newBarChartDlgYAxisPanel" ).form('validate form')) {
				$(this).dialog('close');
			} // if validate panel's form
		},
		"Cancel" : function() { $(this).dialog('close'); },
 	}, // dialog buttons
	initPanel: function() {
		$('#yAxisFieldSelection').dropdown()
		$('#yAxisSummarySelection').dropdown()
	},	
}


function newBarChart(barChartParams) {
	
	newBarChartParams = {
		containerParams: barChartParams,
		barChartCreated: false,
		placeholderID: barChartParams.containerID,
		dialogBox: $( "#newBarchartDialog" )
	}
		
	var firstPanelConfig = barChartXAxisPanelConfig
			
    newBarChartParams.dialogBox.dialog({
      autoOpen: false,
      height: 500, width: 550,
	  resizable: false,
      modal: true,
      buttons: firstPanelConfig.dlgButtons,
      close: function() {
		  console.log("Close dialog")
		  if(!newBarChartParams.barChartCreated)
		  {
			  // If the the text box creation is not complete, remove the placeholder
			  // from the canvas.
			  $('#'+newBarChartParams.placeholderID).remove()
		  }
      }
    });
 
    newBarChartParams.dialogBox.find( "form" ).on( "submit", function( event ) {
      	event.preventDefault();
		//saveNewTextBox()
		// TODO - reimplement save with enter key
		console.log("Save not implemented with enter key")
    });
	
	$( ".wizardPanel" ).hide() // hide all the panels
	$( "#" + firstPanelConfig.divID).show() // show the first panel
	newBarChartParams.dialogBox.dialog("option","buttons",firstPanelConfig.dlgButtons)
	
	// Clear any previous entries validation errors. The message blocks by 
	// default don't clear their values with 'clear', so any remaining error
	// messages need to be removed from the message blocks within the panels.
	$('.wizardPanel').form('clear') // clear any previous entries
	$('.wizardErrorMsgBlock').empty()
	
	
	$('#newBarChartProgress').progress({percent:0});

	barChartXAxisPanelConfig.initPanel()
	barChartYAxisPanelConfig.initPanel()

	newBarChartParams.dialogBox.dialog("open")
	
} // newBarChart


function initNewBarChartDialog() {
	// Initialize the newBarchart dialog with the minimum parameters. This is necessary
	// to hide the dialog from view when the document is initially loaded. The
	// dialog is fully re-initialized just prior to it being opened.
    $( "#newBarchart" ).dialog({ autoOpen: false })

	
}
