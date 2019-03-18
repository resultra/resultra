// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function loadRecordIntoGauge($gaugeContainer, recordRef) {
		
	var gaugeObjectRef = getContainerObjectRef($gaugeContainer)
	
	var gaugeFieldID = gaugeObjectRef.properties.fieldID
	
	function setGaugeVal(gaugeVal) {
		
		var gaugeControl = $gaugeContainer.data("gaugeControl")
		gaugeControl.redraw(gaugeVal)
	}
		
	// Populate the "intersection" of field values in the record
	// with the fields shown by the layout's containers.
	if(recordRef.fieldValues.hasOwnProperty(gaugeFieldID)) {
		var fieldVal = recordRef.fieldValues[gaugeFieldID]
		setGaugeVal(fieldVal)

	} // If record has a value for the current container's associated field ID.
	else
	{
		setGaugeVal(0.0)
	}	
	
}


function initGaugeRecordEditBehavior($gauge,componentContext,recordProxy, gaugeObjectRef) {
		
	$gauge.data("viewFormConfig", {
		loadRecord: loadRecordIntoGauge
	})	
}

