package main

import (
	"fmt"
	"time"
)

/*
	Channel in golang are very important , because there are used to communicate between go routines.
	channels are bi-directional . i.e we can send and receive data from channel.
	(<- chan) means receiving from channel
	(chan <- 5) means sending data in channel

	When a go routine wants to receive data from channel , that thread is blocked until a data is available in the channel
	A go routine which sends data in the channel is not blocked.const

	So , A deadlock state happens when there is go routine which is listening from a channel,
	but no other routine is found sending the data into the channel.

	Since because of the this , the receiver go routine will be blocked forever.

	Important - When a channel is marked close then the blocked routine is also freed by returning value 0.

*/

func PushNumber(a int, mychan chan int) {
	fmt.Println("Start of Pushnumber")

	mychan <- a
	time.Sleep(20000)
	fmt.Println("End of Pushnumber")
	time.Sleep(20000)
	fmt.Println("Extra of Pushnumber")
}

func main() {
	var startTime = time.Now()
	fmt.Println("main function start")
	mychan := make(chan int)
	go PushNumber(5, mychan)
	fmt.Println(time.Until(startTime).Milliseconds())
	fmt.Println(<-mychan)
	//Using close function we can close a channel and no other value can be send in this channel
	close(mychan)
	fmt.Println("main function finished")

}
