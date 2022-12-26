package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSnafu2Dec(t *testing.T) {
	for _, data := range []struct {
		Input    string
		Expected int
	}{
		{"1", 1},
		{"2", 2},
		{"1=", 3},
		{"1-", 4},
		{"10", 5},
		{"11", 6},
		{"12", 7},
		{"2=", 8},
		{"2-", 9},
		{"20", 10},
		{"1=0", 15},
		{"1-0", 20},
		{"1=11-2", 2022},
		{"1-0---0", 12345},
		{"1121-1110-1=0", 314159265},
	} {
		require.Equal(t, data.Expected, snafu2Dec(data.Input))
	}
}

func TestDec2Snafu(t *testing.T) {
	for _, data := range []struct {
		Expected string
		Input    int
	}{
		{"1", 1},
		{"2", 2},
		{"1=", 3},
		{"1-", 4},
		{"10", 5},
		{"11", 6},
		{"12", 7},
		{"2=", 8},
		{"2-", 9},
		{"20", 10},
		{"1=0", 15},
		{"1-0", 20},
		{"1=112", 2022},
		{"1-0---0", 12345},
		{"1121-1110-1=0", 314159265},
	} {
		require.Equal(t, data.Expected, dec2Snafu(data.Input))
	}
}
