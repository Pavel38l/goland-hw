package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	front *ListItem
	back  *ListItem
	len   int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.front == nil {
		l.front = &ListItem{Value: v}
		l.back = l.front
	} else {
		l.front.Prev = &ListItem{Value: v, Next: l.front}
		l.front = l.front.Prev
	}
	l.len++
	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.back == nil {
		l.back = &ListItem{Value: v}
		l.front = l.back
	} else {
		l.back.Next = &ListItem{Value: v, Prev: l.back}
		l.back = l.back.Next
	}
	l.len++
	return l.back
}

func (l *list) Remove(i *ListItem) {
	if i.Prev == nil {
		l.front = i.Next
	} else {
		i.Prev.Next = i.Next
	}

	if i.Next == nil {
		l.back = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}

	if i.Next == nil {
		i.Prev.Next = nil
		l.back = i.Prev
	} else {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	i.Next = l.front
	l.front = i
}

func NewList() List {
	return new(list)
}
