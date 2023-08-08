package state

type Cloneable[T any] interface {
	Clone() T
}

func ShallowCopy[T any](source *T) *T {
	newAllocatedSource := *source
	return &newAllocatedSource
}
