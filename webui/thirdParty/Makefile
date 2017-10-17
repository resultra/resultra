DEPTH = ../..
include $(DEPTH)/build/common.mk
include $(DEPTH)/webui/build/webui.mk

all: prebuild package

prebuild:

all: prebuild package
	
clean: 
	$(RM) *_gen.json
	$(RM) -Rf node_modules/

prebuild:
	npm install
			
package:
	$(EXPORT_THIRD_PARTY_ASSETS) $(abspath ./thirdPartyAssets.json)
	$(EXPORT_THIRD_PARTY_ASSETS) $(abspath ./aceAssets.json)
	