
var paletteItemsEditConfig = {
	paletteItemTextBox: textBoxDesignFormConfig,
	paletteItemCheckBox: checkBoxDesignFormConfig
}

var formDesignCanvasSelector = "#layoutCanvas"

$(document).ready(function() {
					
	var paletteConfig = {
		draggableItemHTML: function(placeholderID,paletteItemID) {
			return paletteItemsEditConfig[paletteItemID].draggableHTMLFunc(placeholderID)
		},
		
		dropComplete: function(droppedItemInfo) {
			console.log("designForm: Dashboard design pallete: drop item: " + JSON.stringify(droppedItemInfo))			
			
			// After the drag operation is complete, the draggable and resizable
			// properties need to be initialized.
			//
			// At this point, the placholder div for the bar chart will have just been inserted. However, the DOM may 
			// not be completely updated at this point. To ensure this, a small delay is needed before
			// drawing the dummy bar charts. See http://goo.gl/IloNM for more.
			var objEditConfig = paletteItemsEditConfig[droppedItemInfo.paletteItemID]
			
			setTimeout(function() {
				initObjectEditBehavior(layoutID,droppedItemInfo.placeholderID,objEditConfig) 
			}, 50);
			
			
			// "repackage" the dropped item paramaters for creating a new layout element. Also add the layoutID
			// to the parameters.
			var layoutContainerParams = {
				parentLayoutID: layoutID,
				geometry: droppedItemInfo.geometry,
				containerID: droppedItemInfo.placeholderID,
				};
				
			objEditConfig.createNewItemAfterDropFunc(layoutContainerParams)
		},
		
		dropDestSelector: formDesignCanvasSelector,
		paletteSelector: "#gallerySidebar",
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
		showEditorFunc:showFormulaEditPane,
		hideEditorFunc:hideFormulaEditPanel
	}
	
	function initFormComponentDesignBehavior(objectRef, designFormConfig) {
		initObjectEditBehavior(objectRef.uniqueID.parentID,
			objectRef.uniqueID.objectID,designFormConfig)
		initObjectSelectionBehavior($("#"+objectRef.uniqueID.objectID), 
				formDesignCanvasSelector,function(objectID) {
			console.log("form design object selected: " + objectID)
			var selectedObjRef	= getElemObjectRef(objectID)
			designFormConfig.selectionFunc(selectedObjRef)
		})	
		
	}	  
	  
	
	loadFormComponents({
		formParentElemID: formDesignCanvasSelector,
		initTextBoxFunc: function(textBoxObjectRef) {
			initFormComponentDesignBehavior(textBoxObjectRef,textBoxDesignFormConfig)
		},
		initCheckBoxFunc: function(checkBoxObjectRef) {
			initFormComponentDesignBehavior(checkBoxObjectRef,checkBoxDesignFormConfig)
		},
		doneLoadingFormDataFunc: function() {} // no-op	
	}); 
	
	console.log("Initializing form design plug-ins/configurations ...")
	initObjectCanvasSelectionBehavior(formDesignCanvasSelector, function() {
		console.log("Select form canvas")
		hideSiblingsShowOne('#formProps')
		closeFormulaEditor()
	})
	$( '#formProps' ).accordion();	
	
	initFormulaEditor(formulaEditorParams)
	
});
