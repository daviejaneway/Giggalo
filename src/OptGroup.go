package Giggalo

import "fmt"
import "errors"

type OptGroup struct {
  Options []Option
}

func (opts OptGroup) Parse(args []string) error {
  for i := 0; i < len(args); i++ {
    for j := 0; j < len(opts.Options); j++ {
      opt := &opts.Options[j]
      
      consume(args, i, opt)
    }
  }

  return nil
}

func consume(args []string, i int, opt *Option) (interface{}, error) {
  ok := Is(args[i], opt.Id)
  
  if ok {
    opt.state = true
    
    if opt.canConsume(args, i) {
      i += 1
      opt.Value = args[i]
      
      return opt.Value, nil
    } else if opt.Consume && len(args) <= (i + 1) {
        return nil, errors.New(fmt.Sprintf("Command line option '%s' requires a value", args[i]))
    }
  }
  
  return nil, errors.New("Option not found")
}

//Gets the specified option by short or long name.
//Returns an error if either the option does not exist or,
//if the option consumes, a value has not been provided.
func (opts OptGroup) Get(s string) (interface{}, error) {
  for i := 0; i < len(opts.Options); i++ {
    if opt := opts.Options[i]; Is(opt.Value, s) && opt.state {
      if opt.Consume && opt.Value != "" {
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