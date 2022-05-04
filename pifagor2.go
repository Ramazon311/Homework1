package main
import "fmt"
func main(){
 var n,m int64
    fmt.Scan(&n)
    for i:=n;i>0;i/=10{
        if(i%10<4){
            m=m*10+(i%10)*(i%10)
            fmt.Println(m,i%10)
        }else{m=m*100+(i%10)*(i%10)}
              fmt.Println(m,i%10)
    }
    fmt.Print(m)
}
