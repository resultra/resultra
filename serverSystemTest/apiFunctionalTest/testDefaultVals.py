#!/usr/bin/env python
#
# This file is part of the Resultra project.
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.


import unittest
import json
import datetime
import time

from testCommon import TestHelperMixin


class TestRecordWithDefaultVals(unittest.TestCase,TestHelperMixin):
    def setUp(self):
        self.createTestSession()
        self.databaseID = self.newDatabase('TestRecordWithChangeSet: Test Database')
        self.boolFieldID = self.newBoolField(self.databaseID,"Bool Field","BoolField")
        self.textFieldID = self.newTextField(self.databaseID,"Text Field","TextField")
        
    def testDefaultValWithChangeSet(self):
        recordID = self.newRecord(self.databaseID)
        
        initialBoolVal = True
        
        recordRef = self.setBoolRecordValue(self.databaseID,recordID,self.boolFieldID,initialBoolVal)
                
        self.assertEquals(self.getRecordFieldVal(recordRef,self.boolFieldID),
            initialBoolVal,"record after being initially set has same value")
            
        initialTextVal = "Hello World!"
        recordRefAfterTextSet = self.setTextRecordValue(self.databaseID,recordID,self.textFieldID,initialTextVal)

        self.assertEquals(self.getRecordFieldVal(recordRefAfterTextSet,self.boolFieldID),
             initialBoolVal,"record after being initially set has same value")
        self.assertEquals(self.getRecordFieldVal(recordRefAfterTextSet,self.textFieldID),
              initialTextVal,"record after being initially set has same value")
              
        # Set the default values under a change set ID. This is how it's done with modal dialogs.
        jsonResp = self.apiRequest('record/allocateChangeSetID',{})
        changeSetID = jsonResp[u'changeSetID']
        if len(changeSetID) <= 0:
            self.fail("ERROR getting change set ID - expecting length > 0")
        print "Change set ID: ",changeSetID
        
              
        defaultVals = [{'fieldID':self.boolFieldID,'defaultValueID':"false"}]
        defaultValParams = {'parentDatabaseID':self.databaseID,'recordID':recordID,'changeSetID':changeSetID,
            'defaultVals':defaultVals}
        recordRefAfterDefaultVals = self.apiRequest('recordUpdate/setDefaultValues',defaultValParams)

        self.assertEquals(self.getRecordFieldVal(recordRefAfterDefaultVals,self.boolFieldID),
            False,"record after default value applied")
        self.assertEquals(self.getRecordFieldVal(recordRefAfterDefaultVals,self.textFieldID),
            initialTextVal,"text field is untouched by default values")
                    


# Allow the tests in this file to be run stand-alone
if __name__ == '__main__':
    unittest.main()