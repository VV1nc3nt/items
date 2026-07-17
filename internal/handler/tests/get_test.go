package tests

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/VV1nc3nt/items/internal/handler"
	mock "github.com/VV1nc3nt/items/internal/handler/mocks"
	"github.com/VV1nc3nt/items/internal/model"
	pb "github.com/VV1nc3nt/items/internal/pb/items"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestHandlerGet(t *testing.T) {
	now := time.Now()

	req := &pb.GetRequest{
		Id: 1,
	}

	wantInput := &model.ItemGetInput{
		ID: 1,
	}

	tests := []struct {
		name      string
		setupMock func(m *mock.MockItemManagerService)
		wantResp  *pb.GetResponse
		wantErr   bool
	}{
		{
			name: "success",
			setupMock: func(m *mock.MockItemManagerService) {
				m.EXPECT().
					Get(context.Background(), wantInput).
					Return(&model.Item{
						ID:          1,
						Category:    "TestCategory",
						Title:       "TestTitle",
						Description: "TestDescription",
						ImageKey:    "TestImageKey",
						Quantity:    1000,
						Price:       10000,
						CreatedAt:   now,
						UpdatedAt:   now,
					}, nil).
					Once()
			},
			wantResp: &pb.GetResponse{
				Item: &pb.Item{
					Id:          1,
					Category:    "TestCategory",
					Title:       "TestTitle",
					Description: "TestDescription",
					ImageKey:    "TestImageKey",
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
			setupMock: func(m *mock.MockItemManagerService) {
				m.EXPECT().
					Get(context.Background(), wantInput).
					Return(nil, errors.New("failed")).
					Once()
			},
			wantResp: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := mock.NewMockItemManagerService(t)
			tt.setupMock(mockService)

			h := handler.New(mockService)

			resp, err := h.Get(context.Background(), req)
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
