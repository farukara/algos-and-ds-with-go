package main
//mean heap implementaiton

import (
    "fmt"
)

type Heap struct {
    Items []int
}

func (h *Heap) Peek() int {
    if len(h.Items)==0 {
        return -1
    } else {
        return h.Items[0]
    }
}

func (h *Heap) Push(n int) {
    h.Items = append(h.Items, n)
    index := len(h.Items)-1
    parentIndex := (index-1)/2
    for h.Items[index] < h.Items[parentIndex] {
        h.Items[index], h.Items[parentIndex] = h.Items[parentIndex], h.Items[index]
        index , parentIndex = parentIndex, (parentIndex-1)/2
    }
}

func (h *Heap) Pop() int {
    if len(h.Items) == 0 {return -1}
    result := h.Items[0]
    h.Items[0], h.Items[len(h.Items)-1] = h.Items[len(h.Items)-1], h.Items[0]
    h.Items = h.Items[:len(h.Items)-1]
    index := 0
    for {
        leftChild :=index*2+1
        rightChild := index*2+2
        if leftChild>len(h.Items)-1 {
            //we are at the bottom of the heap 
            return result
        } else if rightChild>len(h.Items)-1 {
            //there is only left child
            if h.Items[index] > h.Items[leftChild] {
                h.Items[index], h.Items[leftChild] = h.Items[leftChild], h.Items[index]
                return result //because we are at the bottom
            } else {
                return result
            }
        } else {
            //both childs present
            if h.Items[leftChild] < h.Items[rightChild] {
                //left child is smaller
                if h.Items[index] > h.Items[leftChild] {
                    h.Items[index], h.Items[leftChild] = h.Items[leftChild], h.Items[index]
                    index = leftChild
                } else {
                    return result
                }
            } else {
                //right child is smaller
                if h.Items[index] > h.Items[rightChild] {
                    h.Items[index], h.Items[rightChild] = h.Items[rightChild], h.Items[index]
                    index = rightChild
                } else {
                    return result
                }
            }
        }
    }
}

func main () {
    test()
}

func test() {
    var h Heap
    items := []int{1,3,2,-4,2,1,-8}
    for _,item := range items {
        h.Push(item)
        fmt.Println("pushing", item)
        fmt.Println("after push min is:", h.Peek())
        fmt.Println(h.Items)
        fmt.Println("")
    }
    for i:=0;i<8;i++ {
        fmt.Println("popping from", h.Items)
        h.Pop()
        fmt.Println("after pop min is:", h.Peek())
        fmt.Println("")
    }
}
