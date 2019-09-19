package main
 
import "unsafe"
const (
    a = "刘权辉"
    b = len(a)
    c = unsafe.Sizeof(a)
)
 
func main(){
    println(a, b, c)
}