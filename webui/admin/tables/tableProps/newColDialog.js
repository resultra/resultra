// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function openNewTableColDialog(pageContext,tableRef) {
		
	var panelConfig = {
		databaseID: tableRef.parentDatabaseID,
		tableID: tableRef.tableID
	}
	
	
	var newOrExistingPanel = createNewTableColNewOrExistingDialogPanelConfig(panelConfig)
	var newFieldPanel = createNewTableColNewFieldDialogPanelConfig(panelConfig)
	var colTypePanel = createNewTableColColTypeDialogPanelConfig(pageContext,tableRef,panelConfig)
		
	openWizardDialog({
		closeFunc: function() {},
		dialogDivID: '#newTableColDialog',
		panels: [newOrExistingPanel,newFieldPanel,colTypePanel],
		progressDivID: '#tableProps_WizardDialogProgress',
	})
	
}