package main

import "math/rand"

type RandomizedSet struct {
	vals  []int
	index map[int]int
}

func Constructor() RandomizedSet {
	this := RandomizedSet{}
	this.vals = make([]int, 0)
	this.index = make(map[int]int)
	return this
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.index[val]; ok {
		return false
	}
	this.vals = append(this.vals, val)
	this.index[val] = len(this.vals) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.index[val]; !ok {
		return false
	}
	index := this.index[val]
	last := this.vals[len(this.vals)-1]
	this.vals[index] = last
	this.index[last] = index
	this.vals = this.vals[:len(this.vals)-1]
	delete(this.index, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	randomIndex := rand.Intn(len(this.vals))
	return this.vals[randomIndex]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
