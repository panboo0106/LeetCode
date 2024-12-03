package main

import "math/rand"

// @leet start
type RandomizedSet struct {
	valMap map[int]int
}

func Constructor() RandomizedSet {
	mapArr := make(map[int]int)
	return RandomizedSet{valMap: mapArr}
}

func (this *RandomizedSet) Insert(val int) bool {
	if this.valMap[val] == 1 {
		return false
	}
	this.valMap[val] = 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if this.valMap[val] == 1 {
		this.valMap[val] = 0
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	keys := make([]int, 0, len(this.valMap))
	for k, v := range this.valMap {
		if v != 0 {
			keys = append(keys, k)
		}
	}
	randomKey := keys[rand.Intn(len(keys))]
	return randomKey
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
