package algorithms

import "fmt"

var tableSize = 11

type HashTableSearch struct {
	tableSize int
	table     []*LinkedList[string]
	Count     int // các element trong hash table
}

type Word struct {
	word string
	hash int
}

func InitializeHashTableSearch(tableSize int) HashTableSearch {
	hash := HashTableSearch{tableSize: tableSize, Count: 0}
	hash.table = make([]*LinkedList[string], tableSize)
	return hash
}

func (hash *HashTableSearch) Load(c []string) {
	// hash.tableSize = tableSize
	// hash.table = make([]*LinkedList[string], hash.tableSize)
	for _, e := range c {
		// word := Word{word: e}
		// h := hash.hashFunction(word)
		// if hash.table[h] == nil {
		// 	hash.table[h] = new(LinkedList[string])
		// }
		// hash.table[h].AddLastNode(e)
		// hash.Count++
		hash.Add(e, e)
	}
}

func (hash *HashTableSearch) Add(k, v string) {
	if hash.Count == 0 {
		hash.tableSize = tableSize
		hash.table = make([]*LinkedList[string], hash.tableSize)
	}

	h := hash.hashFunction(Word{word: k})
	if hash.table[h] == nil {
		hash.table[h] = new(LinkedList[string])
	}
	hash.table[h].AddLastNode(v)
	hash.Count++
}

func (hash *HashTableSearch) Get(k string) bool {
	h := hash.hashFunction(Word{word: k})
	linkedList := hash.table[h]
	if linkedList == nil {
		return false
	}
	return linkedList.Contains(k)
}

func (hash *HashTableSearch) Delete(k string) (bool, error) {
	h := hash.hashFunction(Word{word: k})
	linkedList := hash.table[h]
	if linkedList == nil {
		return false, fmt.Errorf("not contains element")
	}
	if linkedList.Contains(k) {
		deleteNode := linkedList.DeleteNode(k)
		if deleteNode {
			hash.Count--
		}
		return deleteNode, nil
	}
	return false, nil
}

// hash code của string thành khóa k
func (w Word) hashCode() int {
	h := w.hash
	if h == 0 {
		for i := range w.word {
			h = 31*h + int(w.word[i])
		}
		// w.hash = h
	}
	return h
}
func (hash *HashTableSearch) hashFunction(word Word) int {
	h := word.hashCode()
	if h < 0 {
		h = -h
	}
	return h % hash.tableSize
}
