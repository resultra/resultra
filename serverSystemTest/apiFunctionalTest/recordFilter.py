#!/usr/bin/env python

import unittest
import json
import datetime

from testCommon import TestHelperMixin

class TestRecordFilter(unittest.TestCase,TestHelperMixin):
    
    def setUp(self):
        databaseID = self.newDatabase('TestRecordFilter: Test Database')
        self.tableID = self.newTable(databaseID,"TestRecordFilter: Test Table")
        self.timeFieldID = self.newTimeField(self.tableID,"TestRecordFilter - Time Field","TimeField")
        self.numberFieldID = self.newNumberField(self.tableID,"TestRecordFilter - Number Field","NumberField") 
        self.textFieldID = self.newTextField(self.tableID,"TestRecordFilter - Text Field","TextField") 
    
    def testSimpleFilter(self):        
        filterParams = {'parentTableID':self.tableID,'name':'Simple Filter'}
        jsonResp = self.apiRequest('filter/new',filterParams)
        filterID = jsonResp[u'filterID']
        
        with self.assertRaises(AssertionError):
            # Inavlid Parent Table ID
            filterParams = {'parentTableID':self.timeFieldID,'name':'Simple Filter'}
            jsonResp = self.apiRequest('filter/new',filterParams)
 
        with self.assertRaises(AssertionError):
            # Empty filter name
            filterParams = {'parentTableID':self.tableID,'name':''}
            jsonResp = self.apiRequest('filter/new',filterParams)
 

    def testDuplicateFilterName(self):        
        filterParams = {'parentTableID':self.tableID,'name':'Simple Filter'}
        self.apiRequest('filter/new',filterParams)
        
        print "testDuplicateFilterName: test with different name: should be OK"
        params = {'parentTableID':self.tableID,'name':'My filter'}
        self.apiRequest('filter/new',params)
                 
        with self.assertRaises(AssertionError):
            print "testDuplicateFilterName: duplicate name: should fail"
            self.apiRequest('filter/new',filterParams)



# Allow the tests in this file to be run stand-alone
if __name__ == '__main__':
    unittest.main()