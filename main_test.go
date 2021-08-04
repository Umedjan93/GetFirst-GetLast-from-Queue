package main

import "testing"

func TestForEmptyList_GetFirstElement(t *testing.T) {
	l := List{}
	if l.first != nil {
		want := ""
		got := l.first
		t.Errorf("Method getFirstElement is not working correctly! want: %v\n got: %v\n", want, got)
	}
}



func TestForEmptyList_GetLastElement(t *testing.T) {
	l := List{}
	if l.first != nil {
		want := ""
		got := l.first
		t.Errorf("Method getLastElement is not working correctly! want: %v\n got: %v\n", want, got)
	}
}

func TestForCorrectData_GetFirstElement(t *testing.T) {
	person := ListNode{
		Prev:     nil,
		Next:     nil,
		Name:     "HojiBobo",
		Purchase: 99,
	}
	if person.Name != "HojiBobo" && person.Purchase != 99 {
		NameWant := "Hojibobo"
		PurhaseWant := 99
		NameGot := person.Name
		PurchaseGot := person.Purchase
	t.Errorf("Method GetFirstElement is working correctly for NameWant: %v, NameGot: %v, PurchaseWant: %v, PurchaseGot: %v\n",
		NameWant,NameGot,PurhaseWant,PurchaseGot)
	}
}