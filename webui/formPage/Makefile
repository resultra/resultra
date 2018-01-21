DEPTH = ../..
include $(DEPTH)/build/common.mk
include $(DEPTH)/webui/build/webui.mk

all: prebuild export
	
clean: 
	$(RM) *_gen.json
	
prebuild:
	$(AUTO_GEN_ASSET_MANIFEST)
	
		
export:
	$(EXPORT_ASSETS)
	$(BUILD_ASSET_INJECTION_LIST) submitFormPageAssetInclude.json $(ASSET_BASE_PATH) > submitFormPageAssets_gen.json
	$(EXPORT_HTML_WITH_INJECTED_ASSETS) $(abspath ./submitFormPageAssets_gen.json)
	$(BUILD_ASSET_INJECTION_LIST) viewFormPageAssetInclude.json $(ASSET_BASE_PATH) > viewFormPageAssets_gen.json
	$(EXPORT_HTML_WITH_INJECTED_ASSETS) $(abspath ./viewFormPageAssets_gen.json)