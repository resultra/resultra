import requests
import json

class TestHelperMixin:
    def apiRequest(self, apiPath,jsonArgs):
        baseURL = 'http://localhost:8080/api/'
        fullURL = baseURL + apiPath
        print "TestHelperMixin: API Request: ",apiPath,": args=",json.dumps(jsonArgs)
        resp = requests.post(fullURL,json=jsonArgs)
        self.assertEqual(resp.status_code,200,"expecting success return code from server")
        print "TestHelperMixin: API Response: ",json.dumps(resp.json())
        return resp.json()

    def newDatabase(self,databaseName):
        jsonResp = self.apiRequest('database/new',{'name': databaseName})
        databaseID = jsonResp[u'databaseID']
        return databaseID
        
    def newTable(self,databaseID,tableName):
        jsonResp = self.apiRequest('table/new',{'databaseID': databaseID, 'name': tableName})
        tableID = jsonResp[u'tableID']
        return tableID
        
    def newTimeField(self,tableID,fieldName,refName):
        fieldParams = {'parentTableID':tableID,'name':fieldName,'type':'time','refName':refName}
        jsonResp = self.apiRequest('field/new',fieldParams)
        fieldID = jsonResp[u'fieldID']
        return fieldID

    def newLongTextField(self,tableID,fieldName,refName):
        fieldParams = {'parentTableID':tableID,'name':fieldName,'type':'longText','refName':refName}
        jsonResp = self.apiRequest('field/new',fieldParams)
        fieldID = jsonResp[u'fieldID']
        return fieldID
 
    def newFileField(self,tableID,fieldName,refName):
        fieldParams = {'parentTableID':tableID,'name':fieldName,'type':'file','refName':refName}
        jsonResp = self.apiRequest('field/new',fieldParams)
        fieldID = jsonResp[u'fieldID']
        return fieldID
 
    
    def newRecord(self,tableID):
        jsonResp = self.apiRequest('record/new',{'tableID':tableID})
        recordID = jsonResp[u'recordID']
        return recordID
    
    def getRecord(self,recordID):
        recordRef = self.apiRequest('record/get',{'recordID':recordID})
        return recordRef
        
    def getRecordFieldVal(self,recordRef,fieldID):
        fieldValues = recordRef[u'fieldValues']
        value = fieldValues[fieldID]
        return value      
        
    def setTimeRecordValue(self,recordID,fieldID,timeVal):
        recordRef = self.apiRequest('recordUpdate/setTimeFieldValue',{'recordID':recordID,'fieldID':fieldID,'value':timeVal})
        return recordRef
 
    def setLongTextRecordValue(self,recordID,fieldID,textVal):
        recordRef = self.apiRequest('recordUpdate/setLongTextFieldValue',{'recordID':recordID,'fieldID':fieldID,'value':textVal})
        return recordRef
    