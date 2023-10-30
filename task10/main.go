package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/* Задача 10.
   Дана последовательность температурных колебаний: -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5
   Объединить данные значения в группы с шагом в 10
   градусов. Последовательность в подмножноствах не важна.
   Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

/* Решение
   1) Мы считываем с помощью bufio со стандартного ввода строки. На вход как взадании
      принимаются значения через "," запятую. Эту строку мы переводим в слайс чисел с плавающей точкой,
      с помощью функции scanFloatSlice
   2) Сортируем слайс floatSlice для удобства.
   3) Создаем словарь/хеш-мапу celsius, где ключи это приведенные целые значения из floatSlice
      где последний знак 0.
   4) Далее добавляем пару ключ и значение из floatSlice в хеш мапу.
*/

func main() {
	floatSlice := scanFloatSlice()
	sort.Float64s(floatSlice)
	celsius := make(map[int][]float64)
	for _, val := range floatSlice {
		decim := int(val) - int(val)%10
		celsius[decim] = append(celsius[decim], val)
	}
	fmt.Printf("%+v\n", celsius)
}

func scanFloatSlice() []float64 {
	var floatSlice []float64
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	strSlice := strings.Split(strings.Trim(strings.Trim(str, "\n"), " "), ", ")
	for _, val := range strSlice {

		i, err := strconv.ParseFloat(val, 64)
		if err != nil {
			log.Fatalf("must be float - %s", err)
		} else {
			floatSlice = append(floatSlice, i)
		}
	}
	return floatSlice
}
