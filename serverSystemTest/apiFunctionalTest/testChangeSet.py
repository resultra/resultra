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


class TestRecordWithChangeSet(unittest.TestCase,TestHelperMixin):
    def setUp(self):
        self.createTestSession()
        self.databaseID = self.newDatabase('TestRecordWithChangeSet: Test Database')
        self.longTextFieldID = self.newLongTextField(self.databaseID,"TestLongTextRecordValues - Long Text Field","TimeField")
        
    def testLongText(self):
        recordID = self.newRecord(self.databaseID)
        someText = "Hello World!"
        
        recordRef = self.setLongTextRecordValue(self.databaseID,recordID,self.longTextFieldID,someText)
        
        jsonResp = self.apiRequest('record/allocateChangeSetID',{})
        changeSetID = jsonResp[u'changeSetID']
        if len(changeSetID) <= 0:
            self.fail("ERROR getting change set ID - expecting length > 0")
        print "Change set ID: ",changeSetID
        
        self.assertEquals(self.getRecordFieldVal(recordRef,self.longTextFieldID),
            someText,"record after being initially set has same value")
            
        # Set the value again with the change set ID.
        valueUnderChangeSet = "Hello World! (with changes)"
        recordRefUnderChangeSet = self.setLongTextRecordValueWithChangeSet(self.databaseID,recordID,
                            self.longTextFieldID,changeSetID,valueUnderChangeSet)
        self.assertEquals(self.getRecordFieldVal(recordRefUnderChangeSet,self.longTextFieldID),
            valueUnderChangeSet,"record returned under change set has value set with change set ID")
 
        recordRefAfterChange = self.getRecord(self.databaseID,recordID)
        self.assertEquals(self.getRecordFieldVal(recordRefAfterChange,self.longTextFieldID),
            someText,"retrieving the record after another an uncommitted value has been set should return value set without change set ID ")
  
        commitChangeSetParams = {'recordID':recordID,'changeSetID':changeSetID}
        jsonResp = self.apiRequest('recordUpdate/commitChangeSet',commitChangeSetParams)

        recordRefAfterCommittingChange = self.getRecord(self.databaseID,recordID)
        self.assertEquals(self.getRecordFieldVal(recordRefAfterCommittingChange,self.longTextFieldID),
            valueUnderChangeSet,"retrieving the record after committing the change set should include the changes")
        


# Allow the tests in this file to be run stand-alone
if __name__ == '__main__':
    unittest.main()