function ItemListLayout(resizeCallback)
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
			spacing_open:16,
			spacing_closed:16,
			togglerClass:			"toggler",
			togglerLength_open:	128,
			togglerLength_closed: 128,
			togglerAlign_closed: "middle",	// align to top of resizer
			togglerAlign_open: "middle"		// align to top of resizer
			
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
	
	var recordLayout = $('#recordsPane').layout({
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
	
	function hideFooterLayout() {
		recordLayout.close("south")
	}
	
	function showFooterLayout() {
		recordLayout.open("south")
	}
	
	this.hideFooterLayout = hideFooterLayout
	this.showFooterLayout = showFooterLayout
	
}
