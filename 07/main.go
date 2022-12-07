package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/matsest/advent2022/utils"
)

// A generic node that can either be a file or a directory
type Node struct {
	Name     string
	Type     string // Directory or File
	Children []Node // A list of pointers to directories or files
	Parent   *Node  // Pointer to a directory
	Size     int    // Size of file. Directory has size 0.
}

// Prints the Node struct
func (node *Node) print(level int) {
	spaces := strings.Repeat(" ", level)
	fmt.Printf("%s- %s (%s)\n", spaces, node.Name, node.Type)
	if node.Children != nil {
		level += 1
		for _, c := range node.Children {
			c.print(level)
		}
	}
}

// Gets the total size of a node
func (node *Node) getTotalSize() (size int) {
	size += node.Size
	if node.Children != nil {
		for _, c := range node.Children {
			size += c.getTotalSize()
		}
	}
	return size
}

// Gets the total size of sub-directories under a threshold
func (node *Node) getTotalSizeUnderT(t int) (size int) {
	for _, c := range node.Children {
		if c.Type == "directory" {
			currentSize := c.getTotalSize()
			if currentSize <= t {
				//fmt.Println("child ", c.Name, " size ", currentSize)
				size += currentSize
			}
			size += c.getTotalSizeUnderT(t)
		}
	}
	return size
}

// Get child directories with size over a certain threshold
func (node *Node) getNodesWithSizeOverT(t int) (nodes []Node) {
	if node.Type != "directory" {
		return nil
	}
	currentSize := node.getTotalSize()
	if currentSize >= t {
		nodes = append(nodes, *node)
	}

	for _, c := range node.Children {
		currentSize = c.getTotalSize()
		if currentSize >= t {
			nodes = append(nodes, c.getNodesWithSizeOverT(t)...)
		}
	}
	return nodes
}

// Creates a new file node
func NewFile(name string, parent *Node, size int) (*Node, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("Name must be defined")
	}
	if parent == nil {
		return nil, fmt.Errorf("Parent is not defined")
	}
	if size <= 0 {
		return nil, fmt.Errorf("Size must be a positive integer")
	}
	return &Node{
		Name:     name,
		Type:     "file",
		Children: nil,
		Parent:   parent,
		Size:     size,
	}, nil
}

// Creates a new directory node
func NewDirectory(name string, parent *Node) (*Node, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("Name must be defined")
	}

	return &Node{
		Name:     name,
		Type:     "directory",
		Children: nil,
		Parent:   parent,
		Size:     0,
	}, nil
}

// Add a child node to current node
func (current *Node) addChild(child Node) {
	current.Children = append(current.Children, child)
}

func parseInput(lines []string) (*Node, error) {
	root, _ := NewDirectory(string(lines[0][5]), nil)
	current := root

	for _, line := range lines[1:] {
		//fmt.Println("\n>>>> current line is: ", line, " with current dir ", current.Name)
		if line[0:4] == "$ cd" { // Go to dir
			if line[5:] == ".." { // Go to parent
				current = current.Parent
			} else { // Go to sub-directory
				dirName := line[5:]
				for i, child := range current.Children {
					if child.Name == dirName {
						current = &current.Children[i]
						//fmt.Println("CHANGED DIR TO ", current.Name)
					}
				}
			}
		} else if line == "$ ls" { // list content of dir
			continue
		} else if line[0:3] == "dir" { // is dir entry
			dirName := line[4:]
			node, err := NewDirectory(dirName, current)
			if err != nil {
				return nil, fmt.Errorf("Couldn't initialize dir")
			}
			current.addChild(*node)
		} else { // is file entry
			parts := strings.Split(line, " ")
			size, _ := strconv.Atoi(parts[0])
			name := parts[1]
			child, err := NewFile(name, current, size)
			if err != nil {
				return nil, fmt.Errorf("Couldn't initialize child")
			}
			current.addChild(*child)
		}
	}
	return root, nil
}

func p1(node *Node) int {
	return node.getTotalSizeUnderT(100000)
}

func p2(node *Node) int {
	// Calculate how much to delete
	totalSize := 70000000
	needsUnused := 30000000
	currentSize := node.getTotalSize()
	unused := totalSize - currentSize
	toDelete := needsUnused - unused

	// Find smallest node to delete
	possibleNodes := node.getNodesWithSizeOverT(toDelete)
	smallest := possibleNodes[0].getTotalSize()
	for _, n := range possibleNodes[1:] {
		if n.getTotalSize() < smallest {
			smallest = n.getTotalSize()
		}
	}
	return smallest
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	node, err := parseInput(lines)
	if err != nil {
		fmt.Println("Could not parse input")
		return
	}

	fmt.Println(p1(node))
	fmt.Println(p2(node))
}
