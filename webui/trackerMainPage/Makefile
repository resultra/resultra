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
	$(BUILD_ASSET_INJECTION_LIST) trackerMainPageAssetInclude.json $(ASSET_BASE_PATH) > trackerMainPageAssetInclude_gen.json
	$(EXPORT_HTML_WITH_INJECTED_ASSETS) $(abspath ./trackerMainPageAssetInclude_gen.json)
