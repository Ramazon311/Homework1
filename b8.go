package main
import "fmt"
func main(){
  var a,b,c int
    fmt.Scan(&a,&b)
    for i:=a;i<=b;i++{
        if i%7==0{
            c=i       
        }//else{k++}
    }
    if c==0{fmt.Println("NO")
    }else{
     fmt.Println(c)}
}
