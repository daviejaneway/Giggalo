package Giggalo

import "fmt"
import "errors"

type OptGroup struct {
  Options []Option
}

func isShortOption(opt, expected string) bool {
  return opt == fmt.Sprintf("-%s", expected)
}

func canConsume(opt *Option, args []string, i int) bool {
  return opt.Consume && len(args) > (i + 1)
}

func (opts OptGroup) Parse(args []string) error {
  for i := 0; i < len(args); i++ {
    for j := 0; j < len(opts.Options); j++ {
      opt := &opts.Options[j]
      
      if isShortOption(args[i], opt.Id) { //args[i] == fmt.Sprintf("-%s", opt.Id) {
        opt.state = true
        
        if canConsume(opt, args, i) { //opt.Consume && len(args) > (i + 1) {
          i += 1
          opt.Value = args[i]
        } else if opt.Consume && len(args) <= (i + 1) {
          return errors.New(fmt.Sprintf("Command line option '%s' requires a value", args[i]))
        }
      }
    }
  }

  return nil
}

func (opts OptGroup) Get(s string) (interface{}, error) {
  for i := 0; i < len(opts.Options); i++ {
    if opt := opts.Options[i]; opt.Id == s && opt.state {
      if opt.Consume && opt.Value != nil {
        return opt.Value, nil
      } else if opt.Consume {
        return false, errors.New(fmt.Sprintf("'%s' option value not set"))
      } else {
        return true, nil
      }
    }
  }
  
  return false, errors.New(fmt.Sprintf("'%s' option not found", s))
}