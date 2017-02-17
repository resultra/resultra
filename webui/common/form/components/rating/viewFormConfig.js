function loadRecordIntoRating(ratingElem, recordRef) {
	
	
	var ratingObjectRef = getContainerObjectRef(ratingElem)
	var $ratingControl = getRatingControlFromRatingContainer(ratingElem)
	
	var ratingFieldID = ratingObjectRef.properties.fieldID

	console.log("loadRecordIntoRating: Field ID to load data:" + ratingFieldID)

	// In other words, we are populating the "intersection" of field values in the record
	// with the fields shown by the layout's containers.
	if(recordRef.fieldValues.hasOwnProperty(ratingFieldID)) {

		var fieldVal = recordRef.fieldValues[ratingFieldID]

		console.log("loadRecordIntoTextBox: Load value into container: " + $(this).attr("id") + " field ID:" + 
					ratingFieldID + "  value:" + fieldVal)
		
		var maxRating = 5
		if((fieldVal >= 0) && (fieldVal <= maxRating)) {
			$ratingControl.rating('rate',fieldVal)	
		} else {
			$ratingControl.rating('rate','')		
		}
		
	} // If record has a value for the current container's associated field ID.
	else
	{
		$ratingControl.rating('rate','')
	}
		
	
}


function initRatingRecordEditBehavior($ratingContainer,componentContext,recordProxy, ratingObjectRef) {

	var $ratingControl = getRatingControlFromRatingContainer($ratingContainer)

	function setRatingValue(ratingVal) {
		
		currRecordRef = recordProxy.getRecordFunc()
		var ratingFieldID = ratingObjectRef.properties.fieldID

		var ratingValueFormat = { context: "rating", format: "star" }
		var setRecordValParams = { 
			parentDatabaseID:currRecordRef.parentDatabaseID,
			recordID:currRecordRef.recordID,
			changeSetID: recordProxy.changeSetID,
			fieldID:ratingFieldID, 
			value:ratingVal,
			valueFormat: ratingValueFormat}
		jsonAPIRequest("recordUpdate/setNumberFieldValue",setRecordValParams,function(replyData) {
			// After updating the record, the local cache of records in currentRecordSet will
			// be out of date. So after updating the record on the server, the locally cached
			// version of the record also needs to be updated.
			recordProxy.updateRecordFunc(replyData)
	
		}) // set record's number field value
		
	}

	$ratingControl.rating({
		extendSymbol: function(rating) {
			var ratingIndex = rating-1 // 0 based index
			if(ratingObjectRef.properties.tooltips[ratingIndex] !== undefined) {
				var tooltipText = ratingObjectRef.properties.tooltips[ratingIndex]
				if(tooltipText.length > 0) {
					var tooltipHTML = '<p class="ratingTooltip">' + escapeHTML(tooltipText) + '</p>'
					$(this).tooltip({
						container: 'body',
						placement: 'bottom',
						title: tooltipHTML,
						html: true 
					});
					
				}
			}
			
		}
	})
	$ratingControl.on('change', function() {
		var ratingVal = Number($(this).val())
		console.log('Rating changed: ' + ratingVal);
		setRatingValue(ratingVal)
	});
	
	// When the user clicks on the control, prevent the click from propagating higher.
	// This allows the user to change the rating without selecting the form component itself.
	// The user can still select the component by clicking on the label or anywwhere outside
	// the control.
	$ratingContainer.find(".formRatingControl").click(function (event){
		event.stopPropagation();
   	 	//   ... your code here
		return false;
	});
		
	$ratingContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoRating
	})
	

}