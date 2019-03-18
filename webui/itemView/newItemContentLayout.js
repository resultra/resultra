// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function NewItemContentLayout() {
	
	var contentLayout = $('#newItemContentLayout').layout({
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
			initClosed:false // panel is initially closed	
		},
		north__showOverflowOnHover:	true,
		south__showOverflowOnHover:	true
	})
		
		
	function setCenterContentHeader(header) {
		var $header = $('#newItemCenterContentHeader')
		$header.text(header)
	//	var $breadcrumb = $('#trackerLocationBreadcrumb')
	//	$breadcrumb.text(header)
	}
	
	initButtonClickHandler("#viewTableOfContentsMenuButton", function() {
		console.log("TOC button clicked")
		theMainWindowLayout.toggleLHSSidebar()
	})
	
	
	this.setCenterContentHeader = setCenterContentHeader
	
}