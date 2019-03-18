#!/usr/bin/env python
#
# This file is part of the Resultra project.
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.



#
# Using the API, populate a test database with a single form, a field of each type, 
# and a single dashboard.
#

import unittest
from testCommon import TestHelperMixin
    

class fieldTest(unittest.TestCase,TestHelperMixin):
        
    def duplicateFieldRefName(self):
        
        self.createTestSession()
        
        jsonResp = self.apiRequest('database/new',{'name': 'Test Database'})
        self.databaseID = jsonResp[u'databaseID']
        print "populateSimpleDB: Database ID: ",self.databaseID

        fieldRefName = 'qty'

        fieldParams = {'parentDatabaseID':self.databaseID,'name':'Quantity','type':'number','refName':fieldRefName}
        jsonResp = self.apiRequest('field/new',fieldParams)
              
        with self.assertRaises(AssertionError):
            fieldParams = {'parentDatabaseID':self.databaseID,'name':'Quantity2','type':'number','refName':fieldRefName}
            jsonResp = self.apiRequest('field/new',fieldParams) 


    def duplicateFieldName(self):
        
        self.createTestSession()
        
        jsonResp = self.apiRequest('database/new',{'name': 'Test Database'})
        self.databaseID = jsonResp[u'databaseID']
        print "populateSimpleDB: Database ID: ",self.databaseID

        fieldName = 'Quantity'

        fieldParams = {'parentDatabaseID':self.databaseID,'name':fieldName,'type':'number','refName':'qty'}
        jsonResp = self.apiRequest('field/new',fieldParams)
              
        with self.assertRaises(AssertionError):
            fieldParams = {'parentDatabaseID':self.databaseID,'name':fieldName,'type':'number','refName':'qty2'}
            jsonResp = self.apiRequest('field/new',fieldParams) 
       
        
if __name__ == '__main__':
    unittest.main()