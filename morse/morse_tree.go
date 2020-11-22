package morse

import (
	"errors"	
	"strings"
)

type MorseTree struct {
	Root  *MorseTreeNode
	nodes map[string]*MorseTreeNode
}

type MorseTreeNode struct {
	Left   *MorseTreeNode
	Right  *MorseTreeNode
	Parent *MorseTreeNode
	Code   string
	Letter string
}

func Make(codes MorseCodes) (*MorseTree, error) {
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

func makeTree(positions map[string]int, preorder [][]string) *MorseTree {
	nodes := map[string]*MorseTreeNode{}
	var preorderIdx = 0
	var helper func(parent *MorseTreeNode, start int, end int) *MorseTreeNode
	helper = func(parent *MorseTreeNode, start int, end int) *MorseTreeNode {
		if start > end {
			return nil
		}
		code, letter := preorder[preorderIdx][0], preorder[preorderIdx][1]
		inorderIdx := positions[makeKey(preorder[preorderIdx])]
		node := MorseTreeNode{Code: code, Letter: letter, Parent: parent}
		nodes[letter] = &node
		preorderIdx += 1
		if start != end {
			node.Left = helper(&node, start, inorderIdx-1)
			node.Right = helper(&node, inorderIdx+1, end)
		}
		return &node
	}
	return &MorseTree{helper(nil, 0, len(preorder)-1), nodes}
}

func (tree *MorseTree) Decode(start string, codes string) string {
	letters := ""
	for _, code := range strings.Fields(codes) {		
		letters += tree.decode(start + code)
	}
	return letters
}

func (tree *MorseTree) decode(code string) string {	
	var helper func(node *MorseTreeNode, code string, codeIdx int) string
	helper = func(node *MorseTreeNode, code string, codeIdx int) string {
		if codeIdx == len(code)-1 {
			return node.Letter
		}
		if string(code[codeIdx]) != node.Code {
			return ""
		}
		if node.Left != nil && string(code[codeIdx+1]) == node.Left.Code {
			return helper(node.Left, code, codeIdx+1)
		}
		if node.Right != nil && string(code[codeIdx+1]) == node.Right.Code {
			return helper(node.Right, code, codeIdx+1)
		}
		return ""
	}
	return helper(tree.Root, code, 0)
}

func (tree *MorseTree) Encode(letter string) string {
	var helper func(node *MorseTreeNode, code string) string
	helper = func(node *MorseTreeNode, code string) string {
		if node == nil {
			return code
		}
		return helper(node.Parent, node.Code+code)
	}
	return helper(tree.nodes[letter], "")
}
