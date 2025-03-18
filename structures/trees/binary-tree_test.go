package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type intLikeStruct struct {
	intVal int
}

func (i *intLikeStruct) Compare(other Comparable) int {
	// assert struct are of the same type
	st, canCast := other.(*intLikeStruct)
	if !canCast {
		return 0
	}

	if i.intVal == st.intVal {
		return 0
	}
	if i.intVal < st.intVal {
		return -1
	}
	return 1
}

type binaryTreeTestSuite struct {
	suite.Suite
}

func (s *binaryTreeTestSuite) TestInsert_WithBFS() {
	tree := NewBinaryTree[*intLikeStruct]()

	tree.Insert(&intLikeStruct{1})
	tree.Insert(&intLikeStruct{2})
	tree.Insert(&intLikeStruct{5})
	tree.Insert(&intLikeStruct{-2})
	tree.Insert(&intLikeStruct{-10})
	tree.Insert(&intLikeStruct{111})

	assert.True(s.T(), tree.FindBFS(&intLikeStruct{111}))
	assert.True(s.T(), tree.FindBFS(&intLikeStruct{2}))

	assert.False(s.T(), tree.FindBFS(&intLikeStruct{222}))
	assert.False(s.T(), tree.FindBFS(&intLikeStruct{-123}))
}

func (s *binaryTreeTestSuite) TestFindDFS() {
	tree := NewBinaryTree[*intLikeStruct]()

	tree.Insert(&intLikeStruct{1})
	tree.Insert(&intLikeStruct{2})
	tree.Insert(&intLikeStruct{5})
	tree.Insert(&intLikeStruct{-2})
	tree.Insert(&intLikeStruct{-10})
	tree.Insert(&intLikeStruct{111})

	assert.True(s.T(), tree.FindDFS(&intLikeStruct{-10}))
	assert.True(s.T(), tree.FindDFS(&intLikeStruct{5}))

	assert.False(s.T(), tree.FindDFS(&intLikeStruct{-222}))
	assert.False(s.T(), tree.FindDFS(&intLikeStruct{-111}))
}

func TestBinaryTree(t *testing.T) {
	suite.Run(t, new(binaryTreeTestSuite))
}
