DEPTH=../..
include $(DEPTH)/build/common.mk

.PHONY:simple

all: systest

systest: simple
	
clean:
	$(RM) *_gen.db *_dump.txt
	
simple:
	sqlite3 ./simple/trackers/trackers.db < dumpTrackerDBInfo.sql > simple_dump.txt
	diff simple_dump.txt simple.gold
	$(RM) simple_export_gen.db
	$(DEPTH)/build/dest/bin/exportTrackerTemplate --sourcedb ./simple/trackers/trackers.db \
		--source-tracker-db-id gMCApJHm1aYG --destdb simple_export_gen.db
	sqlite3 simple_export_gen.db < dumpTrackerDBInfo.sql > simple_export_dump.txt
	diff simple_export_dump.txt simple.gold
	