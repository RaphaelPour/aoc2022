package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPoint(t *testing.T) {
	p1 := Point{0, 0}
	p2 := Point{1, 0}
	p3 := Point{1, 1}

	require.Equal(t, 0, p1.EuclideanDistance(p1))
	require.Equal(t, 1, p1.EuclideanDistance(p2))
	require.Equal(t, 1, p1.EuclideanDistance(p3))
}
