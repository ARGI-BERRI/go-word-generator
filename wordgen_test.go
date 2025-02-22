package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"reflect"
	"testing"
)

type FixedRandSource struct{}

func (d FixedRandSource) get(_ int) int {
	return 0
}

func Test_choiceSound(t *testing.T) {
	tests := []struct {
		name   string
		sounds []string
		want   string
	}{
		{
			name:   "no sound is defined",
			sounds: []string{},
			want:   "",
		},
		{
			name:   "only 1 sound is defined",
			sounds: []string{"a"},
			want:   "a",
		},
		{
			name:   "some sounds are defined",
			sounds: []string{"a", "b", "c"},
			want:   "a",
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if got := choiceSound(tt.sounds, FixedRandSource{}); got != tt.want {
				t.Errorf("choiceSound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateWords(t *testing.T) {
	tests := []struct {
		name   string
		config Config
		want   []string
	}{
		{
			name: "both syllables and sounds are defined",
			config: Config{
				Sounds: map[string][]string{
					"C": {"a", "c"},
					"V": {"b", "d"},
				},
				Pattern: []Pattern{
					{
						Syllable: []string{
							"C", "V",
						},
					},
				},
			},
			want: []string{"ab"},
		},
		{
			name: "patterns are not defined",
			config: Config{
				Sounds: map[string][]string{
					"C": {"a"},
				},
				Pattern: []Pattern{},
			},
			want: []string(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateWords(tt.config, FixedRandSource{}); !reflect.DeepEqual(got, tt.want) {
				t.Logf("got: %#v, want: %#v", got, tt.want)
				t.Errorf("generateWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateWords_Failure(t *testing.T) {
	if os.Getenv("FLAG_BE_CRASHED") == "1" {
		generateWords(Config{
			Sounds: map[string][]string{
				"C": {"a"},
			},
			Pattern: []Pattern{
				{
					Label: "Syllable that isn't defined in Sounds",
					Syllable: []string{
						"U",
					},
				},
			},
		}, FixedRandSource{})

		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=Test_generateWords_Failure")
	cmd.Env = append(os.Environ(), "FLAG_BE_CRASHED=1")

	err := cmd.Run()
	var e *exec.ExitError
	if errors.As(err, &e) && !e.Success() {
		return
	}

	log.Fatalf(err.Error())
}
