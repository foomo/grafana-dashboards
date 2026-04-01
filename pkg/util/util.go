package util

import (
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
	"go.yaml.in/yaml/v3"
)

func ToPtr[T any](v T) *T {
	return &v
}

func MustYamlToMap(s string) map[string]any {
	ret, err := UnmarshalYaml[map[string]any](s)
	if err != nil {
		panic(err)
	}
	return ret
}

func MustJSONToMap(s string) map[string]any {
	ret, err := UnmarshalJSON[map[string]any](s)
	if err != nil {
		panic(err)
	}
	return ret
}

func MustYamlToSlice(s string) []any {
	ret, err := UnmarshalYaml[[]any](s)
	if err != nil {
		panic(err)
	}
	return ret
}

func MustJSONToSlice(s string) []any {
	ret, err := UnmarshalJSON[[]any](s)
	if err != nil {
		panic(err)
	}
	return ret
}

func UnmarshalJSON[T any](s string) (T, error) {
	var ret T
	s = strings.Trim(s, "\n")
	s = strings.ReplaceAll(s, "\t", "  ")
	lines := strings.Split(s, "\n")
	offset := len(lines[0]) - len(strings.TrimPrefix(lines[0], " "))
	for i, line := range lines {
		lines[i] = strings.TrimPrefix(line, strings.Repeat(" ", offset))
	}
	s = strings.Join(lines, "\n")
	err := json.Unmarshal([]byte(s), &ret)
	if err != nil {
		return ret, errors.Wrap(err, s)
	}
	return ret, err
}

func UnmarshalYaml[T any](s string) (T, error) {
	var ret T
	s = strings.Trim(s, "\n")
	s = strings.ReplaceAll(s, "\t", "  ")
	lines := strings.Split(s, "\n")
	offset := len(lines[0]) - len(strings.TrimPrefix(lines[0], " "))
	for i, line := range lines {
		lines[i] = strings.TrimPrefix(line, strings.Repeat(" ", offset))
	}
	s = strings.Join(lines, "\n")
	err := yaml.Unmarshal([]byte(s), &ret)
	if err != nil {
		return ret, errors.Wrap(err, s)
	}
	return ret, err
}
