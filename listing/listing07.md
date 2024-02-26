Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
Сначала нужные значения, а потом нулевые, так как merge не проверяет, закрыты ли каналы.

Помимо этого, в коде есть проблема - если горутина зависнет внутри ветки селект, то возможен дэдлок из-за записи в другой канал Чтобы этого избежать, нужен еще один селект внутри другого и канал отмены, чтобы было больше контроля над исполняемыми горутинами.

