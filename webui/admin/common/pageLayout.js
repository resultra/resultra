function initAdminSettingsPageLayout($pageContainer) {
	
	var zeroPaddingInset = { top:0, bottom:0, left:0, right:0 }
	
	$pageContainer.layout({
			inset: zeroPaddingInset,
			north: fixedUILayoutPaneParams(40),
			west: fixedUILayoutPaneParams(250)
		})
	
}