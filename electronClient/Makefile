# This file is part of the Resultra project.
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.
DEPTH = ..
include $(DEPTH)/build/common.mk

install:
	npm install

clean:
	$(RM_DIR) dist
	
realclean:
	$(RM_DIR) node_modules
	
package:
	npm run dist

winpkg:
	npm run dist -- --windows
	
all: install package
