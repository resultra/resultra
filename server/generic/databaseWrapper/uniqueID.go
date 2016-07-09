package databaseWrapper

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"log"
	"sync"
	"time"
)

const (
	numNodeBits   uint64 = 10
	maxNodeNumber int64  = (1 << numNodeBits) - 1

	numSequenceBits   uint64 = 12
	sequenceMask      uint64 = (1 << numSequenceBits) - 1
	maxSequenceNumber int64  = (1 << numSequenceBits) - 1

	nodeShiftBits uint64 = numSequenceBits // 12

	timeShiftBits uint64 = numNodeBits + numSequenceBits // 22

	// July 1st, 2016 is used as a start time for ID generation
	// The ID is encoded as the difference between the current time and this time.
	startTimeMillisSinceEpoch int64 = 1467356400000
)

type SnowflakeIDGenerator struct {
	sync.Mutex
	nodeID          int64
	currSequenceNum int64
	currTimeMillis  int64
}

func millisSinceEpochToTime(millis int64) time.Time {
	return time.Unix(0, millis*int64(time.Millisecond)).UTC()
}

func maxTimestampMillisSinceEpoch() int64 {

	var maxTimestamp int64
	var maxTimestampOffset int64

	maxTimestampOffset = (1 << 41) - 1

	maxTimestamp = startTimeMillisSinceEpoch + maxTimestampOffset

	return maxTimestamp
}

func timeToMillisecondsSinceEpoch(t time.Time) int64 {
	nanos := t.UnixNano()
	millis := nanos / 1000000 // 1 nano-second = 1e6 milliseconds
	return millis
}

func timeNowMillisSinceEpoch() int64 {
	currTime := time.Now().UTC()
	return timeToMillisecondsSinceEpoch(currTime)
}

func nextTimeMillisSinceEpoch(currTimeMillis int64) int64 {
	nextTimeMillis := timeNowMillisSinceEpoch()
	for nextTimeMillis <= currTimeMillis {
		// time.Sleep() takes the number of nano-seconds to sleep. Since there are 10e6 nanoseconds
		// in a millisecond, pass 10e5 to the sleep function to sleep for 1/10 of a millisecond.
		// This keeps the ID generation from burning too many CPU cycles just waiting for the next
		// millisecond.
		time.Sleep(100000) // sleep 1/10 of a millisecond
		nextTimeMillis = timeNowMillisSinceEpoch()
	}
	return nextTimeMillis
}

func (gen SnowflakeIDGenerator) currSnowflakeID() string {

	timestampOffset := gen.currTimeMillis - startTimeMillisSinceEpoch

	snowflakeID := (timestampOffset << timeShiftBits) | (gen.nodeID << nodeShiftBits) | gen.currSequenceNum

	// Encode the integer ID into a byte array
	snowflakeIDBytes := make([]byte, 8, 8)
	binary.PutVarint(snowflakeIDBytes, snowflakeID)

	// Encode the ID as base64, making it suitable for use in URLs, or as a text
	// string to store in the database. RawURLEncoding will encode the string
	// without the padding characters ('=').
	idEncode := base64.RawURLEncoding.EncodeToString(snowflakeIDBytes)

	return idEncode
}

func (gen *SnowflakeIDGenerator) generateID() string {
	gen.Lock()

	currTimeMillis := timeNowMillisSinceEpoch()

	// Advance the timestamp and/or sequence number on the
	// generator.
	if currTimeMillis > gen.currTimeMillis {
		// Time has advanced beyond the last time used for ID generation,
		// so advance the generator to use the current time and reset
		// its sequence number.
		gen.currTimeMillis = currTimeMillis
		gen.currSequenceNum = 0
	} else {
		// The timestamp hasn't advanced since the last ID which was generated.
		// currTimeMillis <= gen.currTimeMillis.
		nextSequenceNum := gen.currSequenceNum + 1
		if nextSequenceNum > maxSequenceNumber {
			// All the sequence numbers are used up for the current millisecond.
			// Wait until the next millisecond to generate a the next ID.
			gen.currSequenceNum = 0
			gen.currTimeMillis = nextTimeMillisSinceEpoch(currTimeMillis)

		} else {
			gen.currSequenceNum = nextSequenceNum
		}
	}

	snowflakeID := gen.currSnowflakeID()

	gen.Unlock()

	return snowflakeID
}

func NewSnowflakeIDGenerator(nodeID int64) (*SnowflakeIDGenerator, error) {

	if nodeID < 0 || nodeID > maxNodeNumber {
		return nil, fmt.Errorf("NewSnowflakeIDGenerator: Invalid node ID for ID = %v for ID generation. nodeID must be between 0 and %v",
			nodeID, maxNodeNumber)
	}

	// Confirm the current clock is after the start time for the ID generation.
	currTimeMillis := timeNowMillisSinceEpoch()
	if currTimeMillis < startTimeMillisSinceEpoch {
		return nil, fmt.Errorf("NewSnowflakeIDGenerator: ID generation is misconfigured: current timestamp = %v is before start timestamp = %v",
			currTimeMillis, startTimeMillisSinceEpoch)
	}

	idGen := SnowflakeIDGenerator{
		nodeID:          nodeID,
		currSequenceNum: 0,
		currTimeMillis:  timeNowMillisSinceEpoch()}

	log.Printf("Snowflake database ID generator intialized: "+
		"start time = %v (%v), curr time = %v (%v), expiration/max time = %v (%v), node ID = %v",
		startTimeMillisSinceEpoch, millisSinceEpochToTime(startTimeMillisSinceEpoch),
		currTimeMillis, millisSinceEpochToTime(currTimeMillis),
		maxTimestampMillisSinceEpoch(), millisSinceEpochToTime(maxTimestampMillisSinceEpoch()),
		nodeID)

	return &idGen, nil
}

var globalIDGen *SnowflakeIDGenerator

func init() {

	var err error
	globalIDGen, err = NewSnowflakeIDGenerator(1)
	if err != nil {
		log.Fatalf("Failure initializing ID generation for database objects: %v", err)
	}
}

func GlobalUniqueID() string {
	return globalIDGen.generateID()
}
