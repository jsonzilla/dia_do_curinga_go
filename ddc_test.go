package main

import (
	"reflect"
	"testing"
)

func Test_validDate(t *testing.T) {
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
		{"date 10, 1, 1999 should be a valid date", args{10, 1, 1999}, true},
		{"date 0, 1, 1999 should be a invalid date", args{0, 1, 1999}, false},
		{"date 32, 1, 1999 should be a invalid date", args{32, 1, 1999}, false},
		{"date 29, 2, 2012 should be a valid date", args{29, 2, 2012}, true},
		{"date 29, 2, 2011 should be a invalid date", args{29, 2, 2011}, false},
		{"date 1, 0, 1999 should be a invalid date", args{1, 0, 1999}, false},
		{"date 1, 13, 1999 should be a invalid date", args{1, 13, 1999}, false},
		{"date 1, 12, 0    should be a invalid date", args{1, 12, 0}, false},
		{"date 1, 12, 1789 should be a valid date", args{1, 12, 1789}, true},
		{"date 1, 12, 1790 should be a valid date", args{1, 12, 1790}, true},
		{"date 1, 12, 9999 should be a valid date", args{1, 12, 9999}, true},
		{"date 1, 4, 9999 should be a valid date", args{1, 4, 9999}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validDate(tt.args.dia, tt.args.mes, tt.args.ano); got != tt.want {
				t.Errorf("validDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_leapYearInt(t *testing.T) {
	type args struct {
		ano int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2012 should be a leap day", args{2012}, 1},
		{"2011 should not be a leap day", args{2011}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leapYearInt(tt.args.ano); got != tt.want {
				t.Errorf("leapYearInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isLeapYear(t *testing.T) {
	type args struct {
		ano int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"2012 should be a leap day", args{2012}, true},
		{"2011 should not be a leap day", args{2011}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLeapYear(tt.args.ano); got != tt.want {
				t.Errorf("isLeapYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortVersion(t *testing.T) {
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
		{"date 28, 2, 2011 should be KE1O1P1P", args{28, 2, 2011}, "KE1O1P1P"},
		{"date 1, 3, 2011  should be 1O1O1P1P", args{1, 3, 2011}, "1O1O1P1P"},
		{"date 27, 2, 2012 should be QEKEKP2P", args{27, 2, 2012}, "QEKEKP2P"},
		{"date 28, 2, 2012 should be KE1O1P2P", args{28, 2, 2012}, "KE1O1P2P"},
		{"date 29, 2, 2012 should be Jo1O1P2P", args{29, 2, 2012}, "Jo1O1P2P"},
		{"date 1, 3, 2012 should be 1O1O1P2P", args{1, 3, 2012}, "1O1O1P2P"},
		{"date 28, 2, 2013 should be KE1O1P3P", args{28, 2, 2013}, "KE1O1P3P"},
		{"date 1, 3, 2013 should be 1O1O1P3P", args{1, 3, 2013}, "1O1O1P3P"},
		{"date 28, 2, 2014 should be KE1O1P4P", args{28, 2, 2014}, "KE1O1P4P"},
		{"date 1, 3, 2014 should be 1O1O1P4P", args{1, 3, 2014}, "1O1O1P4P"},
		{"date 19, 7, 2017 should be 10C8P6P7P", args{19, 7, 2017}, "10C8P6P7P"},
		{"invalid date should be return a empty string", args{19, 70, 2017}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShortVersion(tt.args.dia, tt.args.mes, tt.args.ano); got != tt.want {
				t.Errorf("ShortVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countDays(t *testing.T) {
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
		{"date 19, 7, 2017 shouldBeCount 200 days", args{19, 7, 2017}, 200},
		{"date 31, 12, 2017 shouldBeCount 365 days", args{31, 12, 2017}, 365},
		{"date 1, 1, 20 shouldBeCount}, 1 days", args{1, 1, 2017}, 1},
		{"date 31, 12, 2107 shouldBeCount 365 days", args{31, 12, 2107}, 365},
		{"date 1, 11, 2012 shouldBeCount 366 days", args{1, 11, 2012}, 306},
		{"date 1, 10, 2012 shouldBeCount 366 days", args{1, 10, 2012}, 275},
		{"date 1, 9, 2012 shouldBeCount 366 days", args{1, 9, 2012}, 245},
		{"date 1, 8, 2012 shouldBeCount 366 days", args{1, 8, 2012}, 214},
		{"date 1, 7, 2012 shouldBeCount 366 days", args{1, 7, 2012}, 183},
		{"date 1, 6, 2012 shouldBeCount 366 days", args{1, 6, 2012}, 153},
		{"date 1, 5, 2012 shouldBeCount 366 days", args{1, 5, 2012}, 122},
		{"date 1, 4, 2012 shouldBeCount 366 days", args{1, 4, 2012}, 92},
		{"date 1, 3, 2012 shouldBeCount 366 days", args{1, 3, 2012}, 61},
		{"date 1, 2, 2012 shouldBeCount 366 days", args{1, 2, 2012}, 32},
		{"date 1, 1, 2012 shouldBeCount 366 days", args{1, 1, 2012}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDays(tt.args.dia, tt.args.mes, tt.args.ano); got != tt.want {
				t.Errorf("countDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dayOfYear(t *testing.T) {
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
		{"date 31, 17, 2017 shouldBeCount 0 days", args{31, 17, 2017}, 0},
		{"date 19, 7, 2017 shouldBeCount 200 days", args{19, 7, 2017}, 200},
		{"date 31, 12, 2017 shouldBeCount 365 days", args{31, 12, 2017}, 365},
		{"date 1, 1, 20 shouldBeCount}, 1 day", args{1, 1, 2017}, 1},
		{"date 31, 12, 2107 shouldBeCount 365 days", args{31, 12, 2107}, 365},
		{"date 31, 12, 2012 shouldBeCount 366 days", args{31, 12, 2012}, 366},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dayOfYear(tt.args.dia, tt.args.mes, tt.args.ano); got != tt.want {
				t.Errorf("dayOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fixDay(t *testing.T) {
	type args struct {
		ano int
		dia int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"date 2011, 200 should be  140", args{2011, 200}, 140},
		{"date 2011, 59 should be  364", args{2011, 59}, 364},
		{"date 2011, 61 should be 1", args{2011, 61}, 1},
		{"date 2011, 59 should be  364", args{2011, 59}, 364},
		{"date 2011, 60 should be  365", args{2011, 60}, 365},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fixDay(tt.args.ano, tt.args.dia); got != tt.want {
				t.Errorf("fixDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongVersion(t *testing.T) {
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
		{"default test with date 19 7 2017 should be ok", args{19, 7, 2017}, "\n\tDia Dez de copas\n\tSemana Oito de paus\n\tMes Seis estacao de paus\n\tAno Sete de paus\n\t19/7/2017 e dia numero 140"},
		{"invalid date should be return a empty string", args{19, 70, 2017}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongVersion(tt.args.dia, tt.args.mes, tt.args.ano); got != tt.want {
				t.Errorf("LongVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fixYear(t *testing.T) {
	type args struct {
		ano int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"year 1789 should be return 1 in conversion", args{1789}, 1},
		{"year 1790 should be return 0 in conversion", args{1790}, 0},
		{"year 1791 should be return 1 in conversion", args{1791}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fixYear(tt.args.ano); got != tt.want {
				t.Errorf("fixYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_seasons(t *testing.T) {
	type args struct {
		dia int
		ano int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"day 62 in year 1790 should be return 1 : club", args{62, 1790}, 1},
		{"day 154 in year 1790 should be return 2 : cup", args{154, 1790}, 2},
		{"day 247 in year 1790 should be return 3 : swords", args{247, 1790}, 3},
		{"day 338 in year 1790 should be return 0 : golds", args{338, 1790}, 0},
		{"day 365 in year 1790 should be return 1 : club", args{365, 1790}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := seasons(tt.args.dia, tt.args.ano); got != tt.want {
				t.Errorf("seasons() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearInput(t *testing.T) {
	type args struct {
		entrada string
	}
	tests := []struct {
		name string
		args args
		want simpleDate
	}{
		{"input 10 10 1970 should be an input correct", args{"10 10 1970"}, simpleDate{10, 10, 1970}},
		{"input 1010 1970 should be an input error", args{"1010 1970"}, simpleDate{0, 0, 0}},
		{"input 10101970 should be an input error", args{"10101970"}, simpleDate{0, 0, 0}},
		{"input 10101970 should be an input error", args{"10101970\n"}, simpleDate{0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clearInput(tt.args.entrada); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clearInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_feb(t *testing.T) {
	type args struct {
		dia int
		ano int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"date 30 february 2012 should be not valid", args{30, 2012}, false},
		{"date 29 february 2011 should be not valid", args{29, 2011}, false},
		{"date 29 february 2012 should be valid", args{29, 2012}, true},
		{"date 28 february 2013 should be valid", args{28, 2013}, true},
		{"date 30 february 2011 should be not valid", args{30, 2011}, false},
	}
	for _, tt := range tests {
		if got := feb(tt.args.dia, tt.args.ano); got != tt.want {
			t.Errorf("%q. feb() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
