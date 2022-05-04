package main
import "fmt"
func main(){
  var a int
    fmt.Scan(&a)
    fmt.Print(a%10,a/100,a/10%10)
}
