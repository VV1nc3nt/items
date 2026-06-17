package handler

import (
	pb "github.com/VV1nc3nt/items/internal/pb/items"
)

type Handler struct {
	pb.UnimplementedItemServiceServer
	service ItemManagerService
}

func New(service ItemManagerService) *Handler {
	return &Handler{service: service}
}
