package main

import (
    "testing"

        "github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
    // require.Equal(t, 0, part1([]string{})
}

func TestStack(t *testing.T) {
	s := Stack[string]{}
	s2 := Stack[string]{}
	s.Push("A")
	s.Push("B")
	s.Push("C")

	require.Equal(t,s.Pop(), "C")
	require.Equal(t,s.PopN(2), []string{"A","B"})
	require.Equal(t,s.items, []string{})

	s.PushN([]string{"A","B"})
	require.Equal(t,s.items, []string{"A","B"})
	require.Equal(t,s.PopN(2), []string{"A","B"})

	s.PushN([]string{"A","B","C"})
	s2.PushN(s.PopN(2))
	require.Equal(t,s2.items, []string{"B","C"})
	require.Equal(t,s.items, []string{"A"})
}
