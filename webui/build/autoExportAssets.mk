# This is a reusable makefile template for directories which only need to export their JS, CSS, HTML assets.
# Before including this makefile, the DEPTH variable needs to be defined.

include $(DEPTH)/build/common.mk
include $(DEPTH)/webui/build/webui.mk

all: prebuild package

clean:
	$(RM) *_gen.json
	
prebuild:
	$(AUTO_GEN_ASSET_MANIFEST)
			
package:
	$(EXPORT_ASSETS)