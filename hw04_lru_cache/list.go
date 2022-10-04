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

type Queue struct {
	size  int
	front *ListItem
	back  *ListItem
}

func (l *Queue) Front() *ListItem {
	// Получение начального элемента списка
	return l.front
}

func (l *Queue) Back() *ListItem {
	// Получение последнего элемента списка
	return l.back
}

func (l *Queue) Len() int {
	// Получение длины списка
	return l.size
}

func (l *Queue) PushFront(v interface{}) *ListItem {
	// Поместить значение в начало списка
	newFront := &ListItem{}

	newFront.Value = v
	newFront.Prev = nil
	newFront.Next = l.front

	if l.front != nil {
		l.front.Prev = newFront
	}
	if l.back == nil {
		l.back = newFront
	}

	l.front = newFront
	l.size++

	return newFront
}

func (l *Queue) PushBack(v interface{}) *ListItem {
	// Поместить значение в конец списка
	newBack := &ListItem{}

	newBack.Value = v
	newBack.Next = nil
	newBack.Prev = l.back

	if l.back != nil {
		l.back.Next = newBack
	}
	if l.front == nil {
		l.front = newBack
	}

	l.back = newBack
	l.size++

	return newBack
}

func (l *Queue) Remove(i *ListItem) {
	// Удалиние элемента из списка
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	if i.Next == nil {
		l.back = i.Prev
	}

	if i.Prev == nil {
		l.front = i.Next
	}

	if l.size != 0 {
		l.size--
	}
}

func (l *Queue) MoveToFront(i *ListItem) {
	// Перемещение элемента в начало списка
	v := i.Value
	l.Remove(i)
	l.PushFront(v)
}

func NewList() *Queue {
	return &Queue{}
}
