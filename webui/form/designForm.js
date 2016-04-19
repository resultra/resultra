
var paletteItemsEditConfig = {
	paletteItemTextBox: textBoxDesignFormConfig,
	paletteItemCheckBox: checkBoxDesignFormConfig
}

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
		
		dropDestSelector: "#layoutCanvas",
		paletteSelector: "#gallerySidebar",
	}
	initDesignPalette(paletteConfig)			
	
	initNewTextBoxDialog()
	initNewCheckBoxDialog()
		
	// Initialize the page layout
	$('#layoutPage').layout({
		north: fixedUILayoutPaneParams(40),
		east: fixedUILayoutPaneParams(300),
		west: fixedUILayoutPaneParams(200),
		west__showOverflowOnHover:	true
	})	  
	  
	
	loadFormComponents({
		formParentElemID: "#layoutCanvas",
		initTextBoxFunc: function(textBoxObjectRef) {
			initObjectEditBehavior(textBoxObjectRef.uniqueID.parentID,textBox.uniqueID.objectID,textBoxDesignFormConfig)
		},
		doneLoadingFormDataFunc: function() {} // no-op	
	}); 
});
