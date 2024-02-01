package aras

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/mstgnz/shipping/config"
)

func TestNewArasCargo(t *testing.T) {
	tests := []struct {
		want cargo.Shipper
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := NewArasCargo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArasCargo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_arasCargo_CancelCargo(t *testing.T) {
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
			a := arasCargo{
				Cargo: tt.fields.Cargo,
			}
			got, err := a.CancelCargo(tt.args.data)
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

func Test_arasCargo_CreateCargo(t *testing.T) {
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
			a := arasCargo{
				Cargo: tt.fields.Cargo,
			}
			got, err := a.CreateCargo(tt.args.data)
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

func Test_arasCargo_WhereIsTheCargo(t *testing.T) {
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
			a := arasCargo{
				Cargo: tt.fields.Cargo,
			}
			got, err := a.WhereIsTheCargo(tt.args.data)
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
