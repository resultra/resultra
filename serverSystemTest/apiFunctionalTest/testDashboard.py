#!/usr/bin/env python

import unittest
import json
import datetime
import time

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
        
        recordID1 = self.newRecord(self.tableID)
        recordID2 = self.newRecord(self.tableID)
                
        self.setTextRecordValue(self.tableID,recordID1,self.textFieldID,"hello world")     
        self.setNumberRecordValue(self.tableID,recordID1,self.numberFieldID,25.2)
        
        self.setTextRecordValue(self.tableID,recordID2,self.textFieldID,"Testing 1,2,3")     
        self.setNumberRecordValue(self.tableID,recordID2,self.numberFieldID,42.5)
        
        
        barChartParams = {'dataSrcTableID':self.tableID,
            'parentDashboardID':dashboardID,
            'xAxisVals': {
                'fieldParentTableID':self.tableID,
                'fieldID': self.numberFieldID,
                'groupValsBy':"none",
                'groupByValBucketWidth':0
            },
            'xAxisSortValues':"asc",
            'yAxisVals': {
                'fieldParentTableID':self.tableID,
                 'fieldID':self.textFieldID,
                'summarizeValsWith':"count"
            },
            'geometry': {
        		"positionTop": 56,
        		"positionLeft": 212,
        		"sizeWidth": 394,
        		"sizeHeight": 394
            }
        }
        jsonResp = self.apiRequest('dashboard/barChart/new',barChartParams)
        barChartID = jsonResp[u'barChartID']
        print "Created bar chart : ", barChartID 
        
        getDataParams = {'parentDashboardID':dashboardID, 'barChartID':barChartID}
        jsonResp = self.apiRequest('dashboard/barChart/getData',getDataParams)
        dataRows = jsonResp[u'dataRows']
        self.assertEquals(len(dataRows),1,"Expecting 1 data row in the bar chart")
        firstDataRowVal = dataRows[0]['value']
        self.assertEquals(firstDataRowVal,2,"Expecting value for first data row to be 2 (count of records)")
               
 # TODO - Do more verification on parameter passing with the new datastore.               
 #       with self.assertRaises(AssertionError):
            # Invalid Parent Database ID - passes table ID instead of database ID
 #           dashboardParams = {'databaseID':self.tableID,'name':'My Dashboard'}
 #           jsonResp = self.apiRequest('dashboard/new',dashboardParams)
            
 #       with self.assertRaises(AssertionError):
            # Invalid Parent Database ID - passes dashboardID instead of barChart ID
 #           getDataParams = {'barChartID':dashboardID}
 #           jsonResp = self.apiRequest('dashboard/barChart/getData',getDataParams)
            
        


# Allow the tests in this file to be run stand-alone
if __name__ == '__main__':
    unittest.main()