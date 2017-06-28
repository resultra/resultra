
// By default, changes made within forms are immediately committed to the "main line". 
// However, if a specific change set ID is used, the changes will be "branched" from the main line.
// This branching is needed for modal forms which may cancel changes made before the "save changes" button
// of the dialog is pressed to fully commit the changes.
var MainLineFullyCommittedChangeSetID = ""

function loadFormViewComponentsIntoOneLayout($parentFormLayout, viewFormContext, recordProxy,componentContext) {
	
	function initFormComponentViewBehavior($component,componentID, selectionFunc) {	
		
		function initComponentFooterButtonBarHoverBehavior($component) {
			// Initialize hover behavior on the container, whereby hovering in will display
			// the div with class componentHoaverFooter, and hovering out will hide this
			// same div. There's a delay to give the user time to click on any buttons in the footer.
			var hoverButtonBarTimer
			var hoverDelay = 500
			var $hoverFooter = $component.find(".componentHoverFooter")			
			$component.hover(function() {
				console.log("Start hover in component footer")
				clearTimeout(hoverButtonBarTimer)
				$hoverFooter.delay(hoverDelay).show()
			}, function() {
				console.log("End hover in component footer")
				hoverButtonBarTimer = setTimeout(function(){
	      		  	$hoverFooter.delay(hoverDelay).hide();
	    		},hoverDelay);
			});
			
		}
		initComponentFooterButtonBarHoverBehavior($component)
	}
	
	var loadFormConfig = {
		$parentFormLayout: $parentFormLayout,
		formContext: viewFormContext,
		initTextBoxFunc: function(componentContext,$textBox,textBoxObjectRef) {			
			initTextBoxRecordEditBehavior($textBox,componentContext,recordProxy,textBoxObjectRef)
			initFormComponentViewBehavior($textBox,
					textBoxObjectRef.textBoxID,initTextBoxViewProperties)
		},
		initNumberInputFunc: function(componentContext,$numberInput,numberInputObjectRef) {			
			initNumberInputFormRecordEditBehavior($numberInput,componentContext,recordProxy,numberInputObjectRef)
			initFormComponentViewBehavior($numberInput,
					numberInputObjectRef.numberInputID,initNumberInputViewProperties)
		},
		initSelectionFunc: function(componentContext,$selection,selectionObjectRef) {			
			initSelectionRecordEditBehavior($selection,componentContext,recordProxy,selectionObjectRef)
			initFormComponentViewBehavior($selection,
					selectionObjectRef.selectionID,initSelectionViewProperties)
		},
		initCommentFunc: function(componentContext,$comment,commentObjectRef) {			
			initCommentBoxFormRecordEditBehavior($comment,componentContext,recordProxy,commentObjectRef)
			initFormComponentViewBehavior($comment,
					commentObjectRef.commentID,initCommentViewProperties)
		},
		initCheckBoxFunc: function(componentContext,$checkBox,checkBoxObjectRef) {
			console.log("Init check box in view form")
			initFormCheckboxEditBehavior($checkBox,componentContext,recordProxy,checkBoxObjectRef)
			initFormComponentViewBehavior($checkBox,
					checkBoxObjectRef.checkBoxID,initCheckBoxViewProperties)
		},
		initToggleFunc: function(componentContext,$toggle,toggleObjectRef) {
			console.log("Init toggle in view form")
			initToggleFormRecordEditBehavior($toggle,componentContext,recordProxy,toggleObjectRef)
			initFormComponentViewBehavior($toggle,
					toggleObjectRef.toggleID,initToggleViewProperties)
		},
		initProgressFunc: function(componentContext,$progress,progressObjectRef) {
			console.log("Init progress indicator in view form")
			initProgressRecordEditBehavior($progress,componentContext,recordProxy,progressObjectRef)
			initFormComponentViewBehavior($progress,
					progressObjectRef.progressID,initProgressViewProperties)
		},
		initGaugeFunc: function(componentContext,$gauge,gaugeObjectRef) {
			console.log("Init progress indicator in view form")
			initGaugeRecordEditBehavior($gauge,componentContext,recordProxy,gaugeObjectRef)
			initFormComponentViewBehavior($gauge,
					gaugeObjectRef.gaugeID,initGaugeViewProperties)
		},
		initRatingFunc: function(componentContext,$rating,ratingObjectRef) {
			console.log("Init rating in view form")
			initRatingFormRecordEditBehavior($rating,componentContext,recordProxy,ratingObjectRef)
			initFormComponentViewBehavior($rating,
					ratingObjectRef.ratingID,initRatingViewProperties)
		},
		initFormButtonFunc: function(componentContext,$button,buttonObjectRef) {
			console.log("Init form button in view form")
			
			// The loadFormViewComponents and loadRecordIntoFormLayout functions
			// need to be passed to initFormButtonRecordEditBehavior in order
			// to avoid a cyclical package dependency.
			var defaultValSrc = "frm=" + buttonObjectRef.buttonID
			initFormButtonRecordEditBehavior($button,componentContext,recordProxy,buttonObjectRef,defaultValSrc,
					loadFormViewComponents,loadRecordIntoFormLayout)
		},
		initUserSelectionFunc: function(componentContext,$userSelection,userSelectionObjectRef) {
			console.log("Init user selection in view form")
			initUserSelectionFormRecordEditBehavior($userSelection,componentContext,recordProxy,userSelectionObjectRef)
			initFormComponentViewBehavior($userSelection,
					userSelectionObjectRef.userSelectionID,initUserSelectionViewProperties)
		},
		initDatePickerFunc: function(componentContext,$datePicker,datePickerObjectRef) {
			console.log("Init date picker in view form")
			initFormDatePickerEditBehavior($datePicker,componentContext,recordProxy,datePickerObjectRef)
			initFormComponentViewBehavior($datePicker,
					datePickerObjectRef.datePickerID,initDatePickerViewProperties)
		},
		initHtmlEditorFunc: function(componentContext,$htmlEditor,htmlEditorObjectRef) {
			console.log("Init html editor in view form")
			initNoteEditorFormRecordEditBehavior($htmlEditor,componentContext,recordProxy,htmlEditorObjectRef)
			initFormComponentViewBehavior($htmlEditor,
					htmlEditorObjectRef.htmlEditorID,initHTMLEditorViewProperties)
		},
		initImageFunc: function(componentContext,$image,imageObjectRef) {
			console.log("Init image in view form")
			initAttachmentFormRecordEditBehavior($image,componentContext,recordProxy,imageObjectRef)
			initFormComponentViewBehavior($image,
					imageObjectRef.imageID,initImageViewProperties)
		},
		initHeaderFunc: function($header,componentContext,headerObjectRef) {
			console.log("Init header in view form")
			initHeaderRecordEditBehavior($header, componentContext,headerObjectRef)
		},
		initCaptionFunc: function($caption,componentContext,captionObjectRef) {
			console.log("Init caption in view form")
			initCaptionRecordEditBehavior($caption, componentContext,captionObjectRef)
		}		
	}
	
	populateOneFormLayoutWithComponents(loadFormConfig,componentContext);
	
}

function loadFormViewComponents($parentFormLayout, viewFormContext, recordProxy,doneLoadingComponentsFunc) {
	
	getFormComponentContext(viewFormContext, function(componentContext) {
												
		loadFormViewComponentsIntoOneLayout($parentFormLayout, viewFormContext, recordProxy,componentContext)	
			
		doneLoadingComponentsFunc()
	})
	
		
}


function loadMultipleFormViewContainers(viewFormContext,containersInfo, doneLoadingContainersFunc) {
	getFormComponentContext(viewFormContext, function(componentContext) {
		
		for (var containerIndex=0; containerIndex < containersInfo.length; containerIndex++) {
			
			var currContainerInfo = containersInfo[containerIndex]
			
			loadFormViewComponentsIntoOneLayout(currContainerInfo.$listItemContainer, viewFormContext, 
					currContainerInfo.recordProxy,componentContext)	
			
		}
		doneLoadingContainersFunc()	
	})
}