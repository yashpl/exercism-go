package reverse



//reverse a string

func String(input string) string {

var str string
  for _,i:= range input {
    str = string(i) + str
    }
    return str
}
