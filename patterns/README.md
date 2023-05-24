# Паттерны

## 2 Строитель (Builder)
Отделяет конструирование сложного объекта от его представления так, что в результате одного и того же процесса конструирования могут получаться разные представления.

### Используют в случаях когда:

### Плюсы
 1) Позволяет изменять внутренне представление объекта;
 2) Изолирует код, реализующий конструирование и представление;
 3) Дает более тонкий контроль над процессом конструирования; 

### Минусы
1) Алгоритм создания сложного объекта не должен зависеть от того, из каких частей состоит обхект и как они стыкуются между собой;
2) Процесс конструирования должен обеспечивать различные представления конструируемого объекта;

## 6 Фабричный метод (Factory Method)
Определяет интерфейс для создания объекта, но оставляет подклассам решение о том, на основании какого класса создавать объект.
Фабричный метод позволяет классу делегировать создание подклассов.

### Используют в случаях когда:
 Классу заранее неизвестно, объекта каких подклассов ему необходимо создавать.
 Класс спроектирован так, чтобы объекты которые он создает, специфицировались подклассами.
 Класс делегирует свои обязанности одному из нескольких вспомогательных подклассов, и планируется локализовать знание о том, какой класс принимает эти обязанности на себя
 
### Плюсы
 1) Позволяет сделать код создания объектов более универсальным, не привязываясь к конкретным классам, а оперируя лишь общим интерфейс;
 2) Позволяет установить связь между параллельными иерархиями классов; 
### Минусы
 1) Необходимость создавать наследника общего интерфейса для каждого нового типа объекта;