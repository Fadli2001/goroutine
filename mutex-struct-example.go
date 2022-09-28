package main

import(
	"fmt"
	"sync"
)


var wg sync.WaitGroup

func main(){

	account := BankAccount {Balance:0}

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go account.topUp(1)
	}

	wg.Wait()
	fmt.Println("My Balance = ",account.Balance)

}

type BankAccount struct {
	Balance int
	mutex sync.Mutex 
}

func (account *BankAccount) topUp(amount int){
	defer wg.Done() 
	defer account.mutex.Unlock()
	account.mutex.Lock()
	account.Balance += amount
}

/*best practice nya function wg.Done() dan mutex.Unlock() selalu dipanggil diakhir meskipun terjadi panic di dalam function  
contohnya yaitu gunakan defer contoh : func (){
											defer wg.Done() 
											defer account.mutex.Unlock()
											anything code
										}
*/
