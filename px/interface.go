package px

import (
	"errors"
	"fmt"
	"github.com/bincooo/sdk"
	"reflect"
)

// @Proxy(target="invoke")
type Echo interface {
	Echo(name string) error
}

func invoke(ctx *sdk.Context) {
	fmt.Println("开始代理...")

	name := ctx.In[0].(string)
	fmt.Println("入参: ", name)

	name = "李白"
	fmt.Println("修改入参: ", name)
	ctx.In[0] = name

	// 反射修改实例
	obj := ctx.Receiver
	ox := reflect.ValueOf(obj)
	if ox.Kind() == reflect.Ptr {
		ox = ox.Elem()
	}

	field := ox.FieldByName("Num")

	i := field.Int()
	fmt.Println("Num: ", i)
	// 非指针类型无法设置值
	if field.CanSet() {
		field.SetInt(2)
	}

	ctx.Do()

	var err error

	// 指针类型
	if o := ctx.Out[0]; o != nil {
		err = o.(error)
	}

	fmt.Println("出参: ", err)
	err = errors.New("panic")
	ctx.Out[0] = err
	fmt.Println("修改出参: ", err)

	fmt.Println("结束代理...")
}
