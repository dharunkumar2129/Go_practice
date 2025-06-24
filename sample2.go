package main

/*func main() {
	cha := make(chan int)
	go gen(cha)
	time.Sleep(3 * time.Second)
	for i := range cha {
		fmt.Println(i)
	}
}*/

func gen(a chan int) {
	for i := 0; i < 10; i++ {
		a <- i
	}
	close(a)
}
