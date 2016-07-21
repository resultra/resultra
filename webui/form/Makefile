DEPTH = ../..
include $(DEPTH)/build/common.mk
include $(DEPTH)/webui/build/webui.mk

all: prebuild
	
clean: 
	$(RM) *_gen.json
		
package:
	$(BUILD_ASSET_LIST) designFormAssets.json $(ASSET_BASE_PATH) > designFormAssets_gen.json
	$(GULP) debugExportAssets exportMinifiedAssets injectHTMLFiles --assets $(abspath ./designFormAssets_gen.json)

