package uniqueID

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"log"
	"resultra/datasheet/server/generic/timestamp"
	"sync"
	"time"
)

const (
	numSequenceBits   uint64 = 12
	sequenceMask      int64  = (1 << numSequenceBits) - 1
	maxSequenceNumber int64  = (1 << numSequenceBits) - 1

	numNodeBits   uint64 = 10
	maxNodeNumber int64  = (1 << numNodeBits) - 1
	nodeShiftBits uint64 = numSequenceBits // 12
	nodeMask      int64  = maxNodeNumber << nodeShiftBits

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

type SnowflakeID struct {
	nodeID             int64
	sequenceNum        int64
	timeMsecSinceEpoch int64
	snowflakeID        int64
}

func (id SnowflakeID) encodeBase64() string {
	// Encode the integer ID into a byte array. Putvarint uses a variable length encoding for the
	// int64, so it could potentially need more (or less) than 8 bytes.
	varEncodeBytes := make([]byte, 16)
	varEncodeLen := binary.PutVarint(varEncodeBytes, id.snowflakeID)

	snowflakeIDBytes := make([]byte, varEncodeLen)
	copy(snowflakeIDBytes, varEncodeBytes)

	// Encode the ID as base64, making it suitable for use in URLs, or as a text
	// string to store in the database. RawURLEncoding will encode the string
	// without the padding characters ('=').
	idEncode := base64.RawURLEncoding.EncodeToString(snowflakeIDBytes)

	return idEncode

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
	currTime := timestamp.CurrentTimestampUTC()
	return timeToMillisecondsSinceEpoch(currTime)
}

func (gen SnowflakeIDGenerator) currSnowflakeID() SnowflakeID {

	timestampOffset := gen.currTimeMillis - startTimeMillisSinceEpoch

	snowflakeID := (timestampOffset << timeShiftBits) | (gen.nodeID << nodeShiftBits) | gen.currSequenceNum

	return SnowflakeID{
		timeMsecSinceEpoch: gen.currTimeMillis,
		nodeID:             gen.nodeID,
		sequenceNum:        gen.currSequenceNum,
		snowflakeID:        snowflakeID}
}

func DecodeFromBase64(encodedID string) (*SnowflakeID, error) {

	decodedBytes, decodeErr := base64.RawURLEncoding.DecodeString(encodedID)
	if decodeErr != nil {
		return nil, fmt.Errorf("SnowflakeID: DecodeFromBase64: Can't decode base64 encoded snowflake ID = %v: error = %v", encodedID, decodeErr)
	}

	snowflakeID, readErr := binary.ReadVarint(bytes.NewBuffer(decodedBytes))
	if readErr != nil {
		return nil, fmt.Errorf("SnowflakeID: DecodeFromBase64: Can't decode base64 encoded snowflake ID = %v: error = %v", encodedID, readErr)
	}

	timeSinceEpochMsec := (snowflakeID >> timeShiftBits) + startTimeMillisSinceEpoch
	sequenceNum := snowflakeID & sequenceMask
	nodeID := (snowflakeID & nodeMask) >> nodeShiftBits

	return &SnowflakeID{
		timeMsecSinceEpoch: timeSinceEpochMsec,
		nodeID:             nodeID,
		sequenceNum:        sequenceNum,
		snowflakeID:        snowflakeID}, nil
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

func (gen *SnowflakeIDGenerator) generateID() SnowflakeID {

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
	} else if currTimeMillis == gen.currTimeMillis {
		// The timestamp hasn't advanced since the last ID which was generated.
		// If the sequence numbers aren't exhausted for the current msec,
		// use the next sequence number. Otherwise, wait until the next msec,
		// and generate an ID with a reset sequence number and next mswec.
		nextSequenceNum := gen.currSequenceNum + 1
		if nextSequenceNum > maxSequenceNumber {
			// All the sequence numbers are used up for the current millisecond.
			// Wait until the next millisecond to generate a the next ID.
			gen.currSequenceNum = 0
			gen.currTimeMillis = nextTimeMillisSinceEpoch(currTimeMillis)

		} else {
			// nextSequenceNum is within the range of sequence numbers for
			// the current msec, so it can be used to generate an ID.
			gen.currSequenceNum = nextSequenceNum
		}
	} else { // currTimeMillis < gen.currTimeMillis
		// Shouldn't get here, since SnowflakeIDGenerator is initialized with the current timestamp
		// in NewSnowflakeIDGenerator(). Time is expected to "monotonically increase" from there.
		//
		// However, if there is clock skew with the system clock
		// to set it backwards, this case could potentially happen. This could potentially happen
		// if the system is re-synchronized via NTP.
		//
		// At this point, the only way to gracefully recover is to wait until the current system
		// clock advances past the time currently in gen.currTimeMillis. This is similar to the case
		// where all the sequence numbers are exhausted for the current msec.
		log.Printf("Warning: SnowflakeIDGenerator: system clock skew detected, "+
			"waiting until after last ID generation time = %v(%v) to generate new IDs",
			gen.currTimeMillis, millisSinceEpochToTime(gen.currTimeMillis))
		gen.currSequenceNum = 0
		gen.currTimeMillis = nextTimeMillisSinceEpoch(gen.currTimeMillis)

	}

	// After the code block above, the sequence number and/or timestamp on the generator will
	// have beeen advanced to the next unique ID. So, the next ID can be generated based upon
	// the generator's current snowflake ID paramater values.
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

	log.Printf("Snowflake unique ID generator intialized: "+
		"start time = %v (%v), curr time = %v (%v), expiration/max time = %v (%v), node ID = %v",
		startTimeMillisSinceEpoch, millisSinceEpochToTime(startTimeMillisSinceEpoch),
		currTimeMillis, millisSinceEpochToTime(currTimeMillis),
		maxTimestampMillisSinceEpoch(), millisSinceEpochToTime(maxTimestampMillisSinceEpoch()),
		nodeID)

	return &idGen, nil
}

var globalIDGen *SnowflakeIDGenerator

type UniqueIDGenFunc func() string

var globalUniqueIDFunc UniqueIDGenFunc

func init() {

	var err error
	globalIDGen, err = NewSnowflakeIDGenerator(1)
	if err != nil {
		log.Fatalf("Failure initializing ID generation for database objects: %v", err)
	}

	globalUniqueIDFunc = generateSnowflakeID
}

func generateSnowflakeID() string {
	return globalIDGen.generateID().encodeBase64()
}

func OverrideProductionUniqueIDFuncWithTestFunc(uniqueIDFunc UniqueIDGenFunc) {
	log.Printf("WARNING: Overriding global (snowflake) ID generator. Use *only* for development testing purposes")
	globalUniqueIDFunc = uniqueIDFunc
}

func GenerateUniqueID() string {
	return globalUniqueIDFunc()
}
