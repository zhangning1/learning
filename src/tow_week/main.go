package main

import "fmt"

func main() {
	name := "张三"
	student, err := SearchStudent(name)
	if err != nil {
		fmt.Printf("get student(name=%s) err(%+v) \n", name, err)
		return
	}

	fmt.Printf("get strudent(name=%s) row(%+v) success \n", name, student)
}
