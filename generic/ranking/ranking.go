package ranking

// Implement a ranking data structure that has the following characteristic:
// Elements:
// - Each element is a pair of (id, score)
// - The id is usually user_id. Id has to be unique.
// - The score is usually the score of the user
// - The rank of a record is defined to be the number of other record with strictly higher score, plus 1
// Operation:
// - Update(id, score): Update the value of an id, insert if necessary
// - AddScore(id, additionaScore): add to the score of an id
//   - if empty then the additionalScore is the score
//   - if total score = 0, it's not inserted.
//   - additionalScore should be > 0
// - RankOf(id): Get the rank of an id if it exist
// - ScoreOf(id): Get the score of an id if it exist
// - GetRange(first, last): Get the (id, score) pairs that has rank from the range [first, last)
// The actual implementation is as follow:
// - Store a map[KeyT]ScoreT to quickly reference score by id.
// - Then store a search data structure with the (key, value) being the ordered (score, id) pairs (compares by score then by id)
// - The multiple keys are allowed.
// - The search data structure chosen is a classical AVL tree map.
// Complexity:
// - O(n) memory (and builtin map memory)
// - O(log(n)) time and stack to look up single element (and one builtin map access)
// - O(log(n) + k) time to get a sorted range
// - This should work fine for this server, as frankly even the official server will strugle to get 1m players at once
// - Howerver it's possible to add a max size limit to cap the complexity if necessary.
import (
	"golang.org/x/exp/constraints"
)

type Ranking[ScoreT constraints.Integer, IdT constraints.Integer] struct {
	root      *Node[ScoreT, IdT]
	scoreById map[IdT]ScoreT
}

func NewRanking[ScoreT constraints.Integer, IdT constraints.Integer]() *Ranking[ScoreT, IdT] {
	return &Ranking[ScoreT, IdT]{
		root:      nil,
		scoreById: map[IdT]ScoreT{},
	}
}

func (r *Ranking[ScoreT, IdT]) Update(id IdT, newScore ScoreT) {
	score, exist := r.scoreById[id]
	if exist {
		if score == newScore {
			return
		}
		r.root = r.root.Delete(score, id)
	}
	r.scoreById[id] = newScore
	r.root = r.root.Insert(newScore, id)
}

func (r *Ranking[ScoreT, IdT]) AddScore(id IdT, additionalScore ScoreT) {
	if additionalScore == 0 {
		return
	}
	score, exist := r.scoreById[id]
	if exist {
		r.root = r.root.Delete(score, id)
	} else {
		score = 0
	}
	score += additionalScore
	r.scoreById[id] = score
	r.root = r.root.Insert(score, id)
}

func (r *Ranking[ScoreT, IdT]) RankOf(id IdT) (int, bool) {
	score, exist := r.scoreById[id]
	if !exist {
		return 0, false
	}
	return r.root.RankOf(score, id) + 1, true
}

func (r *Ranking[ScoreT, IdT]) ScoreOf(id IdT) (ScoreT, bool) {
	score, exist := r.scoreById[id]
	return score, exist
}

func (r *Ranking[ScoreT, IdT]) GetRange(first, last int) []Pair[ScoreT, IdT] {
	return r.root.Range(first-1, last-1)
}
