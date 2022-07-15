package main

import (
	"fmt"
  b64 "encoding/base64"
)

func main() {
  var whatIsIt string
  secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
  sd, _ := b64.StdEncoding.DecodeString(secret)
  whatIsIt = Reverse(string(sd))
  fmt.Println(whatIsIt)
}

func Reverse(s string) (result string) {
  for _,v := range s {
    result = string(v) + result
  }
  return 
}