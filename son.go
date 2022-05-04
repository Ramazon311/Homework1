package main 

import "fmt"

func main(){
  var a string
  var b string
    var n int
  fmt.Scan(&a)
  fmt.Scan(&b)
  n=count(a,string(a[len(a)-1]))
  fmt.Println(n)
  if(len(a)>len(b)){
    for i:=0;i<len(a);i++{
      n=0
     for j:=0;j<len(b);j++{
        if string(a[i])==string(b[j]){
            n+=1
            if n==1{
            fmt.Print(string(b[j])," ")
            }}
    }
    }
    }else if(len(a)<=len(b)){
    for i:=0;i<len(b);i++{
        n=0
     for j:=0;j<len(a);j++{
        if string(b[i])==string(a[j]){
            n+=1
            if(n==1){
            fmt.Print(string(a[j])," ")
            }}
    }
    }
    }
}
