package set

import (
	"testing"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func equalSet[T comparable](s1 Set[T], s2 Set[T]) bool {
	if len(s1.m) != len(s2.m) {
		return false
	}
	for k := range s1.m {
		if _, ok := s2.m[k]; !ok {
			return false
		}
	}
	return true
}

func TestOfInt(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		args []int
		want Set[int]
	}{
		"one element": {
			args: []int{1},
			want: Set[int]{
				m: map[int]struct{}{
					1: {},
				},
			},
		},
		"several elements": {
			args: []int{1, 2, 3},
			want: Set[int]{
				m: map[int]struct{}{
					1: {},
					2: {},
					3: {},
				},
			},
		},
		"empty": {
			args: []int{},
			want: Set[int]{
				m: map[int]struct{}{},
			},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			got := Of(tt.args...)
			if !equalSet(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddInt(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		args  []int
		start Set[int]
		want  Set[int]
	}{
		"initialization and one element": {
			args:  []int{1},
			start: Set[int]{},
			want:  Of(1),
		},
		"initialization and several elements": {
			args:  []int{1, 2, 3},
			start: Set[int]{},
			want:  Of(1, 2, 3),
		},
		"initialization and empty": {
			args:  []int{},
			start: Set[int]{},
			want:  Of[int](),
		},
		"no initialization and one element": {
			args:  []int{1},
			start: Of(-1, -2),
			want:  Of(-1, -2, 1),
		},
		"no initialization and several elements": {
			args:  []int{1, 2, 3},
			start: Of(-1, -2),
			want:  Of(-1, -2, 1, 2, 3),
		},
		"no initialization and empty": {
			args:  []int{},
			start: Of(-1, -2),
			want:  Of(-1, -2),
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			tt.start.Add(tt.args...)
			if !equalSet(tt.start, tt.want) {
				t.Fatalf("got %v, want %v", tt.start, tt.want)
			}
		})
	}
}

func TestAddSet(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		args  Set[int]
		start Set[int]
		want  Set[int]
	}{
		"initialization and one element": {
			args:  Of(1),
			start: Set[int]{},
			want:  Of(1),
		},
		"initialization and several elements": {
			args:  Of(1, 2, 3),
			start: Set[int]{},
			want:  Of(1, 2, 3),
		},
		"initialization and empty": {
			args:  Of[int](),
			start: Set[int]{},
			want:  Of[int](),
		},
		"no initialization and one element": {
			args:  Of(1),
			start: Of(-1, -2),
			want:  Of(-1, -2, 1),
		},
		"no initialization and several elements": {
			args:  Of(1, 2, 3),
			start: Of(-1, -2),
			want:  Of(-1, -2, 1, 2, 3),
		},
		"no initialization and empty": {
			args:  Of[int](),
			start: Of(-1, -2),
			want:  Of(-1, -2),
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			tt.start.AddSet(tt.args)
			if !equalSet(tt.start, tt.want) {
				t.Fatalf("got %v, want %v", tt.start, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		args  []int
		start Set[int]
		want  Set[int]
	}{
		"initialization and one element": {
			args:  []int{1},
			start: Set[int]{},
			want:  Of[int](),
		},
		"initialization and several elements": {
			args:  []int{1, 2, 3},
			start: Set[int]{},
			want:  Of[int](),
		},
		"initialization and empty": {
			args:  []int{},
			start: Set[int]{},
			want:  Of[int](),
		},
		"no initialization and one element": {
			args:  []int{1},
			start: Of(1, 2),
			want:  Of(2),
		},
		"no initialization and several elements": {
			args:  []int{1, 2, 3},
			start: Of(-1, -2, 1, 2, 3),
			want:  Of(-1, -2),
		},
		"no initialization and empty": {
			args:  []int{},
			start: Of(-1, -2),
			want:  Of(-1, -2),
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			tt.start.Remove(tt.args...)
			if !equalSet(tt.start, tt.want) {
				t.Fatalf("got %v, want %v", tt.start, tt.want)
			}
		})
	}
}

func TestRemoveSet(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		args  Set[int]
		start Set[int]
		want  Set[int]
	}{
		"initialization and one element": {
			args:  Of(1),
			start: Set[int]{},
			want:  Of[int](),
		},
		"initialization and several elements": {
			args:  Of(1, 2, 3),
			start: Set[int]{},
			want:  Of[int](),
		},
		"initialization and empty": {
			args:  Of[int](),
			start: Set[int]{},
			want:  Of[int](),
		},
		"no initialization and one element": {
			args:  Of(1),
			start: Of(1, 2),
			want:  Of(2),
		},
		"no initialization and several elements": {
			args:  Of(1, 2, 3),
			start: Of(-1, -2, 1, 2, 3),
			want:  Of(-1, -2),
		},
		"no initialization and empty": {
			args:  Of[int](),
			start: Of(-1, -2),
			want:  Of(-1, -2),
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			tt.start.RemoveSet(tt.args)
			if !equalSet(tt.start, tt.want) {
				t.Fatalf("got %v, want %v", tt.start, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		args  int
		start Set[int]
		want  bool
	}{
		"initialization and false": {
			args:  1,
			start: Set[int]{},
			want:  false,
		},
		"no initialization and true": {
			args:  1,
			start: Of(1, 2),
			want:  true,
		},
		"no initialization and false": {
			args:  1,
			start: Of(2, 3),
			want:  false,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			got := tt.start.Contains(tt.args)
			if tt.want != got {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsAny(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		args  Set[int]
		start Set[int]
		want  bool
	}{
		"initialization and false": {
			args:  Of(1, 2),
			start: Set[int]{},
			want:  false,
		},
		"no initialization and true": {
			args:  Of(-1, 0, 1),
			start: Of(1, 2),
			want:  true,
		},
		"no initialization and false": {
			args:  Of(-1, 0, 1),
			start: Of(2, 3),
			want:  false,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			got := tt.start.ContainsAny(tt.args)
			if tt.want != got {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsAll(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		args  Set[int]
		start Set[int]
		want  bool
	}{
		"initialization and false": {
			args:  Of(1, 2),
			start: Set[int]{},
			want:  false,
		},
		"initialization and true": {
			args:  Set[int]{},
			start: Set[int]{},
			want:  true,
		},
		"no initialization and true": {
			args:  Of(1, 2),
			start: Of(1, 2),
			want:  true,
		},
		"no initialization and false": {
			args:  Of(2, 3, 4),
			start: Of(2, 3),
			want:  false,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			got := tt.start.ContainsAll(tt.args)
			if tt.want != got {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSlice(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		start Set[int]
		want  []int
	}{
		"initialization and false": {
			start: Set[int]{},
			want:  []int{},
		},
		"no initialization and empty": {
			start: Of[int](),
			want:  []int{},
		},
		"no initialization and one element": {
			start: Of(1),
			want:  []int{1},
		},
		"no initialization and several elements": {
			start: Of(1, 2, 3),
			want:  []int{1, 2, 3},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			got := tt.start.ToSlice()
			slices.Sort(got)
			slices.Sort(tt.want)
			if !slices.Equal(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClear(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		start Set[int]
		want  Set[int]
	}{
		"initialization and false": {
			start: Set[int]{},
			want:  Of[int](),
		},
		"no initialization and empty": {
			start: Of[int](),
			want:  Of[int](),
		},
		"no initialization and one element": {
			start: Of(1),
			want:  Of[int](),
		},
		"no initialization and several elements": {
			start: Of(1, 2, 3),
			want:  Of[int](),
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			tt.start.Clear()
			if !equalSet(tt.start, tt.want) {
				t.Fatalf("got %v, want %v", tt.start, tt.want)
			}
		})
	}
}

func TestRetain(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		start Set[int]
		arg   func(int) bool
		want  Set[int]
	}{
		"initialization and false": {
			start: Set[int]{},
			arg: func(int) bool {
				return true
			},
			want: Of[int](),
		},
		"no initialization and empty": {
			start: Of[int](),
			arg: func(int) bool {
				return true
			},
			want: Of[int](),
		},
		"no initialization and one element": {
			start: Of(1),
			arg: func(elem int) bool {
				if elem%2 == 0 {
					return true
				}
				return false
			},
			want: Of[int](),
		},
		"no initialization and several elements": {
			start: Of(1, 2, 3),
			arg: func(elem int) bool {
				if elem%2 == 0 {
					return true
				}
				return false
			},
			want: Of(2),
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			tt.start.Retain(tt.arg)
			if !equalSet(tt.start, tt.want) {
				t.Fatalf("got %v, want %v", tt.start, tt.want)
			}
		})
	}
}

func TestLen(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		start Set[int]
		want  int
	}{
		"initialization and false": {
			start: Set[int]{},
			want:  0,
		},
		"no initialization and empty": {
			start: Of[int](),
			want:  0,
		},
		"no initialization and one element": {
			start: Of(1),
			want:  1,
		},
		"no initialization and several elements": {
			start: Of(1, 2, 3),
			want:  3,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			got := tt.start.Len()
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDo(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		start Set[int]
	}{
		"initialization and false": {
			start: Set[int]{},
		},
		"no initialization and empty": {
			start: Of[int](),
		},
		"no initialization and one element": {
			start: Of(1),
		},
		"no initialization and several elements": {
			start: Of(1, 2, 3),
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			var calledTimes int
			mp := map[int]struct{}{}
			tt.start.Do(func(i int) bool {
				mp[i] = struct{}{}
				calledTimes += 1
				return true
			})
			if !maps.Equal(tt.start.m, mp) {
				t.Fatalf("got %v, want %v", mp, tt.start.m)
			}
			if calledTimes != len(tt.start.m) {
				t.Fatalf("calledTimes: got %v, want %v", calledTimes, len(tt.start.m))
			}
		})
	}
}

func TestDoStoppedInTheMiddle(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		start Set[int]
	}{
		"initialization and false": {
			start: Set[int]{},
		},
		"no initialization and empty": {
			start: Of[int](),
		},
		"no initialization and one element": {
			start: Of(1),
		},
		"no initialization and several elements": {
			start: Of(1, 2, 3, 4, 5),
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			added := map[int]struct{}{}
			left := maps.Clone(tt.start.m)
			tt.start.Do(func(i int) bool {
				added[i] = struct{}{}
				delete(left, i)
				if i%2 == 0 {
					return true
				} else {
					return false
				}
			})
			if len(tt.start.m) != len(added)+len(left) {
				t.Fatalf("set len: got %v, want %v", len(tt.start.m), len(added)+len(left))
			}
			for k := range left {
				added[k] = struct{}{}
			}
			if !maps.Equal(tt.start.m, added) {
				t.Fatalf("set elem: got %v, want %v", added, tt.start.m)
			}
		})
	}
}
