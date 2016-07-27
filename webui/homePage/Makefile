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
	$(BUILD_ASSET_INJECTION_LIST) homePagePublicAssetInclude.json $(ASSET_BASE_PATH) > homePagePublicAssetInclude_gen.json
	$(EXPORT_HTML_WITH_INJECTED_ASSETS) $(abspath ./homePagePublicAssetInclude_gen.json)
	$(BUILD_ASSET_INJECTION_LIST) homePageSignedInAssetInclude.json $(ASSET_BASE_PATH) > homePageSignedInAssetInclude_gen.json
	$(EXPORT_HTML_WITH_INJECTED_ASSETS) $(abspath ./homePageSignedInAssetInclude_gen.json)