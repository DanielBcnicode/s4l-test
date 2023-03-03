package internal

import "time"

type DaySlot struct {
	inSlot  time.Time
	outSlot time.Time
}

// NewDaySlot creates a new DaySlot with the in and out date.
// The constructor will change the hour at 12:00 UTC
func NewDaySlot(inSlot time.Time, outSlot time.Time) DaySlot {
	in := time.Date(
		inSlot.Year(),
		inSlot.Month(),
		inSlot.Day(),
		12, 0, 0, 0, time.UTC)
	out := time.Date(
		outSlot.Year(),
		outSlot.Month(),
		outSlot.Day(),
		12, 0, 0, 0, time.UTC)

	// If out is earlier than in we change the values
	if in.UnixNano() > out.UnixNano() {
		in, out = out, in
	}

	return DaySlot{inSlot: in, outSlot: out}
}

// Duration represents the duration of the slot in whole days, the value returned is an integer
func (s *DaySlot) Duration() int {

	return int(s.outSlot.Sub(s.inSlot).Hours() / 24)

}

// Overlaps return true if the time windows between slot and other overlaps
// op1 -------->                1a###############1b
// op2 -------->            2a#########2b
func (s *DaySlot) Overlaps(other DaySlot) bool {
	op1a := s.inSlot.UnixNano()
	op1b := s.outSlot.UnixNano()
	op2a := other.inSlot.UnixNano()
	op2b := other.outSlot.UnixNano()

	if op1a > +op2a && op1a < op2b { // op1a is inside op2
		return true
	}
	if op1b > +op2a && op1b < op2b { // op1b is inside op2
		return true
	}
	if op2a > +op1a && op2a < op1b { // op2a is inside op1
		return true
	}
	if op2b > +op1a && op2b < op1b { // op2b is inside op1
		return true
	}

	return false
}

func (s *DaySlot) StartDate() time.Time {
	return s.inSlot
}
