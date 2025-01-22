func calculate(a, b int) (int, error) {
  if b == 0 {
    return 0, errors.New("division by zero")
  }
  return a / b, nil
}

func main() {
  result, err := calculate(10, 0)
  if err != nil {
    fmt.Println("Error:", err)
  } else {
    fmt.Println(result)
  }
  
  result, err = calculate(10,2)
  if err != nil {
    fmt.Println("Error:", err)
  } else {
    fmt.Println(result)
  }
} 

import (

    "errors"
    "fmt"
)