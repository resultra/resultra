

// Initialized below in newLayoutContainer
var newTextBoxParams;

function transitionToNextDlgPanel(dialog, currPanelConfig, nextPanelConfig) {
	function showNextPanel() {
		$('#'+ nextPanelConfig.divID).show("slide",{direction:"right"},200);
	}
	$("#" + currPanelConfig.divID).hide("slide",{direction:"left"},200,showNextPanel);
	
	$('#newTextBoxProgress').progress({percent:nextPanelConfig.progressPerc});
	
	
	$(dialog).dialog("option","buttons",nextPanelConfig.dlgButtons)
}

function transitionToPrevDlgPanel(dialog, currPanelConfig, prevPanelConfig) {
	function showPrevPanel() {
		$('#'+ prevPanelConfig.divID).show("slide",{direction:"left"},200);
	}
	$("#" + currPanelConfig.divID).hide("slide",{direction:"right"},200,showPrevPanel);
	
	$('#newTextBoxProgress').progress({percent:prevPanelConfig.progressPerc});
	
	$(dialog).dialog("option","buttons",prevPanelConfig.dlgButtons)
}

function newFieldIsCalcField() {
	return $('#'+newFieldPanelConfig.divID).form('get field','isCalcField').prop('checked')
}

function doCreateNewFieldWithTextBox() {
	
	return $('#'+newOrExistingFieldPanelConfig.divID).form('get field','createNewFieldRadio').prop('checked')
}



function saveNewTextBox()
{
	// Complete the parameters to create the new layout container by including the 
	// fieldID selected in the dialog box.
	
	if(doCreateNewFieldWithTextBox()) {
		console.log("saveNewTextBox() with create new field - not implemented yet, closing dialog")
		newTextBoxParams.dialogBox.dialog("close")
	} else {
		var fieldID = $( "#newTextBoxDlgSelectOrNewFieldPanel" ).form('get value','textBoxFieldSelection')
		console.log("saveNewTextBox: Selected field ID: " + fieldID)
			
		newTextBoxParams.containerParams["fieldID"] = fieldID

		// TODO - saveNewTextBox() depends on containerParams - need to pass in somehow
		jsonAPIRequest("newLayoutContainer",newTextBoxParams.containerParams,function(replyData) {
	          console.log("Done getting new ID:response=" + JSON.stringify(replyData));
			  // TODO - Define some kind of common "validateJSONResponse" function
			  // and possibly write errors back to a server log.

			  if(replyData.hasOwnProperty("layoutContainerID") && 
				  replyData.hasOwnProperty("placeholderID")) {
					  // Replace the placeholder ID with the permanent one generated via
					  // the API call.
				  
					  console.log("fields by ID: " + JSON.stringify(newTextBoxParams.fieldsByID))
				  
					  $('#'+newTextBoxParams.placeholderID).find('label').text(newTextBoxParams.fieldsByID[fieldID].name)
					  $('#'+newTextBoxParams.placeholderID).attr("id",replyData.layoutContainerID)
 				  
					  newTextBoxParams.containerCreated = true
					  newTextBoxParams.dialogBox.dialog("close")
				  }
				  else {
		              console.log("ERROR: Missing properties in newLayoutContainer response:response=" + JSON.stringify(replyData));
					  newTextBoxParams.dialogBox.dialog("close")
				  }
	       }) // newLayoutContainer API request
		
	} // An existing field was selected
} // save new text box


var newTextBoxValidateFormatEntriesPanel = {
	divID: "newTextBoxValidateFormatEntriesPanel",
	progressPerc:90,
	dlgButtons: { 
		"Previous": function() {
			if(doCreateNewFieldWithTextBox()) {
				if(newFieldIsCalcField()){
					transitionToPrevDlgPanel(this,newTextBoxValidateFormatEntriesPanel,calcFieldFormulaPanelConfig)	
				} else {
					// Not a calculated field, skip over panel to enter calculated field formula
					transitionToPrevDlgPanel(this,newTextBoxValidateFormatEntriesPanel,newFieldPanelConfig)				
				}
			}  else {
				// Not creating a new field - skip over new field panels
				transitionToPrevDlgPanel(this,newTextBoxValidateFormatEntriesPanel,newOrExistingFieldPanelConfig)				
			}
		 },
		"Done" : function() { 
			if($( "#newTextBoxDlgCalcFieldFormulaPanel" ).form('validate form')) {
				saveNewTextBox()
			} // if validate panel's form
		},		
		"Cancel" : function() { $(this).dialog('close'); },
 	}, // dialog buttons
	initPanel: function() {},	
}

var calcFieldFormulaPanelConfig = {
	divID: "newTextBoxDlgCalcFieldFormulaPanel",
	progressPerc:60,
	dlgButtons: { 
		"Previous": function() { 
			transitionToPrevDlgPanel(this,calcFieldFormulaPanelConfig,newFieldPanelConfig)
		 },	
		"Next" : function() { 
			if($( "#newTextBoxDlgCalcFieldFormulaPanel" ).form('validate form')) {
				transitionToNextDlgPanel(this,calcFieldFormulaPanelConfig,newTextBoxValidateFormatEntriesPanel)
			} // if validate panel's form
		},
		"Cancel" : function() { $(this).dialog('close'); },
 	}, // dialog buttons
	initPanel: function () {
		// Initialize the semantic ui dropdown menus
		$('.ui.dropdown').dropdown()
		initCalcFieldEditBehavior(newTextBoxParams.fieldsByID)
	},
} // wizard dialog configuration for panel to create new field


var newFieldPanelConfig = {
	divID: "newTextBoxDlgNewFieldPanel",
	progressPerc:40,
	dlgButtons: { 
		"Previous": function() { 
			transitionToPrevDlgPanel(this,newFieldPanelConfig,newOrExistingFieldPanelConfig)	
		 },
		"Next" : function() { 
			
			if($( "#newTextBoxDlgNewFieldPanel" ).form('validate form')) {
				if(newFieldIsCalcField()){
					transitionToNextDlgPanel(this,newFieldPanelConfig,calcFieldFormulaPanelConfig)
				}
				else {
					transitionToNextDlgPanel(this,newFieldPanelConfig,newTextBoxValidateFormatEntriesPanel)
				}	
			} // if validate panel's form	
		},
		"Cancel" : function() { $(this).dialog('close'); },
 	}, // dialog buttons
	initPanel: function () {
		$( "#newTextBoxDlgNewFieldPanel" ).form({
	    	fields: {
		        newFieldName: {
		          rules: [
		            {
		              type   : 'empty',
		              prompt : 'Please enter a field name'
		            }
		          ]
		        }, // newFieldName validation
		        newFieldValTypeSelection: {
		          rules: [
		            {
		              type   : 'empty',
		              prompt : 'Please select a type'
		            }
		          ]
		        }, // newFieldValTypeSelection validation
				newFieldRefName: {
		          rules: [
		            {
		              type   : 'empty',
		              prompt : 'Please enter a reference name'
		            }
		          ]
		        }, // newFieldRefName validation
	     	},
	  	})
		$('#newRefNameHelp').popup({on: 'hover'});
		
	}
} // wizard dialog configuration for panel to create new field


var newOrExistingFieldPanelConfig = {
	divID: "newTextBoxDlgSelectOrNewFieldPanel",
	progressPerc:0,
	dlgButtons: { 
		"Next": function() {
			if($( "#newTextBoxDlgSelectOrNewFieldPanel" ).form('validate form')) {			
				console.log("New Field checked: " + doCreateNewFieldWithTextBox())
				if(doCreateNewFieldWithTextBox()) {
					transitionToNextDlgPanel(this,newOrExistingFieldPanelConfig,newFieldPanelConfig)
				}
				else {
					transitionToNextDlgPanel(this,newOrExistingFieldPanelConfig,newTextBoxValidateFormatEntriesPanel)						
				}
			} // if validate form
		 },		
		"Cancel" : function() { $(this).dialog('close'); },
 	}, // dialog buttons
	
	initPanel: function () {
		
		function enableSelectExistingField() {
			$('#selectExistingFieldField').removeClass("disabled")
			$( "#newTextBoxDlgSelectOrNewFieldPanel" ).form({
		    	fields: {
			        textBoxFieldSelection: {
			          rules: [
			            {
			              type   : 'empty',
			              prompt : 'Please select a field'
			            }
			          ]
			        }, // textBoxFieldSelection validation
		     	},
				inline : true,
		  	})
		}

		function disableSelectExistingField() {
			$('#selectExistingFieldField').addClass("disabled")
			$( "#newTextBoxDlgSelectOrNewFieldPanel" ).form({
		    	fields: {},
				inline : true,
		  	})
			// After changing the validation rules, re-validate the form.
			// This will remove any outstanding errors, which no longer apply
			// since selection of an existing field is no longer required.
			$( "#newTextBoxDlgSelectOrNewFieldPanel" ).form('validate form')
		}
		
		
		disableSelectExistingField();
		$( "#createNewFieldRadio" ).prop( "checked", true );
		$("input[name='newOrExistingRadio']").change(function(){
	        console.log("new or existing radio value:",this.value);
			if(this.value == "new") {
				disableSelectExistingField()
			} else {
				enableSelectExistingField()
			}
	    });
		
	}
} // wizard dialog configuration for panel to create new field


function newLayoutContainer(containerParams,fieldsByID)
{
	newTextBoxParams = {
		containerParams: containerParams,
		containerCreated: false,
		placeholderID: containerParams.containerID,
		fieldsByID: fieldsByID,
		dialogBox: $( "#newTextBox" )
	}
	
	// Enable Semantic UI checkboxes and popups
	$('.ui.checkbox').checkbox();
	$('.ui.radio.checkbox').checkbox();	
			
    newTextBoxParams.dialogBox.dialog({
      autoOpen: false,
      height: 500, width: 550,
	  resizable: false,
      modal: true,
      buttons: newOrExistingFieldPanelConfig.dlgButtons,
      close: function() {
		  console.log("Close dialog")
		  if(!newTextBoxParams.containerCreated)
		  {
			  // If the the text box creation is not complete, remove the placeholder
			  // from the canvas.
			  $('#'+newTextBoxParams.placeholderID).remove()
		  }
      }
    });
 
    newTextBoxParams.dialogBox.find( "form" ).on( "submit", function( event ) {
      	event.preventDefault();
		//saveNewTextBox()
		// TODO - reimplement save with enter key
		console.log("Save not implemented with enter key")
    });
	
	$( ".wizardPanel" ).hide() // hide all the panels
	$( "#newTextBoxDlgSelectOrNewFieldPanel").show() // show the first panel
	newTextBoxParams.dialogBox.dialog("option","buttons",newOrExistingFieldPanelConfig.dlgButtons)
	
	// Clear any previous entries validation errors. The message blocks by 
	// default don't clear their values with 'clear', so any remaining error
	// messages need to be removed from the message blocks within the panels.
	$('.wizardPanel').form('clear') // clear any previous entries
	$('.wizardErrorMsgBlock').empty()
	
	
	$('#newTextBoxProgress').progress({percent:0});

	newOrExistingFieldPanelConfig.initPanel()
	newFieldPanelConfig.initPanel()
	calcFieldFormulaPanelConfig.initPanel()
	newTextBoxValidateFormatEntriesPanel.initPanel()


	newTextBoxParams.dialogBox.dialog("open")
	
} // newLayoutContainer

function initNewTextBoxDialog() {
	// Initialize the newTextBox dialog with the minimum parameters. This is necessary
	// to hide the dialog from view when the document is initially loaded. The
	// dialog is fully re-initialized just prior to it being opened.
    $( "#newTextBox" ).dialog({ autoOpen: false })

	
}


