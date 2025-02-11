//+build wireinject

package service

import (
	"github.com/google/wire"

	"github.com/livekit/protocol/auth"

	"github.com/livekit/livekit-server/pkg/config"
	"github.com/livekit/livekit-server/pkg/routing"
)

func InitializeServer(conf *config.Config, keyProvider auth.KeyProvider, currentNode routing.LocalNode, selector routing.NodeSelector) (*LivekitServer, error) {
	wire.Build(
		ServiceSet,
	)
	return &LivekitServer{}, nil
}
