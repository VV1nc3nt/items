package create

import (
	"context"
	"errors"
	"testing"
	"time"

	mock "github.com/VV1nc3nt/items/internal/handler/create/mocks"
	"github.com/VV1nc3nt/items/internal/model"
	pb "github.com/VV1nc3nt/items/internal/pb/items"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestHandlerCreate(t *testing.T) {
	now := time.Now()

	req := &pb.CreateRequest{
		Category:    "test",
		Title:       "test",
		Description: "test",
		ImageKey:    "test",
		Quantity:    1000,
		Price:       10000,
	}

	wantInput := &model.ItemInput{
		Category:    "test",
		Title:       "test",
		Description: "test",
		ImageKey:    "test",
		Quantity:    1000,
		Price:       10000,
	}

	tests := []struct {
		name      string
		setupMock func(m *mock.MockitemManagerService)
		wantResp  *pb.CreateResponse
		wantErr   bool
	}{
		{
			name: "success",
			setupMock: func(m *mock.MockitemManagerService) {
				m.EXPECT().
					Create(context.Background(), wantInput).
					Return(&model.Item{
						ID:          0,
						Category:    "test",
						Title:       "test",
						Description: "test",
						ImageKey:    "test",
						Quantity:    1000,
						Price:       10000,
						CreatedAt:   now,
						UpdatedAt:   now,
					}, nil).
					Once()
			},
			wantResp: &pb.CreateResponse{
				Item: &pb.Item{
					Id:          0,
					Category:    "test",
					Title:       "test",
					Description: "test",
					ImageKey:    "test",
					Quantity:    1000,
					Price:       10000,
					CreatedAt:   timestamppb.New(now),
					UpdatedAt:   timestamppb.New(now),
				},
			},
			wantErr: false,
		},
		{
			name: "failure",
			setupMock: func(m *mock.MockitemManagerService) {
				m.EXPECT().
					Create(context.Background(), wantInput).
					Return(nil, errors.New("failed")).
					Once()
			},
			wantResp: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := mock.NewMockitemManagerService(t)
			tt.setupMock(mockService)

			h := New(mockService)

			resp, err := h.Create(context.Background(), req)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, tt.wantResp.Item, resp.Item)
		})
	}
}
