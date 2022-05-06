package common

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	fmt.Println(PathConcat("1path", "/2path", "3path/", "/4path/", "/path5", "path6"))
}

func TestQueue(t *testing.T) {
	queue := NewQueue()
	queue.Push("1")
	queue.Push("2")
	queue.Push("3")
	fmt.Println(queue.Len())
	d, _ := queue.Pop()
	fmt.Println(d, queue.Len())
}

func TestPoint(t *testing.T) {

}
