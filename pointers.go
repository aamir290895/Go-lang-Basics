package main

import "fmt"

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func pointers() {

	a := 10

	wg.Add(1)
	go increment(&a) // &a represents the memory 0xc000016088 for variable a

	wg.Wait()

	b := &a

	//var b *int = &a

	fmt.Println(*b) // *b represents the value of 0xc000016088 this memory location

	var st []Student

	st = append(st, Student{Name: "aamir", Age: 28}, Student{Name: "Aakib", Age: 21})

	for i := range st {
		changeName(&st[i])

	}

	fmt.Println(st)

}

func increment(a *int) {

	for *a < 100 {
		*a++

	}

	wg.Done()
}

func changeName(st *Student) {

	st.Name = "Aamir Khan"
}
