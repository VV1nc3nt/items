package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/VV1nc3nt/items/internal/model"
	pb "github.com/VV1nc3nt/items/internal/pb/items"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *Handler) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	if err := h.validateID(req.Id); err != nil {
		return nil, err
	}

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

func (h *Handler) validateID(id int64) error {
	if id <= 0 {
		return &model.ValidationError{
			Time:  time.Now(),
			Field: "id",
			Value: id,
			Err:   fmt.Errorf("id should be positive"),
		}
	}

	return nil
}
