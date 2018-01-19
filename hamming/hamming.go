package hamming

//Test strings and return correct value
import (
  "errors"
)
func Distance(a, b string) (int, error) {
  if len(a) != len(b){
    return -1,errors.New("Not allowed")
  } else {
    count := 0
    for i:=0 ; i < len(a) ; i++ {
      if a[i] != b[i]{
        count += 1
      }
    }
    return count,nil
  }
}
