function initAttachmentContainerPopupGallery($attachContainer) {
	$attachContainer.magnificPopup({
			delegate: 'a',
			type: 'image',
			gallery: { enabled: true },
			mainClass: 'attachmentPopupGallery',
			image: {
				tError: 'The image could not be loaded.',
				titleSrc: function(item) {
					var $attachContainer = $(item.el)
					var attachRef = $attachContainer.data("attachRef")
					return attachmentTitleAndCaptionHTML(attachRef)
				}
			},
			inline: {
				markup: '<div class="attachmentListInlineItem well">'+
				            '<div class="mfp-close"></div>'+
							'<div>'+
								'<a class="attachDownloadLink">'+
									'<i class="attachmentThumbnailIcon glyphicon glyphicon-file"></i> '+
									'<span class="attachmentListLinkText">TBD</span>'+
								'</a>' + 
							'</div>' +
							'<div class="attachCaption"></div>'+
							'<div class="mfp-counter marginTop10"></div>'+
				         '</div>'
			},
			callbacks : {
				  markupParse: function($template, values, item) {
					  
					  if (item.type === "inline") {
							var $attachContainer = $(item.el)
							var attachRef = $attachContainer.data("attachRef")
						  
					  		var $linkText = $template.find(".attachmentListLinkText")
						  	$linkText.text(attachRef.attachmentInfo.title)
						  
						  	var $attachLink = $template.find(".attachDownloadLink")
						  	$attachLink.attr("href",attachRef.url)
						  
						 	var $attachCaption = $template.find(".attachCaption")
						  	$attachCaption.empty()
						  	$attachCaption.append( attachmentCaptionHTML(attachRef))
						  
					  }
				      // Triggers each time when content of popup changes
				      console.log('Parsing:', $template, values, item);
				    },
				}
		});
		
		setTimeout(function() {
			// IMPORTANT NOTE: There is a conflict between document event handling for 
			// magnific popup and Bootstrap. When setting the z-index of the Magnific
			// popup container, a stack overflow is caused in Javascript. The problem and 
			// potential workarounds are described here:
			// 
			// https://github.com/dimsemenov/Magnific-Popup/issues/796
			//
			
			$(document).off('focusin');
		},200)

}


function initSingleAttachmentImagePopupLink($parentContainer,$link,attachmentID) {
	if (attachmentID !== null) {
		var getRefParams = { attachmentID: attachmentID }
		jsonAPIRequest("attachment/getReference", getRefParams, function(attachRef) {
			$link.text(attachRef.attachmentInfo.title)
			$link.attr("href",attachRef.url)
			$link.data("attachRef",attachRef)
			$link.addClass("mfp-image")	
			initAttachmentContainerPopupGallery($parentContainer)
		})		
	} else {
		$link.text("")
		$link.attr("href","")
		$link.attr("attachRef",null)	
	}
	
}


function initSingleAttachmentImagePopupThumbnail($parentContainer,$link,attachmentID,maxWidth,maxHeight) {
	if (attachmentID !== null) {
		var getRefParams = { attachmentID: attachmentID }
		jsonAPIRequest("attachment/getReference", getRefParams, function(attachRef) {
			$link.attr("href",attachRef.url)
			$link.data("attachRef",attachRef)
			$link.addClass("mfp-image")
			
			var $thumbnailImage = $('<img class="attachmentThumbnailImage">')
			$thumbnailImage.attr("src",attachRef.url)
			
			// Allow the thumbnail to resize to the given proportions, but keep the aspect ratio.
			// https://stackoverflow.com/questions/12991351/css-force-image-resize-and-keep-aspect-ratio
			// Other CSS for thumbnail sizing and positino is set on the attachmentThumbnailImage class.
			$thumbnailImage.css("max-width",maxWidth+"px")
			$thumbnailImage.css("max-height",maxHeight+"px")			

			$link.find('img').remove()
			$link.append($thumbnailImage )
			
			initAttachmentContainerPopupGallery($parentContainer)
		})		
	} else {
		$link.text("")
		$link.attr("href","")
		$link.attr("attachRef",null)	
		
		$link.find('img').remove()
	}
	
}