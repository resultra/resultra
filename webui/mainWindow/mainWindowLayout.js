function MainWindowLayout(resizeCallback)
{
	var zeroPaddingInset = { top:0, bottom:0, left:0, right:0 }
	
	// Initialize the page layout
	var mainLayout = $('#layoutPage').layout({
		inset: zeroPaddingInset,
		north: fixedUILayoutPaneParams(40),
		onopen_end: function(pane, $pane, paneState, paneOptions) {
			resizeCallback()				
		},
		onclose_end: function(pane, $pane, paneState, paneOptions) {
			resizeCallback()
		},
		east: {
			size: 300,
			resizable:false,
			slidable: false,
			spacing_open:4,
			spacing_closed:4,
			initClosed:true
			
		},
		west: {
			size: 250,
			resizable:false,
			slidable: false,
			spacing_open:4,
			spacing_closed:4,
			initClosed:true // panel is initially closed	
		}
	})

	
	var contentLayout = $('#mainWindowContentPane').layout({
		onopen_end: function(pane, $pane, paneState, paneOptions) {
			resizeCallback()				
		},
		onclose_end: function(pane, $pane, paneState, paneOptions) {
			resizeCallback()
		},
		onresize_end: function(pane, $pane, paneState, paneOptions) {
			if(pane === 'center'){
				// only propagate the resize event for the center/content pane
				console.log("resize triggered")
				resizeCallback()
			}
		},
		north: fixedUILayoutPaneAutoSizeToFitContentsParams(),
		south: fixedUILayoutPaneAutoSizeToFitContentsParams(),
		north__showOverflowOnHover:	true,
		south__showOverflowOnHover:	true
	})
			
	initButtonClickHandler("#viewTableOfContentsMenuButton", function() {
		console.log("TOC button clicked")
		mainLayout.toggle("west")
	})
	
	
	var $viewListOptionsToggleButton = $("#viewListOptionsButton")
	initButtonControlClickHandler($viewListOptionsToggleButton, function() {
		var layoutState = mainLayout.state
		var $iconSpan = $viewListOptionsToggleButton.find("span")
		if (layoutState.east.isClosed) {
			$iconSpan.removeClass("fa-toggle-left")
			$iconSpan.addClass("fa-toggle-right")
		} else {
			$iconSpan.removeClass("fa-toggle-right")
			$iconSpan.addClass("fa-toggle-left")
		}
		console.log("List options button clicked")
		mainLayout.toggle("east")
	})
	
	
	
	function hideFooterLayout() {
		contentLayout.close("south")
	}
	
	function showFooterLayout() {
		contentLayout.open("south")
	}
	
	function closePropertyPanel() {
		contentLayout.close("east")
	}
	
	function openPropertyPanel() {
		contentLayout.open("east")
	}
	
	function clearCenterContentArea() {
		$('#contentLayoutContainer').find("div").empty()
	}
	
	this.hideFooterLayout = hideFooterLayout
	this.showFooterLayout = showFooterLayout
	
	this.closePropertyPanel = closePropertyPanel
	this.openPropertyPanel = openPropertyPanel	
	this.clearCenterContentArea = clearCenterContentArea
	
}