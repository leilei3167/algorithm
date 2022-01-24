package main

type Stackarray struct {
	array []interface{}
	size  int
}

//入栈
func (this *Stackarray) Push(v interface{}) {
	this.array = append(this.array, v)
	this.size++

}

//出栈
func (this *Stackarray) Pop() interface{} {
	//要出的元素就是最后放入的元素
	if this.size == 0 {
		panic("没有可取元素了！")
	}
	v := this.array[this.size-1]
	//缩容
	arr := make([]interface{}, this.size-1, this.size-1)
	for i := 0; i < this.size-1; i++ {
		arr[i] = this.array[i]

	}
	this.array = arr
	this.size--
	return v

}

func main() {

}
