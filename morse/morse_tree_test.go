package morse

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMakeShouldReturnError_WhenPreoderIsEmpty(t *testing.T) {
	preorder := [][]string{}
	inorder := [][]string{{".", "B"}, {".", "A"}}
	tree, err := make(MorseCodes{preorder, inorder})		
	
	assert.Nil(t, tree)
	assert.NotNil(t, err)
}

func TestMakeShouldReturnError_WhenInorderIsEmpty(t *testing.T) {
	preorder := [][]string{{".", "A"}, {".", "B"}}
	inorder := [][]string{}
	tree, err := make(MorseCodes{preorder, inorder})		
	
	assert.Nil(t, tree)
	assert.NotNil(t, err)
}

func TestMakeShouldReturnError_WhenPreorderAndInorderHaveDifferentLenghts(t *testing.T) {
	preorder := [][]string{{".", "A"}, {".", "B"}}
	inorder := [][]string{{".", "A"}}
	root, err := make(MorseCodes{preorder, inorder})		
	
	assert.Nil(t, root)
	assert.NotNil(t, err)
}

func TestMakeShouldReturnRootWithLeftSubtree(t *testing.T) {
	preorder := [][]string{{".", "A"}, {".", "B"}}
	inorder := [][]string{{".", "B"}, {".", "A"}}
	tree, _ := make(MorseCodes{preorder, inorder})
	root := tree.Root
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

func TestMakeShouldReturnRootWithLeftAndRightSubtrees(t *testing.T) {
	preorder := [][]string{{".", "A"}, {".", "B"}, {".", "D"}, {".", "E"}, {".", "C"}}
	inorder := [][]string{{".", "D"}, {".", "B"}, {".", "E"}, {".", "A"}, {".", "C"}}
	tree, _ := make(MorseCodes{preorder, inorder})		
	root := tree.Root
	left := root.Left
	right := root.Right
	
	assert.Equal(t, "A", root.Letter)	
	assert.Equal(t, "B", left.Letter)
	assert.Equal(t, "C", right.Letter)
	assert.Equal(t, "D", left.Left.Letter)
	assert.Equal(t, "E", left.Right.Letter)	
}

func TestDecodeShouldConvertMorseCodeToLetter(t *testing.T) {
	preorder := [][]string{{".", "A"}, {".", "B"}, {".", "D"}, {"-", "E"}, {"-", "C"}}
	inorder := [][]string{{".", "D"}, {".", "B"}, {".", "E"}, {".", "A"}, {".", "C"}}
	tree, _ := make(MorseCodes{preorder, inorder})
	
	assert.Equal(t, "A", tree.decode("."))
	assert.Equal(t, "B", tree.decode(".."))
	assert.Equal(t, "C", tree.decode(".-"))
	assert.Equal(t, "D", tree.decode("..."))
	assert.Equal(t, "E", tree.decode("..-"))
	assert.Equal(t, "", tree.decode("..--"))
}

func TestDecodeShouldConvertSpaceSeparatedMorseCodesToLetters(t *testing.T) {
	preorder := [][]string{{".", "A"}, {".", "B"}, {".", "D"}, {"-", "E"}, {"-", "C"}}
	inorder := [][]string{{".", "D"}, {".", "B"}, {".", "E"}, {".", "A"}, {".", "C"}}
	tree, _ := make(MorseCodes{preorder, inorder})
	
	assert.Equal(t, "AB", tree.decode(". .."))
	assert.Equal(t, "B", tree.decode(".."))	
	assert.Equal(t, "DE", tree.decode("... ..-"))	
}

func TestEncodeShouldConvertLetterToEquivalentMorseCode(t *testing.T) {
	preorder := [][]string{{".", "A"}, {".", "B"}, {".", "D"}, {"-", "E"}, {"-", "C"}}
	inorder := [][]string{{".", "D"}, {".", "B"}, {".", "E"}, {".", "A"}, {".", "C"}}
	tree, _ := make(MorseCodes{preorder, inorder})

	assert.Equal(t, ".", tree.encode("A"))
	assert.Equal(t, "..", tree.encode("B"))
	assert.Equal(t, "..-", tree.encode("E"))
	assert.Equal(t, "", tree.encode("$$$"))
}