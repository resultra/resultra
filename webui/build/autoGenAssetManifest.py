#!/usr/bin/env python

# This script globs for files with .css, .js, and .html extensions and generates a properly formatted
# "asset manifest" file listing the assets to be exported from the current directory. This script
# is intended to be used for "leaf node" directories.

import os
import sys
import json

currPath = "./"
basePath = os.path.abspath(sys.argv[1])
assetIncludeFilename = os.path.join(currPath,"assetInclude.json")


# Copy the asset lists from the asset manifest file, replacing the local filename with a 
# fully qualified (absolute) name.
jsFileList = []
cssFileList = []
htmlFileList = []
imageFileList = []

for item in os.listdir(currPath):
    if os.path.isfile(os.path.join(currPath,item)):
        fileAbsName = os.path.abspath(item)
        if fileAbsName.endswith('.css'):
            cssFileList.append(fileAbsName)
        if fileAbsName.endswith('.html'):
            htmlFileList.append(fileAbsName)
        if fileAbsName.endswith('.js'):
            jsFileList.append(fileAbsName)
        if fileAbsName.endswith('.png'):
            imageFileList.append(fileAbsName)

destAssetList = {}
destAssetList['basePath'] = basePath
destAssetList['jsFiles'] = jsFileList
destAssetList['cssFiles'] = cssFileList
destAssetList['htmlFiles'] = htmlFileList
destAssetList['imageFiles'] = imageFileList

# The script to recursively generate asset include lists looks for a "subDirs" entry
# which tells the script to recursively traverse the directories below for additional 
# assets. For the "leaf node" directories this script is designed to work with, the 
# subDirs entry is empty. 
destAssetList['subDirs'] = []
if os.path.isfile(assetIncludeFilename):
    with open(assetIncludeFilename) as json_file:
        assetInclude = json.load(json_file)
        destAssetList['subDirs'] = assetInclude['subDirs']

print json.dumps(destAssetList, indent=4, sort_keys=True)