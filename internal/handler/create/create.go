package create

import (
	"context"
	"fmt"

	pb "github.com/VV1nc3nt/items/internal/pb/items"
	"github.com/VV1nc3nt/items/internal/service/item_manager"
)

func (h *Handler) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	input := &item_manager.ServiceItemInput{
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
