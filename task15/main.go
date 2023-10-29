package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

/*
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

/*
Данный фрагмент кода может привести к негативным последствиям, связанным с использованием большого объема памяти.
Если огромная строка будет сохранена в памяти после выполнения функции someFunc() из-за ссылки на нее в переменной justString,
это может привести к нежелательным последствиям, таким как избыточное использование памяти.
Когда код вызыывается несколько раз, каждый раз будет создана новая огромная строка,
и предыдущая огромная строка будет оставаться в памяти до тех пор, пока на нее не будет собран мусор.
*/

/*
Чтобы избежать таких проблем, можно использовать буферизацию строк.
Как вариант, можно использовать тип bytes.Buffer для создания буфера и записи в него огромной строки.
Затем можно взять первые 100 символов этого буфера и присвоить их переменной justString.
*/

var justString string

func someFunc() {
	var buffer bytes.Buffer
	v := createHugeString(1 << 10)
	buffer.WriteString(v)
	justString = buffer.String()[:100]
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatalf("Integer must be positive int type.")
	} else if n <= 0 {
		log.Fatalf("Integer must be positive int type not 0. Received: %d", n)
	}
	someFunc()
	fmt.Println(justString)
}

func createHugeString(n int) string {
	strArr := make([]string, n)
	for i := 0; i < n; i++ {
		strArr[i] = "A"
	}
	strVal := strings.Join(strArr, "")
	return strVal
}
