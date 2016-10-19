function loadRecordIntoRating(ratingElem, recordRef) {
	
	
	var ratingObjectRef = ratingElem.data("objectRef")
	var ratingContainerID = ratingObjectRef.ratingID
	var ratingControlID = ratingControlIDFromElemID(ratingContainerID)
	var ratingControlSelector = '#' + ratingControlID;
	
	var componentLink = ratingObjectRef.properties.componentLink
	
	if(componentLink.linkedValType == linkedComponentValTypeField) {
		var ratingFieldID = componentLink.fieldID
	
		console.log("loadRecordIntoRating: Field ID to load data:" + ratingFieldID)

		// TBD - initialize control
		
	} else {
		var ratingGlobalID = componentLink.globalID
		console.log("loadRecordIntoRating: Global ID to load data:" + ratingGlobalID)

		// TBD - initialize control
		
	}
	
	
}

function initRatingFieldEditBehavior(componentContext,ratingObjectRef) {
	
	var ratingSelector = '#'+ratingControlIDFromElemID(ratingObjectRef.ratingID)
	
	var componentLink = ratingObjectRef.properties.componentLink
	
	var fieldRef = getFieldRef(componentLink.fieldID)
	if(fieldRef.isCalcField) {
		// TBD - Set control to read-only
		return;  // stop initialization, the check box is read only.
	}
	

  	$(ratingSelector).click( function () {
		
			// Get the most recent copy of the object reference. It could have changed between
			// initialization time and the time the checkbox was changed.
			var objectRef = getElemObjectRef(ratingObjectRef.ratingID)
			var componentLink = objectRef.properties.componentLink
			
			var currRecordRef = currRecordSet.currRecordRef()

			// TBD - set field value for rating
	})
	
}

function initRatingGlobalEditBehavior(componentContext,ratingObjectRef) {

	var ratingSelector = '#'+ratingControlIDFromElemID(ratingObjectRef.ratingID)
		
  	$(raterSelector).click( function () {
		
		// Get the most recent copy of the object reference. It could have changed between
		// initialization time and the time the checkbox was changed.
		var objectRef = getElemObjectRef(raterObjectRef.raterID)
		var componentLink = objectRef.properties.componentLink
		
		// TBD 		
/*		var setGlobalValParams = {
			parentDatabaseID: componentContext.databaseID,
			globalID: componentLink.globalID,
			value: isChecked } */
//		console.log("Setting rating value (global): " + JSON.stringify(setGlobalValParams))
//		jsonAPIRequest("global/setBoolValue",setGlobalValParams,function(updatedGlobalVal) {
			
			// TODO - Update global structure with updated value.
///		}) // set record's text field value
		
	})
	
}

function initRatingRecordEditBehavior(componentContext,ratingObjectRef) {

	var $ratingContainer = $('#'+ratingObjectRef.ratingID)
		
	$ratingContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoRating
	})
	
	var componentLink = ratingObjectRef.properties.componentLink

	if(componentLink.linkedValType == linkedComponentValTypeField) {
		initRatingFieldEditBehavior(componentContext,ratingObjectRef)
	} else {
		initRatingGlobalEditBehavior(componentContext,ratingObjectRef)
	}
}