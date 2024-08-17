package algorithms

/**
* 	- LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
* 	- int get(int key) Return the value of the key if the key exists, otherwise return -1.
* 	- void put(int key, int value) Update the value of the key if the key exists.
*	Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
*
* 	The functions get and put must each run in O(1) average time complexity.
*
* Example 1:
*
* Input
* ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
* [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
* Output
* [null, null, null, 1, null, -1, null, -1, 3, 4]
*
* Explanation
* LRUCache lRUCache = new LRUCache(2);
* lRUCache.put(1, 1); // cache is {1=1}
* lRUCache.put(2, 2); // cache is {1=1, 2=2}
* lRUCache.get(1);    // return 1
* lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
* lRUCache.get(2);    // returns -1 (not found)
* lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
* lRUCache.get(1);    // return -1 (not found)
* lRUCache.get(3);    // return 3
* lRUCache.get(4);    // return 4
*
 */

type LRUCache struct {
	mm       map[int]*NodeDouble[int]
	linked   *LinkedListDouble[int]
	capacity int
	// lowest   int // key
	nextRank int
}

// type Cache struct {
// 	Key, Value, Rank int
// }

func Constructor(capacity int) LRUCache {
	mm := make(map[int]*NodeDouble[int])
	l := InitLinkedLinkDouble[int]()
	return LRUCache{
		mm:       mm,
		linked:   l,
		capacity: capacity,
		nextRank: 0,
		// lowest:   0,
	}
}

func (lru *LRUCache) Get(key int) int {
	pointer, ok := lru.mm[key]
	if !ok {
		return -1
	}
	lru.update(pointer)

	// find the next lowest if it is lowest
	// if pointer.Rank == lru.arr[lru.lowest].Rank {
	// 	lru.findLowestRank()
	// }
	// pointer.Rank = lru.nextRank
	// lru.nextRank++
	return pointer.value
}

// func (lru *LRUCache) update(pointer *NodeDouble[int]) {
// 	if lru.linked.last == pointer {
// 		return
// 	}

// 	if lru.linked.first == pointer {
// 		lru.linked.first = lru.linked.first.next
// 	} else {
// 		current := lru.linked.first
// 		for current.next != nil && current.next != pointer {
// 			current = current.next
// 		}
// 		current.next = pointer.next
// 	}
// 	lru.linked.last.next = pointer
// 	lru.linked.last = lru.linked.last.next
// 	lru.linked.last.next = nil
// }

// Dịch chuyển node truyền vào về last node
func (lru *LRUCache) update(pointer *NodeDouble[int]) {
	// cập nhật lại con trỏ của 2 node là node pre và node next để trỏ vào nhau
	pointer.pre.next = pointer.next
	pointer.next.pre = pointer.pre

	// Đưa vị trí node pointer này về last node
	// Như hàm add last node
	pointer.next = lru.linked.header
	pointer.pre = lru.linked.header.pre

	// last node trỏ vào (pre của header là last node)
	lru.linked.header.pre.next = pointer
	lru.linked.header.pre = pointer

	lru.assignNextRank(pointer)
}

// test tính tăng đơn điệu của cache
func (lru *LRUCache) TestIncreaseMonotonic() bool {
	oldRank := lru.linked.header.next.rank
	for start := lru.linked.header.next; start != lru.linked.header; start = start.next {
		if start.rank == -1 || start.rank < oldRank {
			return false
		}
		oldRank = start.rank
	}
	return true
}

// rank tăng đơn điệu
func (lru *LRUCache) assignNextRank(node *NodeDouble[int]) {
	node.rank = lru.nextRank
	lru.nextRank++
}

func (lru *LRUCache) Put(key int, value int) {

	pointer, ok := lru.mm[key]
	var nodeRank *NodeDouble[int]
	if ok {
		// update
		pointer.value = value
		lru.update(pointer)
		nodeRank = pointer
	} else {
		if lru.linked.Length < lru.capacity {
			lru.linked.AddNode(key, value)
			lru.mm[key] = lru.linked.header.pre
			nodeRank = lru.linked.header.pre
		} else {

			// case evict
			nodeRank = lru.evict(key, value)
			// delete(lru.mm, lru.lowest)
			// lru.mm[key] = lru.linked.first
			// lru.linked.first.value = value
			// lru.linked.last.next = lru.linked.first
			// lru.linked.first = lru.linked.first.next
			// lru.linked.last = lru.linked.last.next
			// lru.linked.last.next = nil

		}
	}
	lru.assignNextRank(nodeRank)
}

// evict: Policy là first element là mục tiêu bị loại bỏ khỏi cache
func (lru *LRUCache) evict(key, value int) *NodeDouble[int] {
	// Các bước evict:
	// 1. Loại bỏ first element khỏi hashmap key
	// 2. Loại bỏ first element khỏi double link list
	// 3. Thêm giá trị mới vào tail của double link list
	// 4. Đưa vào key mới vào hashmap
	delete(lru.mm, lru.linked.header.next.key)
	lru.linked.PopFront()
	lru.linked.AddNode(key, value)
	lru.mm[key] = lru.linked.header.pre
	return lru.linked.header.pre
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
