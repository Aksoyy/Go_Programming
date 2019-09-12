type Set interface {
	Add(elements ...interface{})
	Remove(elements ...interface{})
	Contains(elements ...interface{}) bool

	containers.Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}