#!/usr/bin/env python
#
# Using the API, populate a test database with a single form, a field of each type, 
# and a single dashboard.
#

import unittest
from testCommon import TestHelperMixin
    

class CopyToTemplate(unittest.TestCase,TestHelperMixin):
    
    # TODO - include tests where invalid IDs are passed to the different functions to create or retrieve
    # database entities; i.e., invalid type or invalid ID altogether.
    
    def testSimpleTemplate(self):
        
        self.createTestSession()
        
        jsonResp = self.apiRequest('database/new',{'name': 'Test Database'})
        self.databaseID = jsonResp[u'databaseID']
        print "testSimpleTemplate: Database ID: ",self.databaseID
        
        jsonResp = self.apiRequest('table/new',{'databaseID': self.databaseID, 'name': 'Test Table'})
        self.tableID = jsonResp[u'tableID']
        print "testSimpleTemplate: Table ID: ",self.tableID
        
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
        
        globalParams = {'parentDatabaseID':self.databaseID,
            'name':'Global Number','refName':'globalNum',
            'type':'number'}  
        jsonResp = self.apiRequest('global/new',globalParams)
        self.numberGlobal = jsonResp[u'globalID']
        

        fieldParams = {'parentTableID':self.tableID,'name':'Total','type':'number',
                    'refName':'total','formulaText':'42.5'}
        jsonResp = self.apiRequest('calcField/new',fieldParams)
        self.totalFieldID = jsonResp[u'fieldID']
        
        newFormParams = { 'parentTableID':self.tableID,'name':'Purchases'}
        jsonResp = self.apiRequest('frm/new',newFormParams)
        self.formID = jsonResp[u'formID']
        print "testSimpleTemplate: Form ID: ", self.formID
    
        newDashboardParams = {'databaseID':self.databaseID,'name':'Summary'}
        jsonResp = self.apiRequest('dashboard/new',newDashboardParams)
        self.dashboardID = jsonResp[u'dashboardID']
        print "testSimpleTemplate: Dashboard ID: ", self.dashboardID
 
        saveTemplateParams = {'sourceDatabaseID':self.databaseID,'newTemplateName':'My Template'}
        jsonResp = self.apiRequest('database/saveAsTemplate',saveTemplateParams)
        self.templateID = jsonResp[u'databaseID']
        print "testSimpleTemplate: Template Database ID: ", self.templateID
 
if __name__ == '__main__':
    unittest.main()