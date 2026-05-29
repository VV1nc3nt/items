package create

import (
	"context"
	"fmt"

	"github.com/VV1nc3nt/items/internal/model"
	pb "github.com/VV1nc3nt/items/internal/pb/items"
	"google.golang.org/protobuf/types/known/timestamppb"
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
