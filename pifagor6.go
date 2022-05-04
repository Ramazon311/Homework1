package main
import "fmt"
func main() {
    var a,k string
    t:=0
    fmt.Scan(&a)
    for i:=0;i<len(a);i++{
        for j:=0;j<len(a);j++{
            if a[i]==a[j] && i!=j {t++
                         }
        } 
        if t==0 { k+=string(a[i])
                  } else { t=0 }
    }
    fmt.Print(k)
}
