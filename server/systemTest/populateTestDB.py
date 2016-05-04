#!/usr/bin/env python
#
# Using the API, populate a test database with a single form, a field of each type, 
# and a single dashboard.
#

import requests

newDatabaseURL = 'http://localhost:8080/api/database/new'
newDatabaseParams = {'name': 'Test Database'}
resp = requests.post(newDatabaseURL,json=newDatabaseParams)
databaseID = resp.json()[u'databaseID']
print "Create database: response status=",resp.status_code, " id=",resp.json()[u'databaseID']

newTableURL = 'http://localhost:8080/api/table/new'
newTableParams = {'databaseID': databaseID, 'name': 'My Table'}
resp = requests.post(newTableURL,json=newTableParams)
tableID = resp.json()[u'tableID']
print "Create table: response status=",resp.status_code, " id=",resp.json()[u'tableID']

newFieldURL = 'http://localhost:8080/api/field/new'

newFieldParams = {'parentTableID':tableID,'name':'Quantiy','type':'number','refName':'qty'}
resp = requests.post(newFieldURL,json=newFieldParams)
print "Create field: response status=",resp.status_code, " id=",resp.json()[u'fieldID']

newFieldParams = {'parentTableID':tableID,'name':'Price','type':'number','refName':'price'}
resp = requests.post(newFieldURL,json=newFieldParams)
print "Create field: response status=",resp.status_code, " id=",resp.json()[u'fieldID']

newFieldParams = {'parentTableID':tableID,'name':'Good Price?','type':'bool','refName':'goodPrice'}
resp = requests.post(newFieldURL,json=newFieldParams)
print "Create field: response status=",resp.status_code, " id=",resp.json()[u'fieldID']

newCalcFieldURL = 'http://localhost:8080/api/calcField/new'
newCalcFieldParams = {'parentTableID':tableID,'name':'Total','type':'number',
                    'refName':'total','formulaText':'42.5'}
resp = requests.post(newCalcFieldURL,json=newCalcFieldParams)
print "Create calculated field: response status=",resp.status_code, " id=",resp.json()[u'fieldID']

newFormURL = 'http://localhost:8080/api/frm/new'
newFormParams = { 'tableID':tableID,'name':'Purchases'}
resp = requests.post(newFormURL,json=newFormParams)
print "Create form: response status=",resp.status_code, " id=",resp.json()[u'formID']

newDashboardURL = 'http://localhost:8080/api/newDashboard'
newDashboardParams = {'databaseID':databaseID,'name':'Summary'}
resp = requests.post(newDashboardURL,json=newDashboardParams)
print "Create dashboard: response status=",resp.status_code, " id=",resp.json()[u'dashboardID']
