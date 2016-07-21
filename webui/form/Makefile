DEPTH = ../..
include $(DEPTH)/build/common.mk
include $(DEPTH)/webui/build/webui.mk

all: prebuild
	
clean: 
	$(RM) *_gen.json
		
package:
	$(BUILD_ASSET_LIST) designFormAssetInclude.json $(ASSET_BASE_PATH) > designFormAssets_gen.json
	$(GULP) exportIndividualAssets  --assets $(abspath ./designFormAssets_gen.json)
	$(GULP) exportMinifiedAssets  --assets $(abspath ./designFormAssets_gen.json)
	$(GULP) injectHTMLFilesWithIndividualAssets --assets $(abspath ./designFormAssets_gen.json)
	$(GULP) injectHTMLFilesWithMinifiedAssets --assets $(abspath ./designFormAssets_gen.json)

