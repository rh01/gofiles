package server

import "errors"

// 导出类型，且内部字段也必须为导出字段
type Args struct {
	A, B int
}

// 导出字段，且内部字段也必须为导出字段
type Quotient struct {
	Quo, Rem int
}

type Arith int // int的重命名

// 创建可导出方法
//1.方法可导出
//2.方法参数必须为可导出或者内建类型
//3.参数第二个为指针类型
//4.返回值必须为error类型
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}


// 创建可导出方法
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}