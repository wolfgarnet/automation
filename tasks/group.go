package tasks
import "wolfgarnet/automation/system"

// TimedGroup is repeated sequentially for a given duration
type TimedGroup struct {
	duration int32
}

// Group is repeated for a number of times
type Group struct {
	number int32
}

// TinedIntervalGroup is repeated every interval in parallel for a give n duration
type TimedIntervalGroup struct {
	duration int32
	interval int32
}

// IntervalGroup is repeated number of times every interval
type IntervalGroup struct {
	interval int32
	number int32
}

func NewGroup() (system.Task, error) {

}