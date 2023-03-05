package internal

import (
	"reflect"
	"testing"
	"time"
)

func TestNewDaySlot(t *testing.T) {
	y := 2023
	m := 3
	d := 4

	cd := time.Date(y, time.Month(m), d, 12, 0, 0, 0, time.UTC)

	type args struct {
		inSlot  time.Time
		outSlot time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    DaySlot
		wantDur int
	}{
		{
			name: "check in is earlier than out",
			args: args{
				inSlot:  cd.Add(time.Hour * 24),
				outSlot: cd,
			},
			want: DaySlot{
				inSlot:  cd,
				outSlot: cd.Add(time.Hour * 24),
			},
			wantDur: 1,
		},
		{
			name: "check duration",
			args: args{
				inSlot:  cd.Add(time.Hour * 24 * 6),
				outSlot: cd,
			},
			want: DaySlot{
				inSlot:  cd,
				outSlot: cd.Add(time.Hour * 24 * 6),
			},
			wantDur: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDaySlot(tt.args.inSlot, tt.args.outSlot)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDaySlot() = %v, want %v", got, tt.want)
			}
			if got.Duration() != tt.wantDur {
				t.Errorf("Duration: %d. expected: %d", got.Duration(), tt.wantDur)
			}

		})
	}
}

func TestDaySlot_Overlaps(t *testing.T) {
	type fields struct {
		inSlot  time.Time
		outSlot time.Time
	}
	type args struct {
		other DaySlot
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "slots do not overlap",
			fields: fields{
				inSlot:  time.Date(2016, time.Month(2), 1, 12, 0, 0, 0, time.UTC),
				outSlot: time.Date(2016, time.Month(2), 3, 12, 0, 0, 0, time.UTC),
			},
			args: args{
				other: DaySlot{
					inSlot:  time.Date(2016, time.Month(2), 3, 12, 0, 0, 0, time.UTC),
					outSlot: time.Date(2016, time.Month(2), 4, 12, 0, 0, 0, time.UTC),
				},
			},
			want: false,
		},
		{
			name: "slots overlap A include B",
			fields: fields{
				inSlot:  time.Date(2016, time.Month(2), 1, 12, 0, 0, 0, time.UTC),
				outSlot: time.Date(2016, time.Month(2), 5, 12, 0, 0, 0, time.UTC),
			},
			args: args{
				other: DaySlot{
					inSlot:  time.Date(2016, time.Month(2), 2, 12, 0, 0, 0, time.UTC),
					outSlot: time.Date(2016, time.Month(2), 4, 12, 0, 0, 0, time.UTC),
				},
			},
			want: true,
		},
		{
			name: "slots overlap B include A",
			fields: fields{
				inSlot:  time.Date(2016, time.Month(2), 5, 12, 0, 0, 0, time.UTC),
				outSlot: time.Date(2016, time.Month(2), 6, 12, 0, 0, 0, time.UTC),
			},
			args: args{
				other: DaySlot{
					inSlot:  time.Date(2016, time.Month(2), 2, 12, 0, 0, 0, time.UTC),
					outSlot: time.Date(2016, time.Month(2), 14, 12, 0, 0, 0, time.UTC),
				},
			},
			want: true,
		},
		{
			name: "slots overlap A starts inside B",
			fields: fields{
				inSlot:  time.Date(2016, time.Month(2), 5, 4, 0, 0, 0, time.UTC),
				outSlot: time.Date(2016, time.Month(2), 6, 12, 0, 0, 0, time.UTC),
			},
			args: args{
				other: DaySlot{
					inSlot:  time.Date(2016, time.Month(2), 2, 12, 0, 0, 0, time.UTC),
					outSlot: time.Date(2016, time.Month(2), 5, 12, 0, 0, 0, time.UTC),
				},
			},
			want: true,
		},
		{
			name: "slots overlap B starts inside A",
			fields: fields{
				inSlot:  time.Date(2016, time.Month(2), 1, 4, 0, 0, 0, time.UTC),
				outSlot: time.Date(2016, time.Month(2), 6, 12, 0, 0, 0, time.UTC),
			},
			args: args{
				other: DaySlot{
					inSlot:  time.Date(2016, time.Month(2), 2, 12, 0, 0, 0, time.UTC),
					outSlot: time.Date(2016, time.Month(2), 15, 12, 0, 0, 0, time.UTC),
				},
			},
			want: true,
		},
		{
			name: "slots overlap A ends inside B",
			fields: fields{
				inSlot:  time.Date(2016, time.Month(2), 2, 4, 0, 0, 0, time.UTC),
				outSlot: time.Date(2016, time.Month(2), 6, 12, 0, 0, 0, time.UTC),
			},
			args: args{
				other: DaySlot{
					inSlot:  time.Date(2016, time.Month(2), 5, 12, 0, 0, 0, time.UTC),
					outSlot: time.Date(2016, time.Month(2), 7, 12, 0, 0, 0, time.UTC),
				},
			},
			want: true,
		},
		{
			name: "slots overlap B ends inside A",
			fields: fields{
				inSlot:  time.Date(2016, time.Month(2), 2, 4, 0, 0, 0, time.UTC),
				outSlot: time.Date(2016, time.Month(2), 6, 12, 0, 0, 0, time.UTC),
			},
			args: args{
				other: DaySlot{
					inSlot:  time.Date(2016, time.Month(2), 3, 12, 0, 0, 0, time.UTC),
					outSlot: time.Date(2016, time.Month(2), 15, 12, 0, 0, 0, time.UTC),
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slot := &DaySlot{
				inSlot:  tt.fields.inSlot,
				outSlot: tt.fields.outSlot,
			}
			if got := slot.Overlaps(&tt.args.other); got != tt.want {
				t.Errorf("DaySlot.Overlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
