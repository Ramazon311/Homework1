package main
import "fmt"
func main(){
  var n int
    fmt.Scan(&n)
    b:=make([]int,n)
    a:=0
    for i:=0;i<n;i++{
        fmt.Scan(&b[i])
    }
    for i:=0;i<n;i++{
        if b[i]==0 {
        a++
    }
    }
    fmt.Print(a)
}
