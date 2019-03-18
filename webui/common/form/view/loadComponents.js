// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// By default, changes made within forms are immediately committed to the "main line". 
// However, if a specific change set ID is used, the changes will be "branched" from the main line.
// This branching is needed for modal forms which may cancel changes made before the "save changes" button
// of the dialog is pressed to fully commit the changes.
var MainLineFullyCommittedChangeSetID = ""

function loadFormViewComponentsIntoOneLayout($parentFormLayout, viewFormContext, recordProxy,componentContext, formPopulationDoneCallback) {
	
	function initFormComponentViewBehavior($component,componentID, selectionFunc) {	
	}
	
	var loadFormConfig = {
		$parentFormLayout: $parentFormLayout,
		formContext: viewFormContext,
		initTextBoxFunc: function(componentContext,$textBox,textBoxObjectRef) {			
			initTextBoxRecordEditFormBehavior($textBox,componentContext,recordProxy,textBoxObjectRef)
			initFormComponentViewBehavior($textBox,
					textBoxObjectRef.textBoxID,initTextBoxViewProperties)
		},
		initEmailAddrFunc: function(componentContext,$emailAddr,emailAddrObjectRef) {			
			initEmailAddrFormRecordEditBehavior($emailAddr,componentContext,recordProxy,emailAddrObjectRef)
			initFormComponentViewBehavior($emailAddr,
					emailAddrObjectRef.emailAddrID,initEmailAddrViewProperties)
		},
		initFileFunc: function(componentContext,$file,fileObjectRef) {			
			initFileFormRecordEditBehavior($file,componentContext,recordProxy,fileObjectRef)
			initFormComponentViewBehavior($file,
					fileObjectRef.fileID,initFileViewProperties)
		},
		initImageFunc: function(componentContext,$image,imageObjectRef) {			
			initImageFormRecordEditBehavior($image,componentContext,recordProxy,imageObjectRef)
			initFormComponentViewBehavior($image,
					imageObjectRef.imageID,initImageViewProperties)
		},
		initUrlLinkFunc: function(componentContext,$urlLink,urlLinkObjectRef) {			
			initUrlLinkFormRecordEditBehavior($urlLink,componentContext,recordProxy,urlLinkObjectRef)
			initFormComponentViewBehavior($urlLink,
					urlLinkObjectRef.urlLinkID,initUrlLinkViewProperties)
		},
		initNumberInputFunc: function(componentContext,$numberInput,numberInputObjectRef) {			
			initNumberInputFormRecordEditBehavior($numberInput,componentContext,recordProxy,numberInputObjectRef)
			initFormComponentViewBehavior($numberInput,
					numberInputObjectRef.numberInputID,initNumberInputViewProperties)
		},
		initSelectionFunc: function(componentContext,$selection,selectionObjectRef,initDoneCallback) {			
			initSelectionRecordEditBehavior($selection,componentContext,recordProxy,selectionObjectRef,function() {
				initDoneCallback()
			})
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
		initSocialButtonFunc: function(componentContext,$socialButton,socialButtonObjectRef) {
			console.log("Init rating in view form")
			initSocialButtonFormRecordEditBehavior($socialButton,componentContext,recordProxy,socialButtonObjectRef)
			initFormComponentViewBehavior($socialButton,
					socialButtonObjectRef.socialButtonID,initSocialButtonViewProperties)
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
		initUserTagFunc: function(componentContext,$userTag,userTagObjectRef) {
			console.log("Init user selection in view form")
			initUserTagFormRecordEditBehavior($userTag,componentContext,recordProxy,userTagObjectRef)
			initFormComponentViewBehavior($userTag,
					userTagObjectRef.userTagID,initUserTagViewProperties)
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
		initAttachmentFunc: function(componentContext,$image,imageObjectRef) {
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
		},	
		initLabelFunc: function(componentContext,$label,labelObjectRef) {
			initLabelFormRecordEditBehavior($label,componentContext,recordProxy,labelObjectRef)
			initFormComponentViewBehavior($label,
					labelObjectRef.labelID,initLabelViewProperties)
		}		
	}
	
	populateOneFormLayoutWithComponents(loadFormConfig,componentContext,function() {
		formPopulationDoneCallback()
	});
	
}

function loadFormViewComponents($parentFormLayout, viewFormContext, recordProxy,doneLoadingComponentsFunc) {
	
	getFormComponentContext(viewFormContext, function(componentContext) {
												
		loadFormViewComponentsIntoOneLayout($parentFormLayout, viewFormContext, recordProxy,componentContext,function() {
			doneLoadingComponentsFunc()			
		})	
	})
	
		
}


function loadMultipleFormViewContainers(viewFormContext,containersInfo, doneLoadingContainersFunc) {
	
	var formComponentsRemaining = containersInfo.length
	function oneFormLayoutComplete() {
		formComponentsRemaining--
		if(formComponentsRemaining <= 0) {
			doneLoadingContainersFunc()
		}
	}
	
	getFormComponentContext(viewFormContext, function(componentContext) {
		
		for (var containerIndex=0; containerIndex < containersInfo.length; containerIndex++) {
			
			var currContainerInfo = containersInfo[containerIndex]
			
			loadFormViewComponentsIntoOneLayout(currContainerInfo.$listItemContainer, viewFormContext, 
					currContainerInfo.recordProxy,componentContext,function() {
						oneFormLayoutComplete()
					})	
			
		}
	})
}