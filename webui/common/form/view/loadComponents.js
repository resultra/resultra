
// By default, changes made within forms are immediately committed to the "main line". 
// However, if a specific change set ID is used, the changes will be "branched" from the main line.
// This branching is needed for modal forms which may cancel changes made before the "save changes" button
// of the dialog is pressed to fully commit the changes.
var MainLineFullyCommittedChangeSetID = ""

function loadFormViewComponents(canvasSelector, viewFormContext, changeSetID,
	 getCurrentRecordFunc, updateCurrentRecordFunc,
	 doneLoadingComponentsFunc) {
	
	function initFormComponentViewBehavior($component,componentID, selectionFunc) {	
		initObjectSelectionBehavior($component, 
				canvasSelector,function(selectedComponentID) {
			console.log("Form view object selected: " + selectedComponentID)
			var selectedObjRef	= getContainerObjectRef($component)
			selectionFunc(selectedObjRef)
		})
	}
	
	loadFormComponents({
		formParentElemID: canvasSelector,
		formContext: viewFormContext,
		initTextBoxFunc: function(componentContext,$textBox,textBoxObjectRef) {			
			initTextBoxRecordEditBehavior(componentContext, changeSetID,
				getCurrentRecordFunc,updateCurrentRecordFunc,textBoxObjectRef)
			initFormComponentViewBehavior($textBox,
					textBoxObjectRef.textBoxID,initTextBoxViewProperties)
		},
		initSelectionFunc: function(componentContext,$selection,selectionObjectRef) {			
			initSelectionRecordEditBehavior($selection,componentContext,changeSetID,
					getCurrentRecordFunc,updateCurrentRecordFunc,selectionObjectRef)
			initFormComponentViewBehavior($selection,
					selectionObjectRef.selectionID,initSelectionViewProperties)
		},
		initCommentFunc: function(componentContext,$comment,commentObjectRef) {			
			initCommentBoxRecordEditBehavior($comment,componentContext,changeSetID,
					getCurrentRecordFunc,updateCurrentRecordFunc,commentObjectRef)
			initFormComponentViewBehavior($comment,
					commentObjectRef.commentID,initCommentViewProperties)
		},
		initCheckBoxFunc: function(componentContext,$checkBox,checkBoxObjectRef) {
			console.log("Init check box in view form")
			initCheckBoxRecordEditBehavior($checkBox,componentContext,changeSetID,
					getCurrentRecordFunc,updateCurrentRecordFunc,checkBoxObjectRef)
			initFormComponentViewBehavior($checkBox,
					checkBoxObjectRef.checkBoxID,initCheckBoxViewProperties)
		},
		initRatingFunc: function(componentContext,$rating,ratingObjectRef) {
			console.log("Init rating in view form")
			initRatingRecordEditBehavior($rating,componentContext,changeSetID,
					getCurrentRecordFunc,updateCurrentRecordFunc,ratingObjectRef)
			initFormComponentViewBehavior($rating,
					ratingObjectRef.ratingID,initRatingViewProperties)
		},
		initFormButtonFunc: function(componentContext,$button,buttonObjectRef) {
			console.log("Init form button in view form")
			
			// The loadFormViewComponents and loadRecordIntoFormLayout functions
			// need to be passed to initFormButtonRecordEditBehavior in order
			// to avoid a cyclical package dependency.
			initFormButtonRecordEditBehavior($button,componentContext,
					getCurrentRecordFunc,updateCurrentRecordFunc,buttonObjectRef,
					loadFormViewComponents,loadRecordIntoFormLayout)
		},
		initUserSelectionFunc: function(componentContext,$userSelection,userSelectionObjectRef) {
			console.log("Init user selection in view form")
			initUserSelectionRecordEditBehavior(componentContext,changeSetID,
					getCurrentRecordFunc,updateCurrentRecordFunc,userSelectionObjectRef)
			initFormComponentViewBehavior($userSelection,
					userSelectionObjectRef.userSelectionID,initUserSelectionViewProperties)
		},
		initDatePickerFunc: function(componentContext,$datePicker,datePickerObjectRef) {
			console.log("Init date picker in view form")
			initDatePickerRecordEditBehavior($datePicker,componentContext,changeSetID,
					getCurrentRecordFunc,updateCurrentRecordFunc,datePickerObjectRef)
			initFormComponentViewBehavior($datePicker,
					datePickerObjectRef.datePickerID,initDatePickerViewProperties)
		},
		initHtmlEditorFunc: function(componentContext,$htmlEditor,htmlEditorObjectRef) {
			console.log("Init html editor in view form")
			initHtmlEditorRecordEditBehavior($htmlEditor,componentContext,changeSetID,
					getCurrentRecordFunc,updateCurrentRecordFunc,htmlEditorObjectRef)
			initFormComponentViewBehavior($htmlEditor,
					htmlEditorObjectRef.htmlEditorID,initHTMLEditorViewProperties)
		},
		initImageFunc: function(componentContext,$image,imageObjectRef) {
			console.log("Init image in view form")
			initImageRecordEditBehavior($image,componentContext,changeSetID,
				getCurrentRecordFunc,updateCurrentRecordFunc,imageObjectRef)
			initFormComponentViewBehavior($image,
					imageObjectRef.imageID,initImageViewProperties)
		},
		initHeaderFunc: function($header,componentContext,headerObjectRef) {
			console.log("Init header in view form")
			initHeaderRecordEditBehavior($header, componentContext,headerObjectRef)
		},
		
		doneLoadingFormDataFunc: doneLoadingComponentsFunc
	});
}