package skiplist

import (
	"log"
	"testing"
)

func TestSkipList_Insert(t *testing.T) {
	sp := NewSkipList()
	sp.Insert(1, "test")
	sp.Insert(7, "test")
	sp.Insert(4, "test")
	sp.Insert(12, "test")
	sp.Insert(15, "test")
	log.Println(sp.Find(12))
	sp.Delete(12)
	log.Println(sp.Find(12))
}
