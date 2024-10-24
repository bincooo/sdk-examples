package main

import (
	"bincooo/sdk-examples/model"
	"bincooo/sdk-examples/px"
	"bincooo/sdk-examples/wire"
	"github.com/bincooo/sdk"
	"syscall"
)

func main() {
	container := sdk.NewContainer()
	err := wire.Injects(container)
	if err != nil {
		panic(err)
	}

	sdk.ProvideBean[sdk.Initializer](container, "main", Initialized)
	if err = container.Run(syscall.SIGINT, syscall.SIGTERM); err != nil {
		panic(err)
	}
}

func Initialized() (sdk.Initializer, error) {
	return sdk.InitializedWrapper(1, func(container *sdk.Container) (err error) {
		bean, err := sdk.InvokeAs[px.Echo](container, "model.A")
		if err != nil {
			panic(err)
		}

		err = bean.Echo("白居易")
		if err != nil {
			println("error: ", err.Error())
		}

		b, err := sdk.InvokeBean[*model.B](container, "model.B")
		if err != nil {
			panic(err)
		}

		b.Echo()
		return
	}), nil
}
