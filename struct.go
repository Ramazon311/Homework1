package main

import "fmt"

type Cars struct{
  Id int64
  Colour string
  Tip    string
  Model  string
  Number int64
}
func (p *Cars) Number1() int64{
   fmt.Scan(&p.Number)
   return p.Number
}
func main(){
  first := Cars{ 1234,"Black","Pickap","Hunday",128}
  fmt.Println(first)
  first.Number1()
  fmt.Println(first)
}
