package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	fmt.Println(Tokens)
	fmt.Println("len=>", Length)
}

func TestIdToString(t *testing.T) {
	id := 72
	exceptValue := "1a"
	assert.Equal(t, exceptValue, IdToString(id))
}

func TestStringToId(t *testing.T) {
	str := "1a"
	exceptValue := 72
	assert.Equal(t, exceptValue, StringToId(str))
}
