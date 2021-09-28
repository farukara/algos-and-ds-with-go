package main

import "fmt"

type Node struct {
    data int
    next *Node
}

type LinkedList struct {
    size int
    head *Node
    tail *Node
}

func AddToTail (p *LinkedList, data int) {
    if p.size == 0 {
        newNode := Node{data: data}
        p.head = &newNode
        p.tail = &newNode
        p.size++
    } else { 
        newNode := Node{data: data}
        p.tail.next = &newNode
        p.tail = &newNode
        p.size++
    }
}

func AddToHead (p *LinkedList, data int) {
    if p.size == 0 {
        newNode := Node{data: data}
        p.head = &newNode
        p.tail = &newNode
        p.size++
    } else {
        newNode := Node{data: data}
        newNode.next = p.head
        p.head = &newNode
        p.size++
    }
}

func RemoveTail (p *LinkedList) {
    if p.size == 0 {
    } else {
        current := p.head
        for i:=1; i<p.size-1; i++ {
            current = current.next
        }
        p.tail = current
        current.next = nil
        p.size--
    }
}

func RemoveHead (p *LinkedList) {
    if p.size == 0 {
    } else if p.size == 1 {
        p.head = nil
        p.tail = nil
        p.size--
    } else {
        p.head = p.head.next
        p.size--
    }
}
func PrintList (p *LinkedList) {
    if p.size == 0 {
        fmt.Println("linked list is empty")
    }; {
        current := p.head
        for i :=0; i<p.size; i++ {
            fmt.Print(current.data, "-->")
            current = current.next
        }
        if p.size != 0 {
            fmt.Println("nil")
        }
    }
}
func main() {
    mylinked := LinkedList{}
    AddToTail(&mylinked, 1)
    AddToTail(&mylinked, 2)
    AddToTail(&mylinked, 3)
    AddToHead(&mylinked, 0)
    AddToHead(&mylinked, -1)
    AddToHead(&mylinked, -2)
    PrintList(&mylinked)
    RemoveTail(&mylinked)
    PrintList(&mylinked)
    RemoveTail(&mylinked)
    PrintList(&mylinked)
    RemoveHead(&mylinked)
    RemoveHead(&mylinked)
    RemoveHead(&mylinked)
    RemoveHead(&mylinked)
    PrintList(&mylinked)
}
