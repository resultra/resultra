function getGlobalInfoIndexedByID(databaseID, globalInfoIndexCallback) {
	
	var listParams =  { parentDatabaseID: databaseID }
	var globalInfoIndex = {}
	jsonAPIRequest("global/getList",listParams,function(globalsInfo) {
		$.each(globalsInfo, function(index, globalInfo) {
			globalInfoIndex[globalInfo.globalID] = globalInfo	
		})
		globalInfoIndexCallback(globalInfoIndex)
	})
}