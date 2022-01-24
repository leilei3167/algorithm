package main

import (
	"sync"
)

/* 一般很多编程语言库，会把不可重复集合（Collection）命名为 Set，这个 Set 中文直译为集合
不可重复集合 Set 存放数据，特点就是没有数据会重复，会去重。你放一个数据进去，
再放一个数据进去，如果两个数据一样，那么只会保存一份数据。

集合 Set 可以没有顺序关系，也可以按值排序，算一种特殊的列表。

因为我们知道字典的键是不重复的，所以只要我们不考虑字典的值，就可以实现集合，我们来实现存整数的集合 Set： */
//创造集合的结构体
type Set struct {
	m   map[int]struct{} //map的key不能重复，因此用map实现
	len int
	sync.RWMutex
}

//初始化一个空的set，容量为输入的cap
func NewSet(cap int64) *Set {
	temp := make(map[int]struct{}, cap)
	return &Set{
		m: temp,
	}
} //因为我们只需要用到map的key，不需要存值，所以用空结构体（不占内存空间）

//增加一个元素
func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = struct{}{} //往map添加key
	s.len = len(s.m)
}

//删除一个元素
func (s *Set) Remove(item int) {
	s.Lock()
	defer s.Unlock()
	if s.len == 0 {
		return
	}
	delete(s.m, item) //从map删除
	s.len = len(s.m)
}

//查看某个元素是否再
func (s *Set) Has(item int) bool {
	_, ok := s.m[item]
	return ok
}
