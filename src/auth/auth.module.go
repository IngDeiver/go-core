package auth

import "go.uber.org/fx"

func Module() fx.Option {
    return fx.Options(
        fx.Provide(),
        fx.Invoke(),
    )
}