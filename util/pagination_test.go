package util

import "testing"

func TestGetPageCount(t *testing.T) {
	type args struct {
		pageSize  int
		totalData int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic test",
			args: args{
				pageSize:  1,
				totalData: 1,
			},
			want: 1,
		},
		{
			name: "data pagination",
			args: args{
				pageSize:  5,
				totalData: 10,
			},
			want: 2,
		},
		{
			name: "over data pagination",
			args: args{
				pageSize:  6,
				totalData: 5,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPageCount(tt.args.pageSize, tt.args.totalData); got != tt.want {
				t.Errorf("GetPageCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
