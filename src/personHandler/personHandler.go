package personHandler

import (
	"fmt"

)

type Person struct {
	Name string
	Age int
}

func NewPerson(name string,age int)(*Person) {
	return &Person{Name: name,Age: age}
}

type PersonHandler interface {
	Batch(origs <-chan Person) <-chan Person
	Handle(orig *Person)
}

func NewPersonHandler() (PersonHandler){
	return &PersonHandlerImpl{}
}


type PersonHandlerImpl struct {}

func (handle PersonHandlerImpl) Handle(orig* Person) {

}

func (handler PersonHandlerImpl) Batch(origs <-chan Person) <-chan Person {
	dests := make(chan Person,10)
	go func() {
		for {
			p, ok := <-origs
			if !ok {
				close(dests)
				break
			}
			handler.Handle(&p)
			dests <- p
		}
	}()
	return dests
}

func FecthPerson(origs chan<- Person){
	go func() {
		for i := 1; i < 200; i++ {
			p := Person{Name: "ruima", Age: i}
			origs <- p
		}
		close(origs)
	}()

}

func GetPersonHandler() (PersonHandler){
	return PersonHandlerImpl{}
}

func SavePerson(dest <-chan Person) (<-chan int){
	sign := make(chan int)
	go func () {
		for {
			p, ok := <-dest
			if !ok {
				break
			}
			fmt.Printf("SavePerson %v\n", p)
		}
		fmt.Println("dest closed")
		sign <- 1
	}()
	return sign
}