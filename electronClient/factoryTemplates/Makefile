DEPTH = ../..
include $(DEPTH)/build/common.mk

.PHONY:assembleFactoryTemplateDB
	
EXPORT_TEMPLATE=$(DEPTH)/build/dest/bin/exportTrackerTemplate
DEST_FACTORY_TEMPLATE_DB=trackers.db

all: export
	
clean:
	$(RM) $(DEST_FACTORY_TEMPLATE_DB)

# The factory template database is assembled from individual trackers. If a new factory templates needs to be
# added, it can be merged/assembled into the factory template database using the commands below. 
assembleFactoryTemplateDB:
	$(RM) $(DEST_FACTORY_TEMPLATE_DB)
	$(EXPORT_TEMPLATE) --sourcedb ./lightWeightCRM.db --source-tracker-db-id gMCAnMPex7wG --destdb $(DEST_FACTORY_TEMPLATE_DB)
	$(EXPORT_TEMPLATE) --sourcedb ./contentMarketing.db --source-tracker-db-id gMCA7MDYybwG --destdb $(DEST_FACTORY_TEMPLATE_DB)
	$(EXPORT_TEMPLATE) --sourcedb ./scrum.db --source-tracker-db-id gMCAwJrW0LwG --destdb $(DEST_FACTORY_TEMPLATE_DB)
	$(EXPORT_TEMPLATE) --sourcedb ./starter.db --source-tracker-db-id gMCA5Lrm07wG --destdb $(DEST_FACTORY_TEMPLATE_DB)


# Export the factory templates to the package templates build directory.
# When packaging the Electron app, these templates are copied from the build
# directory. When running the Electron app in development/debug mode, the
# templates are referenced directly from the build directory.	
export: assembleFactoryTemplateDB
	$(INSTALL_FILE) $(DEST_FACTORY_TEMPLATE_DB) $(PKG_TEMPLATES_DIR)
