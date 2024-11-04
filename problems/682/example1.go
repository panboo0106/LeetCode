package main

import (
	"container/list"
	"strconv"
)

// @leet start
func calPoints(operations []string) int {
	result := 0
	stackList := list.New()
	for i := 0; i < len(operations); i++ {
		if operations[i] == "C" {
			stackList.Remove(stackList.Front())
		} else if operations[i] == "D" {
			value, _ := stackList.Front().Value.(int)
			stackList.PushFront((value * 2))
		} else if operations[i] == "+" {
			frontFirst, _ := stackList.Front().Value.(int)
			frontSecond, _ := stackList.Front().Next().Value.(int)
			stackList.PushFront(frontFirst + frontSecond)
		} else {
			num, _ := strconv.Atoi(operations[i])
			stackList.PushFront(num)
		}
	}
	for e := stackList.Front(); e != nil; e = e.Next() {
		result += e.Value.(int)
	}
	return result
}

// @leet end

func main() {}
