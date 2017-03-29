// Definition of parameters and callbacks for a progess indicator to be editable within the form editor.
// this javascript file needs to included after the other check box related files, so all the functions
// are already defined.

function initDesignFormGauge() {
	console.log("Init gauge design form behavior")
}

function selectFormGauge($container,gaugeObjRef) {
	console.log("Selected gauge: " + JSON.stringify(gaugeObjRef))
	loadGaugeProperties($container,gaugeObjRef)
}

function resizeGauge($container,geometry) {
	
	var gaugeRef = getContainerObjectRef($container)
	 
	var resizeParams = {
		parentFormID: designFormContext.formID,
		gaugeID: gaugeRef.gaugeID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/gauge/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.gaugeID)
	})	
}


var gaugeDesignFormConfig = {
	draggableHTMLFunc:	gaugeContainerHTML,
	startPaletteDrag: function(placeholderID,$paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewGaugeDialog,
	resizeConstraints: elemResizeConstraints(75,640,30,30),
	resizeFunc: resizeGauge,
	initFunc: initDesignFormGauge,
	selectionFunc: selectFormGauge
}
