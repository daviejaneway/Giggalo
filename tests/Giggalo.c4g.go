package main

import "fmt"
import "github.com/daviejaneway/C4G/src"
import "github.com/daviejaneway/C4GUnit/src"
import "github.com/daviejaneway/Giggalo/src"

var contract = C4GUnit.TestContract{
  Conditions: []C4G.Condition{
    C4G.Condition{"Should not return error"},
    C4G.Condition{"Specified option should be set"},
    C4G.Condition{"Specified option should be set to specified value"}}}

func testNonConsumingShortOption() {
  opts := Giggalo.OptGroup{
    Options: []Giggalo.Option{
      Giggalo.Option{Id:"r", Longid:"recursive"}}}
  
  opts.Parse([]string{"-r"})
  
  v, err := opts.Get("r")
  
  contract.Assert(0, err == nil)
  contract.Assert(1, v == true)
}

func testConsumingShortOption() {
  opts := Giggalo.OptGroup{
      Options: []Giggalo.Option{
        Giggalo.Option{Id:"l", Longid:"limit", Consume:true}}}
  
  opts.Parse([]string{"-l", "200"})
  
  v, err := opts.Get("l")
  
  contract.Assert(0, err == nil)
  contract.Assert(2, v == "200")
}

func main() {
  testNonConsumingShortOption()
  testConsumingShortOption()
  
  fmt.Println(&C4GUnit.Session)
}