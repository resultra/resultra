DEPTH = ..
include $(DEPTH)/build/common.mk

prebuild:
	npm install

clean:
	$(RM_DIR) dist
	
realclean: clean
	$(RM_DIR) node_modules
	
package:
	npm run dist
	
all: prebuild package