package main

import "testing"

func Test_getJSON(t *testing.T) {
	type args struct {
		data []User
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Test get JSON", args{[]User{
			{"Азамат", 23, []Comment{
				{"Первый комментарий"},
				{"Второй комментарий"},
			}},
			{"Алмаз", 15, []Comment{
				{"Комментарий"},
			},
			},
		}}, "[{\"name\":\"Азамат\",\"age\":23,\"comments\":[{\"text\":\"Первый комментарий\"},{\"text\":\"Второй комментарий\"}]},{\"name\":\"Алмаз\",\"age\":15,\"comments\":[{\"text\":\"Комментарий\"}]}]", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getJSON(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("getJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getJSON() = %v, want %v", got, tt.want)
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
