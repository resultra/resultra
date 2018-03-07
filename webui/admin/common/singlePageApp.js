function setSettingsPageContent(contentURL, initContentCallback) {
	jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
	        $('#contentPageSection').html(pageContentData);
			initContentCallback()
	});	
}

function setSettingsPageOffPageContent(contentURL,initCallback) {
	jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
	        $('#offPageContentSection').html(pageContentData);
			initCallback()
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

function initSettingsPageButtonLink(buttonSelector,pageContentID) {
	var $button = $(buttonSelector)
	$button.click(function(e) {
		e.preventDefault()
		$button.blur()
		navigateToSettingsPageContent(pageContentID)	
	})	
}


function setPageContentButtonClickHandler($button,contentURL,initContentFunc) {
	$button.click(function(e) {
		e.preventDefault()
		$button.blur()
		setSettingsPageContent(contentURL,initContentFunc)			
	})
}