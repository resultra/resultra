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
			initClosed:true,
			fxName: "none"
		}
	})

		
	function closeRHSSidebar() {
		mainLayout.close("east")
	}
	
	function openRHSSidebar() {
		mainLayout.open("east")
	}

	function hideRHSSidebar() {
		mainLayout.hide("east")
	}
	
	function disableRHSSidebar() {
		mainLayout.hide("east")
		clearMainWindowRHSSidebarContent()
	}

	function showRHSSidebar() {
		mainLayout.show("east",false)
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

	
	function rhsSidebarIsOpen() {
		return (!mainLayout.state.east.isClosed)
	}

		
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
	this.rhsSidebarIsOpen = rhsSidebarIsOpen
		
}