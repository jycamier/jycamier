package main

import (
	"context"
	"fmt"

	"go.uber.org/fx"
)

type Doer interface {
	Do()
}

var _ Doer = (*DoerImpl)(nil)

type DoerImpl struct{}

type DoerFirstResults struct {
	fx.Out

	First Doer `name:"doer.first"`
	Group Doer `group:"doer"`
}

func NewDoerImpl(lc fx.Lifecycle) DoerFirstResults {
	doer := &DoerImpl{}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println(">>>>> DoerImpl.OnStart()")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println(">>>>> DoerImpl.OnStop()")
			return nil
		},
	})
	return DoerFirstResults{
		First: doer,
		Group: doer,
	}
}

func (d DoerImpl) Do() {
	fmt.Println(">>>>> DoerImpl.Do()")
}

var _ Doer = (*DecorationForDoerImpl)(nil)

type DecorationForDoerImpl struct {
}

type DoerSecondResults struct {
	fx.Out

	Second Doer `name:"doer.second"`
	Group  Doer `group:"doer"`
}

func NewDecorationForDoerImpl(lc fx.Lifecycle) DoerSecondResults {
	doer := &DecorationForDoerImpl{}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println(">>>>> DecorationForDoerImpl.OnStart()")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println(">>>>> DecorationForDoerImpl.OnStop()")
			return nil
		},
	})
	return DoerSecondResults{
		Second: doer,
		Group:  doer,
	}
}

func (d DecorationForDoerImpl) Do() {
	fmt.Println(">>>>> DecorationForDoerImpl.Do()")
}

type DoerParams struct {
	fx.In

	First  Doer `name:"doer.first"`
	Second Doer `name:"doer.second"`
}

func DoerFunc(params DoerParams) {
	fmt.Println(">>>>> CALL EACH SERVICE")
	params.First.Do()
	params.Second.Do()
}

type DoersParams struct {
	fx.In

	Doers []Doer `group:"doer"`
}

func DoerAllFunc(params DoersParams) {
	fmt.Println(">>>>> CALL SLICE OF SERVICES")
	for _, doer := range params.Doers {
		doer.Do()
	}
}

func main() {
	fx.New(
		fx.Provide(
			NewDoerImpl,
			NewDecorationForDoerImpl,
		),
		fx.Invoke(
			DoerFunc,
			DoerAllFunc,
		),
		fx.NopLogger,
	).Run()
}
