package main
import "fmt"
func main(){
 var a,b string
 fmt.Scan(&a)
    for i:=len(a)-1;i>=0;i--{
        b+=string(a[i])
    }
    if a==b{
        fmt.Print("Палиндром")
    }else{fmt.Print("Нет")}
  fmt.Print(a,b)
}
