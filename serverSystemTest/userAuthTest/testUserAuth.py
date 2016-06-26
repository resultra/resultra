#!/usr/bin/env python

import unittest

from userAuthTestHelper import UserAuthTestHelperMixin

class TestUserAuth(unittest.TestCase,UserAuthTestHelperMixin):
        
    def testAuth(self):
        newUserParams = {'emailAddr':'test@example.com','password':'testpw123$'}
        jsonResp = self.authRequest('register',newUserParams)
        self.assertEquals(jsonResp['success'],True,"register new user")
 
        newUserParams = {'emailAddr':'test@example.com','password':'testpw123$'}
        jsonResp = self.authRequest('register',newUserParams)
        self.assertEquals(jsonResp['success'],False,"repeat registration")
 
        
        loginParams = {'emailAddr':'test@example.com','password':'testpw123$'}
        jsonResp = self.authRequest('login',loginParams)
        self.assertEquals(jsonResp['success'],True,"successful login")

        failLoginParams = {'emailAddr':'test@example.com','password':'wrongpw123'}
        jsonResp = self.authRequest('login',failLoginParams)
        self.assertEquals(jsonResp['success'],False,"wrong password")

        invalidUserParams = {'emailAddr':'nonuser@example.com','password':'testpw123'}
        jsonResp = self.authRequest('login',invalidUserParams)
        self.assertEquals(jsonResp['success'],False,"invalid email")
    
        print "TestRegisterUser done"

    def testInvalidRegistration(self):
        newUserParams = {'emailAddr':'test1@example.com','password':'test123'}
        jsonResp = self.authRequest('register',newUserParams)
        self.assertEquals(jsonResp['success'],False,"password too short")
          

# Allow the tests in this file to be run stand-alone
if __name__ == '__main__':
    unittest.main()