# This file is part of the Resultra project.
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.
DEPTH = ../..
include $(DEPTH)/build/common.mk
include $(DEPTH)/webui/build/webui.mk

realclean:
	$(RM) -Rf node_modules/

clean: 
	$(RM) *_gen.json

install:
	npm install
			
export:
	$(EXPORT_THIRD_PARTY_ASSETS) $(abspath ./thirdPartyAssets.json)
	$(EXPORT_THIRD_PARTY_ASSETS) $(abspath ./aceAssets.json)
	$(EXPORT_THIRD_PARTY_ASSETS) $(abspath ./dashboardAssets.json)
	
all: prebuild export
	