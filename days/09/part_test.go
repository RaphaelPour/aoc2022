package main

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/RaphaelPour/stellar/input"
)

func TestPoint(t *testing.T) {
	p1 := Point{0, 0}
	p2 := Point{1, 0}
	p3 := Point{1, 1}

	require.Equal(t, 0, p1.ManhattanDistance(p1))
	require.Equal(t, 1, p1.ManhattanDistance(p2))
	require.Equal(t, 2, p1.ManhattanDistance(p3))

	require.Equal(t, 0.0, p1.EuclideanDistance(p1))
	require.Equal(t, 1.0, p1.EuclideanDistance(p2))
	require.Greater(t, 2.0, p1.EuclideanDistance(p3))

	require.Equal(t, Point{2, 1}, p2.Add(p3))
}

func TestPart1Example(t *testing.T) {
	require.Equal(t, 13, part1(input.LoadString("input1")))
}

func TestPart1(t *testing.T) {
	require.Equal(t, 5878, part1(input.LoadString("input")))
}

func TestPart2Example(t *testing.T) {
	require.Equal(t, 36, part2(input.LoadString("input2")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 2405, part2(input.LoadString("input")))
}
