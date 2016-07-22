# This is a reusable makefile template for directories which only need to export their JS, CSS, HTML assets.
# Before including this makefile, the DEPTH variable needs to be defined.

include $(DEPTH)/build/common.mk
include $(DEPTH)/webui/build/webui.mk
			
package:
	$(EXPORT_ASSETS)