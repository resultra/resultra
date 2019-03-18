// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function AlertListContentLayout() {
	
	var $contentLayout = $('#alertListContentLayout')
	
	var contentLayout = $contentLayout.layout({
		onresize_end: function(pane, $pane, paneState, paneOptions) {
			if(pane === 'center'){
				// only propagate the resize event for the center/content pane
				console.log("resize triggered")
			//	resizeMainWindow()
			}
		},
		north: fixedUILayoutPaneAutoSizeToFitContentsParams(),
		north__showOverflowOnHover:	true
	})
		
		
	function setCenterContentHeader(header) {
		var $header = $('#alertListCenterContentHeader')
		$header.text(header)
	//	var $breadcrumb = $('#trackerLocationBreadcrumb')
	//	$breadcrumb.text(header)
	}
	
	initButtonClickHandler("#viewTableOfContentsMenuButton", function() {
		theMainWindowLayout.toggleLHSSidebar()
	})
	
	
	this.setCenterContentHeader = setCenterContentHeader
	
}