package main
import "fmt"

func Panic(m int8){
  if m==0{
 return  panic("000")}
}
func main(){
 fmt.Print("ssalo")
 Panic(0)
// fmt.Print("slaom")
}
