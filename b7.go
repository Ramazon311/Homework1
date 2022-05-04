package main
import "fmt"
func main(){
 var n,k int
    fmt.Scan(&n)
    for i:=n;i>0;i/=10{
      k+=i%10
      //fmt.Print(k," ")
    }
    n=0
  for i:=k;i>0;i/=10{
      n+=i%10
      //fmt.Print(k," ")
    }
   fmt.Println(n)
}
