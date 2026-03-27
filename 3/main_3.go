package main

import "fmt"

const ArraySize = 10

type StringIntMap struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key   string
	value int
	next  *bucketNode
}

func main() {
	myMap := genMap()

	myMap.Add("apple", 5)
	
  value , found := myMap.Get("apple")
	if found {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Not found")
	}

	myMap.Remove("apple")
	
	value, found = myMap.Get("apple")
	if found {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Not found")
	}

    fmt.Printf("Адрес текущей хэш таблицы %p\n", myMap)
    copiedMap := myMap.Copy()
    fmt.Printf("Адрес копии хэш таблицы %p\n", copiedMap)

	myMap.Add("banana", 10)
	
	if myMap.Exists("banana") {
		fmt.Println("Exists: banana found")
	} else {
		fmt.Println("Exists: banana not found")
	}
	
	myMap.Remove("banana")
	if myMap.Exists("banana") {
		fmt.Println("Exists: banana found")
	} else {
		fmt.Println("Exists: banana not found")
	}
}

func genMap() *StringIntMap {
  m := StringIntMap{}
  for i := range m.array {
      m.array[i] = &bucket{}
  }

  return &m
}

func (h *StringIntMap) Add(key string, value int) {
	_, exists := h.Get(key)
	if !exists {
		index := hash(key)
		newNode := &bucketNode{key: key, value: value}
		newNode.next = h.array[index].head
		h.array[index].head = newNode
	} else {
		index := hash(key)
		currNode := h.array[index].head
		for currNode.key != key {
			currNode = currNode.next
		}
		currNode.value = value
	}
}

func (h *StringIntMap) Remove(key string) {
    index := hash(key)
    h.array[index].remove(key)
}

func (b *bucket) remove(key string) {
    if b.head == nil {
        return
    }
    if b.head.key == key {
        b.head = b.head.next
        return
    }
    prevNode := b.head
    for prevNode.next != nil && prevNode.next.key != key {
        prevNode = prevNode.next
    }
    if prevNode.next != nil {
        prevNode.next = prevNode.next.next
    }
}

func (h *StringIntMap) Get(key string) (int, bool) {
    index := hash(key)
    return h.array[index].get(key)
}

func (h *StringIntMap) Exists(key string) bool {
	index := hash(key)
	return h.array[index].exists(key)
}

func (h *StringIntMap) Copy() *StringIntMap {
	copyMap := genMap()
	
	for i := range h.array {
		currNode := h.array[i].head
		for currNode != nil {
			copyMap.Add(currNode.key, currNode.value)
			currNode = currNode.next
		}
	}
	return copyMap
}

func (b *bucket) get(key string) (int, bool) {
	currNode := b.head
	for currNode != nil {
		if currNode.key == key {
			return currNode.value, true
		}
		currNode = currNode.next
	}
	return 0, false
}

func (b *bucket) exists(key string) bool {
	currNode := b.head
	for currNode != nil {
		if currNode.key == key {
			return true
		}
		currNode = currNode.next
	}
	return false
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}