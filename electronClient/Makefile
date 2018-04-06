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
