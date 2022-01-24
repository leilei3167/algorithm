package main

import (
	"fmt"
	"math"

	"sync"

	"github.com/OneOfOne/xxhash"
)

//设定一个扩容因子，这里设置为0.75
const (
	expandFactor = 0.75
)

//哈希表和键值对的结构
type HashMap struct {
	array        []*keyPairs //哈希表数组，每个元素是一个键值对
	capacity     int         //数组容量
	len          int         //已添加键值对元素数量
	capacityMask int         //掩码，等于capacity-1，用来计算数组的下标
	//并发安全
	lock sync.Mutex
}

//键值对，形成一个链表，就是一个个的节点
type keyPairs struct {
	key   string      //键
	value interface{} //值为任意类型
	next  *keyPairs
}

//初始化哈希表
func NewHashMap(capacity int) *HashMap {
	//默认大小为16
	defaultCapacity := 1 << 4 //左移4位，相当于10000,等于十进制的16
	if capacity <= defaultCapacity {
		//如果传入的大小小于16 则就使用默认大小
		capacity = defaultCapacity
	} else {
		//否则，实际大小为大于capacity的第一个2的k次方
		capacity = 1 << (int(math.Ceil(math.Log2(float64(capacity))))) //????
	}
	//新建一个哈希表
	m := new(HashMap)
	m.array = make([]*keyPairs, capacity, capacity)
	m.capacity = capacity
	m.capacityMask = capacity - 1
	return m

}

//返回哈希表已添加的元素数量
func (m *HashMap) Len() int {
	return m.len
}

//计算key的哈希值，此步可根据不同函数进行变化
var hashAlgorithm = func(key []byte) uint64 {
	h := xxhash.New64()
	h.Write(key)
	return h.Sum64()
}

//对键求哈希值，并计算下标，放在哪个array
func (m *HashMap) hashIndex(key string, mask int) int {
	//求哈希
	hash := hashAlgorithm([]byte(key))
	//求下标
	index := hash & uint64(mask) //也可以直接求余，恒等式 hash % 2^k = hash & (2^k-1)
	return int(index)
}

//增加
func (m *HashMap) Put(key string, value interface{}) {
	//加锁
	m.lock.Lock()
	defer m.lock.Unlock()
	//键值对要放的哈希表数组的下标
	index := m.hashIndex(key, m.capacityMask)
	//取出下标对应的元素
	element := m.array[index] //因为array存放的是地址，为nil表示还没有元素
	if element == nil {
		//说明元素还是空的，空链表，直接赋值
		m.array[index] = &keyPairs{
			key:   key,
			value: value,
		}
	} else {
		//否则表示不是空链表
		//链表最后一个键值对
		var lastPairs *keyPairs
		//遍历链表查看元素是否存在，已存在就更新，否则找到最后一个键值对
		for element != nil {
			if element.key == key {
				//更新值并返回
				element.value = value
				return
			}
			lastPairs = element
			element = element.next

		}
		//找不到键值对，将新键值对加到尾部
		lastPairs.next = &keyPairs{
			key:   key,
			value: value,
		}
	}
	newLen := m.len + 1
	//判断是否需要扩容
	if float64(newLen)/float64(m.capacity) >= expandFactor {
		//新建一个原来两倍大的哈希表
		newM := new(HashMap)
		newM.array = make([]*keyPairs, 2*m.capacity, 2*m.capacity)
		newM.capacity = 2 * m.capacity
		newM.capacity = 2*m.capacity - 1
		//遍历老的哈希表，将键值对赋给新哈希表
		for _, pairs := range m.array {
			for pairs != nil {
				newM.Put(pairs.key, pairs.value)
				pairs = pairs.next
			}
		}
		//替换老的哈希表
		m.array = newM.array
		m.capacity = newM.capacity
		m.capacityMask = newM.capacityMask
	}
	m.len = newLen

}

//根据key获取键值对
func (m *HashMap) Get(key string) (value interface{}, ok bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	//先通过哈希函数计算该key对应的数组的下标
	index := m.hashIndex(key, m.capacityMask)
	fmt.Println(index)
	//找到对应的数组
	elment := m.array[index]
	//遍历看元素是否存在
	for elment != nil {
		if elment.key == key {
			return elment.value, true
		}
		elment = elment.next //不断后移
	}
	return

}

func main() {
	a := NewHashMap(16)
	a.Put("1", 66) //首先会通过哈希函数计算要存放的数组标号
	a.Put("2", 77)
	x, _ := a.Get("1")
	fmt.Println(x)
}
