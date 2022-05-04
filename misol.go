package main 

import "fmt"

func main(){
  var a string
  var b string
  fmt.Scan(&a)
  fmt.Scan(&b)
  if(len(a)>len(b)){
    for i:=0;i<len(a);i++{
     for j:=0;j<len(b);j++{
        if string(a[i])==string(b[j]){
            fmt.Print(string(a[i])," ")
        }
    }
    }
    }else if(len(a)<=len(b)){
    for i:=0;i<len(b);i++{
     for j:=0;j<len(a);j++{
        if string(b[i])==string(a[j]){
            fmt.Print(string(b[i])," ")
        }
    }
    }
    }
}
