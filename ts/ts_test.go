package ts

import (
	"fmt"
	"testing"
	"time"

	"github.com/jimeh/go-tyme/dur"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	testAdd[time.Time, time.Duration](t)
	testAdd[time.Time, dur.Duration](t)

	testAdd[Second, time.Duration](t)
	testAdd[Second, dur.Duration](t)

	testAdd[Millisecond, time.Duration](t)
	testAdd[Millisecond, dur.Duration](t)

	testAdd[Microsecond, time.Duration](t)
	testAdd[Microsecond, dur.Duration](t)

	testAdd[Nanosecond, time.Duration](t)
	testAdd[Nanosecond, dur.Duration](t)
}

func testAdd[T Timestamp, D Duration](t *testing.T) {
	t.Run(
		fmt.Sprintf("[T %T, D %T]", T(time.Time{}), D(0)),
		func(t *testing.T) {
			tests := []struct {
				name string
				t    T
				d    D
				want T
			}{
				{
					name: "add 1s",
					t: T(
						time.Date(2020, 1, 1, 0, 0, 2, 0, time.UTC),
					),
					d: D(time.Second),
					want: T(
						time.Date(2020, 1, 1, 0, 0, 3, 0, time.UTC),
					),
				},
				{
					name: "remove 1s",
					t: T(
						time.Date(2020, 1, 1, 0, 0, 2, 0, time.UTC),
					),
					d: D(-time.Second),
					want: T(
						time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
					),
				},
				{
					name: "add 250ms",
					t: T(
						time.Date(2020, 1, 1, 0, 0, 2, 0, time.UTC),
					),
					d: D(250 * time.Millisecond),
					want: T(
						time.Date(2020, 1, 1, 0, 0, 2, 250000000, time.UTC),
					),
				},
				{
					name: "remove 250ms",
					t: T(
						time.Date(2020, 1, 1, 0, 0, 2, 0, time.UTC),
					),
					d: D(-250 * time.Millisecond),
					want: T(
						time.Date(2020, 1, 1, 0, 0, 1, 750000000, time.UTC),
					),
				},
			}
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := Add(tt.t, tt.d)

					assert.Equal(t, tt.want, got)
				})
			}
		},
	)
}

func TestSub(t *testing.T) {
	testSub[time.Time, time.Time](t)
	testSub[time.Time, Second](t)
	testSub[time.Time, Millisecond](t)
	testSub[time.Time, Microsecond](t)
	testSub[time.Time, Nanosecond](t)

	testSub[Second, time.Time](t)
	testSub[Second, Second](t)
	testSub[Second, Millisecond](t)
	testSub[Second, Microsecond](t)
	testSub[Second, Nanosecond](t)

	testSub[Millisecond, time.Time](t)
	testSub[Millisecond, Millisecond](t)
	testSub[Millisecond, Second](t)
	testSub[Millisecond, Microsecond](t)
	testSub[Millisecond, Nanosecond](t)

	testSub[Microsecond, time.Time](t)
	testSub[Microsecond, Second](t)
	testSub[Microsecond, Millisecond](t)
	testSub[Microsecond, Microsecond](t)
	testSub[Microsecond, Nanosecond](t)

	testSub[Nanosecond, time.Time](t)
	testSub[Nanosecond, Second](t)
	testSub[Nanosecond, Millisecond](t)
	testSub[Nanosecond, Microsecond](t)
	testSub[Nanosecond, Nanosecond](t)
}

func testSub[T, U Timestamp](t *testing.T) {
	t.Run(
		fmt.Sprintf("[T %T, U %T]", T(time.Time{}), U(time.Time{})),
		func(t *testing.T) {
			now := time.Now()

			tests := []struct {
				name string
				t    T
				u    U
				want time.Duration
			}{
				{
					name: "-1s",
					t:    T(now),
					u:    U(now.Add(time.Second)),
					want: -time.Second,
				},
				{
					name: "1s",
					t:    T(now),
					u:    U(now.Add(-time.Second)),
					want: time.Second,
				},
				{
					name: "-250ms",
					t:    T(now),
					u:    U(now.Add(250 * time.Millisecond)),
					want: -250 * time.Millisecond,
				},
				{
					name: "250ms",
					t:    T(now),
					u:    U(now.Add(-250 * time.Millisecond)),
					want: 250 * time.Millisecond,
				},
			}
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := Sub(tt.t, tt.u)

					assert.Equal(t, tt.want, time.Duration(got))
				})
			}

			t1 := T(now)
			u1 := U(now.Add(-time.Second))
			got1 := Sub(t1, u1)
			assert.Equal(t, time.Second, time.Duration(got1))

			t2 := T(now)
			u2 := U(now.Add(time.Second))
			got2 := Sub(t2, u2)
			assert.Equal(t, -time.Second, time.Duration(got2))
		},
	)
}

func TestAfter(t *testing.T) {
	testAfter[time.Time, time.Time](t)
	testAfter[time.Time, Second](t)
	testAfter[time.Time, Millisecond](t)
	testAfter[time.Time, Microsecond](t)
	testAfter[time.Time, Nanosecond](t)

	testAfter[Second, time.Time](t)
	testAfter[Second, Second](t)
	testAfter[Second, Millisecond](t)
	testAfter[Second, Microsecond](t)
	testAfter[Second, Nanosecond](t)

	testAfter[Millisecond, time.Time](t)
	testAfter[Millisecond, Millisecond](t)
	testAfter[Millisecond, Second](t)
	testAfter[Millisecond, Microsecond](t)
	testAfter[Millisecond, Nanosecond](t)

	testAfter[Microsecond, time.Time](t)
	testAfter[Microsecond, Second](t)
	testAfter[Microsecond, Millisecond](t)
	testAfter[Microsecond, Microsecond](t)
	testAfter[Microsecond, Nanosecond](t)

	testAfter[Nanosecond, time.Time](t)
	testAfter[Nanosecond, Second](t)
	testAfter[Nanosecond, Millisecond](t)
	testAfter[Nanosecond, Microsecond](t)
	testAfter[Nanosecond, Nanosecond](t)
}

func testAfter[T, u Timestamp](t *testing.T) {
	t.Run(
		fmt.Sprintf("[T %T, U %T]", T(time.Time{}), u(time.Time{})),
		func(t *testing.T) {
			t1 := T(time.Now())
			o1 := u(time.Now().Add(-time.Second))
			got1 := After(t1, o1)
			assert.True(t, got1)

			t2 := T(time.Now())
			o2 := u(time.Now().Add(time.Second))
			got2 := After(t2, o2)
			assert.False(t, got2)
		},
	)
}

func TestBefore(t *testing.T) {
	testBefore[time.Time, time.Time](t)
	testBefore[time.Time, Second](t)
	testBefore[time.Time, Millisecond](t)
	testBefore[time.Time, Microsecond](t)
	testBefore[time.Time, Nanosecond](t)

	testBefore[Second, time.Time](t)
	testBefore[Second, Second](t)
	testBefore[Second, Millisecond](t)
	testBefore[Second, Microsecond](t)
	testBefore[Second, Nanosecond](t)

	testBefore[Millisecond, time.Time](t)
	testBefore[Millisecond, Millisecond](t)
	testBefore[Millisecond, Second](t)
	testBefore[Millisecond, Microsecond](t)
	testBefore[Millisecond, Nanosecond](t)

	testBefore[Microsecond, time.Time](t)
	testBefore[Microsecond, Second](t)
	testBefore[Microsecond, Millisecond](t)
	testBefore[Microsecond, Microsecond](t)
	testBefore[Microsecond, Nanosecond](t)

	testBefore[Nanosecond, time.Time](t)
	testBefore[Nanosecond, Second](t)
	testBefore[Nanosecond, Millisecond](t)
	testBefore[Nanosecond, Microsecond](t)
	testBefore[Nanosecond, Nanosecond](t)
}

func testBefore[T, U Timestamp](t *testing.T) {
	t.Run(
		fmt.Sprintf("[T %T, U %T]", T(time.Time{}), U(time.Time{})),
		func(t *testing.T) {
			t1 := T(time.Now())
			o1 := U(time.Now().Add(-time.Second))
			got1 := Before(t1, o1)
			assert.False(t, got1)

			t2 := T(time.Now())
			o2 := U(time.Now().Add(time.Second))
			got2 := Before(t2, o2)
			assert.True(t, got2)
		},
	)
}

func TestEqual(t *testing.T) {
	testEqual[time.Time, time.Time](t)
	testEqual[time.Time, Second](t)
	testEqual[time.Time, Millisecond](t)
	testEqual[time.Time, Microsecond](t)
	testEqual[time.Time, Nanosecond](t)

	testEqual[Second, time.Time](t)
	testEqual[Second, Second](t)
	testEqual[Second, Millisecond](t)
	testEqual[Second, Microsecond](t)
	testEqual[Second, Nanosecond](t)

	testEqual[Millisecond, time.Time](t)
	testEqual[Millisecond, Millisecond](t)
	testEqual[Millisecond, Second](t)
	testEqual[Millisecond, Microsecond](t)
	testEqual[Millisecond, Nanosecond](t)

	testEqual[Microsecond, time.Time](t)
	testEqual[Microsecond, Second](t)
	testEqual[Microsecond, Millisecond](t)
	testEqual[Microsecond, Microsecond](t)
	testEqual[Microsecond, Nanosecond](t)

	testEqual[Nanosecond, time.Time](t)
	testEqual[Nanosecond, Second](t)
	testEqual[Nanosecond, Millisecond](t)
	testEqual[Nanosecond, Microsecond](t)
	testEqual[Nanosecond, Nanosecond](t)
}

func testEqual[T, U Timestamp](t *testing.T) {
	t.Run(
		fmt.Sprintf("[T %T, U %T]", T(time.Time{}), U(time.Time{})),
		func(t *testing.T) {
			t1 := T(time.Now())
			o1 := U(t1)
			got1 := Equal(t1, o1)
			assert.True(t, got1)

			t2 := T(time.Now())
			o2 := U(time.Now().Add(time.Second))
			got2 := Equal(t2, o2)
			assert.False(t, got2)
		},
	)
}

func TestRound(t *testing.T) {
	testRound[time.Time, time.Duration](t)
	testRound[time.Time, dur.Duration](t)

	testRound[Second, time.Duration](t)
	testRound[Second, dur.Duration](t)

	testRound[Millisecond, time.Duration](t)
	testRound[Millisecond, dur.Duration](t)

	testRound[Microsecond, time.Duration](t)
	testRound[Microsecond, dur.Duration](t)

	testRound[Nanosecond, time.Duration](t)
	testRound[Nanosecond, dur.Duration](t)
}

func testRound[T Timestamp, D Duration](t *testing.T) {
	t.Run(
		fmt.Sprintf("[T %T, D %T]", T(time.Time{}), D(0)),
		func(t *testing.T) {
			tests := []struct {
				name string
				t    T
				d    D
				want T
			}{
				{
					name: "round up to nearest hour",
					t: T(
						time.Date(2020, 1, 1, 0, 30, 0, 0, time.UTC),
					),
					d:    D(time.Hour),
					want: T(time.Date(2020, 1, 1, 1, 0, 0, 0, time.UTC)),
				},
				{
					name: "round down to nearest hour",
					t: T(
						time.Date(2020, 1, 1, 0, 29, 0, 0, time.UTC),
					),
					d:    D(time.Hour),
					want: T(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
				{
					name: "round up to nearest minute",
					t: T(
						time.Date(2020, 1, 1, 0, 0, 30, 0, time.UTC),
					),
					d:    D(time.Minute),
					want: T(time.Date(2020, 1, 1, 0, 1, 0, 0, time.UTC)),
				},
				{
					name: "round down to nearest minute",
					t: T(
						time.Date(2020, 1, 1, 0, 0, 29, 0, time.UTC),
					),
					d:    D(time.Minute),
					want: T(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
				{
					name: "round up to nearest second",
					t: T(
						time.Date(2020, 1, 1, 0, 0, 2, 500000000, time.UTC),
					),
					d:    D(time.Second),
					want: T(time.Date(2020, 1, 1, 0, 0, 3, 0, time.UTC)),
				},
				{
					name: "round down to nearest second",
					t: T(
						time.Date(2020, 1, 1, 0, 0, 2, 499999999, time.UTC),
					),
					d:    D(time.Second),
					want: T(time.Date(2020, 1, 1, 0, 0, 2, 0, time.UTC)),
				},
				{
					name: "round up to nearest millisecond",
					t: T(
						time.Date(2020, 1, 1, 0, 0, 2, 500000, time.UTC),
					),
					d:    D(time.Millisecond),
					want: T(time.Date(2020, 1, 1, 0, 0, 2, 1000000, time.UTC)),
				},
				{
					name: "round down to nearest millisecond",
					t: T(
						time.Date(2020, 1, 1, 0, 0, 2, 499999, time.UTC),
					),
					d:    D(time.Millisecond),
					want: T(time.Date(2020, 1, 1, 0, 0, 2, 0, time.UTC)),
				},
				{
					name: "round up to nearest microsecond",
					t: T(
						time.Date(2020, 1, 1, 0, 0, 2, 500, time.UTC),
					),
					d:    D(time.Microsecond),
					want: T(time.Date(2020, 1, 1, 0, 0, 2, 1000, time.UTC)),
				},
				{
					name: "round down to nearest microsecond",
					t: T(
						time.Date(2020, 1, 1, 0, 0, 2, 499, time.UTC),
					),
					d:    D(time.Microsecond),
					want: T(time.Date(2020, 1, 1, 0, 0, 2, 0, time.UTC)),
				},
			}
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := Round(tt.t, tt.d)

					assert.Equal(t, tt.want, got)
				})
			}
		},
	)
}

func TestTruncate(t *testing.T) {
	testTruncate[time.Time, time.Duration](t)
	testTruncate[time.Time, dur.Duration](t)

	testTruncate[Second, time.Duration](t)
	testTruncate[Second, dur.Duration](t)

	testTruncate[Millisecond, time.Duration](t)
	testTruncate[Millisecond, dur.Duration](t)

	testTruncate[Microsecond, time.Duration](t)
	testTruncate[Microsecond, dur.Duration](t)

	testTruncate[Nanosecond, time.Duration](t)
	testTruncate[Nanosecond, dur.Duration](t)
}

func testTruncate[T Timestamp, D Duration](t *testing.T) {
	t.Run(
		fmt.Sprintf("[T %T, D %T]", T(time.Time{}), D(0)),
		func(t *testing.T) {
			tests := []struct {
				name string
				t    T
				d    D
				want T
			}{
				{
					name: "truncate down to nearest hour",
					t: T(
						time.Date(2020, 1, 1, 0, 29, 0, 0, time.UTC),
					),
					d:    D(time.Hour),
					want: T(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
				{
					name: "truncate down to nearest minute",
					t: T(
						time.Date(2020, 1, 1, 0, 0, 29, 0, time.UTC),
					),
					d:    D(time.Minute),
					want: T(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
				{
					name: "truncate down to nearest second",
					t: T(
						time.Date(2020, 1, 1, 0, 0, 2, 499999999, time.UTC),
					),
					d:    D(time.Second),
					want: T(time.Date(2020, 1, 1, 0, 0, 2, 0, time.UTC)),
				},
				{
					name: "truncate down to nearest millisecond",
					t: T(
						time.Date(2020, 1, 1, 0, 0, 2, 499999, time.UTC),
					),
					d:    D(time.Millisecond),
					want: T(time.Date(2020, 1, 1, 0, 0, 2, 0, time.UTC)),
				},
				{
					name: "truncate down to nearest microsecond",
					t: T(
						time.Date(2020, 1, 1, 0, 0, 2, 499, time.UTC),
					),
					d:    D(time.Microsecond),
					want: T(time.Date(2020, 1, 1, 0, 0, 2, 0, time.UTC)),
				},
			}
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := Truncate(tt.t, tt.d)

					assert.Equal(t, tt.want, got)
				})
			}
		},
	)
}
