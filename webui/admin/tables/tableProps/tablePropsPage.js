// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initTablePropsAdminSettingsPageContent(pageContext,tableInfo) {
		
	
	function initNameProperties(tableRef) {

		var $tableNameForm = $('#tableNamePropertyForm')
		var $nameInput =$tableNameForm.find('input[name=tableNameInput]')
		
		$nameInput.blur() // prevent auto-focus
		
		$nameInput.val(tableRef.name)


		var remoteValidationParams = {
			url: '/api/tableView/validateTableName',
			data: {
				tableID: function() { return tableRef.tableID },
				tableName: function() { return $nameInput.val() }
			}
		}

		var validationSettings = createInlineFormValidationSettings({
			rules: {
				tableNameInput: {
					minlength: 3,
					required: true,
					remote: remoteValidationParams
				} // newRoleNameInput
			}
		})


		var validator = $tableNameForm.validate(validationSettings)

		initInlineInputControlValidationOnBlur(validator,$nameInput,
			remoteValidationParams, function(validatedName) {
				var setNameParams = {
					tableID:tableRef.tableID,
					newTableName:validatedName
				}
				jsonAPIRequest("tableView/setName",setNameParams,function(tableInfo) {
					console.log("Done changing table name: " + validatedName)
				})
		})

		validator.resetForm()

	} // initItemListNameProperties
	
	initFieldInfo(tableInfo.parentDatabaseID, function() {
		var getTableParams = { tableID: tableInfo.tableID }
		jsonAPIRequest("tableView/get",getTableParams,function(tableRef) {
			initNameProperties(tableRef)
			initTableViewColsProperties(pageContext,tableRef)
		})
		
	})
	
	
	initSettingsPageButtonLink('#tablePropsBackToTableListLink',"tables")
	
}
