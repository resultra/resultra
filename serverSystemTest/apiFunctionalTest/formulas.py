#!/usr/bin/env python

import unittest
import json

from testCommon import TestHelperMixin
    

class TestFormulas(unittest.TestCase,TestHelperMixin):
    
    def verifyFormula(self,resultFieldID,formulaText,whatTested):
        jsonResp = self.apiRequest('calcField/validateFormula',{'fieldID':resultFieldID,'formulaText':formulaText})
        isValidFormula = jsonResp[u'isValidFormula']
        if isValidFormula:
            print "PASS: verifyFormula: ", whatTested
        else:
            print "FAIL: verifyFormula: ", whatTested, ": response = ", json.dumps(jsonResp)
            self.fail(msg=whatTested)

    # TODO - Enhance this function to include a string to look for in the expected error message
    def verifyBadFormula(self,resultFieldID,formulaText,whatTested):
        jsonResp = self.apiRequest('calcField/validateFormula',{'fieldID':resultFieldID,'formulaText':formulaText})
        isValidFormula = jsonResp[u'isValidFormula']
        errorMsg = jsonResp[u'errorMsg']
        if not isValidFormula:
            print "PASS: verifyBadFormula: ", whatTested, ": error response = ",json.dumps(jsonResp)
        else:
            print "FAIL: verifyOneFormula: ", whatTested
            self.fail(msg=whatTested)
        
 
    def setUp(self):
        jsonResp = self.apiRequest('database/new',{'name': 'Test Database'})
        self.databaseID = jsonResp[u'databaseID']
        print "testValidateFormula: database ID: ",self.databaseID
        
        jsonResp = self.apiRequest('table/new',{'databaseID': self.databaseID, 'name': 'Test Table'})
        self.tableID = jsonResp[u'tableID']
        print "testValidateFormula: table ID: ",self.tableID
        
        fieldParams = {'parentTableID':self.tableID,'name':'Total','type':'number',
                    'refName':'total','formulaText':'42.5'}
        jsonResp = self.apiRequest('calcField/new',fieldParams)
        self.numberCalcField = jsonResp[u'fieldID']
  
        fieldParams = {'parentTableID':self.tableID,'name':'TextCalc','type':'text',
                  'refName':'textCalc','formulaText':'"hello world"'}
        jsonResp = self.apiRequest('calcField/new',fieldParams)
        self.textCalcField = jsonResp[u'fieldID']
  
        
        fieldParams = {'parentTableID':self.tableID,'name':'Quantity','type':'number','refName':'qty'}
        jsonResp = self.apiRequest('field/new',fieldParams)
        self.numberFieldID = jsonResp[u'fieldID']
         
        fieldParams = {'parentTableID':self.tableID,'name':'Comments','type':'text','refName':'CMT'}
        jsonResp = self.apiRequest('field/new',fieldParams)
        self.textFieldID = jsonResp[u'fieldID']
 
    def testSimpleFormulas(self):
        self.verifyFormula(self.numberCalcField,"52.5","simple number literal")
        self.verifyFormula(self.numberCalcField,"SUM(52.5)","number literal inside function call")
        self.verifyFormula(self.numberCalcField,"-1.5*SUM(52.5)","negative numbers")
        self.verifyBadFormula(self.numberCalcField,"-SUM(52.5)","unary - not supported")
 
    def testFunctionNames(self):
        self.verifyBadFormula(self.numberCalcField, "XYZ()","XYZ is not an undefined function name")
        self.verifyFormula(self.numberCalcField,"sum(52.5)","function names are case insensitive")
        self.verifyFormula(self.numberCalcField,"SUM(52.5)","function names are case insensitive")
        self.verifyFormula(self.numberCalcField,"Sum(52.5)","function names are case insensitive")
        self.verifyFormula(self.numberCalcField,"SuM(52.5)","function names are case insensitive")
        
 
    def testFunctionArgs(self):
        self.verifyBadFormula(self.numberCalcField, "SUM()","SUM function needs at least one argument")
        self.verifyBadFormula(self.numberCalcField, 'SUM("text")',"SUM function takes numberical arguments")
        self.verifyBadFormula(self.textCalcField,"CONCATENATE(25.3)","CONCATENATE function needs text arguments")
        self.verifyBadFormula(self.textCalcField,'CONCATENATE("first arg",25.3)',"CONCATENATE function needs text arguments")
        self.verifyBadFormula(self.textCalcField,"CONCATENATE()","CONCATENATE function needs at least 1 argument")
        self.verifyBadFormula(self.textCalcField,"CONCATENATE([qty])","CONCATENATE qty field is a number but CONCATENATE takes text args")
        self.verifyFormula(self.numberCalcField, "SUM([qty])","qty field should work as an argument to SUM")
        self.verifyFormula(self.textCalcField,"CONCATENATE([CMT])","CMT field should work as an argument to CONCATENATE")
        self.verifyBadFormula(self.numberCalcField, "SUM([CMT])","CMT field should not work as an argument to SUM")

    def testInvalidResultType(self):
        self.verifyBadFormula(self.textCalcField,"52.5","Can't assign number to text field type")
        self.verifyBadFormula(self.numberCalcField,'"abc 123"',"Can't assign text to a number field type")

    def testNonCalcField(self):
        jsonResp = self.apiRequest('calcField/validateFormula',{'fieldID':self.numberFieldID,'formulaText':"42.5"})
        isValidFormula = jsonResp[u'isValidFormula']
        print "testNonCalcField: response = ", json.dumps(jsonResp)
        self.assertFalse(isValidFormula,"Formulas shouldn't work with non-calculated fields")
        
    # TODO - Test setting of formulas on fields, including:
    #    - Trying to set a formula on a non-calculated field.
 
if __name__ == '__main__':
    unittest.main()