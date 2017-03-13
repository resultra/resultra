

function populateAttachmentList($attachmentListContainer,attachmentList) {
	
	var getRefParams = { attachmentIDs: attachmentList }
	jsonAPIRequest("attachment/getReferences", getRefParams, function(attachRefs) {
		$attachmentListContainer.empty()
		for(var refIndex = 0; refIndex < attachRefs.length; refIndex++) {
			var attachRef = attachRefs[refIndex]
			
			var $attachListItem = $('#attachmentListItemTemplate').clone()
			$attachListItem.attr("id","")
			
			var $attachLink = $attachListItem.find(".attachmentListItemLink")
			
			$attachLink.text(attachRef.attachmentInfo.title)
			
			$attachLink.data("attachRef",attachRef)
			if(attachRef.dataType === "image") {
				$attachLink.addClass("mfp-image")
				$attachLink.attr("href",attachRef.url)
			} else {
				$attachLink.addClass("mfp-inline")
			}
						
			$attachmentListContainer.append($attachListItem)
		}
		initAttachmentContainerPopupGallery($attachmentListContainer)
			
	})
		
					
}