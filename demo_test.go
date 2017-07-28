package main

import (
	"reflect"
	"testing"
)

func Test_limparEntrada(t *testing.T) {
	type args struct {
		entrada string
	}
	tests := []struct {
		name string
		args args
		want dataSimples
	}{
		{"entrada ok", args{"10 10 1970"}, dataSimples{10, 10, 1970}},
		{"entrada error", args{"1010 1970"}, dataSimples{0, 0, 0}},
		{"entrada error", args{"10101970"}, dataSimples{0, 0, 0}},
		{"entrada error", args{"10101970\n"}, dataSimples{0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := limparEntrada(tt.args.entrada); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("limparEntrada() = %v, want %v", got, tt.want)
			}
		})
	}
}
