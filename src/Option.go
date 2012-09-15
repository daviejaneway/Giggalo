package Giggalo

type Option struct {
  Id, Longid string
  Value interface{}
  Consume, state bool
}

func (o *Option) setValue(v interface{}) {
  o.Value = v
}