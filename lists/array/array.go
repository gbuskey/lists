package array

import "fmt"

type Array struct {
	data []any
}

func New() *Array {
	return &Array{
		data: []any{},
	}
}

func (a *Array) Add(element any) error {
	a.data = append(a.data, element)
	return nil
}

func (a *Array) Remove(element any) (any, error) {
	index, found, err := a.find(element)
	if err != nil {
		return nil, fmt.Errorf("unable to remove element. %w", err)
	}

	a.data[index] = a.data[a.Length() - 1]
	a.data = a.data[:a.Length() - 1]
	return found, nil
}

func (a *Array) Find(element any) (any, error) {
	_, e, err := a.find(element)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (a *Array) find(element any) (int, any, error) {
	for i, e := range a.data {
		if e == element {
			return i, e, nil
		}
	}
	return 0, nil, fmt.Errorf("unable to find element '%s'", element)
}

func (a *Array) Length() (int) {
	return len(a.data)
}
