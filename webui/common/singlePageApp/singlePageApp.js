function getMainWindowLinkIDAnchorName() {
	var linkID = window.location.hash.substr(1);
	if (linkID === null || linkID.length === 0) {
		return ""
	}
	return linkID
}


function setMainWindowContent(contentURL, initContentCallback) {
	jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
	        $('#mainWindowContent').html(pageContentData);
			initContentCallback()
	});	
}

function setLHSSidebarContent(contentURL, initContentCallback) {
	jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
	        $('#mainWindowLHSSidebar').html(pageContentData);
			initContentCallback()
	});	
}

function clearMainWindowLHSSidebarContent() {
	$('#mainWindowLHSSidebar').empty()
}
 
function setRHSSidebarContent(contentURL, initContentCallback) {
	jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
	        $('#mainWindowRHSSidebar').html(pageContentData);
			initContentCallback()
	});	
}

function clearMainWindowRHSSidebarContent() {
	$('#mainWindowRHSSidebar').empty()
}

function setMainWindowOffPageContent(contentURL,initCallback) {
	jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
	        $('#mainWindowOffPageContent').html(pageContentData);
			initCallback()
	});	
	
}

function setMainWindowHeaderButtonsContent(contentURL,initCallback) {
	jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
	        $('#mainWindowHeaderMenuButtons').html(pageContentData);
			initCallback()
	});	
	
}

function clearMainWindowHeaderButtonsContent() {
	$('#mainWindowHeaderMenuButtons').empty()
}


var registeredMainWindowContentLoaders = {}

function registerMainWindowContentLoader(linkID, loaderFunc) {
	registeredMainWindowContentLoaders[linkID] = loaderFunc
}

function navigateToMainWindowContent(linkID) {
	var loaderFunc = registeredMainWindowContentLoaders[linkID]
	if (loaderFunc !== undefined) {
		loaderFunc()
	}
}

function setMainWindowPageContent(contentConfig,initDoneCallback) {
	
	var contentSections = []
	
	var numContentSections = 0
	
	if (contentConfig.hasOwnProperty("mainContentURL")) {
		numContentSections++
		contentSections.push({
			contentURL: contentConfig.mainContentURL,
			contentFunc: setMainWindowContent
		})
	}

	if (contentConfig.hasOwnProperty("rhsSidebarContentURL")) {
		numContentSections++
		contentSections.push({
			contentURL: contentConfig.rhsSidebarContentURL,
			contentFunc: setRHSSidebarContent
		})
	}
	
	if (contentConfig.hasOwnProperty("lhsSidebarContentURL")) {
		numContentSections++
		contentSections.push({
			contentURL: contentConfig.lhsSidebarContentURL,
			contentFunc: setLHSSidebarContent
		})
	}
	

	if (contentConfig.hasOwnProperty("offPageContentURL")) {
		numContentSections++
		contentSections.push({
			contentURL: contentConfig.offPageContentURL,
			contentFunc: setMainWindowOffPageContent
		})
	}
	
	var contentSectionsRemaining = numContentSections
	function processOneContentSection() {
		contentSectionsRemaining--
		if(contentSectionsRemaining <= 0) {
			initDoneCallback()
		}
	}
	
	for (var contentIndex = 0; contentIndex < contentSections.length; contentIndex++) {
		var currContent = contentSections[contentIndex]
		currContent.contentFunc(currContent.contentURL,processOneContentSection)
	}
}
