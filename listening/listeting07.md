## 7. Что выведет программа? Объяснить вывод программы.
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
### Ответ
```go
1
2
3
4
5
6
8
7
0
0
0
...
```

В функции `merge` используется оператор `select`, у которого нет проверки на то, возвращается или нет стандартное значение переменной из канала. После закрытия каналов `a` и `b` в выходной канал `c` будут отправляться стандартный значения переменной типа `int` - 0.

В языке нет проверки закрытия канала при вычитании из него значений.

Необходимо использовать `waitgroup` и дополнительные `горутины`.
```go
func merge(a, b <-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    
    wg.Add(2)
    go func(c <-chan int) {
        for value := range c {
            out <- value
        }
        wg.Done()
    }(a)
    
    go func(c <-chan int) {
        for value := range c {
            out <- value
        }
        wg.Done()
    }(b)
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}
```
