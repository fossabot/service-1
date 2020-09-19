package model

import (
	"reflect"
	"testing"
)

func TestNewGroup(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    Group
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGroup(tt.args.name)
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
