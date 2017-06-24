
ASSET_BASE_PATH = $(DEPTH)/webui

# IMPORTANT NOTE: Makefile rules with BUILD_ASSET_INJECTION_LIST need to come in the 'build' or 'package' 
# target of the Makefile. Specifically, it needs to come after any rules with GEN_ASSET_MANIFEST. This
# is because GEN_ASSET_MANIFEST generates the manifest files in the sub-directories which BUILD_ASSET_INJECTION_LIST
# depends upon.
BUILD_ASSET_INJECTION_LIST = $(DEPTH)/webui/build/buildAssetList.py

GEN_ASSET_MANFIEST = $(DEPTH)/webui/build/genAssetManifest.py

AUTO_GEN_ASSET_MANIFEST = $(DEPTH)/webui/build/autoGenAssetManifest.py $(ASSET_BASE_PATH) > ./assetManifest_gen.json

GEN_ASSET_MANIFEST = $(GEN_ASSET_MANFIEST) ./assetManifest.json $(ASSET_BASE_PATH) > ./assetManifest_gen.json

GULP = gulp --gulpfile $(DEPTH)/webui/build/gulpfile.js
THIRD_PARTY_GULP = gulp --gulpfile $(DEPTH)/webui/build/thirdPartyGulpFile.js

# TODO - the EXPORT_ASSETS rule will export *all* the assets found in assetManifest_gen.json
# However, many times an included asset from another package/directory will be exported multiple times
# if it is included from another directory. This means a single included asset may be exported many times,
# and possibily erroneously exported without its own injected asset lists included. There needs to be
# an enhancment to the overall build system to prevent assets from being exported more than once.

ifeq ($(DEBUG),1)
	EXPORT_ASSETS = $(GULP) exportIndividualAssets  --assets $(abspath ./assetManifest_gen.json)
	EXPORT_PKG_ASSETS = $(GULP) exportIndividualPackageAssets  --assets $(abspath ./assetManifest_gen.json)
	INJECT_GULP_TARGETS = injectHTMLFilesWithIndividualAssets
else 
	EXPORT_ASSETS = echo "Export assets: Release build: no individual assets exported"
	EXPORT_PKG_ASSETS = echo "Export assets: Release build: no individual assets exported"
	INJECT_GULP_TARGETS = exportMinifiedAssets injectHTMLFilesWithMinifiedAssets
endif


EXPORT_HTML_WITH_INJECTED_ASSETS = $(GULP) $(INJECT_GULP_TARGETS) --assets

EXPORT_THIRD_PARTY_ASSETS = $(THIRD_PARTY_GULP) exportThirdPartyAssets injectHTMLFilesWithIndividualPkgAssets --pathprefix ${CURDIR} --pkgassets $(abspath ./pkg_assets.json)