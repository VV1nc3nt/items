package item_manager

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/VV1nc3nt/items/internal/model"
	mock "github.com/VV1nc3nt/items/internal/service/item_manager/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestItemServiceCreate(t *testing.T) {
	now := time.Now()

	input := &model.ItemInput{
		Category:    "test",
		Title:       "test",
		Description: "test",
		ImageKey:    "test",
		Price:       10000,
		Quantity:    1000,
	}

	wantResult := &model.Item{
		ID:          0,
		Category:    "test",
		Title:       "test",
		Description: "test",
		ImageKey:    "test",
		Price:       10000,
		Quantity:    1000,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tests := []struct {
		name       string
		setupMock  func(m *mock.MockRepository)
		wantResult *model.Item
		wantErr    bool
	}{
		{
			name: "success",
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					Create(context.Background(), input).
					Return(model.Item{
						ID:          0,
						Category:    "test",
						Title:       "test",
						Description: "test",
						ImageKey:    "test",
						Price:       10000,
						Quantity:    1000,
						CreatedAt:   now,
						UpdatedAt:   now,
					}, nil).
					Once()
			},
			wantResult: wantResult,
			wantErr:    false,
		},
		{
			name: "failure",
			setupMock: func(m *mock.MockRepository) {
				m.EXPECT().
					Create(context.Background(), input).
					Return(model.Item{}, errors.New("fail")).
					Once()
			},
			wantResult: wantResult,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mock.NewMockRepository(t)
			tt.setupMock(mockRepo)

			service := New(mockRepo)
			res, err := service.Create(context.Background(), input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantResult, res)
			assert.Equal(t, tt.wantResult.ID, res.ID)
			assert.Equal(t, tt.wantResult.Category, res.Category)
			assert.Equal(t, tt.wantResult.Title, res.Title)
			assert.Equal(t, tt.wantResult.Description, res.Description)
			assert.Equal(t, tt.wantResult.ImageKey, res.ImageKey)
			assert.Equal(t, tt.wantResult.Price, res.Price)
			assert.Equal(t, tt.wantResult.Quantity, res.Quantity)
			assert.Equal(t, tt.wantResult.CreatedAt, res.CreatedAt)
			assert.Equal(t, tt.wantResult.UpdatedAt, res.UpdatedAt)
		})
	}
}
