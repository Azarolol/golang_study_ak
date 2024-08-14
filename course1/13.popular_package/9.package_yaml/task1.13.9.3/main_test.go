package main

import "testing"

func Test_writeYAML(t *testing.T) {
	type args struct {
		filePath string
		data     []User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test write YAML", args{"test/file.yaml", []User{
			{"Азамат", 23, []Comment{
				{"Первый комментарий"},
				{"Второй комментарий"},
			}},
			{"Алмаз", 15, []Comment{
				{"Комментарий"},
			},
			},
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeYAML(tt.args.filePath, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("writeYAML() error = %v, wantErr %v", err, tt.wantErr)
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
