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
        
    def signoutCurrUser(self):
        signoutParams = {}
        jsonResp = self.authRequest('signout',signoutParams)
        self.assertEquals(jsonResp['success'],True,"successful signout")
        
    # Initialize the session, create a test user and sign that user in for the session
    def createTestSession(self):
        self.initSession()
        userID = self.newTestUser()
        self.signinTestUser(userID)      

    def newDatabase(self,databaseName):
        jsonResp = self.apiRequest('database/new',{'name': databaseName})
        databaseID = jsonResp[u'databaseID']
        return databaseID
                
    def newTextField(self,databaseID,fieldName,refName):
        fieldParams = {'parentDatabaseID':databaseID,'name':fieldName,'type':'text','refName':refName}
        jsonResp = self.apiRequest('field/new',fieldParams)
        fieldID = jsonResp[u'fieldID']
        return fieldID
            
    def newNumberField(self,databaseID,fieldName,refName):
        fieldParams = {'parentDatabaseID':databaseID,'name':fieldName,'type':'number','refName':refName}
        jsonResp = self.apiRequest('field/new',fieldParams)
        fieldID = jsonResp[u'fieldID']
        return fieldID
        
        
    def newTimeField(self,databaseID,fieldName,refName):
        fieldParams = {'parentDatabaseID':databaseID,'name':fieldName,'type':'time','refName':refName}
        jsonResp = self.apiRequest('field/new',fieldParams)
        fieldID = jsonResp[u'fieldID']
        return fieldID

    def newBoolField(self,databaseID,fieldName,refName):
        fieldParams = {'parentDatabaseID':databaseID,'name':fieldName,'type':'bool','refName':refName}
        jsonResp = self.apiRequest('field/new',fieldParams)
        fieldID = jsonResp[u'fieldID']
        return fieldID


    def newLongTextField(self,databaseID,fieldName,refName):
        fieldParams = {'parentDatabaseID':databaseID,'name':fieldName,'type':'longText','refName':refName}
        jsonResp = self.apiRequest('field/new',fieldParams)
        fieldID = jsonResp[u'fieldID']
        return fieldID
 
    def newFileField(self,databaseID,fieldName,refName):
        fieldParams = {'parentDatabaseID':databaseID,'name':fieldName,'type':'file','refName':refName}
        jsonResp = self.apiRequest('field/new',fieldParams)
        fieldID = jsonResp[u'fieldID']
        return fieldID
 
    
    def newRecord(self,tableID):
        jsonResp = self.apiRequest('recordUpdate/newRecord',{'parentDatabaseID':tableID})
        recordID = jsonResp[u'recordID']
        return recordID
    
    def getRecord(self,parentDatabaseID,recordID):
        recordRef = self.apiRequest('recordRead/getRecordValueResults',
                {'parentDatabaseID':parentDatabaseID,'recordID':recordID})
        return recordRef
        
    def getRecordFieldVal(self,recordRef,fieldID):
        fieldValues = recordRef[u'fieldValues']
        value = fieldValues[fieldID]
        return value
        
    def verifyUndefinedFieldVal(self,recordRef,fieldID):
        fieldValues = recordRef[u'fieldValues']
        if fieldID in fieldValues:
            print "FAIL: verifyOneFormula: field ID should be undefined: ",fieldID
            self.fail(msg="verifyUndefinedFieldVal")
        else:
            print "SUCCESS: verifyOneFormula: field ID not undefined: ",fieldID
           
    def setNumberRecordValue(self,parentDatabaseID,recordID,fieldID,numberVal):
        recordRef = self.apiRequest('recordUpdate/setNumberFieldValue',{'parentDatabaseID':parentDatabaseID,'recordID':recordID,'fieldID':fieldID,'value':numberVal})
        return recordRef

    def setBoolRecordValue(self,parentDatabaseID,recordID,fieldID,boolVal):
        recordRef = self.apiRequest('recordUpdate/setBoolFieldValue',{'parentDatabaseID':parentDatabaseID,'recordID':recordID,'fieldID':fieldID,'value':boolVal})
        return recordRef

    def setTextRecordValue(self,parentDatabaseID,recordID,fieldID,textVal):
        recordRef = self.apiRequest('recordUpdate/setTextFieldValue',{'parentDatabaseID':parentDatabaseID,'recordID':recordID,'fieldID':fieldID,'value':textVal})
        return recordRef
        
    def setTimeRecordValue(self,parentDatabaseID, recordID,fieldID,timeVal):
        recordRef = self.apiRequest('recordUpdate/setTimeFieldValue',{'parentDatabaseID':parentDatabaseID,'recordID':recordID,'fieldID':fieldID,'value':timeVal})
        return recordRef
 
    def setLongTextRecordValue(self,parentDatabaseID,recordID,fieldID,textVal):
        recordRef = self.apiRequest('recordUpdate/setLongTextFieldValue',{'parentDatabaseID':parentDatabaseID,'recordID':recordID,'fieldID':fieldID,'value':textVal})
        return recordRef

    def setLongTextRecordValueWithChangeSet(self,parentDatabaseID,recordID,fieldID,changeSetID,textVal):
        setValParams = {'parentDatabaseID':parentDatabaseID,'recordID':recordID,'fieldID':fieldID,'changeSetID':changeSetID,'value':textVal}
        recordRef = self.apiRequest('recordUpdate/setLongTextFieldValue',setValParams)
        return recordRef
    