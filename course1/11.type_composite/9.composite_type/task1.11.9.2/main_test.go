package main

import "testing"

func TestSamsungTV_switchOn(t *testing.T) {
	tests := []struct {
		name string
		tv   *SamsungTV
		want bool
	}{
		{"Test SamsungTV switch on", &SamsungTV{false, "Samsung XL-100500"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tv.switchOn(); got != tt.want {
				t.Errorf("SamsungTV.switchOn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLGTV_switchOn(t *testing.T) {
	tests := []struct {
		name string
		tv   *LGTV
		want bool
	}{
		{"Test LGTV switch on", &LGTV{false, "LG TV"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tv.switchOn(); got != tt.want {
				t.Errorf("LGTV.switchOn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSamsungTV_switchOFF(t *testing.T) {
	tests := []struct {
		name string
		tv   *SamsungTV
		want bool
	}{
		{"Test SamsungTV switch off", &SamsungTV{true, "Samsung XL-100500"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tv.switchOFF(); got != tt.want {
				t.Errorf("SamsungTV.switchOFF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLGTV_switchOFF(t *testing.T) {
	tests := []struct {
		name string
		tv   *LGTV
		want bool
	}{
		{"Test LGTV switch off", &LGTV{true, "LG TV"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tv.switchOFF(); got != tt.want {
				t.Errorf("LGTV.switchOFF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSamsungTV_GetModel(t *testing.T) {
	tests := []struct {
		name string
		tv   *SamsungTV
		want string
	}{
		{"Test SamsungTV get model", &SamsungTV{true, "Samsung XL-100500"}, "Samsung XL-100500"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tv.GetModel(); got != tt.want {
				t.Errorf("SamsungTV.GetModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLGTV_GetModel(t *testing.T) {
	tests := []struct {
		name string
		tv   *LGTV
		want string
	}{
		{"Test LGTV get model", &LGTV{true, "LG TV"}, "LG TV"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tv.GetModel(); got != tt.want {
				t.Errorf("LGTV.GetModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSamsungTV_GetStatus(t *testing.T) {
	tests := []struct {
		name string
		tv   *SamsungTV
		want bool
	}{
		{"Test SamsungTV get status", &SamsungTV{true, "Samsung XL-100500"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tv.GetStatus(); got != tt.want {
				t.Errorf("SamsungTV.GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLgTV_GetStatus(t *testing.T) {
	tests := []struct {
		name string
		tv   *LGTV
		want bool
	}{
		{"Test LGTV get status", &LGTV{true, "LG TV"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tv.GetStatus(); got != tt.want {
				t.Errorf("LGTV.GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSamsungTV_SamsungHub(t *testing.T) {
	tests := []struct {
		name string
		tv   *SamsungTV
		want string
	}{
		{"Test SamsungTV Samsung Hub", &SamsungTV{true, "Samsung XL-100500"}, "SamsungHub"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tv.SamsungHub(); got != tt.want {
				t.Errorf("SamsungTV.SamsungHub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLGTV_LGHub(t *testing.T) {
	tests := []struct {
		name string
		tv   *LGTV
		want string
	}{
		{"Test LGTV LG Hub", &LGTV{true, "LG TV"}, "LGHub"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tv.LGHub(); got != tt.want {
				t.Errorf("LGTV.LGHub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test main"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
