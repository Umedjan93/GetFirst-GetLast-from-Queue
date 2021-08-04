package main

import (
	"fmt"
)

//Цель: создать очередь и метод FirstElement и LastElement
//создаем структуру очереди:
type List struct {
	first *ListNode //первый в очереди
	last *ListNode //последний в очереди
	quantity int //количество в очереди
}
// создаем структуру данных для каждого человека в очереди:
type ListNode struct {
	Prev *ListNode //тот кто был до
	Next *ListNode //тот кто после
	Name string    //имя
	Purchase int   //сумма покупки
}
//расчет количества в очереди:
func (r *List) len() int {	//просто расчитываем длину
	return r.quantity	//и сохраняем значение в переменной Quantity
}

func (r *List) PrintList() {
	fmt.Println("Список очереди: ")
	current := r.first
	for current != nil {
		fmt.Println(current)
		current = current.Next
	}
}

/*func (r *List) GetFirstElement(node *ListNode) {
	current := r.last.Prev
	for r.quantity >= 1 {
		current = r.last.Prev
	}
	fmt.Println(current)
}*/
func (r *List) GetFirstElement(node *ListNode) {
	FirstElement := r.first
	fmt.Println("Первый в списке: ", FirstElement)
	return
}

func (r *List) GetLastElement(node *ListNode) {
	LastElement := r.last
	fmt.Println("Послдений в списке: ", LastElement)
	return
}

//добавление в очередь
func (r *List) Add(node ListNode) {
	if r.len() == 0 { //если он единственный в очереди
		r.first = &node //он и первый
		r.last = &node  //и последний
		r.quantity++    //количество в очереди =0+1 => 1
		return

	} else { //т.е., если в очереди более 1 чел:
		currentlyLast := r.last //доп переменная, чтобы запомнить последнего на данный момент
		r.last.Next = &node		//последний заменяем из Next списка
		r.last = &node			//
		r.last.Prev = currentlyLast	//
		r.quantity++
	}
}

//удаление из очереди первого в списке
func (r *List) Remove(node ListNode) *ListNode {
	if r.len() == 1 { //если он единственный в очереди
		r.first = nil //всё обнуляем
		r.last = nil  //обнуляем
		r.quantity = 0    //количество в очереди теперь 0
		return r.first
	} else { //т.е., если в очереди более 1 чел:
		/*currentlyFirst := r.first.Next //доп переменная, чтобы запомнить первого на данный момент
		r.first.Prev = r.first.Next		//первого Nextом в списке
		r.first.Next = currentlyFirst	//следующий в списке становится первым
		r.quantity-- не сработало() */
		//следующая попытыка также не сработала:
		r.first.Prev.Next = r.first.Next
		r.first.Next.Prev = r.first.Prev
		r.first.Next = nil
		r.first.Prev = nil
		r.first.Name = ""
		r.first.Purchase = 0
		r.quantity--
		return r.first
	}
}

func main()	{
	person := ListNode{
		Prev:     nil,
		Next:     nil,
		Name:     "Mr.A",
		Purchase: 111,
	}
	Queue := List{}
	Queue.Add(person)

	person = ListNode{
		Prev:     nil,
		Next:     nil,
		Name:     "Mr.B",
		Purchase: 200,
	}
	Queue.Add(person)

	person = ListNode{
		Prev:     nil,
		Next:     nil,
		Name:     "Mr.C",
		Purchase: 300,
	}

	Queue.Add(person)
	/*fmt.Println(Queue.PrintList)
	fmt.Println(Queue.first, Queue.last)
	*/
	//Queue.Remove(person)
	//fmt.Println(Queue)
	Queue.PrintList()
	Queue.GetFirstElement(&person)
	Queue.GetLastElement(&person)
}