package disperser

import (
	"context"

	"github.com/0glabs/0g-data-avail/common"
	"github.com/0glabs/0g-data-avail/core"
	pb "github.com/0glabs/0g-data-avail/disperser/api/grpc/signer"
)

type SignerClient interface {
	BatchSign(ctx context.Context, addr string, data []*pb.SignRequest, log common.Logger) ([]*core.Signature, error)
}
