#!/usr/bin/env python

import unittest
import json

from testCommon import TestHelperMixin
    

class testValidateFormula(unittest.TestCase,TestHelperMixin):
    
    def validateOneFormula(self,formulaText):
        jsonResp = self.apiRequest('calcField/validateFormula',{'fieldID':self.fieldID,'formulaText':formulaText})
        print "validateOneFormula: formula response: ", json.dumps(jsonResp)
        isValidFormula = jsonResp[u'isValidFormula']
        return isValidFormula
        
 
    def setUp(self):
        jsonResp = self.apiRequest('database/new',{'name': 'Test Database'})
        self.databaseID = jsonResp[u'databaseID']
        print "testValidateFormula: database ID: ",self.databaseID
        
        jsonResp = self.apiRequest('table/new',{'databaseID': self.databaseID, 'name': 'Test Table'})
        self.tableID = jsonResp[u'tableID']
        print "testValidateFormula: table ID: ",self.tableID
        
        fieldParams = {'parentTableID':self.tableID,'name':'Quantity','type':'number','refName':'qty'}
        jsonResp = self.apiRequest('field/new',fieldParams)
        self.fieldID = jsonResp[u'fieldID']
        
 
    def testSimpleFormulas(self):
        self.assertTrue(self.validateOneFormula("52.5"),"simple number literal")
        self.assertTrue(self.validateOneFormula("SUM(52.5)"),"number literal inside function call")
        self.assertFalse(self.validateOneFormula("-SUM(52.5)"),"unary - not supported")
        self.assertTrue(self.validateOneFormula("-1.5*SUM(52.5)"),"negative numbers")
 
 
    def testMoreFormulas(self):
         self.assertTrue(self.validateOneFormula("52.5"))
        
        
 
if __name__ == '__main__':
    unittest.main()