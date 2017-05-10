function initAdminSettingsPageLayout($pageContainer) {
	
	var zeroPaddingInset = { top:0, bottom:0, left:0, right:0 }
	
	$pageContainer.layout({
			inset: zeroPaddingInset,
			north: fixedUILayoutPaneParams(40),
			west: {
				size: 250,
				resizable:false,
				slidable: false,
				spacing_open:4,
				spacing_closed:4,
				initClosed:false // panel is initially open	
			}
		})
	
}