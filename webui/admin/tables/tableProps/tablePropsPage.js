$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#tablePropsAdminPage'))	
	initAdminPageHeader()
	initAdminSettingsTOC(tablePropsAdminContext.databaseID,"settingsTOCTables")
	
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
	
	initFieldInfo(tablePropsAdminContext.databaseID, function() {
		var getTableParams = { tableID: tablePropsAdminContext.tableID }
		jsonAPIRequest("tableView/get",getTableParams,function(tableRef) {
			initNameProperties(tableRef)
			initTableViewColsProperties(tableRef)
		})
		
	})
	
})