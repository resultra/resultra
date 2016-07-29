import requests
import json
import time

class TestHelperMixin:
        
    def initSession(self):
        self.session = requests.session()
    
    def authRequest(self, apiPath,jsonArgs):
        baseURL = 'http://localhost:8080/auth/'
        fullURL = baseURL + apiPath
        print "TestHelperMixin: Request: ",apiPath,": args=",json.dumps(jsonArgs)
        resp = self.session.post(fullURL,json=jsonArgs)
        print resp.cookies
        if resp.status_code != 200:
            print "TestHelperMixin: Error response: ",resp.text
        self.assertEqual(resp.status_code,200,"expecting success return code from server")
        print "TestHelperMixin: Response: ",json.dumps(resp.json())
        return resp.json()
    
    
    def apiRequest(self, apiPath,jsonArgs):
        baseURL = 'http://localhost:8080/api/'
        fullURL = baseURL + apiPath
        print "TestHelperMixin: API Request: ",apiPath,": args=",json.dumps(jsonArgs)
        resp = self.session.post(fullURL,json=jsonArgs)
        if resp.status_code != 200:
            print "TestHelperMixin: Error response: ",resp.text
        self.assertEqual(resp.status_code,200,"expecting success return code from server")
        print "TestHelperMixin: API Response: ",json.dumps(resp.json())
        return resp.json()
        
    def newTestUser(self):
        # Create a new user for the test using the timestamp as the userID
        newUserID = "u" +  str(int(time.time() * 1000))
        newUserEmail = newUserID + "@example.com"
        newUserParams = {'emailAddr':newUserEmail,'password':'testpw123$',
            'firstName':'John','lastName':"Smith",
            'userName':newUserID}
        jsonResp = self.authRequest('register',newUserParams)
        self.assertEquals(jsonResp['success'],True,"new test user")
        return newUserID
        
    def signinTestUser(self,userID):
        signinParams = {'emailAddr':userID+'@example.com','password':'testpw123$',
            'firstName':'John','lastName':"Smith",
            'userName':userID}
        jsonResp = self.authRequest('login',signinParams)
        self.assertEquals(jsonResp['success'],True,"successful login")
        
    # Initialize the session, create a test user and sign that user in for the session
    def createTestSession(self):
        self.initSession()
        userID = self.newTestUser()
        self.signinTestUser(userID)      

    def newDatabase(self,databaseName):
        jsonResp = self.apiRequest('database/new',{'name': databaseName})
        databaseID = jsonResp[u'databaseID']
        return databaseID
        
    def newTable(self,databaseID,tableName):
        jsonResp = self.apiRequest('table/new',{'databaseID': databaseID, 'name': tableName})
        tableID = jsonResp[u'tableID']
        return tableID
        
    def newTextField(self,tableID,fieldName,refName):
        fieldParams = {'parentTableID':tableID,'name':fieldName,'type':'text','refName':refName}
        jsonResp = self.apiRequest('field/new',fieldParams)
        fieldID = jsonResp[u'fieldID']
        return fieldID
            
    def newNumberField(self,tableID,fieldName,refName):
        fieldParams = {'parentTableID':tableID,'name':fieldName,'type':'number','refName':refName}
        jsonResp = self.apiRequest('field/new',fieldParams)
        fieldID = jsonResp[u'fieldID']
        return fieldID
        
        
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
        jsonResp = self.apiRequest('recordUpdate/newRecord',{'parentTableID':tableID})
        recordID = jsonResp[u'recordID']
        return recordID
    
    def getRecord(self,parentTableID,recordID):
        recordRef = self.apiRequest('recordValue/getRecordValueResults',{'parentTableID':parentTableID,'recordID':recordID})
        return recordRef
        
    def getRecordFieldVal(self,recordRef,fieldID):
        fieldValues = recordRef[u'fieldValues']
        value = fieldValues[fieldID]
        return value   
           
    def setNumberRecordValue(self,parentTableID,recordID,fieldID,numberVal):
        recordRef = self.apiRequest('recordUpdate/setNumberFieldValue',{'parentTableID':parentTableID,'recordID':recordID,'fieldID':fieldID,'value':numberVal})
        return recordRef

    def setTextRecordValue(self,parentTableID,recordID,fieldID,textVal):
        recordRef = self.apiRequest('recordUpdate/setTextFieldValue',{'parentTableID':parentTableID,'recordID':recordID,'fieldID':fieldID,'value':textVal})
        return recordRef
        
    def setTimeRecordValue(self,parentTableID, recordID,fieldID,timeVal):
        recordRef = self.apiRequest('recordUpdate/setTimeFieldValue',{'parentTableID':parentTableID,'recordID':recordID,'fieldID':fieldID,'value':timeVal})
        return recordRef
 
    def setLongTextRecordValue(self,parentTableID,recordID,fieldID,textVal):
        recordRef = self.apiRequest('recordUpdate/setLongTextFieldValue',{'parentTableID':parentTableID,'recordID':recordID,'fieldID':fieldID,'value':textVal})
        return recordRef
    