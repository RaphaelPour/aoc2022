package main

import (
	"testing"

	"github.com/RaphaelPour/stellar/input"
	"github.com/stretchr/testify/require"
)

func TestPart1Example(t *testing.T) {
	require.Equal(t, 2, part1(input.LoadString("input1")))
}

func TestPart2Example(t *testing.T) {
	require.Equal(t, 4, part2(input.LoadString("input1")))
}

func TestPart1(t *testing.T) {
	require.Equal(t, 450, part1(input.LoadString("input")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 837, part2(input.LoadString("input")))
}
