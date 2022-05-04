package main
import "fmt"
func main(){
  var a,b,c int
    //fmt.Scan(&a)
    fmt.Scan(&a,&b,&c)
    if a<b && a<c && a*a+b*b==c*c{
    fmt.Println("Yes")
}else if b<a && b<c && a*a+b*b==c*c{ 
    fmt.Println("Yes")
}else if c<a && c<b && c*c+a*a==b*b{ 
    fmt.Println("Yes")
}else if a<c && a<b && c*c+a*a==b*b{ 
    fmt.Println("Yes")
}else if c<b && c<a && c*c+b*b==a*a{ 
    fmt.Println("Yes")
}else if b<c && b<a && c*c+b*b==a*a{ 
    fmt.Println("Yes")
}else{fmt.Println("NO")}
}
