
// A RecordSet manages the pagination of records in a form.

function RecordSet(recordRefData,windowSize) {
	
	this.currStartWindowIndex = 0
	this.recordRefs = recordRefData
	this.windowSize = windowSize

	this.numRecords = function() {
		return this.recordRefs.length
	}
	
	this.setWindowSize = function(newWindowSize) {
		this.windowSize = newWindowSize
		this.currStartWindowIndex = 0
	}
		
	this.getRecordRefAtWindowIndex = function(windowIndex) {
		var windowIndexIter = this.currStartWindowIndex + windowIndex
		if(windowIndexIter < this.recordRefs.length) {
			return this.recordRefs[windowIndexIter]
		} else {
			return null
		}
	}
	
	this.advanceToNextPage = function() {
		var nextPageStartIter = this.currStartWindowIndex + this.windowSize
		
		// Only advance the starting index of the record set if the window size doesn't 
		// overrun the length of the record set.
		if (nextPageStartIter >= this.recordRefs.length) {
			return false
		} else {
			this.currStartWindowIndex = nextPageStartIter
			return true
		}
	}

	this.advanceToPrevPage = function() {
		var prevPageStartIter = this.currStartWindowIndex - this.windowSize
		
		if((prevPageStartIter >= 0) && (prevPageStartIter < this.recordRefs.length)) {
			this.currStartWindowIndex = prevPageStartIter
			return true
		}
		else {
			return false
		}
	}
	
	this.jumpToRecord = function(recordID) {
		// Use an integer iterator, since the value of the matching record
		// will be used to update this.currStartWindowIndex. Otherwise, a string
		// will be assigned to this.currStartWindowIndex.
		for (recIter = 0; recIter < this.recordRefs.length; recIter++) {
			if(this.recordRefs[recIter].recordID == recordID) {
				this.currStartWindowIndex = recIter
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
		return this.currStartWindowIndex + 1 // iterator is 0 based
	}
		
	this.recPageLabel = function () {
		var startWindowRecLabel = (this.currStartWindowIndex + 1).toString()
		
		var pageLabel
		if (this.windowSize > 1) {
			var endIndex = this.currStartWindowIndex + this.windowSize
			var endPageRecLabel
			if (endIndex < this.numRecords()) {
				endPageRecLabel = endIndex.toString() 
			} else {
				endPageRecLabel = this.numRecords().toString()
			}
			pageLabel = startWindowRecLabel + "-" + endPageRecLabel
			
		} else {
			pageLabel = startWindowRecLabel
		}
		
		var totalRecsLabel = this.numRecords().toString()
				
		var recLabel = pageLabel + " of " + totalRecsLabel
		
		return recLabel
	}
	
}
