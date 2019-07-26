package datastore

import (
	"reflect"
	"testing"

	m "github.com/vshelvankar/frontiercg/carmgmt/models"
)

var c1 = m.Car{ID: "123", Make: "Volvo", Model: "XC60", Year: "1999"}
var c2 = m.Car{ID: "456", Make: "Honda", Model: "Civic", Year: "2015"}
var c3 = m.Car{ID: "789", Make: "Hyndai", Model: "I20", Year: "2000"}

var newCar = m.Car{Make: "NewMake", Model: "NewModel", Year: "2020"}

var cars = []m.Car{c1, c2, c3}

var carsCache = map[string]int{
	c1.ID: 0,
	c2.ID: 1,
	c3.ID: 2,
}

var noCars = make([]m.Car, 0)
var noCarsCache = make(map[string]int)

func TestCarsDataStore_GetAll(t *testing.T) {
	type fields struct {
		Cars          []m.Car
		CarsCacheByID map[string]int
	}
	tests := []struct {
		name    string
		fields  fields
		want    []m.Car
		wantErr bool
	}{
		{
			name: "Positive case with pre filled data",
			fields: fields{
				Cars:          cars,
				CarsCacheByID: carsCache,
			},
			want:    cars,
			wantErr: false,
		},
		{
			name: "Negetive case where datastore does not have any data",
			fields: fields{
				Cars:          noCars,
				CarsCacheByID: noCarsCache,
			},
			want:    noCars,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cds := &CarsDataStore{
				Cars:          tt.fields.Cars,
				CarsCacheByID: tt.fields.CarsCacheByID,
			}
			got, err := cds.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("CarsDataStore.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CarsDataStore.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarsDataStore_GetByID(t *testing.T) {
	type fields struct {
		Cars          []m.Car
		CarsCacheByID map[string]int
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *m.Car
		wantErr bool
	}{
		{
			name: "Positive case with existing car",
			fields: fields{
				Cars:          cars,
				CarsCacheByID: carsCache,
			},
			args:    args{"123"},
			want:    &c1,
			wantErr: false,
		},
		{
			name: "Negetive case without car by id",
			fields: fields{
				Cars:          cars,
				CarsCacheByID: carsCache,
			},
			args:    args{"123456"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cds := &CarsDataStore{
				Cars:          tt.fields.Cars,
				CarsCacheByID: tt.fields.CarsCacheByID,
			}
			got, err := cds.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CarsDataStore.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CarsDataStore.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarsDataStore_Create(t *testing.T) {
	type fields struct {
		Cars          []m.Car
		CarsCacheByID map[string]int
	}
	type args struct {
		car *m.Car
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Positive case with valid car",
			fields: fields{
				Cars:          cars,
				CarsCacheByID: carsCache,
			},
			args:    args{&newCar},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cds := &CarsDataStore{
				Cars:          tt.fields.Cars,
				CarsCacheByID: tt.fields.CarsCacheByID,
			}
			got, err := cds.Create(tt.args.car)
			if (err != nil) != tt.wantErr {
				t.Errorf("CarsDataStore.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//GetById
			car, _ := cds.GetByID(got)
			if got != car.ID {
				t.Errorf("CarsDataStore.Create() = %v, want %v", got, car.ID)
			}
		})
	}
}

func TestCarsDataStore_Delete(t *testing.T) {
	type fields struct {
		Cars          []m.Car
		CarsCacheByID map[string]int
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Positive case with existing car",
			fields: fields{
				Cars:          cars,
				CarsCacheByID: carsCache,
			},
			args:    args{"123"},
			wantErr: false,
		},
		{
			name: "Negetive case without car by id",
			fields: fields{
				Cars:          cars,
				CarsCacheByID: carsCache,
			},
			args:    args{"123456"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cds := &CarsDataStore{
				Cars:          tt.fields.Cars,
				CarsCacheByID: tt.fields.CarsCacheByID,
			}
			if err := cds.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("CarsDataStore.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
