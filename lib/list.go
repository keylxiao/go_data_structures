package lib

import (
    "errors"
)

type ElemType int

const (
    MAX = 20
    ADD = 20
)

// List 线性表
type List struct {
    num    [MAX]ElemType // 储存切片
    length int // 长度
}

// InitList 新建线性表
func InitList() *List {
    newList := new(List)
    newList.length = 0
    return newList
}

// Length 返回线性表的长度
func (list *List)Length() int {
    return list.length
}

// InsertElem 在线性表中插入一个数据
func (list *List)InsertElem(location int, data *ElemType) error {
    // 线性表满
    if list.length == MAX {
        return errors.New("线性表已满")
    }
    // 位置错误
    if location < 1 || location > list.length+1 {
        return errors.New("插入位置错误")
    }
    // 插入数据, 后面的数据向后延位
    for j := list.length - 2; j >= location; j-- {
        list.num[j+1] = list.num[j]
    }
    list.num[location-1] = *data
    list.length++
    return nil
}

// DeleteElem 在线性表中删除一个数据
func (list *List)DeleteElem(location int) error {
    // 线性表没有数据
    if list.length == 0 {
        return errors.New("线性表为空")
    }
    // 位置错误
    if location < 1 || location > list.length+1 {
        return errors.New("删除位置错误")
    }
    // 删除数据, 后面的数据向前顺位
    for j := location - 1; j < list.length; j++ {
        list.num[j] = list.num[j+1]
    }
    list.length--
    return nil
}

// ClearList 清空线性表
func (list *List)ClearList() error {
    // 线性表中没有数据
    if list.length == 0 {
        return errors.New("线性表为空")
    }
    for i := range list.num {
        list.num[i] = 0
    }
    return nil
}

// EmptyList 检查线性表是否为空表
func (list *List)EmptyList() bool {
    if list.length == 0 {
        return true
    }
    return false
}

// GetElem 获取某位置的数据
func (list *List)GetElem(location int) (ElemType, error) {
    // 没有该元素
    if list.length < location-1 {
        return 0, errors.New("没有该元素")
    }
    return list.num[location-1], nil
}

// LocateElem 匹配某数据的位置
func (list *List)LocateElem(data *ElemType) (int, error) {
    for i := range list.num {
        if list.num[i] == *data {
            return i, nil
        }
    }
    return 0, errors.New("该数据不存在")
}


