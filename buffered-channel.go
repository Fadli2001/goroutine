package main 


/**
 * APA ITU BUFFERED  CHANNEL ?
 * Channel yang mempunyai kapasitas dan tidak akan melakukan blocking selama data yang ada tidak
 * melebihi kapasitas 
 * 
 * KARAKTERISTIK BUFFERED CHANNEL 
 * - Mempunyai kapasitas dan dapat ditentukan kapasitasnya saat dekalrasi
 * - Mempunyai function cap() untuk mengecek kapasitas dari buffered channel
 * - mempunyai function len() untuk mengecek jumlah data yang ada dalam buffered
 * cannel 
 * - tidakakan melakukan bloking goroutine lain selama kapasitas belum penuh
 * - data yang dikirim dan diterima akan bersifat seperti antrian FIFO (First in First out)
 */		

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

 func main(){
	// // size 3
	// channel := make(chan int,3)
	// channel <- 1
	// channel <- 2
	// channel <- 3
	// fmt.Println( <- channel)
	// fmt.Println("Panjang channel :", len(channel))
	// // channel <- 3
	
	// // fmt.Println( <- channel)
	// // fmt.Println( <- channel)
	// fmt.Println("kapasitas channel :",cap(channel))
	wg.Add(3)
	channel := make(chan int,3)
	
	go send(channel,1)
	go send(channel,2)
	go send(channel,3)	

	// go print(channel)
	// go print(channel)
	// go print(channel)
	wg.Wait()

	sumChannel(channel)

	fmt.Println("Finish")


	
 }

 func send(channel chan  int, number int){
	defer wg.Done()
	channel <- number
 }

 func print(channel chan int){
	defer wg.Done()
	fmt.Println(<- channel)
 }

 func sumChannel(channel chan int){
	result := 0
	for len(channel) > 0 {
		result += <-channel
	}
	fmt.Println(result)
 }