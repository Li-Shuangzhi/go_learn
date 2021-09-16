package main

type MyHashSet struct {
	nums []bool
}

// Constructor /** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{nums: make([]bool, 1000001)}
}

func (this *MyHashSet) Add(key int) {
	this.nums[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.nums[key] = false
}

// Contains /** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.nums[key]
}
