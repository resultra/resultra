#!/usr/bin/env python
#
# This file is part of the Resultra project.
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.


import os
import sys
import json

currPath = "./"
assetListFileName = sys.argv[1]
basePath = os.path.abspath(sys.argv[2])


# Copy the asset lists from the asset manifest file, replacing the local filename with a 
# fully qualified (absolute) name.
jsFileList = []
cssFileList = []
htmlFileList = []
imageFileList = []

with open(assetListFileName) as json_file:
    srcAssetList = json.load(json_file)
    for cssFile in srcAssetList['cssFiles']:
        cssFileAbsPath = os.path.abspath(currPath + cssFile)
        cssFileList.append(cssFileAbsPath)
    for jsFile in srcAssetList['jsFiles']:
        jsFileAbsPath = os.path.abspath(currPath + jsFile)
        jsFileList.append(jsFileAbsPath)
    for htmlFile in srcAssetList['htmlFiles']:
        htmlFileAbsPath = os.path.abspath(currPath + htmlFile)
        htmlFileList.append(htmlFileAbsPath)
    for imageFile in srcAssetList['imageFiles']:
        imageFileAbsPath = os.path.abspath(currPath + imageFile)
        htmlFileList.append(imageFileAbsPath)


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

print json.dumps(destAssetList, indent=4, sort_keys=True)