package main
import "fmt"

func main(){
  var workArray[10],a[6],b uint
    for i:=0;i<10;i++{
        fmt.Scan(&workArray[i])
    }
    for i:=0;i<6;i++{
       fmt.Scan(&a[i])
    }
    for i=0;i<6;i+=2{
    b=workArray[a[i]]
    workArray[a[i]]=workArray[a[i+1]]
    workArray[a[i+1]]=b
   }
}
