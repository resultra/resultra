#!/usr/bin/env python

# This script implements a phased build based upon the makefiles in the development tree.
# Within each build phase, make is run on each directory in no particular order. So,
# the build process from each directory is expected to not depend on other directories 
# within a a given phase.

import os
import sys

failedDirs = []

def runMakePhase(makeTargetName):
    for root, dirs, files in os.walk(".."):
        for file in files:
            if (file == 'Makefile') and (not root.startswith("../webui/build")):
                print os.path.join(root,file)
                retCode = os.system("make -C %s %s" % (root, makeTargetName))
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
                
