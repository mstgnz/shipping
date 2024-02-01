package rest

import (
	"reflect"
	"strconv"
	"testing"

	cargo "github.com/mstgnz/shipping/config"
)

func TestCreateAbroad(t *testing.T) {
	type args struct {
		current cargo.Current
		data    cargo.ShippingData
	}
	tests := []struct {
		args    args
		want    *cargo.Response
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := CreateAbroad(tt.args.current, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAbroad() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAbroad() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateDomestic(t *testing.T) {
	type args struct {
		current cargo.Current
		data    cargo.ShippingData
	}
	tests := []struct {
		args    args
		want    *cargo.Response
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := CreateDomestic(tt.args.current, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateDomestic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDomestic() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWhereIsTheCargoAbroad(t *testing.T) {
	type args struct {
		current cargo.Current
		data    cargo.ShippingData
	}
	tests := []struct {
		args    args
		want    *cargo.Response
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := WhereIsTheCargoAbroad(tt.args.current, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("WhereIsTheCargoAbroad() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WhereIsTheCargoAbroad() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWhereIsTheCargoDomestic(t *testing.T) {
	type args struct {
		current cargo.Current
		data    cargo.ShippingData
	}
	tests := []struct {
		args    args
		want    *cargo.Response
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := WhereIsTheCargoDomestic(tt.args.current, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("WhereIsTheCargoDomestic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WhereIsTheCargoDomestic() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCancelCargoAbroad(t *testing.T) {
	type args struct {
		current cargo.Current
		data    cargo.ShippingData
	}
	tests := []struct {
		args    args
		want    *cargo.Response
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := CancelCargoAbroad(tt.args.current, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("CancelCargoAbroad() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CancelCargoAbroad() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCancelCargoDomestic(t *testing.T) {
	type args struct {
		current cargo.Current
		data    cargo.ShippingData
	}
	tests := []struct {
		args    args
		want    *cargo.Response
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := CancelCargoDomestic(tt.args.current, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("CancelCargoDomestic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CancelCargoDomestic() got = %v, want %v", got, tt.want)
			}
		})
	}
}