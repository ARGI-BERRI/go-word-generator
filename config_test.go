package main

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_loadConfig(t *testing.T) {
	tests := []struct {
		name string
		json string
		want Config
	}{
		{
			name: "valid config with empty content",
			json: "{}",
			want: Config{},
		},
		{
			name: "valid config with contents",
			json: `{"sounds": {"dummy": ["a"]}, "patterns": []}`,
			want: Config{
				Sounds: map[string][]string{
					"dummy": []string{"a"},
				},
				Pattern: make([]Pattern, 0),
			},
		},
		{
			name: "invalid config with corrupted json",
			json: "...",
			want: Config{},
		},
		{
			name: "invalid config with null json",
			json: "",
			want: Config{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.CreateTemp("", "")

			defer func(file *os.File) {
				_ = file.Close()
			}(file)

			_, _ = file.WriteString(tt.json)

			if strings.HasPrefix(tt.name, "invalid") {
				_, err := loadConfig(file.Name())

				if err == nil {
					t.Errorf("loadConfig() should return error")
				}
			}

			if got, _ := loadConfig(file.Name()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
