#!/usr/bin/env python

import unittest
import json
import datetime

from testCommon import TestHelperMixin

class TestDashboard(unittest.TestCase,TestHelperMixin):
    
    def setUp(self):
        self.databaseID = self.newDatabase('TestRecordFilter: Test Database')
        self.tableID = self.newTable(self.databaseID,"TestRecordFilter: Test Table")
        self.timeFieldID = self.newTimeField(self.tableID,"TestRecordFilter - Time Field","TimeField")
        self.numberFieldID = self.newNumberField(self.tableID,"TestRecordFilter - Number Field","NumberField") 
        self.textFieldID = self.newTextField(self.tableID,"TestRecordFilter - Text Field","TextField") 
    
    def testSimpleDashboard(self):      
        dashboardParams = {'databaseID':self.databaseID,'name':'My Dashboard'}
        jsonResp = self.apiRequest('dashboard/new',dashboardParams)
        dashboardID = jsonResp[u'dashboardID']
        
        with self.assertRaises(AssertionError):
            # Invalid Parent Database ID - passes table ID instead of database ID
            dashboardParams = {'databaseID':self.tableID,'name':'My Dashboard'}
            jsonResp = self.apiRequest('dashboard/new',dashboardParams)
        


# Allow the tests in this file to be run stand-alone
if __name__ == '__main__':
    unittest.main()