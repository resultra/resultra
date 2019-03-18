// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function DashboardContentLayout() {
	
	var $dashboardLayout = $('#dashboardContentLayout')
	
	var contentLayout = $dashboardLayout.layout({
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
		var $header = $('#dashboardCenterContentHeader')
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