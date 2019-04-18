#!/usr/bin/env python

# This is a simple helper script for use with the build. It tests whether the current
# path is inside the GOPATH. This is needed to use Go modules when a repository
# is outside GOPATH, and the dep tool otherwise.

import os

goSrcPath = os.environ['GOPATH'] + '/src'
currPath = os.getcwd()

if currPath.startswith(goSrcPath):
	print "TRUE"
else:
	print "FALSE"
