package main 
import "fmt"
func main(){
 var a string
    var c byte
    fmt.Scan(&a)
    c=a[0]
    for i:=0;i<len(a);i++{
       //fmt.Println(a[i])
        if a[i]>c {
            c=a[i]
        }}
        fmt.Print(string(c))
}
