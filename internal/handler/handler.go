package handler

import (
	pb "github.com/VV1nc3nt/items/internal/pb/items"
	"github.com/VV1nc3nt/items/internal/service"
)

type ItemHandler struct {
	pb.UnimplementedItemServiceServer
	service service.ItemService
}

func NewItemHandler(service service.ItemService) *ItemHandler {
	return &ItemHandler{service: service}
}
