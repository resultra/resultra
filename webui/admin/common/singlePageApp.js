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

function clearSettingsPageOffPageContent() {
	$('#offPageContentSection').empty()
}

var registeredSettingsPageContentLoaders = {}

function registerPageContentLoader(linkID, contentURL, initCallbackFunc) {
	var contentInfo = {
		contentURL: contentURL,
		initContentFunc: initCallbackFunc
	}
	registeredSettingsPageContentLoaders[linkID] = contentInfo
}

function resetSettingsPageLayoutForStandardSettingsPages() {
		// For regular settings pages, the default behavior is the show the LHS
		// table of contents for top-level settings pages. If the pages are being navigated
		// to from the dashboard or form designer, the RHS sidebar also needs to be disabled,
		// and the offpage content needs to be cleared out.
		theMainWindowLayout.openLHSSidebar()
		theMainWindowLayout.disableRHSSidebar()
		clearSettingsPageOffPageContent()
}

function navigateToSettingsPageContent(linkID) {
	var contentInfo = registeredSettingsPageContentLoaders[linkID]
	if (contentInfo !== undefined) {
		setSettingsPageContent(contentInfo.contentURL,contentInfo.initContentFunc)
		resetSettingsPageLayoutForStandardSettingsPages()
		
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
