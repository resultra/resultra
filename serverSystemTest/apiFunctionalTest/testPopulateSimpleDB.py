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
        
        self.createTestSession()
        
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
        
        newFormParams = { 'parentTableID':self.tableID,'name':'Purchases'}
        jsonResp = self.apiRequest('frm/new',newFormParams)
        self.formID = jsonResp[u'formID']
        print "populateSimpleDB: Form ID: ", self.formID
    
        newDashboardParams = {'databaseID':self.databaseID,'name':'Summary'}
        jsonResp = self.apiRequest('dashboard/new',newDashboardParams)
        self.dashboardID = jsonResp[u'dashboardID']
        print "populateSimpleDB: Dashboard ID: ", self.dashboardID


    def testInvalidAdminCreateTableOrDashboard(self):
        
        self.initSession()
        
        adminUser = self.newTestUser()
        self.signinTestUser(adminUser)          
        
        jsonResp = self.apiRequest('database/new',{'name': 'Test Database'})
        self.databaseID = jsonResp[u'databaseID']
        print "populateSimpleDB: Database ID: ",self.databaseID
        
        # Completely sign out the current user. This will cause the attempt to create a new 
        # table to fail, since there won't be a signed in user with admin privileges.
        self.signoutCurrUser()
        
        with self.assertRaises(AssertionError):
            jsonResp = self.apiRequest('table/new',{'databaseID': self.databaseID, 'name': 'Test Table'})
            self.tableID = jsonResp[u'tableID']
            print "testInvalidAdminCreateTable: Table ID: ",self.tableID
                        
        with self.assertRaises(AssertionError):
            newDashboardParams = {'databaseID':self.databaseID,'name':'Summary'}
            jsonResp = self.apiRequest('dashboard/new',newDashboardParams)
            self.dashboardID = jsonResp[u'dashboardID']
            print "populateSimpleDB: Dashboard ID: ", self.dashboardID
        
        # Create a completely different user, but not the same one who created the database and has
        # admin privileges. Try again to create a table while signed in as this user, and it should also fail.
        anotherUser = self.newTestUser()
        self.signinTestUser(anotherUser)      
        with self.assertRaises(AssertionError):
            jsonResp = self.apiRequest('table/new',{'databaseID': self.databaseID, 'name': 'Test Table'})
            self.tableID = jsonResp[u'tableID']
            print "testInvalidAdminCreateTable: Table ID: ",self.tableID
                  
        with self.assertRaises(AssertionError):
            newDashboardParams = {'databaseID':self.databaseID,'name':'Summary'}
            jsonResp = self.apiRequest('dashboard/new',newDashboardParams)
            self.dashboardID = jsonResp[u'dashboardID']
            print "populateSimpleDB: Dashboard ID: ", self.dashboardID
            
            
        # Sign out, then sign back in as the est user. Creation of the table should work now.
        self.signoutCurrUser()
        self.signinTestUser(adminUser)  
            
        jsonResp = self.apiRequest('table/new',{'databaseID': self.databaseID, 'name': 'Test Table'})
        self.tableID = jsonResp[u'tableID']
        print "testInvalidAdminCreateTable: Table ID: ",self.tableID
        
        newDashboardParams = {'databaseID':self.databaseID,'name':'Summary'}
        jsonResp = self.apiRequest('dashboard/new',newDashboardParams)
        self.dashboardID = jsonResp[u'dashboardID']
        print "populateSimpleDB: Dashboard ID: ", self.dashboardID
 
 
    def testInvalidAdminCreateFieldOrForm(self):
 
        self.createTestSession()
 
        jsonResp = self.apiRequest('database/new',{'name': 'Test Database'})
        self.databaseID = jsonResp[u'databaseID']
        print "populateSimpleDB: Database ID: ",self.databaseID
         
        jsonResp = self.apiRequest('table/new',{'databaseID': self.databaseID, 'name': 'Test Table'})
        self.tableID = jsonResp[u'tableID']
        print "populateSimpleDB: Table ID: ",self.tableID
         
        # Completely sign out the current user. This will cause the attempt to create a new
        # table to fail, since there won't be a signed in user with admin privileges.
        self.signoutCurrUser()
         
        with self.assertRaises(AssertionError):
            fieldParams = {'parentTableID':self.tableID,'name':'Quantity','type':'number','refName':'qty'}
            jsonResp = self.apiRequest('field/new',fieldParams)
            self.qtyFieldID = jsonResp[u'fieldID']
            
        with self.assertRaises(AssertionError):
            newFormParams = { 'parentTableID':self.tableID,'name':'Purchases'}
            jsonResp = self.apiRequest('frm/new',newFormParams)
            self.formID = jsonResp[u'formID']
            print "populateSimpleDB: Form ID: ", self.formID
         
        # Create a completely different user, but not the same one who created the database and has
        # admin privileges. Try again to create a table while signed in as this user, and it should also fail.
        anotherUser = self.newTestUser()
        self.signinTestUser(anotherUser)
         
        with self.assertRaises(AssertionError):
            fieldParams = {'parentTableID':self.tableID,'name':'Quantity','type':'number','refName':'qty'}
            jsonResp = self.apiRequest('field/new',fieldParams)
            self.qtyFieldID = jsonResp[u'fieldID']
            
        with self.assertRaises(AssertionError):
            newFormParams = { 'parentTableID':self.tableID,'name':'Purchases'}
            jsonResp = self.apiRequest('frm/new',newFormParams)
            self.formID = jsonResp[u'formID']
            print "populateSimpleDB: Form ID: ", self.formID
       
        
if __name__ == '__main__':
    unittest.main()