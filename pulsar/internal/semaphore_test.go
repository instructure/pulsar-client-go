package internal

import "testing"

func TestSemaphoreTryAcquire(t *testing.T) {
	tests := []struct {
		name string
		s    Semaphore
		want bool
	}{
		{
			name: "should succeed when there is capacity",
			s: make(Semaphore, 1),
			want: true,
		},
		{
			name: "should fail when no capacity",
			s: make(Semaphore, 0),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.TryAcquire(); got != tt.want {
				t.Errorf("TryAcquire() = %v, want %v", got, tt.want)
			}
		})
	}
}