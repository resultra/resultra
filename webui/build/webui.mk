
ASSET_BASE_PATH = $(DEPTH)/webui
BUILD_ASSET_INJECTION_LIST = $(DEPTH)/webui/build/buildAssetList.py

GEN_ASSET_MANFIEST = $(DEPTH)/webui/build/genAssetManifest.py

AUTO_GEN_ASSET_MANIFEST = $(DEPTH)/webui/build/autoGenAssetManifest.py $(ASSET_BASE_PATH) > ./assetManifest_gen.json

GEN_ASSET_MANIFEST = $(GEN_ASSET_MANFIEST) ./assetManifest.json $(ASSET_BASE_PATH) > ./assetManifest_gen.json

GULP = gulp --gulpfile $(DEPTH)/webui/build/gulpfile.js


ifeq ($(DEBUG),1)
	EXPORT_ASSETS = $(GULP) exportIndividualAssets  --assets $(abspath ./assetManifest_gen.json)
	INJECT_GULP_TARGETS = injectHTMLFilesWithIndividualAssets
else 
	EXPORT_ASSETS = echo "Export assets: Release build: no individual assets exported"
	INJECT_GULP_TARGETS = exportMinifiedAssets injectHTMLFilesWithMinifiedAssets
endif


EXPORT_HTML_WITH_INJECTED_ASSETS = $(GULP) $(INJECT_GULP_TARGETS) --assets