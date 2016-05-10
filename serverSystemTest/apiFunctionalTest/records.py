#!/usr/bin/env python

import unittest
import json
import datetime

from testCommon import TestHelperMixin

class TestTimeRecordValues(unittest.TestCase,TestHelperMixin):
    def setUp(self):
        databaseID = self.newDatabase('TestTimeRecordValues: Test Database')
        self.tableID = self.newTable(databaseID,"TestTimeRecordValues: Test Table")
        self.timeFieldID = self.newTimeField(self.tableID,"TestTimeRecordValues - Time Field","TimeField")
        
    def getRecordFieldVal(self,recordRef,fieldID):
        fieldValues = recordRef[u'fieldValues']
        value = fieldValues[fieldID]
        return value
        
    
    def testSimpleDates(self):
        recordID = self.newRecord(self.tableID)        
        recordRef = self.getRecord(recordID)
        
        timeVal = "2016-10-12T00:00:00Z" # RFC 3339 date & time format with Z at end for UTC
        
        recordRef = self.setTimeRecordValue(recordID,self.timeFieldID,timeVal)
        # Round-trip comparison on the value set in the record.
        self.assertEquals(self.getRecordFieldVal(recordRef,self.timeFieldID),timeVal, "updated record has time value")       
          
        # Get the record straight from the database
        recordRef = self.getRecord(recordID)
        self.assertEquals(self.getRecordFieldVal(recordRef,self.timeFieldID),timeVal,"retrieved record has time value")
        
        with self.assertRaises(AssertionError):
            self.setTimeRecordValue(recordID,self.timeFieldID,"ABC") # Invalid time format
            
        with self.assertRaises(AssertionError):
            self.setTimeRecordValue(recordID,self.timeFieldID,"") # Invalid time format
              
        
# TODO - Verify a new record is not created with an invalid table ID
# TODO - Verify the field ID given for the record update has the same field ID as the record.
# TODO - If 'value' parameter is ommitted altogether, an error should be generated. For time values, this is currently initializing to a default time.

# Allow the tests in this file to be run stand-alone
if __name__ == '__main__':
    unittest.main()