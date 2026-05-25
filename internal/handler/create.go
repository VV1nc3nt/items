package handler

import (
	"context"
	"fmt"

	pb "github.com/VV1nc3nt/items/internal/pb/items"
	"github.com/VV1nc3nt/items/internal/service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *ItemHandler) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	input := service.CreateItemInput{
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
		Id:          res.ID,
		Category:    res.Category,
		Title:       res.Title,
		Description: res.Description,
		ImageKey:    res.ImageKey,
		Price:       res.Price,
		Quantity:    res.Quantity,
		CreatedAt:   timestamppb.New(res.CreatedAt),
		UpdatedAt:   timestamppb.New(res.UpdatedAt),
	}

	return &pb.CreateResponse{Item: item}, nil
}
