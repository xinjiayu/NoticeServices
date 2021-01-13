package tools

import "testing"

func TestCreateCommunicateId(t *testing.T) {
	type args struct {
		fromUser int
		toUser   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1", args{fromUser: 15820123456, toUser: 15699586652}, "3eb5e3d5cd2c71ef6fce3f391c9e9019"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateCommunicateId(tt.args.fromUser, tt.args.toUser); got != tt.want {
				t.Errorf("CreateCommunicateId() = %v, want %v", got, tt.want)
			}
		})
	}
}
