// Definition of parameters and callbacks for a progess indicator to be editable within the form editor.
// this javascript file needs to included after the other check box related files, so all the functions
// are already defined.

function initDesignFormGauge() {
	console.log("Init gauge design form behavior")
}

function selectFormGauge($container,gaugeObjRef) {
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

function initDummyGaugeControlForPaletteDrag($paletteItemContainer) {
	
	console.log("initDummyGaugeControlForPaletteDrag")
	var gaugeConfig = 
	{
		size: 120, // same as default width (from gaugeComponent.css) for the component's container
		min: 0,
		max: 100,
		minorTicks: 5,
	}
	initGaugeComponentControl($paletteItemContainer,gaugeConfig)
}


var gaugeDesignFormConfig = {
	draggableHTMLFunc:	gaugeContainerHTML,
	initDummyDragAndDropComponentContainer: initDummyGaugeControlForPaletteDrag,
	createNewItemAfterDropFunc: openNewGaugeDialog,
	resizeConstraints: elemResizeConstraints(75,640,30,30),
	resizeFunc: resizeGauge,
	initFunc: initDesignFormGauge,
	selectionFunc: selectFormGauge
}
