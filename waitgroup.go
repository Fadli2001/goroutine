package main

import (
	"fmt"
	"time"
	"sync"
)

/**
 FUNCTION DALAM WAITGROUP 
 1. Add(int) digunakan untuk menambah counter dalam waitgroup
 2. Done() digunakan untuk mengurangi counter dalam waitgroup
 3. Wait() digunakan untuk memblokir eksekusi kode kita sampai counternya bernilai 0 
*/



func main(){
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("test1")
	wg.Wait()
	go repeat("dogs",500,&wg) 	
	fmt.Println("test2")
	fmt.Println("test3")
	go repeat("cat",500, &wg)	 
	fmt.Println("finish")
}


func repeat(word string, delay time.Duration, wg *sync.WaitGroup){
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println(i,word)
		time.Sleep(time.Millisecond * delay)
	}
}

// func sum(arr []int){
// 	 res := 0;
// 	for _, v := range arr {
// 		res += v;
// 	}
// 	fmt.Println(res);
// 	wg.Done()
// }