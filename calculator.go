package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "unicode"
)

var (
    operators = map[string]func(int, int) int{
       "+": func(a, b int) int { return a + b },
       "-": func(a, b int) int { return a - b },
       "*": func(a, b int) int { return a * b },
       "/": func(a, b int) int { return a / b },
    }
    acceptableArabic = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
    acceptableRoman  = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
)

func Contains[T comparable](s []T, e T) bool {
    for _, v := range s {
       if v == e {
          return true
       }
    }
    return false
}

func GetIndex[T comparable](s []T, e T) int {
    for i, v := range s {
       if v == e {
          return i
       }
    }
    return -1
}

func InputClear(input string) string {

    cleanInput := strings.Map(func(r rune) rune {
       if unicode.IsPrint(r) {
          return r
       }
       return -1
    }, input)
    return cleanInput
}

func InputParse(input string) (string, []string, error) {
    var operator string
    var count int
    for op := range operators {
       if strings.Contains(input, op) {
          operator = op
          count += strings.Count(input, operator)
       }
    }

    if operator == "" {
       return "", nil, fmt.Errorf("введённая строка не является математической операцией")
    }
    if count > 1 {
       return "", nil, fmt.Errorf("калькулятор работает только для одного оператора. Вы ввели %v", count)

    }

    numbers := strings.Split(input, operator)

    if Contains(acceptableArabic, numbers[0]) && Contains(acceptableArabic, numbers[1]) {
       numbers = append(numbers, "a")
    } else if Contains(acceptableRoman, numbers[0]) && Contains(acceptableRoman, numbers[1]) {
       if GetIndex(acceptableRoman, numbers[0]) < GetIndex(acceptableRoman, numbers[1]) &&
          (operator == "-" || operator == "/") {
          return "", nil, fmt.Errorf("результат меньше или равен 0. В римской системе нет отрицательных чисел")
       }
       numbers = append(numbers, "r")
    } else if Contains(acceptableArabic, numbers[0]) && Contains(acceptableRoman, numbers[1]) ||
       Contains(acceptableRoman, numbers[0]) && Contains(acceptableArabic, numbers[1]) {
       return "", nil, fmt.Errorf("операнды в разных системах счисления.\nКалькулятор умеет работать только " +
          "в одной системе счисления одновременно.")
    } else {
       return "", nil, fmt.Errorf("недопустимые операнды.\nОперандами могут быть только целые арабские " +
          "и римские числа от 1 до 10.")
    }
    return operator, numbers, nil
}

func CalculateArabic(operator, str1, str2 string) int {
    arg1, _ := strconv.Atoi(str1)
    arg2, _ := strconv.Atoi(str2)

    result := operators[operator](arg1, arg2)
    return result
}

func RomanToArabic(str string) int {
    i := GetIndex(acceptableRoman, str)
    str = acceptableArabic[i]
    result, _ := strconv.Atoi(str)
    return result
}

func CalculateRoman(operator, str1, str2 string) int {
    return operators[operator](RomanToArabic(str1), RomanToArabic(str2))
}

func ArabicToRoman(str string) string {
    num, _ := strconv.Atoi(str)

    arabic := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
    roman := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

    result := ""
    for i := 0; i < len(arabic); i++ {
       for num >= arabic[i] {
          result += roman[i]
          num -= arabic[i]
       }
    }
    return result
}

func main() {
    for {
       fmt.Print("Введите выражение (или 'q' для выхода из программы): ")
       in := bufio.NewReader(os.Stdin)
       input, _ := in.ReadString('\n')
       input = InputClear(input)
       if input == "q" {
          break
       }
       input = strings.ReplaceAll(input, " ", "")
       operator, args, err := InputParse(input)
       if err == nil {
          if args[2] == "a" {
             fmt.Printf("Результат: %v\n", CalculateArabic(operator, args[0], args[1]))
          } else if args[2] == "r" {
             result := CalculateRoman(operator, args[0], args[1])
             resultString := strconv.Itoa(result)
             fmt.Printf("Результат: %v\n", ArabicToRoman(resultString))
          }
       } else {
          fmt.Println("Ошибка:", err)
          break
       }
    }
}
