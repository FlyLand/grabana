package encoder

import (
	"github.com/FlyLand/grabana/encoder/golang"
	"github.com/K-Phoen/sdk"
	"go.uber.org/zap"
)

func ToGolang(logger *zap.Logger, dashboard sdk.Board) (string, error) {
	golangEncoder := golang.NewEncoder(logger)

	return golangEncoder.EncodeDashboard(dashboard)
}
