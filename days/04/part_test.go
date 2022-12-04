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

func TestRangeContains(t *testing.T) {
	for _, data := range []struct {
		from, to Range
		expected bool
	}{
		{Range{1, 3}, Range{1, 3}, true},
		{Range{1, 4}, Range{2, 3}, true},
		{Range{1, 3}, Range{2, 2}, true},
		{Range{1, 3}, Range{4, 6}, false},
	} {
		require.Equal(t, data.from.Contains(data.to), data.expected)
		require.Equal(t, data.to.Contains(data.from), data.expected)
	}
}

func TestRangeOverlaps(t *testing.T) {
	for _, data := range []struct {
		from, to Range
		expected bool
	}{
		{Range{1, 3}, Range{1, 3}, true},
		{Range{1, 3}, Range{3, 5}, true},
		{Range{1, 3}, Range{2, 3}, true},
		{Range{1, 6}, Range{2, 3}, true},
	} {
		require.Equal(t, data.from.Overlaps(data.to), data.expected)
		require.Equal(t, data.to.Overlaps(data.from), data.expected)
	}
}
