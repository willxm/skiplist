package skiplist

import (
	"math/rand"
	"time"
)

const (
	maxLevel    = 16
	probability = 0.5
)

type Node struct {
	Score   int64
	Value   interface{}
	Forward []*Node
}

func newNode(score int64, value interface{}, level int) *Node {
	return &Node{
		Score:   score,
		Value:   value,
		Forward: make([]*Node, level),
	}
}

type SkipList struct {
	Head   *Node
	Size   int
	Levels int
}

func NewSkipList() *SkipList {
	return &SkipList{
		Head: &Node{Forward: make([]*Node, maxLevel)},
	}
}

func randomLevel() int {
	level := 1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for r.Float32() < probability && level < maxLevel {
		level++
	}
	return level
}

func (s *SkipList) Find(score int64) *Node {
	p := s.Head
	for i := s.Levels - 1; i >= 0; i-- {
		for p.Forward[i] != nil && p.Forward[i].Score < score {
			p = p.Forward[i]
		}
	}
	p = p.Forward[0]
	if p != nil && p.Score == score {
		return p
	}
	return nil
}

func (s *SkipList) Insert(score int64, value interface{}) *Node {
	p := s.Head
	fi := make([]*Node, maxLevel)

	for i := s.Levels - 1; i >= 0; i-- {
		for p.Forward[i] != nil && p.Forward[i].Score < score {
			p = p.Forward[i]
		}
		fi[i] = p
	}

	p = p.Forward[0]

	if p != nil && p.Score == score {
		p.Value = value
		return p
	}

	level := randomLevel()
	if level > s.Levels {
		level = s.Levels + 1
		fi[s.Levels] = s.Head
		s.Levels = level
	}
	n := newNode(score, value, level)

	for i := 0; i < level; i++ {
		n.Forward[i] = fi[i].Forward[i]
		fi[i].Forward[i] = n
	}
	s.Size++
	return n
}

func (s *SkipList) Delete(score int64) *Node {
	p := s.Head
	fi := make([]*Node, maxLevel)

	for i := s.Levels - 1; i >= 0; i-- {
		for p.Forward[i] != nil && p.Forward[i].Score < score {
			p = p.Forward[i]
		}
		fi[i] = p
	}
	p = p.Forward[0]

	if p != nil && p.Score == score {
		for i := 0; i < s.Levels; i++ {
			if fi[i].Forward[i] != p {
				return nil
			}
			fi[i].Forward[i] = p.Forward[i]
		}
		s.Levels--
	}
	return p
}
