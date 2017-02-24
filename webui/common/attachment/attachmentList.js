

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
			$attachLink.attr("href",attachRef.url)
			$attachLink.data("attachRef",attachRef)
			
			$attachmentListContainer.append($attachListItem)
		}
		
		$attachmentListContainer.magnificPopup({
				delegate: 'a',
				type: 'image',
				image: {
					tError: '<a href="The image could not be loaded.',
					titleSrc: function(item) {
						var $attachContainer = $(item.el)
						var attachRef = $attachContainer.data("attachRef")
						return attachmentCaptionHTML(attachRef)
					}
				}
			});
		
	})
					
}