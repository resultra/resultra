DEPTH = ../..
include $(DEPTH)/build/common.mk
include $(DEPTH)/webui/build/webui.mk

all: prebuild package

prebuild:

all: prebuild package
	
clean: 
	$(RM) *_gen.json
	
prebuild:
	$(GEN_ASSET_MANFIEST) ./thirdPartyAssetInclude.json $(ASSET_BASE_PATH) > ./assetManifest_gen.json
		
package:
	$(EXPORT_PKG_ASSETS)
	$(BUILD_ASSET_INJECTION_LIST) thirdPartyAssetInclude.json $(ASSET_BASE_PATH) > thirdPartyAssetInclude_gen.json
	$(EXPORT_PKG_HTML_WITH_INJECTED_ASSETS) $(abspath ./thirdPartyAssetInclude_gen.json)
	