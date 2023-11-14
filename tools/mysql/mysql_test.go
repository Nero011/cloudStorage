package mysql

import (
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestGetMysql(t *testing.T) {
	tests := []struct {
		name string
		want *gorm.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMysql(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMysql() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initMysql(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initMysql()
		})
	}
}
