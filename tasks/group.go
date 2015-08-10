package tasks
import "wolfgarnet/automation/system"

// TimedGroup is repeated sequentially for a given duration
type TimedGroup struct {
	duration uint32
}

// Group is repeated for a number of times
type Group struct {
	number uint32
}

// TinedIntervalGroup is repeated every interval in parallel for a give n duration
type TimedIntervalGroup struct {
	duration uint32
	interval uint32
}

// IntervalGroup is repeated number of times every interval
type IntervalGroup struct {
	interval uint32
	number uint32
}

func NewGroup(config map[string]interface{}) (system.Task, error) {

}