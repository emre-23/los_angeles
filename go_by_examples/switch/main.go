package main
import (
	"fmt"
	"time"
)

func main() {
	i :=2
	fmt.Print("write ", i, " as ")
	switch i {
		case 1: 
		fmt.Println("one")
		case 2: 
		fmt.Println("two")
		case 3: 
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's a week day")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	case t.Hour() > 12:
		fmt.Println("It's after noon")	
	}
	whatamI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Printf("%v is a bool\n",t)
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatamI(true)
	whatamI(1)
	whatamI("hey")

	//fmt.Println(time.Now().Weekday())
	//fmt.Println(time.Now())
}