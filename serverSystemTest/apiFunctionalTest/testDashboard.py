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

class TestDashboard(unittest.TestCase,TestHelperMixin):
    
    def setUp(self):
        self.createTestSession()
        self.databaseID = self.newDatabase('TestRecordFilter: Test Database')
        self.timeFieldID = self.newTimeField(self.databaseID,"TestRecordFilter - Time Field","TimeField")
        self.numberFieldID = self.newNumberField(self.databaseID,"TestRecordFilter - Number Field","NumberField") 
        self.textFieldID = self.newTextField(self.databaseID,"TestRecordFilter - Text Field","TextField") 
    
    def testSimpleDashboard(self):      
        dashboardParams = {'databaseID':self.databaseID,'name':'My Dashboard'}
        jsonResp = self.apiRequest('dashboard/new',dashboardParams)
        dashboardID = jsonResp[u'dashboardID']
        
        recordID1 = self.newRecord(self.databaseID)
        recordID2 = self.newRecord(self.databaseID)
                
        self.setTextRecordValue(self.databaseID,recordID1,self.textFieldID,"hello world")     
        self.setNumberRecordValue(self.databaseID,recordID1,self.numberFieldID,25.2)
        
        self.setTextRecordValue(self.databaseID,recordID2,self.textFieldID,"Testing 1,2,3")     
        self.setNumberRecordValue(self.databaseID,recordID2,self.numberFieldID,42.5)
        
        getRecordParams = {
        	"databaseID": self.databaseID,
        	"preFilterRules": {
        		"matchLogic": "all",
        		"filterRules": []
        	},
        	"filterRules": {
        		"matchLogic": "all",
        		"filterRules": []
        	},
        	"sortRules": []
        }
        jsonResp = self.apiRequest('recordRead/getFilteredSortedRecordValues',getRecordParams)
        self.assertEquals(len(jsonResp),2,"Expecting 2 records in the database")
        
        
        barChartParams = {
            'parentDashboardID':dashboardID,
            'xAxisVals': {
                'groupValsByFieldID':self.textFieldID,
                'groupValsBy':"none",
                'groupByValBucketWidth':0
            },
            'xAxisSortValues':"asc",
            'yAxisVals': {
                 'fieldID':self.textFieldID,
                'summarizeValsWith':"count"
            },
            'geometry': {
        		"positionTop": 0,
        		"positionLeft": 0,
        		"sizeWidth": 200,
        		"sizeHeight": 200
            }        
        }
        jsonResp = self.apiRequest('dashboard/barChart/new',barChartParams)
        barChartID = jsonResp[u'barChartID']
        print "Created bar chart : ", barChartID 
        
        defaultDashboardDataParams = {
            "dashboardID":dashboardID
        }
        jsonResp = self.apiRequest('dashboardController/getDefaultData',defaultDashboardDataParams)
        barChartsData = jsonResp[u'barChartsData']
        self.assertEquals(len(barChartsData),1,"Expecting data for 1 bar chart")
        dataRows = barChartsData[0][u'groupedSummarizedVals'][u'groupedDataRows']
        self.assertEquals(len(dataRows),2,"Expecting 2 data rows in the bar chart")
        overallRow = barChartsData[0][u'groupedSummarizedVals'][u'overallDataRow']
        summaryVal = overallRow[u'summaryVals'][0]
        self.assertEquals(summaryVal,2,"Expecting overall summary val (count) to be 2")
        
        getDataParams = {'parentDashboardID':dashboardID, 'barChartID':barChartID}
        jsonResp = self.apiRequest('dashboardController/getBarChartData',getDataParams)
        # TODO - Get working with updated result format (currently returning an empty data set)
 #       dataRows = jsonResp[u'dataRows']
  #      self.assertEquals(len(dataRows),1,"Expecting 1 data row in the bar chart")
   #     firstDataRowVal = dataRows[0]['value']
    #    self.assertEquals(firstDataRowVal,2,"Expecting value for first data row to be 2 (count of records)")
               
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