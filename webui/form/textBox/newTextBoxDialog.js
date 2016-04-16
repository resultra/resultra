

// Initialized below in newLayoutContainer
var newTextBoxParams;

var textBoxProgressDivID = '#newTextBoxProgress'


function newFieldIsCalcField() {
	return $(newFieldPanelConfig.divID).form('get field','isCalcField').prop('checked')
}

function doCreateNewFieldWithTextBox() {
	
	return $(newOrExistingFieldPanelConfig.divID).form('get field','createNewFieldRadio').prop('checked')
}



function saveNewTextBox()
{
	// Complete the parameters to create the new textBox by including the 
	// fieldID selected in the dialog box.
	
	if(doCreateNewFieldWithTextBox()) {
		console.log("saveNewTextBox() with create new field - not implemented yet, closing dialog")
		newTextBoxParams.dialogBox.dialog("close")
	} else {
		var fieldID = $( "#newTextBoxDlgSelectOrNewFieldPanel" ).form('get value','textBoxFieldSelection')
		console.log("saveNewTextBox: Selected field ID: " + fieldID)
					
		var newTextBoxAPIParams = {
			parentID: newTextBoxParams.containerParams.parentLayoutID,
			geometry: newTextBoxParams.containerParams.geometry,
			fieldID: fieldID
		}

		// TODO - saveNewTextBox() depends on containerParams - need to pass in somehow
		jsonAPIRequest("frm/textBox/new",newTextBoxAPIParams,function(newTextBoxObjectRef) {
	          console.log("Done getting new ID:response=" + JSON.stringify(newTextBoxObjectRef));
			  
			  $('#'+newTextBoxParams.placeholderID).find('label').text(newTextBoxObjectRef.fieldRef.fieldInfo.name)
			  $('#'+newTextBoxParams.placeholderID).attr("id",newTextBoxObjectRef.uniqueID.objectID)
			  
			  // Store the newly created object reference in the DOM element. This is needed for follow-on
			  // property setting, resizing, etc.
			  setElemObjectRef(newTextBoxObjectRef.uniqueID.objectID,newTextBoxObjectRef)
			
			  newTextBoxParams.containerCreated = true				  
					  
			  newTextBoxParams.dialogBox.dialog("close")

	       }) // newLayoutContainer API request
		
	} // An existing field was selected
} // save new text box


var newTextBoxValidateFormatEntriesPanel = {
	divID: "#newTextBoxValidateFormatEntriesPanel",
	progressPerc:90,
	dlgButtons: { 
		"Previous": function() {
			if(doCreateNewFieldWithTextBox()) {
				if(newFieldIsCalcField()){
					transitionToPrevWizardDlgPanel(this,textBoxProgressDivID,
						newTextBoxValidateFormatEntriesPanel,calcFieldFormulaPanelConfig)	
				} else {
					// Not a calculated field, skip over panel to enter calculated field formula
					transitionToPrevWizardDlgPanel(this,textBoxProgressDivID,
						newTextBoxValidateFormatEntriesPanel,newFieldPanelConfig)				
				}
			}  else {
				// Not creating a new field - skip over new field panels
				transitionToPrevWizardDlgPanel(this,textBoxProgressDivID,
						newTextBoxValidateFormatEntriesPanel,newOrExistingFieldPanelConfig)				
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
	divID: "#newTextBoxDlgCalcFieldFormulaPanel",
	progressPerc:60,
	dlgButtons: { 
		"Previous": function() { 
			transitionToPrevWizardDlgPanel(this,textBoxProgressDivID,
						calcFieldFormulaPanelConfig,newFieldPanelConfig)
		 },	
		"Next" : function() { 
			if($( "#newTextBoxDlgCalcFieldFormulaPanel" ).form('validate form')) {
				transitionToNextWizardDlgPanel(this,textBoxProgressDivID,
						calcFieldFormulaPanelConfig,newTextBoxValidateFormatEntriesPanel)
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
	divID: "#newTextBoxDlgNewFieldPanel",
	progressPerc:40,
	dlgButtons: { 
		"Previous": function() { 
			transitionToPrevWizardDlgPanel(this,textBoxProgressDivID,
						newFieldPanelConfig,newOrExistingFieldPanelConfig)	
		 },
		"Next" : function() { 
			
			if($( "#newTextBoxDlgNewFieldPanel" ).form('validate form')) {
				if(newFieldIsCalcField()){
					transitionToNextWizardDlgPanel(this,textBoxProgressDivID,
						newFieldPanelConfig,calcFieldFormulaPanelConfig)
				}
				else {
					transitionToNextWizardDlgPanel(this,textBoxProgressDivID,
						newFieldPanelConfig,newTextBoxValidateFormatEntriesPanel)
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
	divID: "#newTextBoxDlgSelectOrNewFieldPanel",
	progressPerc:0,
	dlgButtons: { 
		"Next": function() {
			if($( "#newTextBoxDlgSelectOrNewFieldPanel" ).form('validate form')) {			
				console.log("New Field checked: " + doCreateNewFieldWithTextBox())
				if(doCreateNewFieldWithTextBox()) {
					transitionToNextWizardDlgPanel(this,textBoxProgressDivID,
						newOrExistingFieldPanelConfig,newFieldPanelConfig)
				}
				else {
					transitionToNextWizardDlgPanel(this,textBoxProgressDivID,
						newOrExistingFieldPanelConfig,newTextBoxValidateFormatEntriesPanel)						
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


function newLayoutContainer(containerParams)
{
	// Many of the dialog panels depends on field information, so this needs to be loaded before opening the actual dialog
	loadFieldInfo(function(fieldsByID) { 
		
		newTextBoxParams = {
			containerParams: containerParams,
			containerCreated: false,
			parentID: containerParams.parentID,
			placeholderID: containerParams.containerID,
			fieldsByID: fieldsByID,
			dialogBox: $( "#newTextBox" )
		}
	
		// Enable Semantic UI checkboxes and popups
		$('.ui.checkbox').checkbox();
		$('.ui.radio.checkbox').checkbox();	
	
		populateFieldSelectionMenu(fieldsByID, "#textBoxFieldSelection")
	
		openWizardDialog({
			closeFunc: function() {
				console.log("Close dialog")
				if(!newTextBoxParams.containerCreated)
				{
				  // If the the text box creation is not complete, remove the placeholder
				  // from the canvas.
					$('#'+newTextBoxParams.placeholderID).remove()
				}
	      	},
			width: 500, height: 500,
			dialogDivID: '#newTextBox',
			panels: [newOrExistingFieldPanelConfig, newFieldPanelConfig,
					calcFieldFormulaPanelConfig, newTextBoxValidateFormatEntriesPanel],
			progressDivID: '#newTextBoxProgress',
		})
		
	}) // loadFieldInfo
		
} // newLayoutContainer

function initNewTextBoxDialog() {
	// Initialize the newTextBox dialog with the minimum parameters. This is necessary
	// to hide the dialog from view when the document is initially loaded. The
	// dialog is fully re-initialized just prior to it being opened.
	initWizardDialog('#newTextBox')
}


