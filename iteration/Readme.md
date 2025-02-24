В этом разделе рассматривается, как выполнять повторяющиеся действия в Go с использованием цикла for. В отличие от других языков, в Go нет ключевых слов while, do, until — используется только for.

Написание теста
Для начала напишем тест для функции Repeat, которая повторяет символ 5 раз.

go
Copy
package iteration

import "testing"

func TestRepeat(t *testing.T) {
    repeated := Repeat("a")
    expected := "aaaaa"

    if repeated != expected {
        t.Errorf("expected %q but got %q", expected, repeated)
    }
}
Запуск теста
При попытке запуска теста получим ошибку, так как функция Repeat еще не определена.

bash
Copy
./repeat_test.go:6:14: undefined: Repeat
Минимальная реализация для компиляции
Создадим минимальную реализацию функции Repeat, чтобы тест мог скомпилироваться.

go
Copy
package iteration

func Repeat(character string) string {
    return ""
}
Тест теперь запускается, но завершается с ошибкой:

bash
Copy
repeat_test.go:10: expected 'aaaaa' but got ''
Реализация функции Repeat
Используем цикл for для повторения символа 5 раз.

go
Copy
func Repeat(character string) string {
    var repeated string
    for i := 0; i < 5; i++ {
        repeated = repeated + character
    }
    return repeated
}
Особенности цикла for в Go
В Go нет круглых скобок вокруг условий цикла for.

Фигурные скобки { } обязательны.

Использование var repeated string объявляет переменную без инициализации.

Рефакторинг
Введем константу repeatCount и оператор += для улучшения читаемости кода.

go
Copy
const repeatCount = 5

func Repeat(character string) string {
    var repeated string
    for i := 0; i < repeatCount; i++ {
        repeated += character
    }
    return repeated
}
Бенчмаркинг
Go предоставляет встроенную поддержку бенчмарков. Напишем бенчмарк для функции Repeat.

go
Copy
func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}
Запустим бенчмарк:

bash
Copy
go test -bench=.
Результат:

bash
Copy
goos: darwin
goarch: amd64
pkg: github.com/quii/learn-go-with-tests/for/v4
10000000           136 ns/op
PASS
Оптимизация с использованием strings.Builder
Строки в Go неизменяемы, что может негативно сказываться на производительности при частой конкатенации. Используем strings.Builder для улучшения производительности.

go
Copy
import "strings"

const repeatCount = 5

func Repeat(character string) string {
    var repeated strings.Builder
    for i := 0; i < repeatCount; i++ {
        repeated.WriteString(character)
    }
    return repeated.String()
}
Запустим бенчмарк с флагом -benchmem для анализа использования памяти:

bash
Copy
go test -bench=. -benchmem
Результат:

bash
Copy
goos: darwin
goarch: amd64
pkg: github.com/quii/learn-go-with-tests/for/v4
10000000           25.70 ns/op           8 B/op           1 allocs/op
PASS
Параметры бенчмарка
ns/op: время выполнения одной операции в наносекундах.

B/op: количество байт, выделенных за одну операцию.

allocs/op: количество выделений памяти за одну операцию.

Заключение
Использование strings.Builder значительно улучшает производительность при работе с конкатенацией строк. Бенчмарки помогают оценить производительность и оптимизировать код.