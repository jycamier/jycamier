package main

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/fx"
)

type Doer interface {
	Do()
}

var _ Doer = (*DoerImpl)(nil)

type DoerImpl struct{}

func NewDoerImpl(lc fx.Lifecycle) Doer {
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
	return doer
}

func (d DoerImpl) Do() {
	fmt.Println(">>>>> DoerImpl.Do()")
}

var _ Doer = (*DecorationForDoerImpl)(nil)

type DecorationForDoerImpl struct {
	inner Doer
}

func NewDecorationForDoerImpl(inner Doer, lc fx.Lifecycle) Doer {
	doer := &DecorationForDoerImpl{inner: inner}
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
	return doer
}

func (d DecorationForDoerImpl) Do() {
	fmt.Println(">>>>> DecorationForDoerImpl.Do()")
	d.inner.Do()
}

func DoerFunc(doer Doer) {
	for i := 0; i < 3; i++ {
		doer.Do()
		fmt.Println(">>>>> Sleeping 1 second")
		time.Sleep(time.Second * 1)
	}
}

func main() {
	err := fx.New(
		fx.Provide(NewDoerImpl),
		fx.Decorate(NewDecorationForDoerImpl),
		fx.Invoke(DoerFunc),
	).Start(context.Background())
	if err != nil {
		return
	}
}
