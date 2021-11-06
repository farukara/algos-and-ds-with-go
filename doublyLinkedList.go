package main

import (
    "fmt"
)

type Node struct {
    data int
    prev *Node
    next *Node
}

type DLinkedList struct {
    head *Node
    tail *Node
    size int
}

func AddToTail(dll DLinkedList, data int) DLinkedList{
    newNode := Node{data:data}
    if dll.size == 0 {
        dll.head = &newNode
        dll.tail = &newNode
        dll.size++
    } else {
        dll.tail.next = &newNode
        newNode.prev = dll.tail
        dll.tail = &newNode
        dll.size++
    }
    return dll
}

func AddToHead(dll DLinkedList, data int) DLinkedList{
    newNode := Node{data:data}
    // newNode.data = data
    if dll.size == 0 {
        dll.head = &newNode
        dll.tail = &newNode
        dll.size++
    } else {
        dll.head.prev = &newNode
        newNode.next = dll.head
        dll.head = &newNode
        dll.size++
    }
    return dll
}

func PopTail(dll *DLinkedList) int {
    node := new(Node)
    if dll.size == 0 {
        fmt.Println("doubly linked list is empty")
    } else {
        node = dll.tail
        dll.tail = dll.tail.prev
        dll.tail.next = nil
        node.prev = nil
        dll.size--
    }
    return node.data
}

func PopHead(dll *DLinkedList) int {
    node := new(Node)
    if dll.size == 0 {
        fmt.Println("doubly linked list is empty")
    } else {
        node = dll.head
        dll.head = dll.head.next
        dll.head.prev = nil
        node.next = nil
        dll.size--
    }
    return node.data
}

func DelItem (dll *DLinkedList, item int) {
    if dll.size == 0 {
        fmt.Println("doubly linked list is already empty")
        return
    } else {
        current := dll.head
        for  i:=0; i<dll.size; i++ {
            if current.data == item {
                current.prev.next = current.next
                current.next.prev = current.prev
                dll.size--
                break
            } else if current.next !=nil {
                current = current.next
            } else {
                fmt.Println(item, "is not in doubly linked list")
            }
        }
    }
}

func Contains(dll DLinkedList, item int) bool {
    if dll.size == 0 {
        return false
    } else {
        current := dll.head
        for i:=0; i<dll.size; i++ {
            if current.data == item {
                return true
            } else {
                current = current.next
            }
        }
        return false
    }
}
func PrintDLL (dll *DLinkedList) {
    if dll.size == 0 {
        fmt.Println("Doubly Linked List is empty!!!")
    } else {
        current := dll.head
        for i:=0; i<dll.size; i++ {
            fmt.Print(current.data, "<->")
            current = current.next
        }
    }
    fmt.Println("\n")
}
func testDLL () {
    mydll := DLinkedList{}
    fmt.Println("adding tail 10-90 inclusive")
    for i:=10; i<100; i+=10 {
        mydll = AddToTail(mydll, i)
    }
    PrintDLL(&mydll)
    fmt.Println("adding head 200-240 inclusive")
    fmt.Println("Size:", mydll.size)
    fmt.Println("Head:", mydll.head)
    fmt.Println("Tail:", mydll.tail)
    for i:=200; i<250; i+=10 {
        mydll = AddToHead(mydll, i)
    }
    PrintDLL(&mydll)

    fmt.Println("popping tail")
    tail := PopTail(&mydll)
    PrintDLL(&mydll)
    fmt.Println("tail value:", tail, "\n")

    fmt.Println("popping tail")
    tail = PopTail(&mydll)
    PrintDLL(&mydll)
    fmt.Println("tail value:", tail, "\n")

    fmt.Println("popping head")
    head := PopHead(&mydll)
    PrintDLL(&mydll)
    fmt.Println("head value:", head, "\n")

    fmt.Println("deleting 10")
    DelItem(&mydll, 10)
    PrintDLL(&mydll)

    fmt.Println("deleting 120")
    DelItem(&mydll, 120)
    PrintDLL(&mydll)

    fmt.Println("is 220 in dll?")
    fmt.Println(Contains(mydll, 220))
    PrintDLL(&mydll)

    fmt.Println("is 333 in dll?")
    fmt.Println(Contains(mydll, 333))
    PrintDLL(&mydll)

    fmt.Println("Size:", mydll.size)
    fmt.Println("Size:", mydll.size)
    fmt.Println("Head:", mydll.head)
    fmt.Println("Tail:", mydll.tail)

}
func main() {
    /* dll := DLinkedList{}
    PrintDLL(&dll) */
    testDLL()
}
