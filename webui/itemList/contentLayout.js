function ItemListContentLayout() {
	
	var contentLayout = $('#listViewContentLayout').layout({
		onresize_end: function(pane, $pane, paneState, paneOptions) {
			if(pane === 'center'){
				// only propagate the resize event for the center/content pane
				console.log("resize triggered")
			//	resizeMainWindow()
			}
		},
		north: fixedUILayoutPaneAutoSizeToFitContentsParams(),
		south: {
			size: 44,
			resizable:false,
			slidable: false,
			spacing_open:0,
			spacing_closed:0,
			initClosed:true // panel is initially closed	
		},
		north__showOverflowOnHover:	true,
		south__showOverflowOnHover:	true
	})
	
	function hideFooterLayout() {
		contentLayout.close("south")
	}
	
	function showFooterLayout() {
		contentLayout.open("south")
	}
	
	
	var $refreshButton = $('#listViewRefreshButton')
	
	function disableRefreshButton() {
		$refreshButton.hide()
	}
	
	function enableRefreshButton(refreshCallback) {
		initButtonControlClickHandler($refreshButton,refreshCallback)
		$refreshButton.show()
	}
	
	function setCenterContentHeader(header) {
		var $header = $('#itemListCenterContentHeader')
		$header.text(header)
	//	var $breadcrumb = $('#trackerLocationBreadcrumb')
	//	$breadcrumb.text(header)
	}
	
	initButtonClickHandler("#viewTableOfContentsMenuButton", function() {
		console.log("TOC button clicked")
		theMainWindowLayout.toggleLHSSidebar()
	})
	
	
	this.showFooterLayout = showFooterLayout
	this.hideFooterLayout = hideFooterLayout
	this.enableRefreshButton = enableRefreshButton
	this.setCenterContentHeader = setCenterContentHeader
	
}