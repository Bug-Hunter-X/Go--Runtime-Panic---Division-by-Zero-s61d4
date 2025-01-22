func calculate(a, b int) int {
  return a / b
}

func main() {
  result := calculate(10, 0)
  fmt.Println(result) // This will cause a panic
}