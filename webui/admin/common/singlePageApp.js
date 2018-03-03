function setSettingsPageContent(contentURL, initContentCallback) {
	jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
	        $('#contentPageSection').html(pageContentData);
			initContentCallback()
	});	
}

var registeredSettingsPageContentLoaders = {}

function registerPageContentLoader(linkID, contentURL, initCallbackFunc) {
	var contentInfo = {
		contentURL: contentURL,
		initContentFunc: initCallbackFunc
	}
	registeredSettingsPageContentLoaders[linkID] = contentInfo
}

function navigateToSettingsPageContent(linkID) {
	var contentInfo = registeredSettingsPageContentLoaders[linkID]
	if (contentInfo !== undefined) {
		setSettingsPageContent(contentInfo.contentURL,contentInfo.initContentFunc)
	}
}