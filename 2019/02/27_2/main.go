package main
/* https://twitter.com/bot_golang/status/1100803875299254273
* runtime deadlock error
* main exited
* hello exitedmain exited
 */
import (
	"fmt"
	"time"
)

func hello(done chan int) {
	done <- 5
	done <- 5
	fmt.Println("hello exited")

}
func main() {
	done := make(chan int)
	go hello(done)
	time.Sleep(3 * time.Second)
	fmt.Println("main exited")
}
