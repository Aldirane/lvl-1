package main

import (
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
Данный фрагмент кода может привести к панике
во время выполнения программы из-за выхода за пределы массива "v"
при попытке присвоить его подстроку "v[:100]" переменной "justString".
Это произойдет если длина строки "v" будет меньше 100 символов.
Чтобы исправить эту ошибку, необходимо проверить длину строки "v"
перед присваиванием и убедиться, что она достаточно длинная.
Корректный пример реализации:
*/

var justString string

func someFunc(n int) {
	v := createHugeString(n << 10)
	if len(v) >= 100 {
		justString = v[:100]
	} else {
		justString = v
	}
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatalf("Integer must be positive int type.")
	} else if n <= 0 {
		log.Fatalf("Integer must be positive int type not 0. Received: %d", n)
	}
	someFunc(n)
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

/*
В этом примере, мы добавляем проверку на длину строки "v"
перед выполнением присваивания. Если длина "v" больше или равна 100,
мы присваиваем первые 100 символов "v[:100]" переменной "justString".
В противном случае, мы присваиваем всю строку "v" без обрезания.
Таким образом, мы предотвращаем выход за пределы массива и избегаем паники при выполнении программы.
*/
