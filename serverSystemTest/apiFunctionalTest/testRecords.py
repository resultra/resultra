#!/usr/bin/env python

import unittest
import json
import datetime
import time

from testCommon import TestHelperMixin

class TestTimeRecordValues(unittest.TestCase,TestHelperMixin):
    
    def setUp(self):
        self.createTestSession()
        self.databaseID = self.newDatabase('TestTimeRecordValues: Test Database')
        self.timeFieldID = self.newTimeField(self.databaseID,"TestTimeRecordValues - Time Field","TimeField")
        
        fieldParams = {'parentDatabaseID':self.databaseID,'name':'Future','type':'time',
              'refName':'calcTest','formulaText':'DATEADD([TimeField],1)'}
        jsonResp = self.apiRequest('calcField/new',fieldParams)
        self.calcFieldID = jsonResp[u'fieldID']
  
        fieldParams = {'parentDatabaseID':self.databaseID,'name':'Past','type':'time',
            'refName':'pastCalcTest','formulaText':'DATEADD([TimeField],-2)'}
        jsonResp = self.apiRequest('calcField/new',fieldParams)
        self.pastCalcFieldID = jsonResp[u'fieldID']
            
    
    def testSimpleDates(self):
        recordID = self.newRecord(self.databaseID)
                         
        recordRef = self.getRecord(self.databaseID,recordID)
        
        timeVal = "2016-10-12T00:00:00Z" # RFC 3339 date & time format with Z at end for UTC
        calcTimeVal = "2016-10-13T00:00:00Z" # Date is offset by 1 day
        pastCalcTimeVal = "2016-10-10T00:00:00Z"
        
        recordRef = self.setTimeRecordValue(self.databaseID,recordID,self.timeFieldID,timeVal)
        # Round-trip comparison on the value set in the record.
        self.assertEquals(self.getRecordFieldVal(recordRef,self.timeFieldID),timeVal, "updated record has time value")       
          
        # Get the record straight from the database
        recordRef = self.getRecord(self.databaseID,recordID)
        self.assertEquals(self.getRecordFieldVal(recordRef,self.timeFieldID),timeVal,"retrieved record has time value")
        self.assertEquals(self.getRecordFieldVal(recordRef,self.calcFieldID),
                    calcTimeVal,"retrieved record has calculated time value")
        self.assertEquals(self.getRecordFieldVal(recordRef,self.pastCalcFieldID),
                    pastCalcTimeVal,"retrieved record has calculated time value")
        
        with self.assertRaises(AssertionError):
            self.setTimeRecordValue(self.databaseID,recordID,self.timeFieldID,"ABC") # Invalid time format
            
        with self.assertRaises(AssertionError):
            self.setTimeRecordValue(self.databaseID,recordID,self.timeFieldID,"") # Invalid time format



class TestNumberRecordValues(unittest.TestCase,TestHelperMixin):
    
    def setUp(self):
        self.createTestSession()
        self.databaseID = self.newDatabase('TestNumberRecordValues: Test Database')
        self.numberFieldID = self.newNumberField(self.databaseID,"TestNumberRecordValues - Number Field","numField") 
        
        fieldParams = {'parentDatabaseID':self.databaseID,'name':'Total','type':'number',
                    'refName':'calcTest','formulaText':'[numField]*2'}
        jsonResp = self.apiRequest('calcField/new',fieldParams)
        self.calcFieldID = jsonResp[u'fieldID']
           
    
    def testSimpleNumbers(self):
        recordID = self.newRecord(self.databaseID)
                         
        recordRef = self.getRecord(self.databaseID,recordID) 
        
        numberVal = 25.2
        calcFieldExpectedVal = numberVal * 2
 
        self.setNumberRecordValue(self.databaseID,recordID,self.numberFieldID,numberVal)
 
        # Get the record straight from the database
        recordRef = self.getRecord(self.databaseID,recordID)
        self.assertEquals(self.getRecordFieldVal(recordRef,self.numberFieldID),numberVal,"retrieved record has number value")
        self.assertEquals(self.getRecordFieldVal(recordRef,
                    self.calcFieldID),calcFieldExpectedVal,"calculated field has expected value")
        
        # Test what happens when clearing the value. This is done by setting the value to 'null'
        self.setNumberRecordValue(self.databaseID,recordID,self.numberFieldID,None)
        recordRef = self.getRecord(self.databaseID,recordID)
        self.assertEquals(self.getRecordFieldVal(recordRef,self.numberFieldID),None,"retrieved record has null value")
        self.verifyUndefinedFieldVal(recordRef,self.calcFieldID)
                      

class TestLongTextRecordValues(unittest.TestCase,TestHelperMixin):
    def setUp(self):
        self.createTestSession()
        self.databaseID = self.newDatabase('TestLongTextRecordValues: Test Database')
        self.longTextFieldID = self.newLongTextField(self.databaseID,"TestLongTextRecordValues - Long Text Field","TimeField")
        
    def testLongText(self):
        recordID = self.newRecord(self.databaseID)
        someText = "Hello World!"
        
        recordRef = self.setLongTextRecordValue(self.databaseID,recordID,self.longTextFieldID,someText)
        
        self.assertEquals(self.getRecordFieldVal(recordRef,self.longTextFieldID),
            someText,"record after being initially set has same value")
          

# TODO - Set record values with invalid types - e.g. set a bool for a text field
# TODO - Verify a new record is not created with an invalid table ID
# TODO - Verify the field ID given for the record update has the same field ID as the record.
# TODO - If 'value' parameter is ommitted altogether, an error should be generated. For time values, this is currently initializing to a default time.

# Allow the tests in this file to be run stand-alone
if __name__ == '__main__':
    unittest.main()