DEPTH = ..
include $(DEPTH)/build/common.mk

prebuild:
	npm install

clean:
	$(RM_DIR) dist
	
realclean: clean
	$(RM_DIR) node_modules
	
	
# Export the factory templates to the package templates build directory.
# When packaging the Electron app, these templates are copied from the build
# directory. When running the Electron app in development/debug mode, the
# templates are referenced directly from the build directory.	
export:
	$(INSTALL_FILE) factoryTemplates/trackers.db $(PKG_TEMPLATES_DIR)

	
package:
	npm run dist
	
all: prebuild package