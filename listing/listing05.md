Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error
```
Сам интерфейс error реализован, тк туда упал nil-pointer на тип структуры. Если интерфейс реализован, то он уже не nil.

В этой программе неправильно возвращается ошибка, всегда лучше возвращать интерфейс error, а также если сообщение простое, то использовать errors.New("...")