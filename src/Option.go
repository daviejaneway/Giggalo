package Giggalo

import "fmt"

type OPTION_TYPE uint8

const (
  OPTION_TYPE_SHORT = 1 << iota
  OPTION_TYPE_LONG  = 1 << iota
)

type Option struct {
  Id, Longid, Value string
  Consume, state bool
  optiontype OPTION_TYPE
}

/*********** UTILITY FUNCTIONS ***********/
func Is(opt, expected string) bool {
  return (isShortOption(opt, expected) || isLongOption(opt, expected))
}

func isShortOption(opt, expected string) bool {
  return opt == fmt.Sprintf("-%s", expected)
}

func isLongOption(opt, expected string) bool {
  return opt == fmt.Sprintf("--%s", expected)
}

func (opt *Option) canConsume(args []string, i int) bool {
  return opt.Consume && len(args) > (i + 1)
}