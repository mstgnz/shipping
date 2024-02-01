package turk

import (
	"reflect"
	"strconv"
	"testing"

	cargo "github.com/mstgnz/shipping/config"
)

func TestNewTurkCargo(t *testing.T) {
	tests := []struct {
		want cargo.Shipper
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := NewTurkCargo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTurkCargo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_turkCargo_CancelCargo(t1 *testing.T) {
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
		t1.Run(strconv.Itoa(i), func(t1 *testing.T) {
			t := turkCargo{
				Cargo: tt.fields.Cargo,
			}
			got, err := t.CancelCargo(tt.args.data)
			if (err != nil) != tt.wantErr {
				t1.Errorf("CancelCargo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("CancelCargo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_turkCargo_CreateCargo(t1 *testing.T) {
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
		t1.Run(strconv.Itoa(i), func(t1 *testing.T) {
			t := turkCargo{
				Cargo: tt.fields.Cargo,
			}
			got, err := t.CreateCargo(tt.args.data)
			if (err != nil) != tt.wantErr {
				t1.Errorf("CreateCargo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("CreateCargo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_turkCargo_WhereIsTheCargo(t1 *testing.T) {
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
		t1.Run(strconv.Itoa(i), func(t1 *testing.T) {
			t := turkCargo{
				Cargo: tt.fields.Cargo,
			}
			got, err := t.WhereIsTheCargo(tt.args.data)
			if (err != nil) != tt.wantErr {
				t1.Errorf("WhereIsTheCargo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("WhereIsTheCargo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
