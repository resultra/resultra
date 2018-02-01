package dashboardController

import (
	"resultra/datasheet/server/dashboard/values"
	"time"
)

type timeIncrementIterFunc func(currTime time.Time) time.Time

func prevMonthEnd(currMonthTime time.Time) time.Time {

	firstOfMonth := time.Date(currMonthTime.Year(), currMonthTime.Month(),
		1, 0, 0, 0, 0, currMonthTime.Location())
	lastOfMonth := firstOfMonth.Add(-1 * time.Nanosecond)

	return lastOfMonth

}

func prevDayEnd(currMonthTime time.Time) time.Time {

	startOfDay := time.Date(currMonthTime.Year(), currMonthTime.Month(), currMonthTime.Day(),
		0, 0, 0, 0, currMonthTime.Location())
	lastDay := startOfDay.Add(-1 * time.Nanosecond)

	return lastDay

}

func prevWeekEnd(currMonthTime time.Time) time.Time {

	startOfDay := time.Date(currMonthTime.Year(), currMonthTime.Month(), currMonthTime.Day(),
		0, 0, 0, 0, currMonthTime.Location())

	// Back up to the most recent Sunday
	currDay := startOfDay

	for currDay.Weekday() != time.Sunday {
		currDay = currDay.AddDate(0, 0, -1)
	}

	// Back up to the last nanosecon on Saturday
	lastDay := currDay.Add(-1 * time.Nanosecond)

	return lastDay

}

const timeIntervalDays string = "days"
const timeIntervalWeeks string = "weeks"
const timeIntervalMonths string = "months"

var timeIntervalIncrementFuncs = map[string]timeIncrementIterFunc{
	timeIntervalDays:   prevDayEnd,
	timeIntervalWeeks:  prevWeekEnd,
	timeIntervalMonths: prevMonthEnd}

func generateTimeIncrements(startTime time.Time, beginningTime time.Time, incrementIterFunc timeIncrementIterFunc) []time.Time {

	increments := []time.Time{}

	// Always have at least one increment
	firstIncrementTime := incrementIterFunc(startTime)
	increments = append(increments, firstIncrementTime)

	currIncrTime := incrementIterFunc(firstIncrementTime)
	for currIncrTime.After(beginningTime) {
		increments = append(increments, currIncrTime)
		currIncrTime = incrementIterFunc(currIncrTime)
	}

	// The code above returns the increments in reverse chronological order.
	// However, the client of this function expects the increments in chronological/reverse order.
	chronologicalIncrements := []time.Time{}
	for i := len(increments) - 1; i >= 0; i-- {
		chronologicalIncrements = append(chronologicalIncrements, increments[i])
	}

	return chronologicalIncrements

}

const timeRangeLast6Months string = "last6Months"
const timeRangeLast3Months string = "last3Months"
const timeRangeLastWeek string = "lastWeek"
const timeRangeLast2Weeks string = "last2Weeks"
const timeRangeLastMonth string = "lastMonth"
const timeRangeLastYear string = "lastYear"
const timeRangeYTD string = "YTD"

func generateTimeRange(timeRangeLabel string, startTime time.Time) time.Time {

	switch timeRangeLabel {
	case timeRangeLastWeek:
		return startTime.AddDate(0, 0, -7)
	case timeRangeLast2Weeks:
		return startTime.AddDate(0, 0, -14)
	case timeRangeLastMonth:
		return startTime.AddDate(0, -1, 0)
	case timeRangeLast3Months:
		return startTime.AddDate(0, -3, 0)
	case timeRangeLast6Months:
		return startTime.AddDate(0, -6, 0)
	case timeRangeYTD:
		startOfYear := time.Date(startTime.Year(), 1, 1, 0, 0, 0, 0, startTime.Location())
		return startOfYear
	case timeRangeLastYear:
		return startTime.AddDate(-1, 0, 0)
	default:
		return startTime.AddDate(0, 0, -7) // minimum
	}

}

func generateTimeIncrementsForValGrouping(valGrouping values.ValGrouping) []time.Time {

	timeRange := timeRangeLastWeek
	if valGrouping.TimeRange != nil {
		timeRange = *valGrouping.TimeRange
	}

	currTime := time.Now().UTC()
	startRangeTime := generateTimeRange(timeRange, currTime)

	timeInterval := timeIntervalWeeks
	if valGrouping.GroupValsByTimeIncrement != nil {
		timeInterval = *valGrouping.GroupValsByTimeIncrement
	}
	timeIncrFunc, funcFound := timeIntervalIncrementFuncs[timeInterval]
	if !funcFound {
		timeIncrFunc = prevWeekEnd
	}

	return generateTimeIncrements(currTime, startRangeTime, timeIncrFunc)

}
