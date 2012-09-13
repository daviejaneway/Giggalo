package Giggalo

import "fmt"
import "errors"

type Option struct {
  Id, Longid string
  Value interface{}
  Consume, state bool
}

func (o *Option) setValue(v interface{}) {
  o.Value = v
}

type OptGroup struct {
  Options []Option
}

func (opts OptGroup) Parse(args []string) error {
  for i := 0; i < len(args); i++ {
    for j := 0; j < len(opts.Options); j++ {
      opt := &opts.Options[j]
      
      if args[i] == fmt.Sprintf("-%s", opt.Id) {
        opt.state = true
        
        if opt.Consume && len(args) > (i + 1) {
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

/*var opts = OptGroup{
  options: []Option{
    Option{id:"r", longid:"recursive"}}}

func main() {
  err := opts.parse([]string{"-r"})
  
  if err != nil {
    fmt.Println(err)
  } else {
    v, err := opts.Get("r")
    
    if err != nil {
      fmt.Println(err)
    } else {
      fmt.Printf("-r : %v", v)
    }
  }
}*/