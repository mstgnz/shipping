package vatan

import (
	"reflect"
	"strconv"
	"testing"

	cargo "github.com/mstgnz/shipping/config"
)

func TestNewVatanCargo(t *testing.T) {
	tests := []struct {
		want cargo.Shipper
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := NewVatanCargo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVatanCargo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vatanCargo_CancelCargo(t *testing.T) {
	type fields struct {
		Cargo *cargo.Cargo
	}
	type args struct {
		data cargo.ShippingData
	}
	tests := []struct {
		fields  fields
		args    args
		want    *cargo.Response
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			v := vatanCargo{
				Cargo: tt.fields.Cargo,
			}
			got, err := v.CancelCargo(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("CancelCargo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CancelCargo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vatanCargo_CreateCargo(t *testing.T) {
	type fields struct {
		Cargo *cargo.Cargo
	}
	type args struct {
		data cargo.ShippingData
	}
	tests := []struct {
		fields  fields
		args    args
		want    *cargo.Response
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			v := vatanCargo{
				Cargo: tt.fields.Cargo,
			}
			got, err := v.CreateCargo(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCargo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCargo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vatanCargo_WhereIsTheCargo(t *testing.T) {
	type fields struct {
		Cargo *cargo.Cargo
	}
	type args struct {
		data cargo.ShippingData
	}
	tests := []struct {
		fields  fields
		args    args
		want    *cargo.Response
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			v := vatanCargo{
				Cargo: tt.fields.Cargo,
			}
			got, err := v.WhereIsTheCargo(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("WhereIsTheCargo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WhereIsTheCargo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
