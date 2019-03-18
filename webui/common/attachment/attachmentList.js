// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


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