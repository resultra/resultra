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

parser = argparse.ArgumentParser(description='Main build script.')
parser.add_argument('--release',default=False,action='store_true',
                    help='perform a release build')
args = parser.parse_args()

failedDirs = []

debugBuild = 1
if(args.release):
    debugBuild = 0
    

def runMakePhase(makeTargetName):
    for root, dirs, files in os.walk(".."):
        for file in files:
            if (file == 'Makefile') and (not root.startswith("../webui/build")):
                print os.path.join(root,file)
                retCode = os.system("make -C %s DEBUG=%s %s" % (root, debugBuild, makeTargetName))
                if retCode != 0:
                    failedDirs.append(makeTargetName + ":" + root)

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
                
