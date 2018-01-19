package reverse

import (
  "strings"
)

//reverse a string

func String(input string) string {

  length := len(input)
  buffer := make([]string,length)
  j,i:= 0,length
  for ; i > 0 ; i,j=i-1,j+1{
    buffer[j] = string(input[i])
    }
    return strings.Join(buffer,"")
}
