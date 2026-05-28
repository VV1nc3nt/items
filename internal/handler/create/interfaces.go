package create

import (
	"context"

	"github.com/VV1nc3nt/items/internal/model"
	pb "github.com/VV1nc3nt/items/internal/pb/items"
)

type ItemManagerService interface {
	Create(ctx context.Context, req *model.ItemInput) (*pb.CreateResponse, error)
}
