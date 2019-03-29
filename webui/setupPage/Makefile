# This file is part of the Resultra project.
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.
DEPTH = ../..
include $(DEPTH)/build/common.mk
include $(DEPTH)/webui/build/webui.mk

clean: 
	$(RM) *_gen.json
	
prebuild:
	$(AUTO_GEN_ASSET_MANIFEST)
	
		
export:
	$(EXPORT_ASSETS)
	$(BUILD_ASSET_INJECTION_LIST) adminSetupPageAssetInclude.json $(ASSET_BASE_PATH) > adminSetupPageAssetInclude_gen.json
	$(EXPORT_HTML_WITH_INJECTED_ASSETS) $(abspath ./adminSetupPageAssetInclude_gen.json)
	
	
all: prebuild export
