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
            buildAssetList(jsFileList,cssFileList,htmlFileList,subDirPath,"assetManifest_gen.json")
        for cssFile in assetList['cssFiles']:
            cssFileList.append(cssFile)
        for jsFile in assetList['jsFiles']:
            jsFileList.append(jsFile)
        for htmlFile in assetList['htmlFiles']:
             htmlFileList.append(htmlFile)

# The 'minJSFile' and 'minCSSFile' properties set in the root asset file sets the name of the 
# minified JS and CSS files respectively.
jsFileList = []
cssFileList = []
htmlFileList = []
currPath = "./"
assetIncludeFileName = sys.argv[1]
basePath = os.path.abspath(sys.argv[2])

#buildAssetList(jsFileList,cssFileList,htmlFileList,currPath,assetListFileName)

destAssetList = {}
destAssetList['jsFiles'] = jsFileList
destAssetList['cssFiles'] = cssFileList
destAssetList['htmlFiles'] = htmlFileList
destAssetList['basePath'] = basePath

# Copy the individual properties from the source asset list to the destination, then 
# recursively expanded set of included assets.
with open(assetIncludeFileName) as json_file:
    srcAssetList = json.load(json_file)
    destAssetList['minJSFile'] = srcAssetList['minJSFile']
    destAssetList['minCSSFile'] = srcAssetList['minCSSFile']
    destAssetList['injectPlaceholderName']  = srcAssetList['injectPlaceholderName']
    for subDir in srcAssetList['subDirs']:
        subDirPath = currPath + subDir + "/"
        buildAssetList(jsFileList,cssFileList,htmlFileList,subDirPath,"assetManifest_gen.json")
    # Lastly, include any local assets from the current directory. The file names on these
    # assets need to be converted to absolute paths.
    for cssFile in srcAssetList['cssFiles']:
        cssFileAbsPath = os.path.abspath(cssFile)
        cssFileList.append(cssFileAbsPath)
    for jsFile in srcAssetList['jsFiles']:
        jsFileAbsPath = os.path.abspath(jsFile)
        jsFileList.append(jsFileAbsPath)
    for htmlFile in srcAssetList['htmlFiles']:
        htmlFileAbsPath = os.path.abspath(htmlFile)
        htmlFileList.append(htmlFileAbsPath)
    

print json.dumps(destAssetList, indent=4, sort_keys=True)