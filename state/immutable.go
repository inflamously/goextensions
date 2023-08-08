package state

type Immutable struct {
	value any
}

func (i *Immutable) Get() any {
	// TODO: Clone to avoid changing original value
	return i.value
}

func NewImmutable(value any) Immutable {
	return Immutable{
		value: value,
	}
}
