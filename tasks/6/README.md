## 6
Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:

    -f - "fields" - выбрать поля (колонки)
    -d - "delimiter" - использовать другой разделитель
    -s - "separated" - только строки с разделителем

### Запуск

`go run main.go "./test/test.txt"`

`go run main.go -f="1-3" -d=" " "./test/test.txt"`

`go run main.go -f="1,3" -d=" " "./test/test.txt"`