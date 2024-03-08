package main

import (
	"math/rand"
)

func main() {

}

// example
// {100: 0, 200: 1, 300: 2}
// [100, 200, 300]
type RandomizedSet struct {
	m    map[int]int // num to index (index is for nums)
	nums []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m:    map[int]int{},
		nums: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; !ok {
		this.m[val] = len(this.nums)
		this.nums = append(this.nums, val)

		return true
	}

	return false
}

// example
// {100: 0, 200: 1, 300: 2}
// [100, 200, 300]
// want to remove 200 (index 1)
// 1. [100, 200, 300] -> [100, 300, 300]
// 2. {100: 0, 200: 1, 300: 2} -> {100: 0, 200: 1, 300: 1}
// 3. [100, 300, 300] -> [100, 300]
// 4. {100: 0, 200: 1, 300: 1} -> {100: 0, 300: 1}
func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.m[val]; ok {
		// overwrite the delete one with the last one
		this.nums[index] = this.nums[len(this.nums)-1]

		// update map index
		this.m[this.nums[len(this.nums)-1]] = index

		// remove last one in slice
		this.nums = this.nums[:len(this.nums)-1]

		delete(this.m, val)

		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.nums)) // 0 <= index < len(this.nums)
	return this.nums[index]
}
