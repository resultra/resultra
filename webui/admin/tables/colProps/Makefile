DEPTH = ../../../..
include $(DEPTH)/build/common.mk
include $(DEPTH)/webui/build/webui.mk

all: prebuild export
	
clean: 
	$(RM) *_gen.json
	
prebuild:
	$(AUTO_GEN_ASSET_MANIFEST)
	
export:
	$(EXPORT_ASSETS)
	$(BUILD_ASSET_INJECTION_LIST) colPropsPageAssetInclude.json $(ASSET_BASE_PATH) > colPropsPageAssets_gen.json
	$(EXPORT_HTML_WITH_INJECTED_ASSETS) $(abspath ./colPropsPageAssets_gen.json)