// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package uniqueID

import (
	"strconv"
	"testing"
	"time"
)

// TestSnowflakeIDBase verifies all the masking and shifting parameters which go into formulating a snowflake ID.
func TestSnowflakeIDBase(t *testing.T) {

	nodeIDMaskBase2 := strconv.FormatInt(nodeMask, 2)
	t.Logf("     Node ID mask:%v", nodeIDMaskBase2)
	if nodeIDMaskBase2 != `1111111111000000000000` { // 12 right-most bits for sequence, then 10 left-most bits for node ID
		t.Errorf("Invalid node ID mask %v", nodeIDMaskBase2)
	}

	seqMaskBase2 := strconv.FormatInt(sequenceMask, 2)
	t.Logf("Sequence num mask:          %v", seqMaskBase2)
	if seqMaskBase2 != `111111111111` { // 12 right-most bits
		t.Errorf("Invalid sequence num mask %v", seqMaskBase2)
	}

	base2MaxNodeNodeNum := strconv.FormatInt(maxNodeNumber, 2)
	t.Logf("Max node number: %v (base 2 = %v)", maxNodeNumber, base2MaxNodeNodeNum)
	if base2MaxNodeNodeNum != `1111111111` { // 10 bits for node ID
		t.Errorf("Invalid max node number %v", base2MaxNodeNodeNum)
	}

	maxSeqNumBase2 := strconv.FormatInt(maxSequenceNumber, 2)
	t.Logf("Max sequence number: %v (base 2 = %v)", maxSequenceNumber, maxSeqNumBase2)
	if maxSeqNumBase2 != `111111111111` { // 12 right-most bits
		t.Errorf("Invalid sequence num mask %v", maxSeqNumBase2)
	}

	maxMsecSinceEpoch := maxTimestampMillisSinceEpoch()
	t.Logf("Max timestamp (msec since epoch): %v (base 2 = %v)", maxMsecSinceEpoch, strconv.FormatInt(maxMsecSinceEpoch, 2))
	t.Logf("Max timestamp date: %v ", millisSinceEpochToTime(maxMsecSinceEpoch))
	maxOffsetVsStart := maxMsecSinceEpoch - startTimeMillisSinceEpoch // time encoded vs start time - 41 bits
	maxOffsetVsStartBase2 := strconv.FormatInt(maxOffsetVsStart, 2)
	t.Logf("Max timestamp - offset vs start time: %v (base 2 = %v)", maxOffsetVsStart, maxOffsetVsStartBase2)

	if maxOffsetVsStartBase2 != `11111111111111111111111111111111111111111` {
		t.Errorf("Invalid maximum ID generation time %v", millisSinceEpochToTime(maxMsecSinceEpoch))
	}
	if len(maxOffsetVsStartBase2) != 41 {
		t.Errorf("Invalid maximum ID generation time %v", millisSinceEpochToTime(maxMsecSinceEpoch))
	}

}

// TestSnowflakeIDConv is a helper function for TestSnowflakeIDConv(). It test one ID generation and compares
// against expected results.
func verifyOneSnowflakeIDGen(t *testing.T, msecSinceEpoch int64, nodeID int64, seqNum int64, expectedBase2 string) {

	idGen := SnowflakeIDGenerator{
		nodeID:          nodeID,
		currSequenceNum: seqNum,
		currTimeMillis:  msecSinceEpoch}

	id := idGen.currSnowflakeID()
	base64ID := id.encodeBase64()
	base2ID := strconv.FormatInt(id.snowflakeID, 2)
	t.Logf("verifyOneSnowflakeIDGen: %+v -- generates --> %v (base64), %v (base 2)", idGen, base64ID, base2ID)

	if base2ID != expectedBase2 {
		t.Errorf("verifyOneSnowflakeIDGen: Binary version of snowflake ID doesn't match expected: got = %v, expected = %v ",
			base2ID, expectedBase2)
	}

}

// TestSnowflakeIDConv tests the boundary cases for ID generation
func TestSnowflakeIDConv(t *testing.T) {

	minSeqNumBase2 := `000000000000`
	maxSeqNumBase2 := `111111111111`

	minNodeIDBase2 := `0000000000`
	maxNodeIDBase2 := `1111111111`

	maxTimeBase2 := `11111111111111111111111111111111111111111`
	maxMSecTimestamp := maxTimestampMillisSinceEpoch()

	// Everything should be 0 at the start
	verifyOneSnowflakeIDGen(t, startTimeMillisSinceEpoch, 0, 0, `0`)

	// When only maximum sequence number set, only the 12 bit sequence number is set
	verifyOneSnowflakeIDGen(t, startTimeMillisSinceEpoch, 0, maxSequenceNumber, maxSeqNumBase2)

	// When only maximum node ID set, only the 10 bit node ID is set (sequence number in right-most bits is 0)
	verifyOneSnowflakeIDGen(t, startTimeMillisSinceEpoch, maxNodeNumber, 0, maxNodeIDBase2+minSeqNumBase2)

	// Time and nodeID maxed out, minimum sequence number
	verifyOneSnowflakeIDGen(t, maxMSecTimestamp, maxNodeNumber, 0, maxTimeBase2+maxNodeIDBase2+minSeqNumBase2)

	// Time and sequence number maxed out, minimum nodeID
	verifyOneSnowflakeIDGen(t, maxMSecTimestamp, 0, maxSequenceNumber, maxTimeBase2+minNodeIDBase2+maxSeqNumBase2)

	// Only time maxed out
	verifyOneSnowflakeIDGen(t, maxMSecTimestamp, 0, 0,
		maxTimeBase2+minNodeIDBase2+minSeqNumBase2)

	// Everything maxed out
	verifyOneSnowflakeIDGen(t, maxMSecTimestamp, maxNodeNumber, maxSequenceNumber,
		maxTimeBase2+maxNodeIDBase2+maxSeqNumBase2)

}

// verifyRoundTripDecode is a helper function to verify encoding and decoding an ID from base64 yields the
// same ID.
func verifyRoundTripDecode(t *testing.T, generatedID SnowflakeID) {

	encodedID := generatedID.encodeBase64()

	decodedID, decodeErr := DecodeFromBase64(encodedID)
	if decodeErr != nil {
		t.Errorf("verifyRoundTripDecode: decode ID error: %v", decodeErr)
	}

	if *decodedID != generatedID {
		t.Errorf("verifyRoundTripDecode: decode ID error: encoded = %+v not equal decoded = %+v", generatedID, decodedID)
	}

}

func TestSnowflakeIDGenSimple(t *testing.T) {
	var nodeID int64 = 12
	idGen, err := NewSnowflakeIDGenerator(nodeID)

	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(2 * time.Millisecond)

	id1 := idGen.generateID()
	id2 := idGen.generateID()
	id3 := idGen.generateID()

	verifyRoundTripDecode(t, id1)

	t.Logf("id2 = %v, %+v", id2.encodeBase64(), id2)
	t.Logf("id3 = %v, %+v", id3.encodeBase64(), id3)

	time.Sleep(2 * time.Millisecond)
	id4 := idGen.generateID()
	t.Logf("id4 = %v, %+v", id1.encodeBase64(), id4)
}

func TestSnowflakeIDGenAdvanceSeqNumber(t *testing.T) {
	var nodeID int64 = 12

	idGen, err := NewSnowflakeIDGenerator(nodeID)
	if err != nil {
		t.Fatal(err)
	}

	// Set the sequence number to its maximum value. The next sequence number must be 0
	idGen.currSequenceNum = maxSequenceNumber
	timestampBeforeIDGen := idGen.currTimeMillis

	idWithResetSeqNum := idGen.generateID()
	t.Logf("idWithResetSeqNum = %v, %+v", idWithResetSeqNum.encodeBase64(), idWithResetSeqNum)
	if idWithResetSeqNum.sequenceNum != 0 {
		t.Errorf("TestSnowflakeIDGenAdvanceSeqNumber: Invalid sequence number for ID detected, expecting reset after max sequence number")
	}
	if idWithResetSeqNum.timeMsecSinceEpoch <= timestampBeforeIDGen {
		t.Errorf("TestSnowflakeIDGenAdvanceSeqNumber: Invalid timestamp for ID detected, must wait until next msec if sequence number exhaustedw")
	}

}

func TestSnowflakeIDGenClockSkew(t *testing.T) {
	var nodeID int64 = 12

	idGen, err := NewSnowflakeIDGenerator(nodeID)
	if err != nil {
		t.Fatal(err)
	}

	// Push the current "ID generation time forward". This simulates a scenario
	// where the system clock goes backward vs the time last used to generate IDs.
	// Since the timestamp used to generate IDs must always be increasing, the generator
	// must wait until after the last timestamp used to generate a new ID.
	idGen.currTimeMillis += 1000
	timestampBeforeIDGen := idGen.currTimeMillis

	id1 := idGen.generateID()
	t.Logf("id1 = %v, %+v", id1.encodeBase64(), id1)
	if id1.timeMsecSinceEpoch <= timestampBeforeIDGen {
		t.Errorf("TestSnowflakeIDGenClockSkew: Invalid timestamp for ID detected, must increase vs last timstamp before clock skew")
	}

}

func TestSnowflakeIDGenStressTest(t *testing.T) {

	var nodeID int64 = 12
	idGen, err := NewSnowflakeIDGenerator(nodeID)
	if err != nil {
		t.Fatal(err)
	}

	// Generate 1M IDs ... they all must be unique, and a round-trip encode-decode
	// must be successful.
	uniqueIDMap := map[int64]bool{}
	for i := 0; i < 1000000; i++ {
		uniqueID := idGen.generateID()
		snowflakeID := uniqueID.snowflakeID
		if _, idFound := uniqueIDMap[snowflakeID]; idFound == true {
			t.Errorf("Duplicate ID found %+v", uniqueID)
		}
		uniqueIDMap[snowflakeID] = true
		verifyRoundTripDecode(t, uniqueID)
		if i > 10000 {
			// Let the first 10K go full-speed, then slow it down. This will test the logic to
			// generate at most 1023 sequence numbers per msec, then wait until the next msec.
			time.Sleep(200)
		}
	}

}
