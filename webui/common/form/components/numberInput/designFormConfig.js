// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Definition of parameters and callbacks for a text box to be editable within the form editor.
// this javascript file needs to included after the other text box related files, so all the functions
// are already defined.


function initDesignFormNumberInput() {
	console.log("Init number input design form behavior")
	initNewNumberInputDialog()
}

function selectFormNumberInput ($container,numberInputRef) {
	console.log("Select number input: " + JSON.stringify(numberInputRef))
	loadNumberInputProperties($container,numberInputRef)
}

function resizeNumberInput($container,geometry) {
	
	var numberInputRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		numberInputID: numberInputRef.numberInputID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/numberInput/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.numberInputID)
	})	
}


var numberInputDesignFormConfig = {
	draggableHTMLFunc:	numberInputContainerHTML,
	initDummyDragAndDropComponentContainer: function($paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewNumberInputDialog,
	resizeConstraints: elemResizeConstraintsWidthOnly(75,600),
	resizeFunc: resizeNumberInput,
	initFunc: initDesignFormNumberInput,
	selectionFunc: selectFormNumberInput
}
