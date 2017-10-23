DEPTH = ../../..
include $(DEPTH)/build/common.mk
include $(DEPTH)/webui/build/webui.mk

all: prebuild package
	
clean: 
	$(RM) *_gen.json
	
prebuild:
	$(GEN_ASSET_MANIFEST)
	
		
package:
	$(EXPORT_ASSETS)