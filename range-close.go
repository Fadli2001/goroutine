package main


/**
 * RANGE & CLOSE (CHANNEL)
 * - proses menerima data dari banyak channel bisa dikombinasikan menggunakan perulangan range 
 * - Dan perulangan tidak akan berhenti jika channelnya belum diclose 
 */

 import "fmt"

 func main(){
	channel := make(chan int);

	go func(){
		channel <- 22;
		channel <- 33;
		channel <- 112;
		close(channel)
	}()
	fmt.Println("mulai")
	average(channel)
	fmt.Println("Finish")

	

 } 
 func average(channel chan int){
	sum := 0
	count := 0
	for value := range channel{
		sum += value;
		count++
		fmt.Println("Rata-rata", sum/count)
	}
	
 }