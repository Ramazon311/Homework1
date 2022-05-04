package main
import "fmt"
func main() {
str1 := "тестестестест"
var str2 string
for i:=len(str1)-1;i<=0;i--{
 str2 += string([]rune(str1)[i])}
fmt.Println(str2)
}
