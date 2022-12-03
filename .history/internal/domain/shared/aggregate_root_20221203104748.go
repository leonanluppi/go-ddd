package domain

type (
	eventRecorder struct {
		events []interface{}
	}
	AggregateRoot struct {
		eventRecorder eventRecorder
	}
)

func (e *eventRecorder) Record(event interface{}) {
	e.events = append(e.events, event)
}

func (e *eventRecorder) Events() []interface{} {
	return e.events
}

func (e *eventRecorder) Clear() {
	e.events = []interface{}{}
}

func (root *AggregateRoot) AddEvent(event interface{}) {
	root.eventRecorder.Record(event)
}

func (root *AggregateRoot) Clear() {
	root.eventRecorder.Clear()
}

func (root *AggregateRoot) Events() []interface{} {
	return root.eventRecorder.Events()
}
