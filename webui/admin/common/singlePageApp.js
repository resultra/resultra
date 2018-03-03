function setSettingsPageContent(contentURL, initContentCallback) {
	jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
	        $('#contentPageSection').html(pageContentData);
			initContentCallback()
	});	
}