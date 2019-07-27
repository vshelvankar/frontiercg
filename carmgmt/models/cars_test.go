package models

import (
	"net/http"
	"reflect"
	"testing"
)

var c1 = Car{ID: "123", Make: "Volvo", Model: "XC60", Year: "1999"}

var carRequest = CarRequest{Make: "NewMake", Model: "NewModel", Year: "2020"}

func TestNewCarResponse(t *testing.T) {
	type args struct {
		car *Car
	}
	tests := []struct {
		name string
		args args
		want *CarResponse
	}{
		{
			name: "New CarResponse from Car. Will not have ID in it",
			args: args{&c1},
			want: &CarResponse{
				Make:  c1.Make,
				Model: c1.Model,
				Year:  c1.Year,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCarResponse(tt.args.car); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCarResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCar(t *testing.T) {
	type args struct {
		cr *CarRequest
	}
	tests := []struct {
		name string
		args args
		want *Car
	}{
		{
			name: "New Car from CarRequest",
			args: args{&carRequest},
			want: &Car{
				Make:  carRequest.Make,
				Model: carRequest.Model,
				Year:  carRequest.Year,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCar(tt.args.cr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarRequest_Bind(t *testing.T) {
	type fields struct {
		Make  string
		Model string
		Year  string
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Bind post process validation success case",
			fields:  fields{"abc", "def", "2000"},
			args:    args{&http.Request{}},
			wantErr: false,
		},
		{
			name:    "Bind post process validation failure case where make is empty",
			fields:  fields{"", "def", "2000"},
			args:    args{&http.Request{}},
			wantErr: true,
		},
		{
			name:    "Bind post process validation failure case where model is empty",
			fields:  fields{"abc", "", "2000"},
			args:    args{&http.Request{}},
			wantErr: true,
		},
		{
			name:    "Bind post process validation failure case where year is empty",
			fields:  fields{"abc", "def", ""},
			args:    args{&http.Request{}},
			wantErr: true,
		},
		{
			name:    "Bind post process validation failure case where year is not a number",
			fields:  fields{"abc", "def", "aaaa"},
			args:    args{&http.Request{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CarRequest{
				Make:  tt.fields.Make,
				Model: tt.fields.Model,
				Year:  tt.fields.Year,
			}
			if err := c.Bind(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("CarRequest.Bind() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
