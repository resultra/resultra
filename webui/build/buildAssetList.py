#!/usr/bin/env python

import os
import sys
import json


# Recursively traverse the asset list/manifest, starting with the
# "root asset list" given on the command-line. This will recursively expand the 
# JSS and CSS dependencies for a given asset list.
def buildAssetList(jsFileList, cssFileList, htmlFileList, rootPath, assetListFileName):
    with open(rootPath + assetListFileName) as json_file:
        assetList = json.load(json_file)
        for subDir in assetList['subDirs']:
            subDirPath = rootPath + subDir + "/"
            buildAssetList(jsFileList,cssFileList,htmlFileList,subDirPath,"assetManifest.json")
        for cssFile in assetList['cssFiles']:
            cssFileAbsPath = os.path.abspath(rootPath + cssFile)
            cssFileList.append(cssFileAbsPath)
        for jsFile in assetList['jsFiles']:
            jsFileAbsPath = os.path.abspath(rootPath + jsFile)
            jsFileList.append(jsFileAbsPath)
        for htmlFile in assetList['htmlFiles']:
            htmlFileAbsPath = os.path.abspath(rootPath + htmlFile)
            htmlFileList.append(htmlFileAbsPath)

# The 'minJSFile' and 'minCSSFile' properties set in the root asset file sets the name of the 
# minified JS and CSS files respectively.
jsFileList = []
cssFileList = []
htmlFileList = []
currPath = "./"
assetListFileName = sys.argv[1]
basePath = os.path.abspath(sys.argv[2])

buildAssetList(jsFileList,cssFileList,htmlFileList,currPath,assetListFileName)

destAssetList = {}
destAssetList['jsFiles'] = jsFileList
destAssetList['cssFiles'] = cssFileList
destAssetList['htmlFiles'] = htmlFileList
destAssetList['basePath'] = basePath

# Copy the individual properties from the source asset list to the destination.
# The final output will contain a recursively expanded set of assets and these
# properties.
with open(assetListFileName) as json_file:
    srcAssetList = json.load(json_file)
    destAssetList['minJSFile'] = srcAssetList['minJSFile']
    destAssetList['minCSSFile'] = srcAssetList['minCSSFile']
    destAssetList['injectPlaceholderName']  = srcAssetList['injectPlaceholderName']




print json.dumps(destAssetList, indent=4, sort_keys=True)