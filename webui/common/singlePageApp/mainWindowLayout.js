function MainWindowLayout()
{
	var zeroPaddingInset = { top:0, bottom:0, left:0, right:0 }
	
	
	var resizeMainWindowPanesEventName = "resize-main-window-panes"
	function resizeMainWindow() {
		console.log("Resizing main window")
		$(window).trigger(resizeMainWindowPanesEventName)
	}
	
	// Initialize the page layout
	var mainLayout = $('#layoutPage').layout({
		inset: zeroPaddingInset,
		north: fixedUILayoutPaneParams(40),
		onopen_end: function(pane, $pane, paneState, paneOptions) {
			resizeMainWindow()				
		},
		onclose_end: function(pane, $pane, paneState, paneOptions) {
			resizeMainWindow()
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
			initClosed:false // TOC sidebar initially open	
		}
	})

	
	var contentLayout = $('#mainWindowContentPane').layout({
		onresize_end: function(pane, $pane, paneState, paneOptions) {
			if(pane === 'center'){
				// only propagate the resize event for the center/content pane
				console.log("resize triggered")
				resizeMainWindow()
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
			
	initButtonClickHandler("#viewTableOfContentsMenuButton", function() {
		console.log("TOC button clicked")
		mainLayout.toggle("west")
	})
	
	
	
	
	var $viewListOptionsToggleButton = $("#viewListOptionsButton")
	var $iconSpan = $viewListOptionsToggleButton.find("span")
	
	initButtonControlClickHandler($viewListOptionsToggleButton, function() {
		var layoutState = mainLayout.state
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
	
	function disablePropertyPanelToggleButton() {
		$viewListOptionsToggleButton.prop("disabled",true)
	}
	function enablePropertyPanelToggleButton() {
		$viewListOptionsToggleButton.prop("disabled",false)
	}
	
	
	
	function hideFooterLayout() {
		contentLayout.close("south")
	}
	
	function showFooterLayout() {
		contentLayout.open("south")
	}
	
	var $refreshButton = $('#mainWindowViewRefreshButton')
	
	function disableRefreshButton() {
		$refreshButton.hide()
	}
	
	function enableRefreshButton(refreshCallback) {
		initButtonControlClickHandler($refreshButton,refreshCallback)
		$refreshButton.show()
	}
	
	
	function closePropertyPanel() {
		$iconSpan.removeClass("fa-toggle-right")
		$iconSpan.addClass("fa-toggle-left")
		mainLayout.close("east")
	}
	
	function openPropertyPanel() {
		$iconSpan.removeClass("fa-toggle-left")
		$iconSpan.addClass("fa-toggle-right")
		mainLayout.open("east")
	}
	
	
	function clearCenterContentArea() {
		// Clear any event handlers which are attached to the layout-specific events.
		$(window).off(resizeMainWindowPanesEventName)
		$('#contentLayoutContainer').find(".clearableViewContent").empty()
	}
	
	function clearSidebarNavigationSelection() {
		$('#viewFormTocLayout').find('li').removeClass("active")
	}
	
	function setCenterContentHeader(header) {
		var $header = $('#mainWindowCenterContentHeader')
		$header.text(header)
		var $breadcrumb = $('#trackerLocationBreadcrumb')
		$breadcrumb.text(header)
	}
	
	function disablePropertySidebar() {
		mainLayout.hide("east")
		$viewListOptionsToggleButton.hide()
	}

	function enablePropertySidebar() {
		mainLayout.show("east",false)
		$viewListOptionsToggleButton.show()
	}

	
	this.hideFooterLayout = hideFooterLayout
	this.showFooterLayout = showFooterLayout
	
	this.closePropertyPanel = closePropertyPanel
	this.openPropertyPanel = openPropertyPanel	
	this.clearCenterContentArea = clearCenterContentArea
	this.setCenterContentHeader = setCenterContentHeader
	this.enablePropertySidebar = enablePropertySidebar
	this.disablePropertySidebar = disablePropertySidebar
	this.disablePropertyPanelToggleButton = disablePropertyPanelToggleButton
	this.enablePropertyPanelToggleButton = enablePropertyPanelToggleButton
	this.enableRefreshButton = enableRefreshButton
	this.disableRefreshButton = disableRefreshButton
	
	this.clearSidebarNavigationSelection = clearSidebarNavigationSelection
	
}