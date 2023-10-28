package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := NewStack(128)

	s.Push(1)
	s.Push(2)

	value := s.Pop()

	assert.Equal(t, value, 2)

	value2 := s.Pop()

	assert.Equal(t, value2, 1)
}

func TestStackBytes(t *testing.T) {
	s := NewStack(128)
	s.Push(2)
	s.Push(0x61)
	s.Push(0x61)
}

func TestVM(t *testing.T) {
	// 1 + 2 = 3
	// 1
	// push stack
	// 2
	// push stack
	// add
	// 3
	// push stack
	state := NewState()

	data := []byte{0x03, 0x0a, 0x02, 0x0a, 0x0b}
	vm := NewVM(data, state)
	err := vm.Run()

	assert.Nil(t, err)

	// fmt.Println(vm.stack.data[:6])
	assert.Equal(t, int(5), vm.stack.Pop())
}

func TestVMPack(t *testing.T) {
	state := NewState()

	data := []byte{0x61, 0x0c, 0x61, 0x0c, 0x61, 0x0c, 0x03, 0x0a, 0x0d}
	vm := NewVM(data, state)
	err := vm.Run()

	assert.Nil(t, err)

	// fmt.Println(vm.stack.data[:6])
	assert.Equal(t, []byte{0x61, 0x61, 0x61}, vm.stack.Pop())
}

func TestVMStore(t *testing.T) {
	state := NewState()

	// pack FOO in bytes array -> key then push 3 as value then store it
	data := []byte{0x03, 0x0a, 0x4f, 0x0c, 0x4f, 0x0c, 0x46, 0x0c, 0x03, 0x0a, 0x0d, 0x0f}
	vm := NewVM(data, state)

	err := vm.Run()
	assert.Nil(t, err)

	v, err := state.Get([]byte("FOO"))
	assert.Nil(t, err)

	desValue := deserializedInt64(v)
	assert.Equal(t, int64(3), desValue)
}

func TestVM2(t *testing.T) {
	state := NewState()

	data := []byte{0x02, 0x0a, 0x03, 0x0a, 0x0b, 0x4f, 0x0c, 0x4f, 0x0c, 0x46, 0x0c, 0x03, 0x0a, 0x0d, 0x0f}
	dataMore := []byte{0x02, 0x0a, 0x03, 0x0a, 0x0b, 0x4d, 0x0c, 0x4f, 0x0c, 0x46, 0x0c, 0x03, 0x0a, 0x0d, 0x0f}
	data = append(data, dataMore...)

	vm := NewVM(data, state)

	err := vm.Run()
	assert.Nil(t, err)

	// fmt.Printf("%+v \n", vm.stack.data[:10])
	fmt.Printf("%+v \n", state.data)

	v, err := state.Get([]byte("FOO"))
	assert.Nil(t, err)

	desValue := deserializedInt64(v)
	assert.Equal(t, int64(5), desValue)
}

func TestVMGetFromState(t *testing.T) {
	pushFoo := []byte{0x4f, 0x0c, 0x4f, 0x0c, 0x46, 0x0c, 0x03, 0x0a, 0x0d}
	data := []byte{0x02, 0x0a, 0x03, 0x0a, 0x0b, 0x4f, 0x0c, 0x4f, 0x0c, 0x46, 0x0c, 0x03, 0x0a, 0x0d, 0x0f}

	data = append(data, pushFoo...)
	data = append(data, 0xae)

	state := NewState()
	vm := NewVM(data, state)

	err := vm.Run()
	assert.Nil(t, err)

	// fmt.Printf("%+v \n", vm.stack.data[:10])
	// fmt.Printf("%+v \n", state.data)

	value := vm.stack.Pop().([]byte)
	valueDeserialized := deserializedInt64(value)

	assert.Equal(t, valueDeserialized, int64(5))
}

func TestVMMul(t *testing.T) {
	data := []byte{0x02, 0x0a, 0x02, 0x0a, 0xea}

	state := NewState()
	vm := NewVM(data, state)

	err := vm.Run()
	assert.Nil(t, err)

	result := vm.stack.Pop()
	assert.Equal(t, 4, result)

	// fmt.Printf("%+v\n", vm.stack.data[:10])
}

func TestVMDiv(t *testing.T) {
	data := []byte{0x04, 0x0a, 0x02, 0x0a, 0xfd}

	state := NewState()
	vm := NewVM(data, state)

	err := vm.Run()
	assert.Nil(t, err)

	result := vm.stack.Pop()
	assert.Equal(t, 2, result)

	// fmt.Printf("%+v\n", vm.stack.data[:10])
}
