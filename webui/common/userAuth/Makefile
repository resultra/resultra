DEPTH = ../../..
include $(DEPTH)/build/common.mk
include $(DEPTH)/webui/build/webui.mk

all: prebuild build package
	
# This pacakge includes a hand-written AssetManifest.json, since the HTML template for the sign in page, signInPage.html,
# needs to be excluded from the standard $(EXPORT_ASSETS) rule. 
# 
# Other pages which include the assets from this package need signInPage.html with the assets injected. However, 
# if this packages's asset manifest includes signInPage.html, other packages which include this package will 
# erroneously (re)export signInPage.html without its own asset list injected into the file.
	
clean: 
	$(RM) *_gen.json
	
prebuild:
	$(GEN_ASSET_MANIFEST)

build:
	$(BUILD_ASSET_INJECTION_LIST) signinPageAssetInclude.json $(ASSET_BASE_PATH) > signinPageAssetInclude_gen.json
	
		
package:
	$(EXPORT_ASSETS)
	$(EXPORT_HTML_WITH_INJECTED_ASSETS) $(abspath ./signinPageAssetInclude_gen.json)