package service

import (
	"github.com/sirupsen/logrus/hooks/test"
	"order-service/internal/model"
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
)

func getMockLogger() (*logrus.Entry, *test.Hook) {
	logger, hook := test.NewNullLogger()
	logEntry := logrus.NewEntry(logger)
	return logEntry, hook
}

func TestNewCloudPillarRepositoryImpl(t *testing.T) {

	mockLogger, _ := getMockLogger()
	type args struct {
		logger *logrus.Entry
	}

	tests := []struct {
		name string
		args args
		want OrderService
	}{
		{
			name: "Test to set a nil logger",
			args: args{
				logger: nil,
			},
			want: &orderServiceImpl{logger: nil},
		}, {
			name: "Test to set a mock logger",
			args: args{
				logger: mockLogger,
			},
			want: &orderServiceImpl{logger: mockLogger},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrderService(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCloudPillarRepositoryImpl() = %v, want %v", got, tt.want)
			}
		})

	}
}

func TestOrderServiceImpl_CreateOrder(t *testing.T) {

	mockOrder := model.Order{
		BankID: "test",
		Amount: 2000,
	}
	mockLogger, _ := getMockLogger()

	tests := []struct {
		name    string
		mock    model.Order
		wantErr bool
	}{
		{
			name:    "expect nil,error when gorm query returns ErrInvalidSQL",
			mock:    mockOrder,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderService := orderServiceImpl{
				logger: mockLogger,
			}
			_, err := orderService.CreateOrder(tt.mock)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
