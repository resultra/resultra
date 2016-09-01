
var paletteItemsEditConfig = {
	paletteItemTextBox: textBoxDesignFormConfig,
	paletteItemCheckBox: checkBoxDesignFormConfig,
	paletteItemDatePicker: datePickerDesignFormConfig,
	paletteItemHtmlEditor: htmlEditorDesignFormConfig,
	paletteItemImage: imageDesignFormConfig
}

var formDesignCanvasSelector = "#layoutCanvas"

$(document).ready(function() {
					
	var paletteConfig = {
		draggableItemHTML: function(placeholderID,paletteItemID) {
			return paletteItemsEditConfig[paletteItemID].draggableHTMLFunc(placeholderID)
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
				initObjectGridEditBehavior(formID,droppedItemInfo.placeholderID,objEditConfig) 
			}, 50);
					
			// "repackage" the dropped item paramaters for creating a new layout element. Also add the formID
			// to the parameters.
			var containerParams = {
				parentFormID: formID,
				geometry: droppedItemInfo.geometry,
				containerID: droppedItemInfo.placeholderID,
				finalizeLayoutIncludingNewComponentFunc: droppedItemInfo.finalizeLayoutIncludingNewComponentFunc
				};
				
			objEditConfig.createNewItemAfterDropFunc(designFormContext.databaseID,formID,tableID,containerParams)
		},
		
		dropDestSelector: formDesignCanvasSelector,
		paletteSelector: "#paletteSidebar",
	}
	initDesignPalette(paletteConfig)			
	
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
		tableID: tableID,
		showEditorFunc:showFormulaEditPane,
		hideEditorFunc:hideFormulaEditPanel
	}
	
	loadFormComponents({
		formParentElemID: formDesignCanvasSelector,
		formContext: designFormContext,
		initTextBoxFunc: function(componentContext,textBoxObjectRef) {
			var componentIDs = { formID: formID, componentID: textBoxObjectRef.textBoxID }
			initFormComponentDesignBehavior(componentIDs,textBoxObjectRef,textBoxDesignFormConfig)
		},
		initCheckBoxFunc: function(componentContext,checkBoxObjectRef) {
			var componentIDs = { formID: formID, componentID: checkBoxObjectRef.checkBoxID }
			initFormComponentDesignBehavior(componentIDs,checkBoxObjectRef,checkBoxDesignFormConfig)
		},
		initDatePickerFunc: function(componentContext,datePickerObjectRef) {
			var componentIDs = { formID: formID, componentID: datePickerObjectRef.datePickerID }
			initFormComponentDesignBehavior(componentIDs,datePickerObjectRef,datePickerDesignFormConfig)
		},
		initHtmlEditorFunc: function(componentContext,htmlEditorObjectRef) {
			var componentIDs = { formID: formID, componentID: htmlEditorObjectRef.htmlEditorID }
			initFormComponentDesignBehavior(componentIDs,htmlEditorObjectRef,htmlEditorDesignFormConfig)
		},
		initImageFunc: function(componentContext,imageObjectRef) {
			var componentIDs = { formID: formID, componentID: imageObjectRef.imageID }
			initFormComponentDesignBehavior(componentIDs,imageObjectRef,imageDesignFormConfig)
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
		initDesignFormProperties(designFormContext.tableID,formID)
	})
	
	
});
