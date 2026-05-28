package item_manager

import (
	"context"
	"fmt"

	"github.com/VV1nc3nt/items/internal/model"
	pb "github.com/VV1nc3nt/items/internal/pb/items"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ItemService) Create(ctx context.Context, req *model.ItemInput) (*pb.CreateResponse, error) {
	repoInput := &model.ItemInput{
		Category:    req.Category,
		Title:       req.Title,
		Description: req.Description,
		ImageKey:    req.ImageKey,
		Price:       req.Price,
		Quantity:    req.Quantity,
	}

	row, err := s.repo.Create(ctx, repoInput)
	if err != nil {
		return &pb.CreateResponse{}, fmt.Errorf("service create item: %w", err)
	}

	item := &pb.CreateResponse{
		Item: &pb.Item{
			Id:          row.ID,
			Category:    row.Category,
			Title:       row.Title,
			Description: row.Description,
			ImageKey:    row.ImageKey,
			Price:       row.Price,
			Quantity:    row.Quantity,
			CreatedAt:   timestamppb.New(row.CreatedAt),
			UpdatedAt:   timestamppb.New(row.UpdatedAt),
		},
	}

	return item, nil
}
