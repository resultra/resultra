DEPTH = ../../..
include $(DEPTH)/build/common.mk
include $(DEPTH)/webui/build/webui.mk

all: prebuild export
	
clean: 
	$(RM) *_gen.json
	
prebuild:
	$(AUTO_GEN_ASSET_MANIFEST)
	
export:
	$(EXPORT_ASSETS)
	$(BUILD_ASSET_INJECTION_LIST) userManagementAssetInclude.json $(ASSET_BASE_PATH) > userManagementAssets_gen.json
	$(EXPORT_HTML_WITH_INJECTED_ASSETS) $(abspath ./userManagementAssets_gen.json)
	$(BUILD_ASSET_INJECTION_LIST) userPropsAssetInclude.json $(ASSET_BASE_PATH) > userPropsAssets_gen.json
	$(EXPORT_HTML_WITH_INJECTED_ASSETS) $(abspath ./userPropsAssets_gen.json)