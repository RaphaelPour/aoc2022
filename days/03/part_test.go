package main

import (
	"testing"

	"github.com/RaphaelPour/stellar/input"
	"github.com/stretchr/testify/require"
)

func test_input() []string {
	return []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}
}

func TestPriority(t *testing.T) {
	require.Equal(t, 1, priority('a'))
	require.Equal(t, 26, priority('z'))
	require.Equal(t, 27, priority('A'))
	require.Equal(t, 52, priority('Z'))
}

func TestRucksack(t *testing.T) {
	r := NewRucksack("abcdea")
	require.NotNil(t, r)
	require.Equal(t, "abcdea", r.total)
	require.Equal(t, "abc", r.left)
	require.Equal(t, "dea", r.right)
	p, ok := r.Duplicate()
	require.True(t, ok)
	require.Equal(t, 1, p)
}

func TestGroup(t *testing.T) {
	group := Group{
		NewRucksack("abc"),
		NewRucksack("cda"),
		NewRucksack("dea"),
	}

	p, ok := group.Duplicate()
	require.True(t, ok)
	require.Equal(t, 1, p)
}

func TestExamplePart1(t *testing.T) {
	require.Equal(t, 157, part1(test_input()))
}

func TestPart1(t *testing.T) {
	require.Equal(t, 8515, part1(input.LoadString("input")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 2434, part2(input.LoadString("input")))
}

func TestExamplePart2(t *testing.T) {
	require.Equal(t, 70, part2(test_input()))
}
