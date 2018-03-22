DEPTH=../..
include $(DEPTH)/build/common.mk

.PHONY:simple blogTopic

all: systest

systest: simple blogTopic
	
clean:
	$(RM) *_gen.db *_dump.txt
	
simple:
	sqlite3 ./simple/trackers/trackers.db < dumpTrackerDBInfo.sql > simple_dump.txt
	diff simple_dump.txt simple.gold
	$(RM) simple_export_gen.db
	$(DEPTH)/build/dest/bin/exportTrackerTemplate --test-ids --sourcedb ./simple/trackers/trackers.db \
		--source-tracker-db-id gMCApJHm1aYG --destdb simple_export_gen.db
	sqlite3 simple_export_gen.db < dumpTrackerDBInfo.sql > simple_export_dump.txt
	diff simple_export_dump.txt simple.gold
	sqlite3 simple_export_gen.db < dumpRawDBInfo.sql > simple_raw_dump.txt
	diff simple_raw_dump.txt simple_raw.gold
	
blogTopic:
	sqlite3 ./blogTopic/trackers/trackers.db < dumpTrackerDBInfo.sql > blogTopic_dump.txt
	diff blogTopic_dump.txt blogTopic.gold
	$(RM) blogTopic_export_gen.db
	$(DEPTH)/build/dest/bin/exportTrackerTemplate --test-ids --sourcedb ./blogTopic/trackers/trackers.db \
		--source-tracker-db-id gMCA4PjEp6oG --destdb blogTopic_export_gen.db
	sqlite3  blogTopic_export_gen.db < dumpTrackerDBInfo.sql > blogTopic_export_dump.txt
	diff blogTopic_export_dump.txt blogTopic.gold
	sqlite3 blogTopic_export_gen.db < dumpRawDBInfo.sql > blogTopic_raw_dump.txt
	diff blogTopic_raw_dump.txt blogTopic_raw.gold
	