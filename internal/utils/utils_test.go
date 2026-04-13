package utils_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/protoenum/internal/utils"
)

func TestGetValuePointer(t *testing.T) {
	v := 42
	p := utils.GetValuePointer(v)
	require.NotNil(t, p)
	require.Equal(t, 42, *p)

	// Changing the pointer should not affect the source value
	*p = 99
	require.Equal(t, 42, v)
}

func TestGetValuePointer_String(t *testing.T) {
	s := "hello"
	p := utils.GetValuePointer(s)
	require.NotNil(t, p)
	require.Equal(t, "hello", *p)
}

func TestGetPointerValue(t *testing.T) {
	v := 42
	result := utils.GetPointerValue(&v)
	require.Equal(t, 42, result)
}

func TestGetPointerValue_Nil(t *testing.T) {
	var p *int
	result := utils.GetPointerValue(p)
	require.Equal(t, 0, result)
}

func TestGetPointerValue_BoolNil(t *testing.T) {
	var p *bool
	result := utils.GetPointerValue(p)
	require.Equal(t, false, result)
}

func TestGetPointerValue_BoolTrue(t *testing.T) {
	v := true
	result := utils.GetPointerValue(&v)
	require.Equal(t, true, result)
}

func TestZero(t *testing.T) {
	require.Equal(t, 0, utils.Zero[int]())
	require.Equal(t, "", utils.Zero[string]())
	require.Equal(t, false, utils.Zero[bool]())
	require.Nil(t, utils.Zero[*int]())
}
