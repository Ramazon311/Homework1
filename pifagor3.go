package main
import (
    "bufio"
	"fmt"
	"os"

    
)
func main(){
 //var a string
    a,_:= bufio.NewReader(os.Stdin).ReadString('\n')
    if ((a[0]>=224 && a[0]<=255) || (a[0]) && a[len(a)-1]=='.'{
        fmt.Print("Right")
    }else{fmt.Print("Wrong")}
}
