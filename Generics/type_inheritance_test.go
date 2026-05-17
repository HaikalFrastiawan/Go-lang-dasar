package golang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](parameter T) string {
	return parameter.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m MyManager) GetName() string {
	return m.Name
}

func (m MyManager) GetManagerName() string {
	return "Manager of " + m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (m *MyVicePresident) GetName() string {
	return m.Name
}

func (m *MyVicePresident) GetVicePresidentName() string {
	return "Vice President of " + m.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Alice", GetName[MyManager](MyManager{Name: "Alice"}))
	assert.Equal(t, "Bob", GetName[*MyVicePresident](&MyVicePresident{Name: "Bob"}))
}