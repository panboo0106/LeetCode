package main

import "math/rand"

// @leet start
type RandomizedSet struct {
	valMap map[int]int
	values []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		valMap: make(map[int]int),
		values: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.valMap[val]; exists {
		return false
	}
	this.valMap[val] = len(this.values)
	this.values = append(this.values, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, exists := this.valMap[val]
	if !exists {
		return false
	}
	lastIdx := len(this.values) - 1
	lastVal := this.values[lastIdx]

	if idx != lastIdx {
		this.values[idx] = lastVal
		this.valMap[lastVal] = idx
	}

	this.values = this.values[:lastIdx]
	delete(this.valMap, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.values[rand.Intn(len(this.values))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @leet end

func main() {}
