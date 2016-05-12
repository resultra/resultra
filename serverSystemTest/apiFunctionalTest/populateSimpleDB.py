#!/usr/bin/env python
#
# Using the API, populate a test database with a single form, a field of each type, 
# and a single dashboard.
#

import unittest
from testCommon import TestHelperMixin
    

class populateSimpleDB(unittest.TestCase,TestHelperMixin):
    
    # TODO - include tests where invalid IDs are passed to the different functions to create or retrieve
    # database entities; i.e., invalid type or invalid ID altogether.
    
    def testPopulate(self):
        jsonResp = self.apiRequest('database/new',{'name': 'Test Database'})
        self.databaseID = jsonResp[u'databaseID']
        print "populateSimpleDB: Database ID: ",self.databaseID
        
        jsonResp = self.apiRequest('table/new',{'databaseID': self.databaseID, 'name': 'Test Table'})
        self.tableID = jsonResp[u'tableID']
        print "populateSimpleDB: Table ID: ",self.tableID
        
        fieldParams = {'parentTableID':self.tableID,'name':'Quantity','type':'number','refName':'qty'}
        jsonResp = self.apiRequest('field/new',fieldParams)
        self.qtyFieldID = jsonResp[u'fieldID']

        fieldParams = {'parentTableID':self.tableID,'name':'Price','type':'number','refName':'price'}
        jsonResp = self.apiRequest('field/new',fieldParams)
        self.priceFieldID = jsonResp[u'fieldID']
        
        fieldParams = {'parentTableID':self.tableID,'name':'Good Price?','type':'bool','refName':'goodPrice'}
        jsonResp = self.apiRequest('field/new',fieldParams)
        self.goodPriceField = jsonResp[u'fieldID']
        
        self.purchaseDateField = self.newTimeField(self.tableID,"Purchase Date","PurchDate")
        self.purchaseCommentsField = self.newLongTextField(self.tableID,"Purchase Comments","PurchComment")
        self.entryChartField = self.newFileField(self.tableID,"Entry Chart","EntryChart")

        fieldParams = {'parentTableID':self.tableID,'name':'Total','type':'number',
                    'refName':'total','formulaText':'42.5'}
        jsonResp = self.apiRequest('calcField/new',fieldParams)
        self.totalFieldID = jsonResp[u'fieldID']
        
        newFormParams = { 'tableID':self.tableID,'name':'Purchases'}
        jsonResp = self.apiRequest('frm/new',newFormParams)
        self.formID = jsonResp[u'formID']
        print "populateSimpleDB: Form ID: ", self.formID
    
        newDashboardParams = {'databaseID':self.databaseID,'name':'Summary'}
        jsonResp = self.apiRequest('newDashboard',newDashboardParams)
        self.dashboardID = jsonResp[u'dashboardID']
        print "populateSimpleDB: Dashboard ID: ", self.dashboardID
        
        
if __name__ == '__main__':
    unittest.main()