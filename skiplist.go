package skiplist

import "math"

const (
	maxLevel    = 1 << 10
	probability = 0.5 //TODO:???
)

type Node struct {
	Score int64
	Val   interface{}
	Next  *Node
	Pre   *Node
	Up    *Node
	Down  *Node
}

// example: 双向链表的跳表
type SkipList struct {
	Head   *Node
	Tail   *Node
	Size   int
	Levels int
}

func NewSkipList() *SkipList {
	skipList := &SkipList{
		Head: new(Node),
		Tail: new(Node),
	}

	skipList.Head.Score = math.MinInt64
	skipList.Tail.Score = math.MaxInt64

	skipList.Head.Next = skipList.Tail
	skipList.Tail.Pre = skipList.Head

	skipList.Size = 0
	skipList.Levels = 1

	return skipList
}

func (s *SkipList) newLevel() {
	nhead := &Node{Score: math.MinInt64}
	ntail := &Node{Score: math.MaxInt64}
	nhead.Next = ntail
	ntail.Pre = nhead

	s.Head.Up = nhead
	nhead.Down = s.Head
	s.Tail.Up = ntail
	ntail.Down = s.Tail

	s.Head = nhead
	s.Tail = ntail
	s.Levels++
}

func (s *SkipList) Insert(score int64, val interface{}) {
	//TODO:
}

func (s *SkipList) Remove(score int64) interface{} {
	//TODO:
	return nil
}

// 先横向查找 再向下查找
func (s *SkipList) findNode(score int64) *Node {
	p := s.Head

	for p != nil {
		if p.Score == score {
			if p.Down == nil {
				return p
			}
			p = p.Down
		} else if p.Score < score {
			if p.Next.Score > score {
				if p.Down == nil {
					return p
				}
				p = p.Down
			} else {
				p = p.Next
			}
		}
	}
	return p
}
