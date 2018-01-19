package isogram

import (
  "strings"
)
//check if given string is isogram

func IsIsogram(input string) bool {

input = strings.ToLower(input)

index := make(map[string]bool)

for i:= 0 ; i < len(input) ; i++ {
  //skip any non alphabets
  if !strings.ContainsAny(string(input[i]),"abcdefghijklmnopqrstuvwxyz") {
    continue;
  }
  _,letterFound:=index[string(input[i])]
  if letterFound {
    return false
  } else {
    index[string(input[i])] = true
  }
}
  return true
}
