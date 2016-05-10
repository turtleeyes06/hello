package main

import "fmt"
type Unicorn struct{
	Name string
	Age int
	weight int
}

type TurtleIsland struct{
	Name string
	Age int
	weight int
}

func (u *Unicorn) SayName()(){
	fmt.Println(u.Name,"\t",u.Age,"\t",u.weight)	
}
func (t *TurtleIsland) SayName()(){
	fmt.Println(t.Name,"\t",t.Age,"\t",t.weight)	
}
func main(){
	chandara := new(Unicorn)
	chandara.Name = "Channie"
	chandara.Age = 21
	chandara.weight = 130

	//same as
	nick := &TurtleIsland{
	Name:	"Nick Carroll",
	Age:	28,
	weight:	200,
}
	nick.SayName()
	chandara.SayName()
}


