
var paletteItemsEditConfig = {
	paletteItemTextBox: textBoxDesignFormConfig,
	paletteItemCheckBox: checkBoxDesignFormConfig,
	paletteItemDatePicker: datePickerDesignFormConfig,
	paletteItemHtmlEditor: htmlEditorDesignFormConfig,
	paletteItemImage: imageDesignFormConfig,
	paletteItemHeader: formHeaderDesignFormConfig,
	paletteItemRating: ratingDesignFormConfig,
	paletteItemSelection: selectionDesignFormConfig,
	paletteItemUserSelection: userSelectionDesignFormConfig,
	paletteItemComment: commentDesignFormConfig,
	paletteItemButton: formButtonDesignFormConfig
}

var formDesignCanvasSelector = "#layoutCanvas"


function createFormLayoutDesignConfig() {
	function saveUpdatedFormComponentLayout(updatedLayout) {
		console.log("saveUpdatedFormComponentLayout: component layout = " + JSON.stringify(updatedLayout))		
		var setLayoutParams = {
			formID: designFormContext.formID,
			layout: updatedLayout
		}
		jsonAPIRequest("frm/setLayout", setLayoutParams, function(formInfo) {
		})
	}		
	
	
	var designFormLayoutConfig =  {
		parentLayoutSelector: formDesignCanvasSelector,
		saveLayoutFunc: saveUpdatedFormComponentLayout
	}
	
	return designFormLayoutConfig
}

$(document).ready(function() {
	
	initUserDropdownMenu()
	
					
	var paletteConfig = {
		draggableItemHTML: function(placeholderID,paletteItemID) {
			return paletteItemsEditConfig[paletteItemID].draggableHTMLFunc(placeholderID)
		},
		
		startPaletteDrag: function(placeholderID,paletteItemID,$paletteItemContainer) {
			// If a palette item needs to initialize the dragged item after it's been
			// inserted into the DOM, then this is done in the startPaletteDrag function
			return paletteItemsEditConfig[paletteItemID].startPaletteDrag(placeholderID,$paletteItemContainer)			
		},
		
		dropComplete: function(droppedItemInfo) {			
			// After the drag operation is complete, the resizable
			// properties need to be initialized.
			//
			// At this point, the placholder div for the bar chart will have just been inserted. However, the DOM may 
			// not be completely updated at this point. To ensure this, a small delay is needed before
			// drawing the dummy bar charts. See http://goo.gl/IloNM for more.
			var objEditConfig = paletteItemsEditConfig[droppedItemInfo.paletteItemID]
			
			setTimeout(function() {
				// TODO - need to pass "layoutDesignConfig" parameter to initObjectGridEditBehavior
				initObjectGridEditBehavior(droppedItemInfo.droppedElem,objEditConfig) 
			}, 50);
					
			// "repackage" the dropped item paramaters for creating a new layout element. Also add the formID
			// to the parameters.
			var containerParams = {
				parentFormID: formID,
				geometry: droppedItemInfo.geometry,
				containerID: droppedItemInfo.placeholderID,
				containerObj: droppedItemInfo.droppedElem,
				finalizeLayoutIncludingNewComponentFunc: droppedItemInfo.finalizeLayoutIncludingNewComponentFunc
				};
				
			objEditConfig.createNewItemAfterDropFunc(designFormContext.databaseID,formID,containerParams)
		},
		
		dropDestSelector: formDesignCanvasSelector,
		paletteSelector: "#paletteSidebar",
	}
	
	var designFormPaletteLayoutConfig =  {
		parentLayoutSelector: formDesignCanvasSelector,
		saveLayoutFunc: function(updatedLayout) { } // no-op: layout gets saved after placeholder replaced with real component.
	}
	
	
	initDesignPalette(paletteConfig,designFormPaletteLayoutConfig)			
	
	// Initialize all the different plug-ins/configurations for text boxes, check boxes, etc.
	console.log("designForm: Initializing form design plug-ins/configurations ...")
	$.each(paletteItemsEditConfig, function (i, designFormConfig) {
		designFormConfig.initFunc()
	})
	console.log("designForm: Done initializing form design plug-ins/configurations.")
			
	// Initialize the page layout
	var formDesignLayout = $('#layoutPage').layout({
		north: fixedUILayoutPaneParams(40),
		east: fixedUILayoutPaneParams(300),
		west: fixedUILayoutPaneParams(200),
		south: fixedInitiallyHiddenUILayoutPaneAutoSizeToFitContentsParams(),
		// Important: The 'showOverflowOnHover' options give a higher
		// z-index to sidebars and other panels with controls, etc. Otherwise
		// popups and other controlls will not be shown on top of the rest
		// of the layout.
		west__showOverflowOnHover:	true,
		south__showOverflowOnHover:	true 
	})
	function showFormulaEditPane() { formDesignLayout.open("south") }
	function hideFormulaEditPanel() { formDesignLayout.close("south")}
	var formulaEditorParams = {
		databaseID: designFormContext.databaseID,
		showEditorFunc:showFormulaEditPane,
		hideEditorFunc:hideFormulaEditPanel
	}
		
	var designFormLayoutConfig =  createFormLayoutDesignConfig()
	var $parentFormLayout = $(formDesignCanvasSelector)
	
	loadFormComponentsIntoSingleLayout({
		$parentFormLayout: $parentFormLayout,
		formContext: designFormContext,
		initTextBoxFunc: function(componentContext,$textBox,textBoxObjectRef) {
			var componentIDs = { formID: formID, componentID: textBoxObjectRef.textBoxID }
			initFormComponentDesignBehavior($textBox,componentIDs,textBoxObjectRef,textBoxDesignFormConfig,designFormLayoutConfig)
		},
		initSelectionFunc: function(componentContext,$selection,selectionObjectRef) {
			var componentIDs = { formID: formID, componentID: selectionObjectRef.selectionID }
			initFormComponentDesignBehavior($selection,componentIDs,selectionObjectRef,selectionDesignFormConfig,designFormLayoutConfig)
		},
		initCheckBoxFunc: function(componentContext,$checkBox,checkBoxObjectRef) {
			var componentIDs = { formID: formID, componentID: checkBoxObjectRef.checkBoxID }
			initFormComponentDesignBehavior($checkBox,componentIDs,checkBoxObjectRef,checkBoxDesignFormConfig,designFormLayoutConfig)
		},
		initCommentFunc: function(componentContext,$comment,commentObjectRef) {
			var componentIDs = { formID: formID, componentID: commentObjectRef.commentID }
			initFormComponentDesignBehavior($comment,componentIDs,commentObjectRef,commentDesignFormConfig,designFormLayoutConfig)
		},
		initRatingFunc: function(componentContext,$rating,ratingObjectRef) {
			initRatingDesignControlBehavior($rating,ratingObjectRef)
			var componentIDs = { formID: formID, componentID: ratingObjectRef.ratingID }
			initFormComponentDesignBehavior($rating,componentIDs,ratingObjectRef,ratingDesignFormConfig,designFormLayoutConfig)
		},
		initUserSelectionFunc: function(componentContext,$userSelection,userSelectionObjectRef) {
			initUserSelectionDesignControlBehavior(userSelectionObjectRef)
			var componentIDs = { formID: formID, componentID: userSelectionObjectRef.userSelectionID }
			initFormComponentDesignBehavior($userSelection,componentIDs,userSelectionObjectRef,
						userSelectionDesignFormConfig,designFormLayoutConfig)
		},
		initDatePickerFunc: function(componentContext,$datePicker,datePickerObjectRef) {
			var componentIDs = { formID: formID, componentID: datePickerObjectRef.datePickerID }
			initFormComponentDesignBehavior($datePicker,componentIDs,datePickerObjectRef,datePickerDesignFormConfig,designFormLayoutConfig)
		},
		initHtmlEditorFunc: function(componentContext,$htmlEditor,htmlEditorObjectRef) {
			var componentIDs = { formID: formID, componentID: htmlEditorObjectRef.htmlEditorID }
			initFormComponentDesignBehavior($htmlEditor,componentIDs,htmlEditorObjectRef,htmlEditorDesignFormConfig,designFormLayoutConfig)
		},
		initImageFunc: function(componentContext,$image,imageObjectRef) {
			var componentIDs = { formID: formID, componentID: imageObjectRef.imageID }
			initFormComponentDesignBehavior($image,componentIDs,imageObjectRef,imageDesignFormConfig,designFormLayoutConfig)
		},
		initHeaderFunc: function($header,componentContext,headerObjectRef) {
			var componentIDs = { formID: formID, componentID: headerObjectRef.headerID }
			initFormComponentDesignBehavior($header,componentIDs,headerObjectRef,formHeaderDesignFormConfig,designFormLayoutConfig)
		},
		initFormButtonFunc: function(componentContext,$button,buttonObjectRef) {
			var componentIDs = { formID: formID, componentID: buttonObjectRef.buttonID }
			initFormComponentDesignBehavior($button,componentIDs,buttonObjectRef,formButtonDesignFormConfig,designFormLayoutConfig)
		},
		doneLoadingFormDataFunc: function() {
			// The formula editor depends on the field information first being initialized.
			initFormulaEditor(formulaEditorParams)
			
		} // no-op	
	}); 
		
		
	
	console.log("Initializing form design plug-ins/configurations ...")
	initObjectCanvasSelectionBehavior(formDesignCanvasSelector, function() {
		console.log("Select form canvas")
		hideSiblingsShowOne('#formProps')
		closeFormulaEditor()
		initDesignFormProperties(formID)
	})
	
	
});
