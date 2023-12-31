package main

import (
	"fmt"
	"strings"
)

/* Задание 15.
   К каким негативным последствиям может привести данный фрагмент кода, и как
   это исправить? Приведите корректный пример реализации.

    var justString string
    func someFunc() {
       v := createHugeString(1 << 10)
       justString = v[:100]
    }
    func main() {
       someFunc()
    }
*/

/* Решение
   Данный фрагмент кода может привести к следующим негативным последствиям:
   1) Потенциальная утечка памяти: функция "createHugeString" создает строку размером в 1024 байта (1 << 10),
   но после этого использует только первые 100 байт. Однако весь созданный массив байтов будет храниться в памяти
   до окончания работы программы, что может привести к излишнему использованию памяти.

   2) Ненужное копирование данных: оператор присваивания "justString = v[:100]" создает новую строку,
   скопировав только первые 100 байт из строки "v".
   Это занимает дополнительное время и использует дополнительную память.

   Для исправления кода и избежания данных проблем, можно воспользоваться типом "[]byte" вместо "string"
   и использовать функцию "copy" для копирования данных:
*/

var justBytes []byte

func someFunc() {
	v := createHugeString(1 << 10)   // создаем большую строку
	justBytes = make([]byte, 100)    // создаем срез байтов нужного размера
	copy(justBytes, []byte(v[:100])) // копируем первые 100 байт из строки v в срез justBytes
}

func main() {
	someFunc()
	justString := string(justBytes) // получение строки из байтов
	fmt.Println(justString)
}

func createHugeString(n int) string {
	var strBuild strings.Builder
	for i := 0; i < n; i++ {
		strBuild.WriteString("A")
	}
	return strBuild.String()
}
