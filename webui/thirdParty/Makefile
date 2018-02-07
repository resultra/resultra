DEPTH = ../..
include $(DEPTH)/build/common.mk
include $(DEPTH)/webui/build/webui.mk

realclean:
	$(RM) -Rf node_modules/

clean: 
	$(RM) *_gen.json

prebuild:
	npm install
			
export:
	$(EXPORT_THIRD_PARTY_ASSETS) $(abspath ./thirdPartyAssets.json)
	$(EXPORT_THIRD_PARTY_ASSETS) $(abspath ./aceAssets.json)
	$(EXPORT_THIRD_PARTY_ASSETS) $(abspath ./dashboardAssets.json)
	
all: prebuild export
	