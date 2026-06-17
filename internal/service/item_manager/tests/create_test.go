package tests

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/VV1nc3nt/items/internal/model"
	"github.com/VV1nc3nt/items/internal/service/item_manager"
	mock "github.com/VV1nc3nt/items/internal/service/item_manager/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestItemServiceCreate(t *testing.T) {
	now := time.Now()

	input := &model.ItemCreateInput{
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
		setupMock  func(m *mock.Mockrepository)
		wantResult *model.Item
		wantErr    bool
	}{
		{
			name: "success",
			setupMock: func(m *mock.Mockrepository) {
				m.EXPECT().
					Create(context.Background(), input).
					Return(*wantResult, nil).
					Once()
			},
			wantResult: wantResult,
			wantErr:    false,
		},
		{
			name: "failure",
			setupMock: func(m *mock.Mockrepository) {
				m.EXPECT().
					Create(context.Background(), input).
					Return(model.Item{}, errors.New("fail")).
					Once()
			},
			wantResult: nil,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mock.NewMockrepository(t)
			tt.setupMock(mockRepo)

			service := item_manager.New(mockRepo)
			res, err := service.Create(context.Background(), input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, res)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}
