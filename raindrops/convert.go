package raindrops
import "strconv"
//The answer to life

func Convert (in int) string {

var buffer = ""

if in == 1 {

  return "1"
  } else {
    if in%3 == 0 {
      buffer += "Pling"
    }
    if in%5 == 0 {
      buffer += "Plang"
    }
    if in%7 == 0 {
      buffer += "Plong"
    }
    //check if none of the 3 number are factors
    if buffer == ""{
      return strconv.Itoa(in)
    } else {
      return buffer
    }
  }

}
