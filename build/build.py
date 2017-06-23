#!/usr/bin/env python

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

from multiprocessing import Pool

parser = argparse.ArgumentParser(description='Main build script.')
parser.add_argument('--release',default=False,action='store_true',
                    help='perform a release build')
args = parser.parse_args()

failedDirs = []

debugBuild = 1
if(args.release):
    debugBuild = 0
    
    
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
    
    
def buildOneDir(buildSpec):
    print "Building: dir=", buildSpec.dirName, " phase=", buildSpec.targetName, " debug=", buildSpec.debugBuild
    retCode = os.system("make -C %s DEBUG=%s %s" % (buildSpec.dirName, buildSpec.debugBuild, buildSpec.targetName))
    return buildDirResult(buildSpec.dirName,retCode)
  
def runMakePhase(makeTargetName):
        
    print "Build: Starting phase = ", makeTargetName
    makeDirs = []
    for root, dirs, files in os.walk(".."):
        for file in files:
            if (file == 'Makefile') and (not root.startswith("../webui/build")):
                makeDirs.append(buildDirSpec(root,makeTargetName,debugBuild))
    buildPool = Pool(processes=6)
    results = buildPool.map(buildOneDir,makeDirs)
    buildPool.close()
    buildPool.join()
    print "Build: Done with phase = ", makeTargetName, " results = ",results
    for res in results:
        if res.errCode != 0:
            failedDirs.append(makeTargetName + ":" + res.dirName)
    
            
runMakePhase("prebuild")
runMakePhase("build")
runMakePhase("package")
                
print "\nBuild Results:\n"

if len(failedDirs) > 0:
    print "Build failed on following directories:\n"
    print "\n".join(failedDirs)
    sys.exit(255)
else:
    print "Build succeeded"
    sys.exit(0)
                
