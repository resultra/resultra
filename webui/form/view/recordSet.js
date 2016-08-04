
// A RecordSet manages the pagination of records in a form.

function RecordSet(recordRefData) {
	
	this.currRecordIter = 0
	this.recordRefs = recordRefData

	this.numRecords = function() {
		return this.recordRefs.length
	}
	
	this.currRecordRef = function() {
		if(this.numRecords() > 0) {
			return this.recordRefs[this.currRecordIter]
		}
		else {
			return null
		}
	}
	
	this.advanceToNextRecord = function() {
		var nextRecordIter = this.currRecordIter + 1
		if(nextRecordIter < this.recordRefs.length)
		{
			this.currRecordIter = nextRecordIter
			return true
		}
		else {
			return false
		}
	}

	this.advanceToPrevRecord = function() {
		var prevRecordIter = this.currRecordIter - 1
		if((prevRecordIter >= 0) && (prevRecordIter < this.recordRefs.length)) {
			this.currRecordIter = prevRecordIter
			return true
		}
		else {
			return false
		}
	}
	
	this.jumpToRecord = function(recordID) {
		// Use an integer iterator, since the value of the matching record
		// will be used to update this.currRecordIter. Otherwise, a string
		// will be assigned to this.currRecordIter.
		for (recIter = 0; recIter < this.recordRefs.length; recIter++) {
			if(this.recordRefs[recIter].recordID == recordID) {
				this.currRecordIter = recIter
				return true
			}
		}
		return false // TODO - some type of better error handling/assertion checking needed here
	}
	
	this.updateRecordRef = function(updatedRecordRef) {
		for (recIter in this.recordRefs) {
			if(this.recordRefs[recIter].recordID == updatedRecordRef.recordID) {
				this.recordRefs[recIter] = updatedRecordRef
				return
			}
		}
	}
	
	this.appendNewRecord = function(recordRef) {
		this.recordRefs.push(recordRef)
	}
	
	this.currRecordNum = function() {
		return this.currRecordIter + 1 // iterator is 0 based
	}
		
	this.recPageLabel = function () {
		var recNumLabel = (this.currRecordIter + 1).toString()
		var totalRecsLabel = this.numRecords().toString()		
		var recLabel = recNumLabel + " of " + totalRecsLabel
		return recLabel
	}
	
}
