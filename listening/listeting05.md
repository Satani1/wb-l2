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
