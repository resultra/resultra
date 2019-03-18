# This file is part of the Resultra project.
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.


# This is a reusable makefile template for directories which only need to export their JS, CSS, HTML assets.
# Before including this makefile, the DEPTH variable needs to be defined.

include $(DEPTH)/build/common.mk
include $(DEPTH)/webui/build/webui.mk

all: prebuild package

clean:
	$(RM) *_gen.json
	
prebuild:
	$(GEN_ASSET_MANIFEST)
			
export:
	$(EXPORT_ASSETS)