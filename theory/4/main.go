package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		//close(ch) - для предотвращения deadlock'a
	}()

	// Выведет числа от 0 до 9, после чего случится deadlock:
	// т.к. в вышеуказанная горутина завершит выполнение, канал будет заблокирован до получения значения
	//  т.е. навсегда.
	for n := range ch {
		println(n)
	}
}
