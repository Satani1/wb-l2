# WB-L2

# Устная подготовка
## 1. Что выведет программа? Объяснить вывод программы.

```go
package main
 
import (
    "fmt"
)
 
func main() {
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4]
    fmt.Println(b)
}
```

### Ответ
```go
[77 78 79]
```
В программе создается массив из 5 элементов. Затем на базе этого массива создается слайс [1:4], емкость слайса будет равна 4, а длина 3. В результате окажутся элементы `[77 78 79]`. Так как первый индекс среза является включительным, а второй нет.
## 2. Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и порядок их вызовов.

```go
package main
 
import (
    "fmt"
)
 
func test() (x int) {
    defer func() {
        x++
    }()
    x = 1
    return
}
 
 
func anotherTest() int {
    var x int
    defer func() {
        x++
    }()
    x = 1
    return x
}
 
 
func main() {
    fmt.Println(test())
    fmt.Println(anotherTest())
}
```
## 3. Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.
```go
package main
 
import (
    "fmt"
    "os"
)
 
func Foo() error {
    var err *os.PathError = nil
    return err
}
 
func main() {
    err := Foo()
    fmt.Println(err)
    fmt.Println(err == nil)
}
```
### Ответ
```go
<nil>
false
```
Базовая структура является равно `nil`, однако интерфейс таковым не является. Интерфейс может быть равен `nil` только в случае, если тип и значение равны `nil`.
## Устройство интерфейсов
## 4. Что выведет программа? Объяснить вывод программы.
```go
package main
 
func main() {
    ch := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
    }()
 
    for n := range ch {
        println(n)
    }
}
```
### Ответ
```go
0
1
2
3
4
5
6
7
8
9
fatal error: all goroutines are asleep - deadlock!
```
Произойдет `deadlock` в цикле, где числа вычитаются из канала. Это происходит, потому что после окончания записи в канал данных он не был закрыт. Следовательно цикл будет вечно ждать данные из канала. Необходимо закрыть канал `close(ch)` после цикла в горутине.
## 5. Что выведет программа? Объяснить вывод программы.
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
### Ответ
```go
error
```
Интерфейс равен `nil`, только в случае если тип и значения оба равны `nil`. В данной программме функция `test` возвращает интерфейс, в котором данные будут `nil`, но тип будет определен, как `customError` => в терминале будет `error`

## 6. Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передаче их в качестве аргументов функции.
```go
package main
 
import (
  "fmt"
)
 
func main() {
  var s = []string{"1", "2", "3"}
  modifySlice(s)
  fmt.Println(s)
}
 
func modifySlice(i []string) {
  i[0] = "3"
  i = append(i, "4")
  i[1] = "5"
  i = append(i, "6")
}
```
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
