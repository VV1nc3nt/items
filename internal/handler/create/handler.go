package create

import (
	"context"
	"fmt"

	"github.com/VV1nc3nt/items/internal/model"
	pb "github.com/VV1nc3nt/items/internal/pb/items"
)

type Handler struct {
	pb.UnimplementedItemServiceServer
	service ItemManagerService
}

func New(service ItemManagerService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	input := &model.ItemInput{
		Category:    req.Category,
		Title:       req.Title,
		Description: req.Description,
		ImageKey:    req.ImageKey,
		Price:       req.Price,
		Quantity:    req.Quantity,
	}

	res, err := h.service.Create(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("handler create item: %w", err)
	}

	item := &pb.Item{
		Id:          res.Item.Id,
		Category:    res.Item.Category,
		Title:       res.Item.Title,
		Description: res.Item.Description,
		ImageKey:    res.Item.ImageKey,
		Price:       res.Item.Price,
		Quantity:    res.Item.Quantity,
		CreatedAt:   res.Item.CreatedAt,
		UpdatedAt:   res.Item.UpdatedAt,
	}

	return &pb.CreateResponse{Item: item}, nil
}
