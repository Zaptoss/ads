package dynamicarray

import "fmt"

type DynamicArray struct {
	length   uint
	capacity uint
	array    []string
}

func NewDynamicArray() *DynamicArray {
	return &DynamicArray{capacity: 1, array: make([]string, 1)}
}

func (d *DynamicArray) Append(s string) {
	if d.length == d.capacity {
		nr := make([]string, d.capacity*2)
		for i := 0; i < int(d.length); i++ {
			nr[i] = d.array[i]
		}
		d.array = nr
		d.capacity = uint(cap(d.array))
	}

	d.array[d.length] = s
	d.length += 1
}

func (d *DynamicArray) Pop() (s string, err error) {
	if d.length == 0 {
		return "", fmt.Errorf("array is empty")
	}
	d.length -= 1
	s, d.array[d.length] = d.array[d.length], ""
	return
}

func (d *DynamicArray) InsertAfter(i uint, data string) error {
	if d.length == 0 {
		return fmt.Errorf("array is empty")
	}

	if i >= d.length {
		return fmt.Errorf("out of range")
	}

	if d.length == d.capacity {
		nr := make([]string, d.capacity*2)
		for i := 0; i < int(d.length); i++ {
			nr[i] = d.array[i]
		}
		d.array = nr
		d.capacity = uint(cap(d.array))
	}

	for j := d.length; j > i+1; j-- {
		d.array[j] = d.array[j-1]
	}

	d.array[i+1] = data
	d.length += 1
	return nil
}

func (d *DynamicArray) InsertBefore(i uint, data string) error {
	if d.length == 0 {
		return fmt.Errorf("array is empty")
	}

	if i >= d.length {
		return fmt.Errorf("out of range")
	}

	if d.length == d.capacity {
		nr := make([]string, d.capacity*2)
		for i := 0; i < int(d.length); i++ {
			nr[i] = d.array[i]
		}
		d.array = nr
		d.capacity = uint(cap(d.array))
	}

	for j := d.length; j > i; j-- {
		d.array[j] = d.array[j-1]
	}

	d.array[i] = data
	d.length += 1
	return nil
}

func (d *DynamicArray) Remove(i uint) error {
	if d.length == 0 {
		return fmt.Errorf("array is empty")
	}

	if i >= d.length {
		return fmt.Errorf("out of range")
	}

	for j := i; j < d.length-1; j++ {
		d.array[j] = d.array[j+1]
	}

	d.length -= 1
	d.array[d.length] = ""
	return nil
}

func (d *DynamicArray) String() string {
	return fmt.Sprint(d.array)
}

func (d *DynamicArray) Get(i uint) (string, error) {
	if d.length == 0 {
		return "", fmt.Errorf("array is empty")
	}

	if i >= d.length {
		return "", fmt.Errorf("out of range")
	}

	return d.array[i], nil
}

func (d *DynamicArray) Set(i uint, data string) error {
	if d.length == 0 {
		return fmt.Errorf("array is empty")
	}

	if i >= d.length {
		return fmt.Errorf("out of range")
	}

	d.array[i] = data
	return nil
}

func (d *DynamicArray) Length() uint {
	return d.length
}
