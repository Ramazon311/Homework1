package main
import "fmt"
func main(){
  var a int
    fmt.Scan(&a)
    fmt.Printf("It is %d hours %d minutes.",a/3600,((a-a%60)-(a/3600)*3600)/60)
}
