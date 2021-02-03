package hash_test

import (
	"testing"

	"github.com/one-piece-official/Nimbus/hash"
)

// NOTE: t.Paraller()
// parallel 是一个并行库，用于不改变现有接口声明前提下的业务聚合或者重构。
// 如果需要循环某一个函数只是参数不同，则可以使用 parallel，它可以并行运行，极大提高运行速度。
// 相关阅读：
// - https://github.com/kunwardeep/paralleltest
// - https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721
// - https://blog.csdn.net/u012189747/article/details/77878583
func TestMD5(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{s: "hello"},
			want: "5d41402abc4b2a76b9719d911017c592",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := hash.MD5(tt.args.s); got != tt.want {
				t.Errorf("MD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA1(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{s: "hello"},
			want: "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := hash.SHA1(tt.args.s); got != tt.want {
				t.Errorf("SHA1() = %v, want %v", got, tt.want)
			}
		})
	}
}
