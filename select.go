package main


import (
	"fmt"
)

/**
 * SELECT CHANNEL
 * 
 * - Select mempermudah kontrol komunikasi data dari beberapa channel
 * - select akan memblokir goroutine sampai salahsatu case terpenuhi,yang dimana 
 * ini terjadi ketika data dari channel berhasil diambil, dan setelahnya akan 
 * menjalankan statement yang ada dalam case tersebut
 * 
 */


func main(){
	animalChannel := make(chan string)
	fruitChannel := make(chan string)

	go func(){
		animalChannel <- "Cat" 
		animalChannel <- "Dog" 
		animalChannel <- "Fish" 
		animalChannel <- "Chicken" 
		animalChannel <- "Bird" 
		close(animalChannel)
	}()

	go func(){
		fruitChannel <- "Banana"
		fruitChannel <- "Apple"
		fruitChannel <- "Grapes"
		fruitChannel <- "Orange"
		fruitChannel <- "Watermelon"
		close(fruitChannel)
	}()

	fmt.Println("Proses")

	
	var animalStatus bool
	var fruitStatus bool
	for{
		select{
			case data, status <- animalChannel:
				animalStatus = status
				if animalStatus {
					fmt.Println("Animal : "+data)
				}
			case data, status  <- fruitChannel:
				fruitStatus = status
				if fruitStatus {
					fmt.Println("Fruit : " + data)
				}
		}

		if(!animalStatus && !fruitStatus){
			break;
		}
	}


	fmt.Println("Finsh")
}