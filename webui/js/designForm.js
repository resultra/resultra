

function initLayoutEditFieldInfo(fieldRef)
{
	// Add the ability to select the field from the new layout container
	// dialog. This callback function is called for both number and 
	// text fields; so, text boxes can be used for both numbers and 
	// regular text.
 	var selectFieldOptionHTML = '<option value="' +
 		fieldRef.fieldID + '">' +
 		fieldRef.fieldInfo.name + '</option>'
 	$("#textBoxFieldSelection").append(selectFieldOptionHTML)	
}

function initCanvasComplete() {	
	
} // noop



function initContainerEditBehavior(container)
{

	// While in edit mode, disable input on the container
	container.find('input').prop('disabled',true);


	// TODO - This could be put into a common function, since these
	// properties should be the same for objects loaded with the page
	// and newly added objects.
	container.draggable ({
		grid: [20, 20], // snap to a grid
		cursor: "move",
		containment: "parent",
		clone: "original",				
		stop: function(event, ui) {
				  var layoutPos = {
					  positionLeft: ui.position.left,
					  positionTop: ui.position.top,
				   };
		  
				  console.log("drag stop: id: " + event.target.id);
				  console.log("drag stop: new position: " + JSON.stringify(layoutPos));
				  
				  // TODO: send ajax request to reposition the container
  
		      } // stop function
	})
	container.resizable({
		aspectRatio: false,
		handles: 'e, w', // Only allow resizing horizontally
		maxWidth: 600,
		minWidth: 100,
		grid: 20, // snap to grid during resize

		stop: function(event, ui) {  
				  
	  			var layoutContainerParams = {
	  				parentLayoutID: layoutID,containerID: event.target.id,
	  				geometry: { positionTop: ui.position.top,positionLeft: ui.position.left,
						sizeWidth: ui.size.width, sizeHeight: ui.size.height }
	  			};
		
			 	jsonAPIRequest("resizeLayoutContainer",layoutContainerParams,function(replyData) {})	
				  
			  } // stop function
	});
	
}


function newLayoutContainer(containerParams)
{
	var placeholderID = containerParams.containerID
	var containerCreated = false
	
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
	
	// Enable Semantic UI checkboxes and popups
	$('#newRefNameHelp').popup({on: 'hover'});
	$('.ui.checkbox').checkbox();
	$('.ui.radio.checkbox').checkbox();
	
	function saveNewTextBox()
	{
		// Complete the parameters to create the new layout container by including the 
		// fieldID selected in the dialog box.
		var fieldID = $( "#newTextBoxDlgSelectOrNewFieldPanel" ).form('get value','textBoxFieldSelection')
		console.log("saveNewTextBox: Selected field ID: " + fieldID)
				
		containerParams["fieldID"] = fieldID

		jsonAPIRequest("newLayoutContainer",containerParams,function(replyData) {
	          console.log("Done getting new ID:response=" + JSON.stringify(replyData));
			  // TODO - Define some kind of common "validateJSONResponse" function
			  // and possibly write errors back to a server log.
  
			  if(replyData.hasOwnProperty("layoutContainerID") && 
				  replyData.hasOwnProperty("placeholderID")) {
					  // Replace the placeholder ID with the permanent one generated via
					  // the API call.
					  $('#'+placeholderID).find('label').text(fieldsByID[fieldID].name)
					  $('#'+placeholderID).attr("id",replyData.layoutContainerID)
	 				  
					  containerCreated = true
					  dialog.dialog("close")
				  }
				  else {
		              console.log("ERROR: Missing properties in newLayoutContainer response:response=" + JSON.stringify(replyData));
					  dialog.dialog("close")
				  }
	       }) // newLayoutContainer API request
	} // save new text box
	
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
	} // wizard dialog configuration for panel to create new field
			
    dialog = $( "#newTextBox" ).dialog({
      autoOpen: false,
      height: 500, width: 550,
      modal: true,
      buttons: newOrExistingFieldPanelConfig.dlgButtons,
      close: function() {
		  console.log("Close dialog")
		  if(!containerCreated)
		  {
			  // If the the text box creation is not complete, remove the placeholder
			  // from the canvas.
			  $('#'+placeholderID).remove()
		  }
      }
    });
 
    form = dialog.find( "form" ).on( "submit", function( event ) {
      	event.preventDefault();
		if($( "#newTextBox" ).form('validate form')) {
			saveNewTextBox()
		}
    });
	
	$( ".wizardPanel" ).hide() // hide all the panels
	$( "#newTextBoxDlgSelectOrNewFieldPanel").show() // show the first panel
	$(dialog).dialog("option","buttons",newOrExistingFieldPanelConfig.dlgButtons)
	
	// Clear any previous entries validation errors. The message blocks by 
	// default don't clear their values with 'clear', so any remaining error
	// messages need to be removed from the message blocks within the panels.
	$('.wizardPanel').form('clear') // clear any previous entries
	$('.wizardErrorMsgBlock').empty()
	
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
	
	
	$('#newTextBoxProgress').progress({percent:0});

	$( "#newTextBox" ).dialog("open")
	
} // newLayoutContainer

function droppedObjGeometry(dropDest,droppedObj,ui)
{
	// TODO - Support snapping of the top and left
	var relTop = ui.offset.top-$(dropDest).offset().top
	var relLeft = ui.offset.left-$(dropDest).offset().left
	var objWidth = droppedObj.width()
	var objHeight = droppedObj.height()
		
	return { top: relTop, left: relLeft, width: objWidth, height: objHeight }
}


$(document).ready(function() {
				
	
	$(".newField").draggable({
		
		revert: 'invalid',
		cursor: "move",
		helper: function() { 
			 // Instead of dragging the icon in the gallery, drag a "placeholder image" of the
			 // field itself. This allows the user to see the proporitions of the field relative
			 // to the other elements already on the canvas.
			var newFieldDraggablePlaceholderHTML = fieldContainerHTML(allocNextPlaceholderID());
			var placeholder = $(newFieldDraggablePlaceholderHTML);
						
			 // While in layout mode, disable entry into the fields
			 // Interestingly, there's not CSS for this.
			placeholder.find('input').prop('disabled',true);
			 			 
			 return placeholder
		 },
		// The following are needed to keep the draggable above the other elements. Getting Semantic UI,
		// jQuery UI Layout and jQuery UI drag-and-drop working together is really tricky. Some issues:
		//
		// 1. jQuery UI layout has difficulty with the z-index for items which are opened in one pane,
		//    but overlap another. The recommended solution is to attach items which need to overlap all
		//    the panes to <body>. If this isn't possible, the west__showOverflowOnHover option is provided
		//    for initializing the panes.
		// 2. The Semantic UI form elements inherit some of their styling from the <form> styling. So,
		//    if a form element is not embeded in a <form> tag, the styling isn't correct. To work around this,
		//    the entire page is embedded in a <form> tag. However, it wasn't possible to attach the item to
		//    <body> and still have it inherit Semantic UI's <form> styling.
		//
		// So, the following solution is the only solution which could be found to address both issues.
		// described above.
		stack: "div",
		appendTo: "#gallerySidebar",
		zIndex: 999
		
	});
	
    $("#layoutCanvas").droppable({
        accept: ".newField",
        activeClass: "ui-state-highlight",
        drop: function( event, ui ) {
 			var theClone = $(ui.helper).clone()
			
			// After the drag operation is complete, the draggable and resizable
			// properties need to be initialized.
			initContainerEditBehavior(theClone)
					
			$(this).append(theClone);

			// IMPORTANT - Don't call droppedObjGeometry until after theClone has been appended
			// Otherwise the width and height will be returned as 0.
			var placeholderID = $(theClone).attr("id");
			var droppedObjGeom = droppedObjGeometry(this,theClone,ui)
			
			theClone.css({top: droppedObjGeom.top, left: droppedObjGeom.left, position:"absolute"});
			
		    console.log("End Drag and drop: placeholder ID=" + placeholderID + 
				" drop loc =" + JSON.stringify(droppedObjGeom));
			
			var layoutContainerParams = {
				parentLayoutID: layoutID, containerID: placeholderID,
				geometry: {positionTop: droppedObjGeom.top, positionLeft: droppedObjGeom.left,
				sizeWidth: droppedObjGeom.width,sizeHeight: droppedObjGeom.height}
				};
			
			newLayoutContainer(layoutContainerParams)
						
        }
    }); // #layoutCanvas droppable
	
	
	// Initialize the newTextBox dialog with the minimum parameters. This is necessary
	// to hide the dialog from view when the document is initially loaded. The
	// dialog is fully re-initialized just prior to it being opened.
    $( "#newTextBox" ).dialog({ autoOpen: false })
 
		
	// Initialize the page layout
	$('#layoutPage').layout({
		north: fixedUILayoutPaneParams(40),
		east: fixedUILayoutPaneParams(300),
		west: fixedUILayoutPaneParams(200),
		west__showOverflowOnHover:	true
	})	  
	  
	
	initCanvas(initContainerEditBehavior,initLayoutEditFieldInfo,initCanvasComplete)
	  
	  


});
