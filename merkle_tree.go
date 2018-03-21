package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

type MerkleTree struct {
	RootNode *MerkleNode
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	mNode := MerkleNode{}

	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		mNode.Data = hash[:]
	} else {
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		mNode.Data = hash[:]
	}

	mNode.Left = left
	mNode.Right = right

	return &mNode
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	if len(data) == 0 {
		return &MerkleTree{nil}
	}

	if len(data) == 1 {
		return &MerkleTree{&nodes[0]}
	}

	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	for _, datum := range data {
		node := NewMerkleNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}

	for len(nodes) != 1 {
		if len(nodes)%2 != 0 {
			node := nodes[len(nodes)-1]
			node.Left = nil
			node.Right = nil
			nodes = append(nodes, node)
		}
		var newLevel []MerkleNode
		for j := 0; j < len(nodes); j += 2 {
			node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
			newLevel = append(newLevel, *node)
		}
		nodes = newLevel
	}

	mTree := MerkleTree{&nodes[0]}
	return &mTree
}

func (n *MerkleNode) String() string {
	return fmt.Sprintf("%x", n.Data)
}

func (t *MerkleTree) String() string {
	var lines []string
	var nodes []*MerkleNode
	nodes = append(nodes, t.RootNode)
	for len(nodes) > 0 {
		node := nodes[0]
		lines = append(lines, node.String())
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
		nodes = nodes[1:]
	}

	x := 1
	c := 0
	var linesWithNewline []string
	for _, line := range lines {
		linesWithNewline = append(linesWithNewline, line)
		c++
		if c == x {
			x *= 2
			c = 0
			linesWithNewline = append(linesWithNewline, "\n")
		} else {
			linesWithNewline = append(linesWithNewline, " - ")
		}
	}
	return strings.Join(linesWithNewline, "")
}
