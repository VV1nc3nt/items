package handler

import (
	"context"
	"fmt"

	"github.com/VV1nc3nt/items/internal/model"
	pb "github.com/VV1nc3nt/items/internal/pb/items"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *Handler) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {

	input := &model.ItemGetInput{
		ID: req.Id,
	}

	res, err := h.service.Get(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("handler get item: %w", err)
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

	return &pb.GetResponse{Item: item}, nil
}
