package diffsquares

//calc sum of squares

func SumOfSquares(in int) int {
  result := 0
  for i:=1 ; i <= in ; i++ {
    result = result + (i*i)
  }
  return result
}

func SquareOfSums(in int) int {
  result := 0
  for i:=1 ; i <= in ; i++ {
    result = result + i
  }
  result = result * result
  return result
}

func Difference(in int) int {
  return SquareOfSums(in) - SumOfSquares(in)
}
