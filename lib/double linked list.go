package lib

import "errors"

// DoubleLinkedListNode 双链表节点
type DoubleLinkedListNode struct {
    data ElemType
    prev *DoubleLinkedListNode
    next *DoubleLinkedListNode
}

// DoubleLinkedList 双链表
type DoubleLinkedList struct {
    length int
    head   *DoubleLinkedListNode
    tail   *DoubleLinkedListNode
}

// InitDoubleLinkedList 初始化双链表
func InitDoubleLinkedList() *DoubleLinkedList {
    return new(DoubleLinkedList)
}

// AddElem 在双链表尾部添加一个节点
func (list *DoubleLinkedList) AddElem(value ElemType) {
    node := new(DoubleLinkedListNode)
    node.data = value
    // 前方无节点
    if list.length == 0 {
        list.head = node
    } else {
        // 前方有节点
        oldNode := list.tail
        oldNode.next = node
        node.prev = oldNode
    }
    // 处理尾部和长度
    list.tail = node
    list.length++
}

// InsertElem 在双链表前方插入一个节点
func (list *DoubleLinkedList) InsertElem(value ElemType) {
    node := new(DoubleLinkedListNode)
    node.data = value
    oldNode := list.head
    oldNode.prev = node
    node.next = oldNode
    list.head = node
    list.length++
}

// InsertLocationElem 在双链表某位置插入一个节点
func (list *DoubleLinkedList) InsertLocationElem(location int, value ElemType) error {
    node := new(DoubleLinkedListNode)
    node.data = value
    // 不允许插入在头部以外、尾部以外
    if location > list.length || location < 0 {
        return errors.New("无法插入该位置")
    }
    if location == 0 {
        oldNode := list.head
        oldNode.prev = node
        node.next = oldNode
        list.head = node
    } else {
        preElem := list.head
        // 寻找第location个节点
        for j := 0; j < location-1; j++ {
            preElem = preElem.next
        }
        // 此时preElem是第location个节点, nexElem是第location+2个节点
        nexElem := preElem.next
        preElem.next = node
        node.prev = preElem
        nexElem.prev = node
        node.next = nexElem
        // 尾部处理
        if location == list.length {
            list.tail = node
        }
    }
    list.length++
    return nil
}

// DeleteLocationElem 删除单链表某位置的节点
func (list *DoubleLinkedList) DeleteLocationElem(location int) error {
    // 不允许删除头部以外、尾部以外的节点
    if location > list.length-1 || location < 0 {
        return errors.New("该处无节点")
    }
    // 删除头部
    if location == 0 {
        nexElem := list.head.next
        list.head = nexElem
        nexElem.prev = nil
    } else {
        preElem := list.head
        // 寻找第location个节点
        for j := 0; j < location-1; j++ {
            preElem = preElem.next
        }
        // 此时preElem是第location个节点, nexElem是第location+2个节点
        node := preElem.next
        nexElem := node.next
        preElem.next = nexElem
        nexElem.prev = preElem
        // 尾部处理
        if location == list.length-1 {
            list.tail = preElem
        }
    }
    list.length--
    return nil
}

// GetLocationNode 获取某节点
func (list *DoubleLinkedList) GetLocationNode(location int) *DoubleLinkedListNode {
    if location > list.length-1 {
        return nil
    }
    preElem := list.head
    for j := 0; j < location; j++ {
        preElem = preElem.next
    }
    return preElem
}

// GetLength 获取长度
func (list *DoubleLinkedList) GetLength() int {
    return list.length
}

// DeleteAll 删除全部节点
func (list *DoubleLinkedList) DeleteAll() {
    list.head = nil
    list.tail = nil
    list.length = 0
}
