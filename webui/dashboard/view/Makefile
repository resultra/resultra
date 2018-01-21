DEPTH = ../../..
include $(DEPTH)/build/common.mk
include $(DEPTH)/webui/build/webui.mk

all: prebuild export
	
clean: 
	$(RM) *_gen.json
	
prebuild:
	$(GEN_ASSET_MANIFEST)
	
		
export:
	$(EXPORT_ASSETS)