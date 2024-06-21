package ranking

import (
	"golang.org/x/exp/constraints"
)

type Pair[ScoreT any, IdT any] struct {
	Score ScoreT
	Id    IdT
}

// Implement an AVL tree map where the score is a pair of id
// it's similar to a std::set<std::pair<score_t, id_t>>
// We only have the following operation:
// - Insert(score, id)
// - Delete(score, id)
// The interface we have is as follow:
// - Every modification operation is called with the root, and return the new root.
// - The nil root equal an empty tree
// - Every get operation doesn't change the root and return the desired result.
type Node[ScoreT constraints.Integer, IdT constraints.Integer] struct {
	left  *Node[ScoreT, IdT]
	right *Node[ScoreT, IdT]

	// this is the height of the subtree, used to balance the tree
	// we don't actually need the height and just need the balance factor, but this is easier to reason about
	height int

	// this is the sum of size of the storing maps, used to calculate ranking
	totalSize int
	score     ScoreT
	id        IdT
}

// return true if the score id must come before the node, false otherwise
func before[ScoreT constraints.Integer, IdT constraints.Integer](score ScoreT, id IdT, node *Node[ScoreT, IdT]) bool {
	return (score > node.score) || (score == node.score && id < node.id)
}

// return true if the score id must come after the node, false otherwise
func after[ScoreT constraints.Integer, IdT constraints.Integer](score ScoreT, id IdT, node *Node[ScoreT, IdT]) bool {
	return (score < node.score) || (score == node.score && id > node.id)
}

func newNode[ScoreT constraints.Integer, IdT constraints.Integer](score ScoreT, id IdT) *Node[ScoreT, IdT] {
	var node Node[ScoreT, IdT]
	node.totalSize = 1
	node.height = 0
	node.score = score
	node.id = id
	return &node
}

// recalculate the stuff
func (root *Node[ScoreT, IdT]) fix() {
	root.height = 0
	root.totalSize = 1
	if root.left != nil {
		root.height = root.left.height + 1
		root.totalSize += root.left.totalSize
	}
	if root.right != nil {
		if root.right.height >= root.height {
			root.height = root.right.height + 1
		}
		root.totalSize += root.right.totalSize
	}
}

// bring the right child to be the new root
func (root *Node[ScoreT, IdT]) rotateLeft() *Node[ScoreT, IdT] {
	right := root.right
	rightLeft := right.left
	root.right = rightLeft
	right.left = root
	root.fix()
	right.fix()
	return right
}

// bring the left child to be the new root
func (root *Node[ScoreT, IdT]) rotateRight() *Node[ScoreT, IdT] {
	left := root.left
	leftRight := left.right
	root.left = leftRight
	left.right = root
	root.fix()
	left.fix()
	return left
}

func (root *Node[ScoreT, IdT]) balance() *Node[ScoreT, IdT] {
	leftHeight := 0
	rightHeight := 0
	if root.left != nil {
		leftHeight = root.left.height
	}
	if root.right != nil {
		rightHeight = root.right.height
	}
	if leftHeight < rightHeight {
		if leftHeight+1 < rightHeight {
			return root.rotateLeft()
		} else {
			root.height = rightHeight + 1
		}
	} else {
		if rightHeight+1 < leftHeight {
			return root.rotateRight()
		} else {
			root.height = leftHeight + 1
		}
	}
	root.fix()
	return root
}

// detach the highest score in the tree
func (root *Node[ScoreT, IdT]) detachLast() (*Node[ScoreT, IdT], *Node[ScoreT, IdT]) {
	var deatched *Node[ScoreT, IdT]
	if root.right != nil {
		root.right, deatched = root.right.detachLast()
		return root.balance(), deatched
	} else { // this is the node to be detached, just return this node and replace it with its left child
		return root.left, root
	}
}

// detach the lowest score in the tree
func (root *Node[ScoreT, IdT]) detachFirst() (*Node[ScoreT, IdT], *Node[ScoreT, IdT]) {
	var deatched *Node[ScoreT, IdT]
	if root.left != nil {
		root.left, deatched = root.left.detachFirst()
		return root.balance(), deatched
	} else { // this is the node to be detached, just return this node and replace it with its right child
		return root.right, root
	}
}

// insert a pair of score id
func (root *Node[ScoreT, IdT]) Insert(score ScoreT, id IdT) *Node[ScoreT, IdT] {
	if root == nil { // empty tree, insert it directly into a new node
		return newNode(score, id)
	}
	if before(score, id, root) {
		root.left = root.left.Insert(score, id)
	} else if after(score, id, root) {
		root.right = root.right.Insert(score, id)
	} else {
		return root // already exist, doesn't change anything
	}
	return root.balance()
}

// delete a pair of score id
func (root *Node[ScoreT, IdT]) Delete(score ScoreT, id IdT) *Node[ScoreT, IdT] {
	if root == nil { // item doesn't exist
		return nil
	}
	if after(score, id, root) {
		root.right = root.right.Delete(score, id)
	} else if before(score, id, root) {
		root.left = root.left.Delete(score, id)
	} else { // need to delete this
		if root.right == nil {
			return root.left
		} else if root.left == nil {
			return root.right
		} else {
			var newRoot *Node[ScoreT, IdT]
			if root.left.height > root.right.height {
				root.left, newRoot = root.left.detachLast()
			} else {
				root.right, newRoot = root.right.detachFirst()
			}
			newRoot.left = root.left
			newRoot.right = root.right
			root = newRoot
		}
	}
	root.fix()
	return root.balance()
}

// return the number of items that come before this item, if it were inserted
func (root *Node[ScoreT, IdT]) RankOf(score ScoreT, id IdT) int {
	if root == nil {
		return 0
	}
	if before(score, id, root) {
		return root.left.RankOf(score, id)
	} else if after(score, id, root) {
		if root.right != nil {
			return root.totalSize - root.right.totalSize + root.right.RankOf(score, id)
		} else {
			return root.totalSize
		}
	} else {
		if root.left == nil {
			return 0
		} else {
			return root.left.totalSize
		}
	}
}

// return the (score, id) pair at position p when all pairs are sorted
func (root *Node[ScoreT, IdT]) At(p int) (*ScoreT, *IdT) {
	if (root == nil) || (root.totalSize < p) {
		return nil, nil
	}
	if root.left != nil {
		if root.left.totalSize > p {
			return root.left.At(p)
		} else {
			p -= root.left.totalSize
		}
	}
	if p == 0 {
		return &root.score, &root.id
	}
	return root.right.At(p - 1)
}

// get the range [begin, end) when all the pairs are sorted
// the numbering is from 0
func (root *Node[ScoreT, IdT]) Range(begin, end int) []Pair[ScoreT, IdT] {
	result := []Pair[ScoreT, IdT]{}
	root.internalRange(begin, end, 0, &result)
	return result
}

func (root *Node[ScoreT, IdT]) internalRange(begin, end, index int, result *[]Pair[ScoreT, IdT]) {
	if root == nil {
		return
	}
	if (index >= end) || (index+root.totalSize <= begin) {
		return
	}
	if root.left != nil {
		root.left.internalRange(begin, end, index, result)
		index += root.left.totalSize
	}
	if index >= end {
		return
	}
	if index >= begin {
		*result = append(*result, Pair[ScoreT, IdT]{
			Score: root.score,
			Id:    root.id,
		})
	}
	root.right.internalRange(begin, end, index+1, result)
}
