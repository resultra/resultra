var theMainWindowLayout

function initMainWindowLayout() {
	theMainWindowLayout = new MainWindowLayout()
}

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
			initClosed:true
		}
	})

	
/*			
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
	
	
	
	
	*/
	
	function closeRHSSidebar() {
//		$iconSpan.removeClass("fa-toggle-right")
//		$iconSpan.addClass("fa-toggle-left")
		mainLayout.close("east")
	}
	
	function openRHSSidebar() {
//		$iconSpan.removeClass("fa-toggle-left")
//		$iconSpan.addClass("fa-toggle-right")
		mainLayout.open("east")
	}

	function hideRHSSidebar() {
		mainLayout.hide("east")
//		$viewListOptionsToggleButton.hide()
	}
	
	function disableRHSSidebar() {
		mainLayout.hide("east")
		clearMainWindowRHSSidebarContent()
	}

	function showRHSSidebar() {
		mainLayout.show("east",false)
//		$viewListOptionsToggleButton.show()
	}
	
	function disableLHSSidebar() {
		mainLayout.hide("west")
		clearMainWindowLHSSidebarContent()
	}
	
	function hideLHSSidebar() {
		mainLayout.hide("west")
	}
	function showLHSSidebar() {
		mainLayout.show("west")
	}
	
	function closeLHSSidebar() {
		mainLayout.close("west")
	}
	
	function openLHSSidebar() {
		mainLayout.open("west")
	}
	
	function toggleLHSSidebar() {
		mainLayout.toggle("west")
	}
	
	// Initial configuration of the layout
	hideRHSSidebar()
	hideLHSSidebar()

/*	
	function clearCenterContentArea() {
		// Clear any event handlers which are attached to the layout-specific events.
		$(window).off(resizeMainWindowPanesEventName)
		$('#contentLayoutContainer').find(".clearableViewContent").empty()
	}
	
	function clearSidebarNavigationSelection() {
		$('#viewFormTocLayout').find('li').removeClass("active")
	}
	
	
*/
	

	
//	this.hideFooterLayout = hideFooterLayout
//	this.showFooterLayout = showFooterLayout
	
	this.closeRHSSidebar = closeRHSSidebar
	this.openRHSSidebar = openRHSSidebar	
	this.showRHSSidebar = showRHSSidebar
	this.hideRHSSidebar = hideRHSSidebar
	this.closeLHSSidebar = closeLHSSidebar
	this.openLHSSidebar = openLHSSidebar	
	this.showLHSSidebar = showLHSSidebar
	this.hideLHSSidebar = hideLHSSidebar
	this.toggleLHSSidebar = toggleLHSSidebar
	this.disableLHSSidebar = disableLHSSidebar
	this.disableRHSSidebar = disableRHSSidebar
	
//	this.clearCenterContentArea = clearCenterContentArea
//	this.setCenterContentHeader = setCenterContentHeader
//	this.disablePropertyPanelToggleButton = disablePropertyPanelToggleButton
//	this.enablePropertyPanelToggleButton = enablePropertyPanelToggleButton
//	this.enableRefreshButton = enableRefreshButton
//	this.disableRefreshButton = disableRefreshButton
	
//	this.clearSidebarNavigationSelection = clearSidebarNavigationSelection
	
}