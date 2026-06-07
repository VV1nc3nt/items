package create

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/VV1nc3nt/items/internal/model"
	pb "github.com/VV1nc3nt/items/internal/pb/items"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MockService struct {
	mockItem *model.Item
	mockErr  error
}

func (m *MockService) Create(ctx context.Context, req *model.ItemInput) (*model.Item, error) {
	return m.mockItem, m.mockErr
}

func TestHandler_Create(t *testing.T) {
	tests := []struct {
		name       string
		mockInput  *pb.CreateRequest
		mockOutput *model.Item
		mockErr    error
		wantOutput *pb.CreateResponse
		wantErr    bool
	}{
		{
			name: "success",
			mockInput: &pb.CreateRequest{
				Category:    "Testing",
				Title:       "Test",
				Description: "Testing handler",
				ImageKey:    "test",
				Price:       1,
				Quantity:    1,
			},
			mockOutput: &model.Item{
				ID:          0,
				Category:    "Testing",
				Title:       "Test",
				Description: "Testing handler",
				ImageKey:    "test",
				Price:       1,
				Quantity:    1,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			mockErr: nil,
			wantOutput: &pb.CreateResponse{
				Item: &pb.Item{
					Category:    "Testing",
					Title:       "Test",
					Description: "Testing handler",
					ImageKey:    "test",
					Price:       1,
					Quantity:    1,
				},
			},
			wantErr: false,
		},
		{
			name: "failure",
			mockInput: &pb.CreateRequest{
				Category:    "Testing",
				Title:       "Test",
				Description: "Testing handler",
				ImageKey:    "test",
				Price:       1,
				Quantity:    1,
			},
			mockOutput: nil,
			mockErr:    errors.New("some error"),
			wantOutput: nil,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &MockService{tt.mockOutput, tt.mockErr}
			h := New(service)

			out, err := h.Create(context.Background(), tt.mockInput)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantOutput.Item.Id, out.Item.Id)
			assert.Equal(t, tt.wantOutput.Item.Category, out.Item.Category)
			assert.Equal(t, tt.wantOutput.Item.Title, out.Item.Title)
			assert.Equal(t, tt.wantOutput.Item.Description, out.Item.Description)
			assert.Equal(t, tt.wantOutput.Item.ImageKey, out.Item.ImageKey)
			assert.Equal(t, tt.wantOutput.Item.Price, out.Item.Price)
			assert.Equal(t, tt.wantOutput.Item.Quantity, out.Item.Quantity)
		})
	}
}
