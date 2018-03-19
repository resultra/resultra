# Template System Test

## Creating Templates to Test With

From the project root directory, go to ./build/dest and launch the single user version of the system: e.g.:

% ./bin/resultraLocalBackend --templates-path ./factoryTemplates --tracker-path ../../serverSystemTest/templateTest/simple/trackers

Then, launch the client in the web browser at localhost:43409. This will bring up the client to populate the template. 

After using the client to populate the client, the Makefile is setup to perform a test:

	1. Dump a summary of the tracker database in the original file.
	2. Compare the original summary with the gold file for the summary.
	3. Export the tracker database from the original to an exported database.
	4. Dump the summary from the exported database and compare with the gold file.