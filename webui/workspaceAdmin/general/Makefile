DEPTH = ../../..
include $(DEPTH)/build/common.mk
include $(DEPTH)/webui/build/webui.mk

	
clean: 
	$(RM) *_gen.json
	
prebuild:
	$(AUTO_GEN_ASSET_MANIFEST)
	
		
export:
	$(EXPORT_ASSETS)
	$(BUILD_ASSET_INJECTION_LIST) workspaceAdminPageAssetInclude.json $(ASSET_BASE_PATH) > workspaceAdminPageAssetInclude_gen.json
	$(EXPORT_HTML_WITH_INJECTED_ASSETS) $(abspath ./workspaceAdminPageAssetInclude_gen.json)


all: prebuild export
