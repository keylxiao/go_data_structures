package lib

import "fmt"

// google上机: 利用hash table实现员工管理系统
// 1. 不使用数据库, 尽量节省内存, 速度越快越好
// 2. 添加时, 保证按照雇员的id从低到高插入

const (
    HashNum = 7
)

// Emp 员工结构体
type Emp struct {
    Id   int
    Name string
    Next *Emp
}

// EmpLink 员工链表
type EmpLink struct {
    Head *Emp
}

// HashTable 哈希表
type HashTable struct {
    Array [HashNum]EmpLink
}

// Show 显示员工信息
func (peo *Emp) Show() {
    fmt.Printf("链表%d 找到该员工%s\n", peo.Id%HashNum, peo.Name)
}

// Insert 在链表中插入员工
func (link *EmpLink) Insert(emp *Emp) {
    cur := link.Head // 头指针, 最终为emp后方指针
    pre := new(Emp)  // 辅助指针pre一直在cur的前方

    // 检查EmpLink是否为空
    if cur == nil {
        link.Head = emp
        return
    }

    // EmpLink不为空, 寻找emp的位置
    for {
        if cur != nil {
            // 保证员工的id在链表中按顺序排序
            if cur.Id > emp.Id {
                break
            }
            pre = cur
            cur = cur.Next
        } else {
            break
        }
    }
    pre.Next = emp
    emp.Next = cur
}

// Show 展示链表
func (link *EmpLink) Show(num int) {
    // 空链表
    if link.Head == nil {
        fmt.Printf("链表%d为空\n", num)
        return
    }

    cur := link.Head // 头指针
    for {
        if cur != nil {
            fmt.Printf("链表%d,雇员id=%d,姓名:%s\n", num, cur.Id, cur.Name)
            cur = cur.Next
        } else {
            break
        }
    }
}

// FindById 根据id查找员工
func (link *EmpLink) FindById(id int) *Emp {
    cur := link.Head // 头指针
    for {
        // 如果找到对应节点
        if cur != nil && cur.Id == id {
            return cur
        } else if cur == nil { // 如果结尾还没找到
            break
        }
        cur = cur.Next
    }
    return nil
}

// Insert 员工信息插入哈希表
func (table *HashTable) Insert(emp *Emp) {
    // 确定雇员应该添加到哪个链表
    linkNum := table.HashFunc(emp.Id)
    // 插入
    table.Array[linkNum].Insert(emp)
}

// Show 显示所有雇员
func (table *HashTable) Show() {
    for i := 0; i < HashNum; i++ {
        table.Array[i].Show(i)
    }
}

// HashFunc 哈希散列方法
func (table *HashTable) HashFunc(id int) int {
    return id % 7
}

// FindById 哈希表中id查找员工
func (table *HashTable) FindById(id int) *Emp {
    linkNum := table.HashFunc(id)
    return table.Array[linkNum].FindById(id)
}
