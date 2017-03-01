DEPTH = ../..
include $(DEPTH)/build/common.mk
include $(DEPTH)/webui/build/webui.mk

all: prebuild package
	
clean: 
	$(RM) *_gen.json
	
prebuild:
	$(AUTO_GEN_ASSET_MANIFEST)
	
		
package:
	$(EXPORT_ASSETS)
	$(BUILD_ASSET_INJECTION_LIST) submitFormPageAssetInclude.json $(ASSET_BASE_PATH) > submitFormPageAssets_gen.json
	$(EXPORT_HTML_WITH_INJECTED_ASSETS) $(abspath ./submitFormPageAssets_gen.json)