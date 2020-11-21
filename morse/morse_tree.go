package morse

import (
	"errors"
	"strings"	
)

type MorseTreeNode struct {
	Left *MorseTreeNode
	Right *MorseTreeNode
	Parent *MorseTreeNode
	Code string
	Letter string
}

func make(codes MorseCodes) (*MorseTreeNode, error) {
	if len(codes.Preorder) == 0 {
		return nil, errors.New("Preorder is mandatory")
	}
	if len(codes.Inorder) == 0 {
		return nil, errors.New("Inorder is mandatory")
	}
	if len(codes.Preorder) != len(codes.Inorder) {
		return nil, errors.New("Preorder and Inorder should have same number of nodes")
	}	
	return makeTree(		
		makeInorderPositions(codes.Inorder),
		codes.Preorder), nil
}

func makeInorderPositions(inorder [][]string) map[string]int {
	positions := map[string]int{}
	for i, e := range inorder {
		positions[makeKey(e)] = i 
	}	
	return positions
}

func makeKey(mapping []string) string {
	return strings.Join(mapping[:], "# ")
}

func makeTree(positions map[string]int, preorder [][]string) *MorseTreeNode {
	var preorderIdx = 0
	var helper func(parent *MorseTreeNode, start int, end int) *MorseTreeNode
	helper = func(parent *MorseTreeNode, start int, end int) *MorseTreeNode {
		if start > end {
			return nil
		}		
		code, letter := preorder[preorderIdx][0], preorder[preorderIdx][1]
		inorderIdx := positions[makeKey(preorder[preorderIdx])]
		node := MorseTreeNode{Code: code, Letter: letter, Parent: parent}
		preorderIdx += 1
		if start != end {			
			node.Left = helper(&node, start, inorderIdx - 1)
			node.Right = helper(&node, inorderIdx + 1, end)
		}	
		return &node
	}
	return helper(nil, 0, len(preorder) - 1)		
}