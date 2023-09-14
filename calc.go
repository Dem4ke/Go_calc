package main

import (

    "fmt"
    "os"
    "bufio"
    "strconv"
)
    
func RomanToInt(input string) (int) {
    rom := map[string]int {
        "I": 1,
        "V": 5,
        "X": 10,
    }

    var total int = 0
    var cursore int = 0
    var previously int = 0

    for i := 0; i < len(input); i++ {
        cursore = rom[string(input[i])]
        if cursore > previously {
            total = total + cursore - 2 * previously
        } else {
            total += cursore
        }

        previously = cursore
    }
    
    return total        
}

func IntToRoman(input int) (string) {
    var result string = ""
    if input < 4 {
        for i := 0; i < input; i++ {
            result += "I"
        }
    } else if input == 4 {
        result = "IV"
    } else if input >= 5 && input < 9 {
        result = "V"
        for i := 5; i < input; i++ {
            result += "I" 
        } 
    } else if input == 9 {
        result = "IX"
    } else if input >= 10 && input < 14 {
        result = "X"
        for i := 10; i < input; i++ {
            result += "I"
        }
    } else if input == 14 {
        result = "XIV"
    } else if input >=15 && input < 20 {
        result = "XV"
        for i := 15; i < input; i++ {
            result += "I"
        } 
    } else if input == 20 {
        result = "XX"
    }

    return result

}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var name string
    scanner.Scan()
    name = scanner.Text()

    var first string
    var first_roman bool = false
    var second string
    var second_roman bool = false
    var expr string

    for i := 0; i < len(name); i++ {
        if string(name[i]) == " " {
            continue
        }

        if (string(name[i]) == "+" || string(name[i]) == "-" ||
                string(name[i]) == "*" || string(name[i]) == "/") && expr == "" {
                    expr = string(name[i])
                    continue
                } else if (string(name[i]) == "+" || string(name[i]) == "-" ||
                string(name[i]) == "*" || string(name[i]) == "/") && expr != "" {
                    fmt.Println("Wrong format")
                    return
            } else if expr != "" { second += string(name[i])
        } else { first += string(name[i])
    }

    }

    var new_first int
    var new_second int
        
    if string(first[0]) == "I" || string(first[0]) == "V" || string(first[0]) == "X" {
        new_first = RomanToInt(first)
        first_roman = true
    } else {
        new_first, _ = strconv.Atoi(first)
    }

    if string(second[0]) == "I" || string(second[0]) == "V" || string(second[0]) == "X" {
        new_second = RomanToInt(second)
        second_roman = true
    } else {
        new_second, _ = strconv.Atoi(second)
    }

    if (!first_roman && !second_roman) || (first_roman && second_roman) {

    } else {
        fmt.Println("Different types")
        return
    }
    
    if new_first <= 0 || new_first > 10 || new_second <= 0 || new_second > 10 {
        fmt.Println("Wrong range")
        return
    }

    var answer_roman bool = first_roman && second_roman

    var answer int

    if expr == "+" { 
        answer = new_first + new_second
    } else if expr == "-" {
        answer = new_first - new_second
    } else if expr == "*" {
        answer = new_first * new_second
    } else if expr == "/" {
        answer = new_first / new_second 
    }

    if answer <= 0 && answer_roman {
        fmt.Println("No negative in Rome")
        return
    }

    if answer_roman { 
        fmt.Println(IntToRoman(answer))
    } else {
        fmt.Println(answer)
    }
}
    
