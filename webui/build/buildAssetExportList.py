#!/usr/bin/env python

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


destAssetList = {}
destAssetList['basePath'] = basePath
destAssetList['jsFiles'] = jsFileList
destAssetList['cssFiles'] = cssFileList
destAssetList['htmlFiles'] = htmlFileList
print json.dumps(destAssetList, indent=4, sort_keys=True)