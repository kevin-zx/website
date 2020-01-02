package websitetool

import "testing"

func Test_isDisplayFullDomain(t *testing.T) {
	type args struct {
		displayUrl string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDisplayFullDomain(tt.args.displayUrl); got != tt.want {
				t.Errorf("isDisplayFullDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
