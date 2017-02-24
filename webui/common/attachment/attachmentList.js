

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
			
/*			if(attachRef.dataType === "image") {
				$attachLink.addClass("attachmentImageLink")
			} else {
				$attachLink.addClass("mfp-inline")
			}
*/			
			$attachmentListContainer.append($attachListItem)
		}
	
	/*	TODO - Get attachment list working with a light-box and mixed types.
		$attachmentListContainer.magnificPopup({
				delegate: 'a',
				type: 'image',
				gallery: { enabled: true},
				image: {
					tError: 'The image could not be loaded.',
					titleSrc: function(item) {
						var $attachContainer = $(item.el)
						var attachRef = $attachContainer.data("attachRef")
						return attachmentCaptionHTML(attachRef)
					}
				}
			});
		*/
	})
		
					
}