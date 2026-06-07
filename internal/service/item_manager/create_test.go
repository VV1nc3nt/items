package item_manager

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/VV1nc3nt/items/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MockRepository struct {
	mockItem model.Item
	mockErr  error
}

func (r *MockRepository) Create(ctx context.Context, in *model.ItemInput) (model.Item, error) {
	return r.mockItem, r.mockErr
}

func TestItemService_Create(t *testing.T) {
	tests := []struct {
		name       string
		mockInput  *model.ItemInput
		mockOutput model.Item
		mockErr    error
		wantOutput model.Item
		wantErr    bool
	}{
		{
			name: "success",
			mockInput: &model.ItemInput{
				Category:    "Tests",
				Title:       "Testing",
				Description: "Testing service create",
				ImageKey:    "Test",
				Price:       1000,
				Quantity:    10,
			},
			mockOutput: model.Item{
				ID:          0,
				Category:    "Tests",
				Title:       "Testing",
				Description: "Testing service create",
				ImageKey:    "Test",
				Price:       1000,
				Quantity:    10,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			mockErr: nil,
			wantOutput: model.Item{
				ID:          0,
				Category:    "Tests",
				Title:       "Testing",
				Description: "Testing service create",
				ImageKey:    "Test",
				Price:       1000,
				Quantity:    10,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			wantErr: false,
		},
		{
			name: "failure",
			mockInput: &model.ItemInput{
				Category:    "Tests",
				Title:       "Testing",
				Description: "Testing service create",
				ImageKey:    "Test",
				Price:       1000,
				Quantity:    10,
			},
			mockOutput: model.Item{},
			mockErr:    fmt.Errorf("some error"),
			wantOutput: model.Item{},
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MockRepository{mockItem: tt.mockOutput, mockErr: tt.mockErr}
			service := New(repo)

			row, err := service.Create(context.Background(), tt.mockInput)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantOutput.ID, row.ID)
			assert.Equal(t, tt.wantOutput.Category, row.Category)
			assert.Equal(t, tt.wantOutput.Title, row.Title)
			assert.Equal(t, tt.wantOutput.Description, row.Description)
			assert.Equal(t, tt.wantOutput.ImageKey, row.ImageKey)
			assert.Equal(t, tt.wantOutput.Price, row.Price)
			assert.Equal(t, tt.wantOutput.Quantity, row.Quantity)
		})
	}
}
