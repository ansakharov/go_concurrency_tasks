package main

// code: https://go.dev/play/p/BKqT33xVV0R
func main() {
	ch := make(chan int)

	<-ch
}

/*func main() {
	select {}
	fmt.Println("finish")
}*/
