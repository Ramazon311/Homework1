package main

//import "fmt"

func count(a string,b string)(int){
 var n int
 for i:=0;i<len(a);i++{
  if string(a[i])==b{
   n++
}
}
 return n
}
//func main(){
// var a string
// fmt.Scan(&a)
// b:=0
// b=Count(a,"1")
// fmt.Println(b)
//}
//count(a,b)
