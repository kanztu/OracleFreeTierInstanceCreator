package main

import (
	"github.com/kanztu/OracleFreeTierInstanceCreator/config"
	"github.com/kanztu/OracleFreeTierInstanceCreator/core"
	"github.com/kanztu/OracleFreeTierInstanceCreator/oci"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(config.New),
		fx.Provide(zap.NewProduction),
		fx.Provide(oci.New),
		fx.Provide(core.New),
		fx.Invoke(core.CreateInstance),
	).Run()
}
