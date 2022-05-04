package main
import "fmt"
func main(){
    var n,b int
    fmt.Scan(&n)
    b=123
    a:=make([]int,n)
    for i:=0;i<n;i++{
    fmt.Scan(&a[i])
    if a[i]<b{
       b=a[i]}
}
 d:=0
    for i:=0;i<n;i++{
        if a[i]==b{
           d++}
}
fmt.Println(d,b)
}
