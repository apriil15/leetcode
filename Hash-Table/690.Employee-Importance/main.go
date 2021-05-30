package main

import "fmt"

func main() {
	employees := []*Employee{
		{Id: 1, Importance: 5, Subordinates: []int{2, 3}},
		{Id: 2, Importance: 3, Subordinates: []int{4}},
		{Id: 3, Importance: 4, Subordinates: []int{}},
		{Id: 4, Importance: 1, Subordinates: []int{}}}
	result := getImportance(employees, 1)
	fmt.Println(result)
}

func getImportance(employees []*Employee, id int) int {
	// 建字典
	m := make(map[int]Employee)
	for _, employee := range employees {
		m[employee.Id] = *employee
	}

	importance := 0
	if _, ok := m[id]; ok {
		importance += getImportanceValueOf(id, m)
	}
	return importance
}

// 用遞迴方式處理
func getImportanceValueOf(id int, m map[int]Employee) int {
	employee := m[id]
	value := employee.Importance

	if len(employee.Subordinates) == 0 { // 沒有直屬下屬
		return value
	}

	for _, subId := range employee.Subordinates {
		value += getImportanceValueOf(subId, m) // 可能還會有下屬，所以用遞迴
	}
	return value
}

type Employee struct {
	Id           int
	Importance   int   // 重要點數
	Subordinates []int // 直屬的下屬 ID
}
