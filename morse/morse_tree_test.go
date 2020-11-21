package morse

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnError_WhenPreoderIsEmpty(t *testing.T) {
	preorder := [][]string{}
	inorder := [][]string{{".", "B"}, {".", "A"}}
	root, err := make(MorseCodes{preorder, inorder})		
	
	assert.Nil(t, root)
	assert.NotNil(t, err)
}

func TestShouldReturnError_WhenInorderIsEmpty(t *testing.T) {
	preorder := [][]string{{".", "A"}, {".", "B"}}
	inorder := [][]string{}
	root, err := make(MorseCodes{preorder, inorder})		
	
	assert.Nil(t, root)
	assert.NotNil(t, err)
}

func TestShouldReturnError_WhenPreorderAndInorderHaveDifferentLenghts(t *testing.T) {
	preorder := [][]string{{".", "A"}, {".", "B"}}
	inorder := [][]string{{".", "A"}}
	root, err := make(MorseCodes{preorder, inorder})		
	
	assert.Nil(t, root)
	assert.NotNil(t, err)
}

func TestShouldReturnRootWithLeftSubtree(t *testing.T) {
	preorder := [][]string{{".", "A"}, {".", "B"}}
	inorder := [][]string{{".", "B"}, {".", "A"}}
	root, _ := make(MorseCodes{preorder, inorder})		
	left := root.Left
	right := root.Right
	parent := root.Parent

	assert.Equal(t, ".", root.Code)
	assert.Equal(t, "A", root.Letter)
	assert.Equal(t, ".", left.Code)
	assert.Equal(t, "B", left.Letter)
	assert.Equal(t, root, left.Parent)
	assert.Nil(t, right)
	assert.Nil(t, parent)
}

func TestShouldReturnRootWithLeftAndRightSubtrees(t *testing.T) {
	preorder := [][]string{{".", "A"}, {".", "B"}, {".", "D"}, {".", "E"}, {".", "C"}}
	inorder := [][]string{{".", "D"}, {".", "B"}, {".", "E"}, {".", "A"}, {".", "C"}}
	root, _ := make(MorseCodes{preorder, inorder})		
	left := root.Left
	right := root.Right
	
	assert.Equal(t, "A", root.Letter)	
	assert.Equal(t, "B", left.Letter)
	assert.Equal(t, "C", right.Letter)
	assert.Equal(t, "D", left.Left.Letter)
	assert.Equal(t, "E", left.Right.Letter)	
}