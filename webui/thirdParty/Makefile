DEPTH = ../..
include $(DEPTH)/build/common.mk
include $(DEPTH)/webui/build/webui.mk

prebuild:

clean: 
	$(RM) *_gen.json
	$(RM) -Rf node_modules/

prebuild:
	npm install
			
export:
	$(EXPORT_THIRD_PARTY_ASSETS) $(abspath ./thirdPartyAssets.json)
	$(EXPORT_THIRD_PARTY_ASSETS) $(abspath ./aceAssets.json)
	
all: prebuild export
	