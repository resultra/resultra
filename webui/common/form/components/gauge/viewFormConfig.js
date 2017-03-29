function loadRecordIntoGauge($gaugeContainer, recordRef) {
		
	var gaugeObjectRef = getContainerObjectRef($gaugeContainer)
	
	var gaugeFieldID = gaugeObjectRef.properties.fieldID
	
	function setGaugeVal(gaugeVal) {
		// TBD
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
	
    // TBD
	
}