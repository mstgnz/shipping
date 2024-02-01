package cargo

import (
	"context"
	"reflect"
	"strconv"
	"testing"
)

func TestCargo_AddCredential(t *testing.T) {
	type fields struct {
		ctx         context.Context
		serviceType ServiceType
		mode        Mode
		isDomestic  bool
		endpoints   []*Endpoint
		credentials []*Credential
	}
	type args struct {
		title    string
		username string
		password string
	}
	tests := []struct {
		fields fields
		args   args
		want   *Credential
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Cargo{
				ctx:         tt.fields.ctx,
				serviceType: tt.fields.serviceType,
				mode:        tt.fields.mode,
				isDomestic:  tt.fields.isDomestic,
				endpoints:   tt.fields.endpoints,
				credentials: tt.fields.credentials,
			}
			if got := c.AddCredential(tt.args.title, tt.args.username, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddCredential() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCargo_AddEndpoint(t *testing.T) {
	type fields struct {
		ctx         context.Context
		serviceType ServiceType
		mode        Mode
		isDomestic  bool
		endpoints   []*Endpoint
		credentials []*Credential
	}
	type args struct {
		title       string
		production  string
		development string
	}
	tests := []struct {
		fields fields
		args   args
		want   *Endpoint
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Cargo{
				ctx:         tt.fields.ctx,
				serviceType: tt.fields.serviceType,
				mode:        tt.fields.mode,
				isDomestic:  tt.fields.isDomestic,
				endpoints:   tt.fields.endpoints,
				credentials: tt.fields.credentials,
			}
			if got := c.AddEndpoint(tt.args.title, tt.args.production, tt.args.development); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCargo_DelCredential(t *testing.T) {
	type fields struct {
		ctx         context.Context
		serviceType ServiceType
		mode        Mode
		isDomestic  bool
		endpoints   []*Endpoint
		credentials []*Credential
	}
	type args struct {
		title string
	}
	tests := []struct {
		fields fields
		args   args
		want   []*Credential
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Cargo{
				ctx:         tt.fields.ctx,
				serviceType: tt.fields.serviceType,
				mode:        tt.fields.mode,
				isDomestic:  tt.fields.isDomestic,
				endpoints:   tt.fields.endpoints,
				credentials: tt.fields.credentials,
			}
			if got := c.DelCredential(tt.args.title); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DelCredential() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCargo_DelEndpoint(t *testing.T) {
	type fields struct {
		ctx         context.Context
		serviceType ServiceType
		mode        Mode
		isDomestic  bool
		endpoints   []*Endpoint
		credentials []*Credential
	}
	type args struct {
		title string
	}
	tests := []struct {
		fields fields
		args   args
		want   []*Endpoint
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Cargo{
				ctx:         tt.fields.ctx,
				serviceType: tt.fields.serviceType,
				mode:        tt.fields.mode,
				isDomestic:  tt.fields.isDomestic,
				endpoints:   tt.fields.endpoints,
				credentials: tt.fields.credentials,
			}
			if got := c.DelEndpoint(tt.args.title); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DelEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCargo_GetContext(t *testing.T) {
	type fields struct {
		ctx         context.Context
		serviceType ServiceType
		mode        Mode
		isDomestic  bool
		endpoints   []*Endpoint
		credentials []*Credential
	}
	tests := []struct {
		fields fields
		want   context.Context
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Cargo{
				ctx:         tt.fields.ctx,
				serviceType: tt.fields.serviceType,
				mode:        tt.fields.mode,
				isDomestic:  tt.fields.isDomestic,
				endpoints:   tt.fields.endpoints,
				credentials: tt.fields.credentials,
			}
			if got := c.GetContext(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCargo_GetCredential(t *testing.T) {
	type fields struct {
		ctx         context.Context
		serviceType ServiceType
		mode        Mode
		isDomestic  bool
		endpoints   []*Endpoint
		credentials []*Credential
	}
	type args struct {
		title string
	}
	tests := []struct {
		fields fields
		args   args
		want   *Credential
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Cargo{
				ctx:         tt.fields.ctx,
				serviceType: tt.fields.serviceType,
				mode:        tt.fields.mode,
				isDomestic:  tt.fields.isDomestic,
				endpoints:   tt.fields.endpoints,
				credentials: tt.fields.credentials,
			}
			if got := c.GetCredential(tt.args.title); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCredential() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCargo_GetCredentials(t *testing.T) {
	type fields struct {
		ctx         context.Context
		serviceType ServiceType
		mode        Mode
		isDomestic  bool
		endpoints   []*Endpoint
		credentials []*Credential
	}
	tests := []struct {
		fields fields
		want   []*Credential
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Cargo{
				ctx:         tt.fields.ctx,
				serviceType: tt.fields.serviceType,
				mode:        tt.fields.mode,
				isDomestic:  tt.fields.isDomestic,
				endpoints:   tt.fields.endpoints,
				credentials: tt.fields.credentials,
			}
			if got := c.GetCredentials(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCargo_GetCurrent(t *testing.T) {
	type fields struct {
		ctx         context.Context
		serviceType ServiceType
		mode        Mode
		isDomestic  bool
		endpoints   []*Endpoint
		credentials []*Credential
	}
	tests := []struct {
		fields fields
		want   Current
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Cargo{
				ctx:         tt.fields.ctx,
				serviceType: tt.fields.serviceType,
				mode:        tt.fields.mode,
				isDomestic:  tt.fields.isDomestic,
				endpoints:   tt.fields.endpoints,
				credentials: tt.fields.credentials,
			}
			if got := c.GetCurrent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCurrent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCargo_GetEndpoint(t *testing.T) {
	type fields struct {
		ctx         context.Context
		serviceType ServiceType
		mode        Mode
		isDomestic  bool
		endpoints   []*Endpoint
		credentials []*Credential
	}
	type args struct {
		title string
	}
	tests := []struct {
		fields fields
		args   args
		want   *Endpoint
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Cargo{
				ctx:         tt.fields.ctx,
				serviceType: tt.fields.serviceType,
				mode:        tt.fields.mode,
				isDomestic:  tt.fields.isDomestic,
				endpoints:   tt.fields.endpoints,
				credentials: tt.fields.credentials,
			}
			if got := c.GetEndpoint(tt.args.title); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCargo_GetEndpoints(t *testing.T) {
	type fields struct {
		ctx         context.Context
		serviceType ServiceType
		mode        Mode
		isDomestic  bool
		endpoints   []*Endpoint
		credentials []*Credential
	}
	tests := []struct {
		fields fields
		want   []*Endpoint
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Cargo{
				ctx:         tt.fields.ctx,
				serviceType: tt.fields.serviceType,
				mode:        tt.fields.mode,
				isDomestic:  tt.fields.isDomestic,
				endpoints:   tt.fields.endpoints,
				credentials: tt.fields.credentials,
			}
			if got := c.GetEndpoints(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEndpoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCargo_GetMode(t *testing.T) {
	type fields struct {
		ctx         context.Context
		serviceType ServiceType
		mode        Mode
		isDomestic  bool
		endpoints   []*Endpoint
		credentials []*Credential
	}
	tests := []struct {
		fields fields
		want   Mode
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Cargo{
				ctx:         tt.fields.ctx,
				serviceType: tt.fields.serviceType,
				mode:        tt.fields.mode,
				isDomestic:  tt.fields.isDomestic,
				endpoints:   tt.fields.endpoints,
				credentials: tt.fields.credentials,
			}
			if got := c.GetMode(); got != tt.want {
				t.Errorf("GetMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCargo_GetServiceType(t *testing.T) {
	type fields struct {
		ctx         context.Context
		serviceType ServiceType
		mode        Mode
		isDomestic  bool
		endpoints   []*Endpoint
		credentials []*Credential
	}
	tests := []struct {
		fields fields
		want   ServiceType
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Cargo{
				ctx:         tt.fields.ctx,
				serviceType: tt.fields.serviceType,
				mode:        tt.fields.mode,
				isDomestic:  tt.fields.isDomestic,
				endpoints:   tt.fields.endpoints,
				credentials: tt.fields.credentials,
			}
			if got := c.GetServiceType(); got != tt.want {
				t.Errorf("GetServiceType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCargo_IsDomestic(t *testing.T) {
	type fields struct {
		ctx         context.Context
		serviceType ServiceType
		mode        Mode
		isDomestic  bool
		endpoints   []*Endpoint
		credentials []*Credential
	}
	tests := []struct {
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Cargo{
				ctx:         tt.fields.ctx,
				serviceType: tt.fields.serviceType,
				mode:        tt.fields.mode,
				isDomestic:  tt.fields.isDomestic,
				endpoints:   tt.fields.endpoints,
				credentials: tt.fields.credentials,
			}
			if got := c.IsDomestic(); got != tt.want {
				t.Errorf("IsDomestic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCargo_SetContext(t *testing.T) {
	type fields struct {
		ctx         context.Context
		serviceType ServiceType
		mode        Mode
		isDomestic  bool
		endpoints   []*Endpoint
		credentials []*Credential
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Cargo{
				ctx:         tt.fields.ctx,
				serviceType: tt.fields.serviceType,
				mode:        tt.fields.mode,
				isDomestic:  tt.fields.isDomestic,
				endpoints:   tt.fields.endpoints,
				credentials: tt.fields.credentials,
			}
			c.SetContext(tt.args.ctx)
		})
	}
}

func TestCargo_SetDomestic(t *testing.T) {
	type fields struct {
		ctx         context.Context
		serviceType ServiceType
		mode        Mode
		isDomestic  bool
		endpoints   []*Endpoint
		credentials []*Credential
	}
	type args struct {
		isDomestic bool
	}
	tests := []struct {
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Cargo{
				ctx:         tt.fields.ctx,
				serviceType: tt.fields.serviceType,
				mode:        tt.fields.mode,
				isDomestic:  tt.fields.isDomestic,
				endpoints:   tt.fields.endpoints,
				credentials: tt.fields.credentials,
			}
			c.SetDomestic(tt.args.isDomestic)
		})
	}
}

func TestCargo_SetMode(t *testing.T) {
	type fields struct {
		ctx         context.Context
		serviceType ServiceType
		mode        Mode
		isDomestic  bool
		endpoints   []*Endpoint
		credentials []*Credential
	}
	type args struct {
		mode Mode
	}
	tests := []struct {
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Cargo{
				ctx:         tt.fields.ctx,
				serviceType: tt.fields.serviceType,
				mode:        tt.fields.mode,
				isDomestic:  tt.fields.isDomestic,
				endpoints:   tt.fields.endpoints,
				credentials: tt.fields.credentials,
			}
			c.SetMode(tt.args.mode)
		})
	}
}

func TestCargo_SetServiceType(t *testing.T) {
	type fields struct {
		ctx         context.Context
		serviceType ServiceType
		mode        Mode
		isDomestic  bool
		endpoints   []*Endpoint
		credentials []*Credential
	}
	type args struct {
		serviceType ServiceType
	}
	tests := []struct {
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := &Cargo{
				ctx:         tt.fields.ctx,
				serviceType: tt.fields.serviceType,
				mode:        tt.fields.mode,
				isDomestic:  tt.fields.isDomestic,
				endpoints:   tt.fields.endpoints,
				credentials: tt.fields.credentials,
			}
			c.SetServiceType(tt.args.serviceType)
		})
	}
}
