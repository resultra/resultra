

function loadFormViewComponents(canvasSelector, viewFormContext, 
	 getCurrentRecordFunc, updateCurrentRecordFunc,
	 doneLoadingComponentsFunc) {
	
	function initFormComponentViewBehavior($component,componentID, selectionFunc) {	
		initObjectSelectionBehavior($component, 
				canvasSelector,function(selectedComponentID) {
			console.log("Form view object selected: " + selectedComponentID)
			var selectedObjRef	= getElemObjectRef(selectedComponentID)
			selectionFunc(selectedObjRef)
		})
	}
	
	loadFormComponents({
		formParentElemID: canvasSelector,
		formContext: viewFormContext,
		initTextBoxFunc: function(componentContext,$textBox,textBoxObjectRef) {			
			initTextBoxRecordEditBehavior(componentContext,
				getCurrentRecordFunc,updateCurrentRecordFunc,textBoxObjectRef)
			initFormComponentViewBehavior($textBox,
					textBoxObjectRef.textBoxID,initTextBoxViewProperties)
		},
		initSelectionFunc: function(componentContext,$selection,selectionObjectRef) {			
			initSelectionRecordEditBehavior(componentContext,
					getCurrentRecordFunc,updateCurrentRecordFunc,selectionObjectRef)
			initFormComponentViewBehavior($selection,
					selectionObjectRef.selectionID,initSelectionViewProperties)
		},
		initCommentFunc: function(componentContext,$comment,commentObjectRef) {			
			initCommentBoxRecordEditBehavior(componentContext,
					getCurrentRecordFunc,updateCurrentRecordFunc,commentObjectRef)
			initFormComponentViewBehavior($comment,
					commentObjectRef.commentID,initCommentViewProperties)
		},
		initCheckBoxFunc: function(componentContext,$checkBox,checkBoxObjectRef) {
			console.log("Init check box in view form")
			initCheckBoxRecordEditBehavior(componentContext,
					getCurrentRecordFunc,updateCurrentRecordFunc,checkBoxObjectRef)
			initFormComponentViewBehavior($checkBox,
					checkBoxObjectRef.checkBoxID,initCheckBoxViewProperties)
		},
		initRatingFunc: function(componentContext,$rating,ratingObjectRef) {
			console.log("Init rating in view form")
			initRatingRecordEditBehavior(componentContext,
					getCurrentRecordFunc,updateCurrentRecordFunc,ratingObjectRef)
			initFormComponentViewBehavior($rating,
					ratingObjectRef.ratingID,initRatingViewProperties)
		},
		initFormButtonFunc: function(componentContext,$button,buttonObjectRef) {
			console.log("Init form button in view form")
			
			// The loadFormViewComponents and loadRecordIntoFormLayout functions
			// need to be passed to initFormButtonRecordEditBehavior in order
			// to avoid a cyclical package dependency.
			initFormButtonRecordEditBehavior(componentContext,
					getCurrentRecordFunc,updateCurrentRecordFunc,buttonObjectRef,
					loadFormViewComponents,loadRecordIntoFormLayout)
		},
		initUserSelectionFunc: function(componentContext,$userSelection,userSelectionObjectRef) {
			console.log("Init user selection in view form")
			initUserSelectionRecordEditBehavior(componentContext,
					getCurrentRecordFunc,updateCurrentRecordFunc,userSelectionObjectRef)
			initFormComponentViewBehavior($userSelection,
					userSelectionObjectRef.userSelectionID,initUserSelectionViewProperties)
		},
		initDatePickerFunc: function(componentContext,$datePicker,datePickerObjectRef) {
			console.log("Init date picker in view form")
			initDatePickerRecordEditBehavior(componentContext,
					getCurrentRecordFunc,updateCurrentRecordFunc,datePickerObjectRef)
			initFormComponentViewBehavior($datePicker,
					datePickerObjectRef.datePickerID,initDatePickerViewProperties)
		},
		initHtmlEditorFunc: function(componentContext,$htmlEditor,htmlEditorObjectRef) {
			console.log("Init html editor in view form")
			initHtmlEditorRecordEditBehavior(componentContext,
					getCurrentRecordFunc,updateCurrentRecordFunc,htmlEditorObjectRef)
			initFormComponentViewBehavior($htmlEditor,
					htmlEditorObjectRef.htmlEditorID,initHTMLEditorViewProperties)
		},
		initImageFunc: function(componentContext,$image,imageObjectRef) {
			console.log("Init image in view form")
			initImageRecordEditBehavior(componentContext,
				getCurrentRecordFunc,updateCurrentRecordFunc,imageObjectRef)
			initFormComponentViewBehavior($image,
					imageObjectRef.imageID,initImageViewProperties)
		},
		initHeaderFunc: function(componentContext,headerObjectRef) {
			console.log("Init header in view form")
			initHeaderRecordEditBehavior(componentContext,headerObjectRef)
		},
		
		doneLoadingFormDataFunc: doneLoadingComponentsFunc
	});
}