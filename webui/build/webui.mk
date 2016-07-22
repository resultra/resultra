
ASSET_BASE_PATH = $(DEPTH)/webui
BUILD_ASSET_INJECTION_LIST = $(DEPTH)/webui/build/buildAssetList.py
BUILD_ASSET_EXPORT_LIST = $(DEPTH)/webui/build/buildAssetExportList.py

GULP = gulp --gulpfile $(DEPTH)/webui/build/gulpfile.js

ifeq ($(DEBUG),1)
	EXPORT_ASSETS = $(BUILD_ASSET_EXPORT_LIST) ./assetManifest.json $(ASSET_BASE_PATH) > ./assetManifest_gen.json && \
			$(GULP) exportIndividualAssets  --assets $(abspath ./assetManifest_gen.json)
	INJECT_GULP_TARGETS = injectHTMLFilesWithIndividualAssets
else 
	EXPORT_ASSETS = echo "Export assets: Release build: no individual assets exported"
	INJECT_GULP_TARGETS = exportMinifiedAssets injectHTMLFilesWithMinifiedAssets
endif


EXPORT_HTML_WITH_INJECTED_ASSETS = $(GULP) $(INJECT_GULP_TARGETS) --assets