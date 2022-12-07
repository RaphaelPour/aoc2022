package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/RaphaelPour/stellar/input"
	s_strings "github.com/RaphaelPour/stellar/strings"
)

type DirectoryMatcher func(Directory) bool

type Directory struct {
	parent      *Directory
	name        string
	files       map[string]int
	directories map[string]*Directory
	cached_size int
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name:        name,
		files:       make(map[string]int),
		directories: make(map[string]*Directory),
		cached_size: -1,
	}
}

func (d Directory) Path() string {
	if d.parent == nil {
		return "/"
	}

	return filepath.Join(d.parent.Path(), d.name)
}

func (d Directory) Ls(ident int) string {
	result := ""

	if d.parent == nil {
		result += fmt.Sprintf("%s (%d)\n", d.name, d.Size())
		ident += 1
	}

	spaces := strings.Repeat(" ", ident*2)
	for _, dir := range d.directories {
		result += fmt.Sprintf("%s%s (%d)\n", spaces, dir.name, dir.Size())
		result += dir.Ls(ident + 1)
	}

	for file, size := range d.files {
		result += fmt.Sprintf("%s%s (%d)\n", spaces, file, size)
	}
	return result
}

func (d Directory) Size() int {
	/*if d.cached_size > -1 {
		return d.cached_size
	}*/

	sum := 0
	for _, size := range d.files {
		sum += size
	}

	for _, dir := range d.directories {
		sum += dir.Size()
	}
	d.cached_size = sum
	return sum
}

func (d Directory) Collect(matcher DirectoryMatcher) []*Directory {
	result := make([]*Directory, 0)
	for _, dir := range d.directories {
		if matcher(*dir) {
			result = append(result, dir)
		}
		result = append(result, dir.Collect(matcher)...)
	}
	return result
}

func part1(data []string) int {
	var root, current *Directory

	re := regexp.MustCompile(`^(\d+)\s(.*)$`)

	for _, line := range data {
		if strings.HasPrefix(line, "$ cd ..") {
			if current.parent == nil {
				panic(fmt.Sprintf("error going up. Current directory %s has no parent", current.name))
			}
			// fmt.Println("cd", current.name, "to", current.parent.name)
			current = current.parent
			continue
		} else if strings.HasPrefix(line, "$ cd") {
			name := line[strings.LastIndex(line, " ")+1:]
			newDir := NewDirectory(name)

			/* initialize root of the fs tree */
			if root == nil {
				root = newDir
				current = newDir
				// fmt.Println("init tree with", name)
				continue
			}

			newDir.parent = current
			current.directories[name] = newDir
			// fmt.Println("cd", current.name, "to", newDir.name)
			current = newDir
			continue
		} else if match := re.FindStringSubmatch(line); len(match) == 3 {
			current.files[match[2]] = s_strings.ToInt(match[1])
			// fmt.Println("add file", match[2], "to", current.name)
			continue
		} else if strings.HasPrefix(line, "dir") || strings.HasPrefix(line, "$ ls") {
			continue
		}

		panic(fmt.Sprintf("Unknown line '%s'", line))
	}

	fmt.Println(root.Ls(0))

	total := 0
	for _, dir := range root.Collect(func(d Directory) bool { return d.Size() <= 100000 }) {
		fmt.Println(dir)
		total += dir.Size()
	}
	return total
}

func part2(data []string) int {
	var root, current *Directory

	re := regexp.MustCompile(`^(\d+)\s(.*)$`)

	for _, line := range data {
		if strings.HasPrefix(line, "$ cd ..") {
			if current.parent == nil {
				panic(fmt.Sprintf("error going up. Current directory %s has no parent", current.name))
			}
			// fmt.Println("cd", current.name, "to", current.parent.name)
			current = current.parent
			continue
		} else if strings.HasPrefix(line, "$ cd") {
			name := line[strings.LastIndex(line, " ")+1:]
			newDir := NewDirectory(name)

			/* initialize root of the fs tree */
			if root == nil {
				root = newDir
				current = newDir
				// fmt.Println("init tree with", name)
				continue
			}

			newDir.parent = current
			current.directories[name] = newDir
			// fmt.Println("cd", current.name, "to", newDir.name)
			current = newDir
			continue
		} else if match := re.FindStringSubmatch(line); len(match) == 3 {
			current.files[match[2]] = s_strings.ToInt(match[1])
			// fmt.Println("add file", match[2], "to", current.name)
			continue
		} else if strings.HasPrefix(line, "dir") || strings.HasPrefix(line, "$ ls") {
			continue
		}

		panic(fmt.Sprintf("Unknown line '%s'", line))
	}

	var candidate *Directory
	freeSpace := 70000000 - root.Size()
	goal := 30000000
	for _, dir := range root.Collect(func(d Directory) bool { return freeSpace+d.Size() >= goal }) {
		fmt.Println(dir.name, dir.Size())
		if candidate == nil {
			candidate = dir
			continue
		}
		if dir.Size() < candidate.Size() {
			candidate = dir
		}
	}
	return candidate.Size()
}

func main() {
	data := input.LoadString("input")

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(data))

	fmt.Println("== [ PART 2 ] ==")
	total := part2(data)
	if total > 13431333 {
		fmt.Println("too high:", total)
	} else {
		fmt.Println(total)
	}
}
