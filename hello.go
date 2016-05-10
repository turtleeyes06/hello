package main

import "fmt"
type Unicorn struct{
	Name string
	Age int
	weight int
}

func (u *Unicorn) SayName()(){
	fmt.Println(u.Name,"\t",u.Age,"\t",u.weight)	
}

func main(){
	chandara := new(Unicorn)
	chandara.Name = "Channie"
	chandara.Age = 21
	chandara.weight = 130

	//same as
	nick := &Unicorn{
	Name:	"Nick Carroll",
	Age:	28,
	weight:	200,
}
	nick.SayName()
	chandara.SayName()
}


