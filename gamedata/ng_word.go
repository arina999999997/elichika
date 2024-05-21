package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

// Bad word checker implemented using Aho-Corasick
// also note that this one doesn't report which "word" is bad, just whether there's a bad word, so we don't need to compute
// the actual matching words, just whether there's a word with current suffix
// also it's possible to prune the tree after building stuff, but it only reduce memory usage and don't really affect performance

// The matching is done without "space"
var isSkipedRune = map[rune]bool{}

func init() {
	// https://github.com/Suyooo/sifas-badwords/blob/main/check.js
	skipList := ` .,;-_~'´!?"&<>()[]{}　（）．‥…，；ー＿〜’！？”＜＞［］｛}` + "`"
	for _, r := range skipList {
		isSkipedRune[r] = true
	}
}

type NgWordNode struct {
	Parent  *NgWordNode
	Rune    rune
	HasWord bool
	Child   map[rune]*NgWordNode
	Suffix  *NgWordNode
}

func (node *NgWordNode) GetChild(r rune) *NgWordNode {
	result, exist := node.Child[r]
	if exist {
		return result
	}
	node.Child[r] = new(NgWordNode)
	node.Child[r].Child = map[rune]*NgWordNode{}
	node.Child[r].Parent = node
	node.Child[r].Rune = r
	return node.Child[r]
}

func (root *NgWordNode) AddWord(s string) {
	node := root
	for _, r := range s {
		node = node.GetChild(r)
	}
	node.HasWord = true
}

// once built, adding word will require a rebuild to be correct
// adding word without rebuilding can cause wrong check or even null reference
func (root *NgWordNode) Build() {
	queue := []*NgWordNode{}
	queue = append(queue, root)
	len := 1
	root.Suffix = root
	for i := 0; i < len; i++ {
		node := queue[i]
		for _, child := range node.Child {
			queue = append(queue, child)
			len++
		}
		if node == root {
			continue
		}
		node.Suffix = node.Parent.Suffix
		for node.Suffix != root {
			_, exist := node.Suffix.Child[node.Rune]
			if exist {
				break
			} else {
				node.Suffix = node.Suffix.Suffix
			}
		}
		suffix, exist := node.Suffix.Child[node.Rune]
		if exist && (suffix != node) {
			node.Suffix = suffix
		}
		// fmt.Println(node.GetString(), "->", node.Suffix.GetString())
		node.HasWord = node.HasWord || node.Suffix.HasWord
	}
}

func (root *NgWordNode) GetString() string {
	if root.Parent != nil {
		return root.Parent.GetString() + string(root.Rune)
	} else {
		return string(root.Rune)
	}
}
func (root *NgWordNode) HasMatch(s string) bool {
	fmt.Println(s)
	matched := root
	for _, r := range s {
		if isSkipedRune[r] {
			continue
		}
		for matched != root {
			_, exist := matched.Child[r]
			if exist {
				break
			} else {
				matched = matched.Suffix
			}
		}
		next, exist := matched.Child[r]
		if exist {
			matched = next
			if matched.HasWord {
				return true
			}
		}
	}
	return false
}

func loadNgWord(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading NgWord")
	gamedata.NgWord = new(NgWordNode)
	gamedata.NgWord.Child = map[rune]*NgWordNode{}
	words := []string{}
	err := serverdata_db.Table("s_ng_word").Cols("word").Find(&words)
	utils.CheckErr(err)

	for _, word := range words {
		gamedata.NgWord.AddWord(word)
	}
	gamedata.NgWord.Build()
}

func init() {
	addLoadFunc(loadNgWord)
}
