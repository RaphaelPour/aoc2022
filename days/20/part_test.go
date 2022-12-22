package main

import (
    "testing"

  "github.com/stretchr/testify/require"
	"github.com/RaphaelPour/stellar/input"
)

func TestNeighbor(t *testing.T) {
	data := input.LoadInt("input1")
	ring := NewRing(data)
	
	require.Equal(t, len(data), len(ring.values))
	require.NotNil(t, ring.zero)

	for _, value := range ring.values {
		require.NotNil(t, value)
		require.NotNil(t, value.left)
		require.NotNil(t, value.right)
	}

	require.Equal(t, 1, ring.values[0].number)
	require.Equal(t, ring.values[0].right, ring.values[0].RightNeighborN(1))
	require.Equal(t, ring.values[0].right, ring.values[0].GetNeighborN(1))
	require.Equal(t, ring.values[0].right.right, ring.values[0].GetNeighborN(2))
	require.Equal(t, ring.values[0].left, ring.values[0].LeftNeighborN(1))
	require.Equal(t, ring.values[0].left, ring.values[0].GetNeighborN(-1))

	require.Equal(t, ring.values[0].left.left, ring.values[0].GetNeighborN(-2))
	require.Equal(t, ring.values[0].left.left.left, ring.values[0].GetNeighborN(-3))
	require.Equal(t, ring.values[0].left.left.left.left, ring.values[0].GetNeighborN(-4))
	require.Equal(t, ring.values[0].left.left.left.left.left, ring.values[0].GetNeighborN(-5))
	require.Equal(t, ring.values[0].left.left.left.left.left.left, ring.values[0].GetNeighborN(-6))
}

func TestMove1(t *testing.T) {
	ring := NewRing([]int{0, -3,  0,  0, 0,  0,  0})
	val := ring.values[1]
	ring.Mix()
	Dump(ring.values[0])
	require.Equal(t,val, ring.values[0].right.right.right.right)
}

func TestMove2(t *testing.T) {
	ring := NewRing([]int{0, -2,  0,  0, 0,  0,  0})
	val := ring.values[1]
	ring.Mix()
	Dump(ring.values[0])
	require.Equal(t,val, ring.values[0].right.right.right.right.right)
}

func TestMove3(t *testing.T) {
	ring := NewRing([]int{0, -1,  0,  0, 0,  0,  0})
	val := ring.values[1]
	ring.Mix()
	require.Equal(t,val, ring.values[0].left)
}

func TestMove4(t *testing.T) {
	ring := NewRing([]int{0, 0,  0,  0, 2,  0,  0})
	val := ring.values[1]
	ring.Mix()
	require.Equal(t,val, ring.values[0].right)
}

func TestExamplePart1(t *testing.T) {
	require.Equal(t, 3, part1(input.LoadInt("input1")))
}
