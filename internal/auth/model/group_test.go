package model

import (
	"github.com/google/uuid"
	"reflect"
	"testing"
	"time"
)

type MockTime struct{}

func (mt MockTime) Now() time.Time {
	testTime, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	return testTime
}

type MockUUID struct{}

func (mu MockUUID) New() uuid.UUID {
	testUUID, _ := uuid.FromBytes([]byte("seed"))
	return testUUID
}

func TestNewGroup(t *testing.T) {
	wantID := MockUUID{}.New()
	wantTime := MockTime{}.Now()

	type args struct {
		name string
		uuid UUID
		time Time
	}
	tests := []struct {
		name    string
		args    args
		want    Group
		wantErr bool
	}{
		{
			name: "successfully creates a new group",
			args: args{name: "Example Group", uuid: MockUUID{}, time: MockTime{}},
			want: Group{
				ID:        wantID,
				CreatedAt: wantTime,
				UpdatedAt: wantTime,
				Name:      "Example Group",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGroup(tt.args.name, tt.args.uuid, tt.args.time)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
