package create

import (
	pb "github.com/VV1nc3nt/items/internal/pb/items"
	"github.com/VV1nc3nt/items/internal/service/item_manager"
)

type Handler struct {
	pb.UnimplementedItemServiceServer
	service item_manager.ItemManagerService
}

func New(service item_manager.ItemManagerService) *Handler {
	return &Handler{service: service}
}
