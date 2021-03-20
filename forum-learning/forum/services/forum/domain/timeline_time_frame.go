package domain

import "fmt"

// TimelineTimeFrame is an Object to handle Timeline TimeFrame format
type TimelineTimeFrame struct {
	allowedValue []string
	value        string
}

// NewTimelineTimeFrame is an Contructor for TimelineTimeFrame Object
func NewTimelineTimeFrame(timeframe string) (*TimelineTimeFrame, error) {

	timelineTimeFrameObj := &TimelineTimeFrame{
		allowedValue: []string{"year", "month", "week"},
		value:        timeframe,
	}

	if !timelineTimeFrameObj.validateTimelineTimeFrame() {
		return nil, fmt.Errorf("Timeline TimeFrame Invalid")
	}

	return timelineTimeFrameObj, nil

}

// GetValue is a Getter Function for Value
func (obj *TimelineTimeFrame) GetValue() string {
	return obj.value
}

func (obj *TimelineTimeFrame) validateTimelineTimeFrame() bool {

	for i := 0; i < len(obj.allowedValue); i++ {
		if obj.value == obj.allowedValue[i] {
			return true
		}
	}

	return false

}
