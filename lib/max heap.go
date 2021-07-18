package lib

// MaxHeap 最大堆
type MaxHeap struct {
    length int        // 长度
    array  []ElemType // 储存数组
}

// InitMaxHeap 初始化最大堆
func InitMaxHeap() *MaxHeap {
    heap := new(MaxHeap)
    heap.Sort()
    return heap
}

// Insert 插入元素
func (heap *MaxHeap) Insert(i ElemType) {
    // 检查是否需要扩容
    if heap.length < len(heap.array) {
        heap.array[heap.length] = i
    } else {
        heap.array = append(heap.array, i)
    }
    // 获取父子的位置
    parent := Parent(heap.length)
    child := heap.length
    // 将插入节点交换到它的位置
    for heap.array[parent] < heap.array[child] {
        heap.array[parent], heap.array[child] = heap.array[child], heap.array[parent]
        child = parent
        parent = Parent(parent)
    }
    heap.length++
}

// Delete 删除堆顶元素
func (heap *MaxHeap) Delete() {
    parent := 0
    // 用最后一个节点代替根节点
    heap.array[parent] = heap.array[heap.length-1]
    // 对新根节点进行下沉, 直至它的位置
    heap.Down(parent)
    heap.length--
}

// Down 对parent位置的元素进行下沉操作
func (heap *MaxHeap) Down(parent int) {
    var max int
    for {
        left := LeftChild(parent)
        switch {
        // 没有左边数据, 说明已沉到底
        case left+1 > heap.length:
            return
        // 没有右边数据, 只需比较左节点
        case left+2 > heap.length:
            if heap.array[parent] < heap.array[left] {
                heap.array[parent], heap.array[left] = heap.array[left], heap.array[parent]
            }
            return
        // 左右节点数据都有

        // 左节点大于右节点
        case heap.array[left] >= heap.array[left+1]:
            max = left
        // 左节点小于右节点
        case heap.array[left] < heap.array[left+1]:
            max = left + 1
        }
        // 根节点比左右节点都小
        if heap.array[parent] >= heap.array[max] {
            return
        }
        heap.array[parent], heap.array[max] = heap.array[max], heap.array[parent]
        parent = max
    }
}

// Sort 二叉堆排序
func (heap *MaxHeap) Sort() {
    end := Parent(heap.length - 1)
    for end >= 0 {
        heap.Down(end)
        end--
    }
}

// Parent 获取父节点索引
func Parent(i int) int {
    if i == 0 {
        return 0
    }
    return (i - 1) / 2
}

// LeftChild 获取左节点索引
func LeftChild(i int) int {
    return 2*i + 1
}

// RightChild 获取右节点索引
func RightChild(i int) int {
    return 2*i + 2
}

// Length 获取堆长度
func (heap *MaxHeap) Length() int {
    return heap.length
}

// IsEmpty 是否为空
func (heap *MaxHeap) IsEmpty() bool {
    return heap.length == 0
}

// GetAllElem 获取最大堆中所有元素
func (heap *MaxHeap) GetAllElem() []ElemType {
    return heap.array
}

// GetMaxElem 获取最大堆中最大元素
func (heap *MaxHeap) GetMaxElem() ElemType {
    if heap.IsEmpty() {
        return 0
    }
    return heap.array[0]
}
