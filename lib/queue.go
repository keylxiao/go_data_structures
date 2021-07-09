package lib

import "errors"

// Queue 队列
type Queue struct {
    maxLength int              // 队列最长
    end    int              // 队尾(当前栈高度)
    list   *SingleLinkedList // 单链表模拟队列
}

// InitQueue 初始化队列
func InitQueue() *Queue {
    queue := new(Queue)
    queue.maxLength = MAX
    queue.end = -1
    queue.list = new(SingleLinkedList)
    return queue
}

// Push 入队
func (queue *Queue) Push(value ElemType) error {
    // 队满
    if queue.end == queue.maxLength-1 {
        return errors.New("队列满")
    }
    queue.end++
    queue.list.AddElem(value)
    return nil
}

// Pop 出队
func (queue *Queue) Pop() (ElemType, error) {
    // 队空
    if queue.end == -1 {
        return 0, errors.New("队列空")
    }
    queue.end--
    node := queue.list.GetLocationNode(0)
    _ = queue.list.DeleteLocationElem(0)
    return node.data, nil
}

// List 遍历队列
func (queue *Queue) List() ([]ElemType, error) {
    if queue.end == -1 {
        return nil, errors.New("队列空")
    }
    var array []ElemType
    for i := 0; i < queue.end; i++ {
        array = append(array, queue.list.GetLocationNode(i).data)
    }
    return array, nil
}

// Length 队列长度
func (queue *Queue) Length() int {
    return queue.end + 1
}

// Clear 清空队列
func (queue *Queue) Clear() error {
    if queue.end == -1 {
        return errors.New("队列空")
    }
    queue.list.DeleteAll()
    return nil
}
