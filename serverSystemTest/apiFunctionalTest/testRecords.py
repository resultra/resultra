#!/usr/bin/env python

import unittest
import json
import datetime
import time

from testCommon import TestHelperMixin

class TestTimeRecordValues(unittest.TestCase,TestHelperMixin):
    
    def setUp(self):
        databaseID = self.newDatabase('TestTimeRecordValues: Test Database')
        self.tableID = self.newTable(databaseID,"TestTimeRecordValues: Test Table")
        self.timeFieldID = self.newTimeField(self.tableID,"TestTimeRecordValues - Time Field","TimeField")    
    
    def testSimpleDates(self):
        recordID = self.newRecord(self.tableID)
                         
        recordRef = self.getRecord(self.tableID,recordID)
        
        timeVal = "2016-10-12T00:00:00Z" # RFC 3339 date & time format with Z at end for UTC
        
        recordRef = self.setTimeRecordValue(self.tableID,recordID,self.timeFieldID,timeVal)
        # Round-trip comparison on the value set in the record.
        self.assertEquals(self.getRecordFieldVal(recordRef,self.timeFieldID),timeVal, "updated record has time value")       
          
        # Get the record straight from the database
        recordRef = self.getRecord(self.tableID,recordID)
        self.assertEquals(self.getRecordFieldVal(recordRef,self.timeFieldID),timeVal,"retrieved record has time value")
        
        with self.assertRaises(AssertionError):
            self.setTimeRecordValue(self.tableID,recordID,self.timeFieldID,"ABC") # Invalid time format
            
        with self.assertRaises(AssertionError):
            self.setTimeRecordValue(self.tableID,recordID,self.timeFieldID,"") # Invalid time format
              

class TestLongTextRecordValues(unittest.TestCase,TestHelperMixin):
    def setUp(self):
        databaseID = self.newDatabase('TestLongTextRecordValues: Test Database')
        self.tableID = self.newTable(databaseID,"TestLongTextRecordValues: Test Table")
        self.longTextFieldID = self.newLongTextField(self.tableID,"TestLongTextRecordValues - Long Text Field","TimeField")
        
    def testLongText(self):
        recordID = self.newRecord(self.tableID)
        someText = "Hello World!"
        
        recordRef = self.setLongTextRecordValue(self.tableID,recordID,self.longTextFieldID,someText)
        
        self.assertEquals(self.getRecordFieldVal(recordRef,self.longTextFieldID),
            someText,"record after being initially set has same value")
          

# TODO - Set record values with invalid types - e.g. set a bool for a text field
# TODO - Verify a new record is not created with an invalid table ID
# TODO - Verify the field ID given for the record update has the same field ID as the record.
# TODO - If 'value' parameter is ommitted altogether, an error should be generated. For time values, this is currently initializing to a default time.

# Allow the tests in this file to be run stand-alone
if __name__ == '__main__':
    unittest.main()