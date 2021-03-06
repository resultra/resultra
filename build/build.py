#!/usr/bin/env python
#
# This file is part of the Resultra project.
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.



# This script implements a phased build based upon the makefiles in the development tree.
# Within each build phase, make is run on each directory in no particular order. So,
# the build process from each directory is expected to not depend on other directories 
# within a a given phase.
#
# By default a debug build is performed. However to perform a release build, pass the --release
# option on the command line.
import os
import sys
import argparse
import time
import sys

from multiprocessing import Pool

parser = argparse.ArgumentParser(description='Main build script.')
parser.add_argument('--release',default=False,action='store_true',
                    help='perform a release build')
parser.add_argument('--realcleanonly',default=False,action='store_true',
                    help='only run the clean and realclean targets across the build')
parser.add_argument('--verbose',default=False,action='store_true',
                    help='more verbose build output')

# Building a Docker distribution and cross-compiling the Windows
# Electron client can only be done on a Linux build machine.
isLinuxBuild = sys.platform.startswith('linux')
if isLinuxBuild:  
    parser.add_argument('--windows',default=False,action='store_true',
                    help='cross-compile the Windows Electron client.')

parser.add_argument('--procs',default=4,type=int,
                    help='number of build tasks to run in parallel build on (default = 4)')
args = parser.parse_args()

failedDirs = []

debugBuild = 1
if(args.release):
    debugBuild = 0
    
verboseBuild = 0
if(args.verbose):
    verboseBuild =1
        
class buildDirResult:
    def __repr__(self):
        return "(dir = %s, err = %d) " % (self.dirName,self.errCode)
        
    def __init__(self, dirName,errCode):
        self.dirName = dirName
        self.errCode = errCode

class buildDirSpec:
    def __init__(self, dirName,targetName,debugBuild):
        self.dirName = dirName
        self.targetName = targetName
        self.debugBuild = debugBuild
     
# TODO - One option to consider is to redirect the output from each individual directory's build to a file, then 
# only output the verbose output from the build back out to STDOUT if building that directory fails. This would
# further reduce the size of the build logs.
def buildOneDir(buildSpec):
    if verboseBuild:
        print "Building: dir=", buildSpec.dirName, " phase=", buildSpec.targetName, " debug=", buildSpec.debugBuild
    bldCmd = "make --silent -C %s --jobs=2 DEBUG=%s %s" % (buildSpec.dirName, buildSpec.debugBuild, buildSpec.targetName)
    if verboseBuild:
        print "Build cmd: %s " % (bldCmd)
    retCode = os.system(bldCmd)
    if retCode != 0:
        print "FAIL: failure building dir = %s, target= %s, err = %d" % (buildSpec.dirName,buildSpec.targetName,retCode)
    return buildDirResult(buildSpec.dirName,retCode)
  
def runMakePhase(makeTargetName):
        
    print "Build: Starting phase = ", makeTargetName
    makeDirs = []
        
    for root, dirs, files in os.walk(".."):
        for file in files:
            if (file == 'Makefile') and (not "node_modules" in root) and (not "vendor" in root):
                makeDirs.append(buildDirSpec(root,makeTargetName,debugBuild))
    buildPool = Pool(processes=args.procs)
    results = buildPool.map(buildOneDir,makeDirs)
    buildPool.close()
    buildPool.join()
    print "Build: Done with phase = ", makeTargetName
    for res in results:
        if res.errCode != 0:
            failedDirs.append(makeTargetName + ":" + res.dirName)
    
startTime = time.time() 

if args.realcleanonly:
    runMakePhase("clean")
    runMakePhase("realclean")
else:        
    runMakePhase("install")
    runMakePhase("prebuild")
    runMakePhase("build")
    runMakePhase("export")
    runMakePhase("package")
    runMakePhase("test")
    runMakePhase("systest")
    if isLinuxBuild:
        runMakePhase("dockerdist")
        runMakePhase("dockerpkg")
        if args.windows:
            runMakePhase("windows")
            runMakePhase("winpkg")

endTime = time.time()

print "\n\n--------------------------------------------------------"
print "Build complete: parallel build tasks = %d, elapse time = %d secs " % (args.procs, endTime-startTime)
print "\nBuild Results:\n"

if len(failedDirs) > 0:
    print "Build failed on following directories:\n"
    print "\n".join(failedDirs)
    sys.exit(255)
else:
    print "Build succeeded"
    sys.exit(0)
                
