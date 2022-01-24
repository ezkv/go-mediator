package main

import (
	"fmt"

	m "github.com/ezkv/gomediator/internal/mediator"
	"github.com/ezkv/gomediator/pkg/mediator"
)

type A struct {
	FirstName string
	LastName  string
}
type B struct {
	Name string
}

func main() {

	x := mediator.AddOrGet[A, string]()
	y := mediator.AddOrGet[A, string]()
	_ = mediator.AddOrGet[B, string]()

	mtr := x.NewMediator("test", test)
	mtr2 := y.NewMediator("test", test)
	// a := mediator.New[A, string]()
	// mtr := a.NewMediator("test", test)
	// _ = a.NewMediator("test", test2)
	// _ = a.NewMediator("test", test2)
	// _ = a.NewMediator("test", test2)
	// _ = a.NewMediator("test", test2)
	// _ = a.NewMediator("test", test2)
	// _ = a.NewMediator("test", test2)
	// _ = a.NewMediator("test", test2)

	for i := 0; i < 10; i++ {
		//	go func(mtr mediator.Mediator[A, string], i int) {
		fmt.Println(
			mtr.Mediate(A{FirstName: "mt1: " + fmt.Sprint(i), LastName: fmt.Sprint(i)}))
		fmt.Println(
			mtr2.Mediate(A{FirstName: "mt2: " + fmt.Sprint(i), LastName: fmt.Sprint(i)}))
		//	}(mtr, i)

	}

	// fmt.Scanln()

}
func test(tt *m.MediatePayload[A, string]) {
	//fmt.Printf(" %s from  - %s", tt.Payload.FirstName, "test1")
	//fmt.Print("ack: " + tt.Payload.FirstName + " - ")
	tt.Response = tt.Payload.FirstName
}

func test2(tt *m.MediatePayload[A, string]) {
	//	fmt.Printf(" %s from  - %s", tt.Payload.FirstName, "test2")
	//	fmt.Print("ack2: " + tt.Payload.FirstName + "  ")
	//	fmt.Print("ack: " + tt.Payload.FirstName + " - ")

	//tt.Response = tt.Payload.FirstName

}
