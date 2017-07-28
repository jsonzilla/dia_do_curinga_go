package main

import (
	"reflect"
	"testing"
)

func Test_ehDataValida(t *testing.T) {
	type args struct {
		dia int
		mes int
		ano int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valido", args{10, 1, 1999}, true},
		{"dia 0", args{0, 1, 1999}, false},
		{"dia 32", args{32, 1, 1999}, false},
		{"29 fev", args{29, 2, 2012}, true},
		{"29 fev", args{29, 2, 2011}, false},
		{"mes 0", args{1, 0, 1999}, false},
		{"mes 13", args{1, 13, 1999}, false},
		{"ano 0", args{1, 12, 0}, false},
		{"ano 0", args{1, 12, 1789}, false},
		{"ano 0", args{1, 12, 1790}, true},
		{"ano 9999", args{1, 12, 9999}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ehDataValida(tt.args.dia, tt.args.mes, tt.args.ano); got != tt.want {
				t.Errorf("ehDataValida() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intAnoBissexto(t *testing.T) {
	type args struct {
		ano int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"eh", args{2012}, 1},
		{"nao", args{2011}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intAnoBissexto(tt.args.ano); got != tt.want {
				t.Errorf("intAnoBissexto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ehAnoBissexto(t *testing.T) {
	type args struct {
		ano int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"eh", args{2012}, true},
		{"nao", args{2011}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ehAnoBissexto(tt.args.ano); got != tt.want {
				t.Errorf("ehAnoBissexto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrodeSimples(t *testing.T) {
	type args struct {
		dia int
		mes int
		ano int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"28, 2, 2011 ", args{28, 2, 2011}, "Jo1O1P1P"},
		{"1, 3, 2011 ", args{1, 3, 2011}, "1O1O1P1P"},
		{"27, 2, 2012", args{27, 2, 2012}, "QEKEKP2P"},
		{"28, 2, 2012", args{28, 2, 2012}, "KE1O1P2P"},
		{"29, 2, 2012", args{29, 2, 2012}, "Jo1O1P2P"},
		{"1, 3, 2012", args{1, 3, 2012}, "1O1O1P2P"},
		{"28, 2, 2013", args{28, 2, 2013}, "Jo1O1P3P"},
		{"1, 3, 2013", args{1, 3, 2013}, "1O1O1P3P"},
		{"28, 2, 2014", args{28, 2, 2014}, "Jo1O1P4P"},
		{"1, 3, 2014", args{1, 3, 2014}, "1O1O1P4P"},
		{"19, 7, 2017", args{19, 7, 2017}, "JC8P6P7P"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FrodeSimples(tt.args.dia, tt.args.mes, tt.args.ano); got != tt.want {
				t.Errorf("FrodeSimples() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contaPorMes(t *testing.T) {
	type args struct {
		dia int
		mes int
		ano int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{19, 7, 2017}, 200},
		{"", args{31, 12, 2017}, 365},
		{"", args{1, 1, 2017}, 1},
		{"", args{31, 12, 2107}, 365},
		{"", args{31, 12, 2012}, 366},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contaPorMes(tt.args.dia, tt.args.mes, tt.args.ano); got != tt.want {
				t.Errorf("contaPorMes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_diaDoAno(t *testing.T) {
	type args struct {
		dia int
		mes int
		ano int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{31, 17, 2017}, 0},
		{"", args{19, 7, 2017}, 200},
		{"", args{31, 12, 2017}, 365},
		{"", args{1, 1, 2017}, 1},
		{"", args{31, 12, 2107}, 365},
		{"", args{31, 12, 2012}, 366},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diaDoAno(tt.args.dia, tt.args.mes, tt.args.ano); got != tt.want {
				t.Errorf("diaDoAno() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_corrigeDiaFrode(t *testing.T) {
	type args struct {
		dia int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{200}, 140},
		{"", args{59}, 364},
		{"", args{61}, 1},
		{"", args{59}, 364},
		{"", args{60}, 365},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := corrigeDiaFrode(tt.args.dia); got != tt.want {
				t.Errorf("corrigeDiaFrode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_corrigeDiaFrodeVerificaBissexto(t *testing.T) {
	type args struct {
		dia int
		ano int
	}
	tests := []struct {
		name string
		args args
		want int
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := corrigeDiaFrodeVerificaBissexto(tt.args.dia, tt.args.ano); got != tt.want {
				t.Errorf("corrigeDiaFrodeVerificaBissexto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrode(t *testing.T) {
	type args struct {
		dia int
		mes int
		ano int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"default", args{19, 7, 2017}, "\n\tDia de de Valete de copas\n\tSemana de Oito de paus\n\tMes de de Seis estacao de paus\n\tAno de de Sete de paus\n\t19/7/2017 e dia numero 141"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Frode(tt.args.dia, tt.args.mes, tt.args.ano); got != tt.want {
				t.Errorf("Frode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_corrigeAnoFrode(t *testing.T) {
	type args struct {
		ano int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ano negativo", args{1789}, 1},
		{"ano 1", args{1790}, 0},
		{"ano 2", args{1791}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := corrigeAnoFrode(tt.args.ano); got != tt.want {
				t.Errorf("corrigeAnoFrode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_estacoesFrode(t *testing.T) {
	type args struct {
		dia int
		ano int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"paus", args{62, 1790}, 1},
		{"copas", args{154, 1790}, 2},
		{"espadas", args{247, 1790}, 3},
		{"ouros", args{338, 1790}, 0},
		{"paus", args{365, 1790}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := estacoesFrode(tt.args.dia, tt.args.ano); got != tt.want {
				t.Errorf("estacoesFrode() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
