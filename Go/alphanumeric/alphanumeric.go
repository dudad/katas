package alphanumeric

import (
  "unicode"
  "fmt"
  )

func Alphanumeric(str string) bool {
  fmt.Println(str)
  for _,char := range str {
    if (!unicode.IsDigit(char)) && (!unicode.IsLetter(char)) {
      return false
      }
  }
  return len(str) > 0
}