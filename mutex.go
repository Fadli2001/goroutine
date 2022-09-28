package main 

/*
	1. APA ITU RACE CONDITION ? 
		Race Condition terjadi ketika dua atau lebih goroutine dapat
	mengakses data bersamaan dalam waktu yang sama dan mereka mengubahnya 
	pada saat bersamaan juga

	2. APA ITU MUTEX ? 
		Mutex adalah method yang digunakan untuk melakukan mekanisme locking untuk memastikan hanya ada satu 
	Goroutine yang dapat mengakses suatu section code kita 

			Function dalam Mutex yaitu : lock() dan unlock() => jadi ketika ada 1 Goroutine yang sudan memegang kunci atau 
		lock(). lalu ada goroutine baru mencoba mendapatkan lock itu, maka goroutine baru itu akan dihentikan sampai mutex itu dibuka 
		atau di unlock()

*/
import (
	"fmt"
	"sync"
)

var x = 0

func main(){
	var wg sync.WaitGroup
	var mutex sync.Mutex 	

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go increment(&wg, &mutex )		
	}
	wg.Wait()
	fmt.Println("Counter = ", x)

}

func increment(wg *sync.WaitGroup, mutex *sync.Mutex){
	for i := 1; i <= 100; i++ {
		mutex.Lock()
		x += 1
		mutex.Unlock()
	}

	wg.Done()
}
