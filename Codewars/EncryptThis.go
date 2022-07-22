package Codewars

/*
Number of trailing zeros of N!

You want to create secret messages which can be deciphered by the Decipher this! kata. Here are the conditions:

Your message is a string containing space separated words.
You need to encrypt each word in the message using the following rules:
The first letter must be converted to its ASCII code.
The second letter must be switched with the last letter
Keepin' it simple: There are no special characters in the input.

The following are examples of expected output values:

EncryptThis("Hello") == "72olle"
EncryptThis("good") == "103doo"
EncryptThis("hello world") == "104olle 119drlo"
*/

import (
        "strings"
        "strconv"
       )

func EncryptThis(text string) string {
  if len(text)==0 {return ""}
  result := []string{} 
  crypter := func(s string) string {
    runes := []rune(s)    
    result := strconv.Itoa(int(runes[0]))
    if len(s)>1{
      runes[1], runes[len(runes)-1] =runes[len(runes)-1], runes[1]
      for i:=1; i<len(runes); i++ {
        result+=string(runes[i])
      }      
    }
    return result
  }
  words := strings.Split(text, " ")
  for _, elem := range words {
      line := string(elem)
      newLine := crypter(line)
      result = append(result, newLine)
    }
  return strings.Join(result, " ")    
}
