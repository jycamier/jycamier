package main

import (
	"testing"

	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func TestFXDoer(t *testing.T) {
	var first Doer
	var second Doer
	app := fxtest.New(
		t,
		fx.Provide(
			NewDoerImpl,
			NewDecorationForDoerImpl,
		),
		fx.Invoke(
			DoerFunc,
			DoerAllFunc,
		),
		fx.Populate(
			fx.Annotate(
				&first,
				fx.ParamTags(`name:"doer.first"`),
			),
			fx.Annotate(
				&second,
				fx.ParamTags(`name:"doer.second"`),
			),
		),
	)
	defer app.RequireStart().RequireStop()
	first.Do()
	second.Do()
}
