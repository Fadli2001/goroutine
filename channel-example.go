package main 



/**
 * 1. APA ITU CHANNEL? 
 * 		Channel adalah tempat untuk saling berkomunikasi antar 1 goroutine dengan goroutine lain 
 * KARAKTERISTIK DARI CHANNEL 
 * - Channel hanya bisa menampung satu data, Jika kita ingin menambahkan data lagi, harus menunggu data 
 * yang ada dichannel diambil
 * - channel hanya bisa menerima satu jenis data
 * - Channel bisa diambil atau dikirim oleh lebih dari satu goroutine
 * - Channel harus diclose jika tidak digunakan, karena bisa menyebabkan memory leak 
 * 
 * noted : 
 * - jika ingin mengirim data ke channel maka panahnya berapa pada sisi kanan variable channel,
 * 	jika ingin menerima dat adri channel maka panahnya berapa pada sisi kiri,
 * - sifat channel defaultnya pass by reference meskipun kita terapkan parameter pada function 
 * - deadlock bisa saja terjadi jika ter
 */

 import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

 func main(){
	channel := make(chan string)
	defer close(channel)

	// go func(){
	// 	channel <- "hello"
	// 	// channel <- true
	// }()

	/**
	 * bagaiman jika mengirimkan channel tanpa ada goroutine lain, (ingat ya func main golang termasuk goroutine juga),
	 * contoh disini kita hilangkan anonymous funcnya
	 */
	
		// channel <- "hello"	

	/**
	 * Code diatas akan terjadi error deadlock hal tersebut karena saat goroutine main mencoba mengirim string "Hello"
	 * kedalam channel kemudian tidak ada goroutine lain yang menerimanya 
	 * 
	 * error deadlock bisa terjdi juga karena tidak adanya pengirim dan penerima pada channel yang sudah di define 
	 */
	// wg.Add(1)
	//  go gateway(channel,"Hello, Enimgacamp")
	//  go showYourName("Fadli")

	wg.Add(1)
	go sendMessage(channel,"Hallo Enigma")
	// go getMessage(channel)

	// word := <- channel
	// fmt.Println(word)
	// atau bisa langsung memanggilnya 
	fmt.Println(<- channel)
	wg.Wait()
	fmt.Println("done!!")
 }

 func gateway(channel chan string, word string){
	defer wg.Done()
	time.Sleep(time.Second * 3)
	channel <- word
 }

 func showYourName(name string){
	fmt.Println(name)
 }

 /**
  * Kita juga bisa memubat validasi atau spesifikasikan paramter channel tersebut
  * untuk mengirim atau menerima
  */

  // hanya bisa mengirim data ke channel
  func sendMessage (channel chan<- string, message string){
	defer wg.Done()
	fmt.Println("Sedang mengirim Pesan...")
	time.Sleep(time.Second * 3)
	channel<- message
  }

  // hanya bisa menerima data dari channel
  func getMessage(channel <-chan string){
	defer wg.Done()
	fmt.Println(<- channel)
  }