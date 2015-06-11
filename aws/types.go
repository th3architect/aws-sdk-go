package aws

import (
	"fmt"
	"io"
	"time"
)

// StringPtr returns a pointer of the string value passed in.
func StringPtr(v string) *string {
	return &v
}

// StringPtrValue returns the value of the string pointer passed in or empty
// string if the pointer is nil.
func StringPtrValue(a *string) string {
	if a != nil {
		return *a
	}
	return ""
}

// BoolPtr returns a pointer of the bool value passed in.
func BoolPtr(v bool) *bool {
	return &v
}

// BoolPtrValue returns the value of the bool pointer passed in or false if the
// pointer is nil.
func BoolPtrValue(a *bool) bool {
	if a != nil {
		return *a
	}
	return false
}

// IntPtr returns a pointer of the int value passed in.
func IntPtr(v int) *int {
	return &v
}

// IntPtrValue returns the value of the int pointer passed in or zero if the
// pointer is nil.
func IntPtrValue(a *int) int {
	if a != nil {
		return *a
	}
	return 0
}

// Int64Ptr returns a pointer of the int64 value passed in.
func Int64Ptr(v int64) *int64 {
	return &v
}

// Int64PtrValue returns the value of the int64 pointer passed in or zero if the
// pointer is nil.
func Int64PtrValue(a *int64) int64 {
	if a != nil {
		return *a
	}
	return 0
}

// Float64Ptr returns a pointer of the float64 value passed in.
func Float64Ptr(v float64) *float64 {
	return &v
}

// Float64PtrValue returns the value of the float64 pointer passed in or zero if the
// pointer is nil.
func Float64PtrValue(a *float64) float64 {
	if a != nil {
		return *a
	}
	return 0
}

// TimePtr returns a pointer of the Time value passed in.
func TimePtr(t time.Time) *time.Time {
	return &t
}

// TimePtrValue returns the value of the float64 pointer passed in or zero
// time if the pointer is nil.
func TimePtrValue(a *time.Time) time.Time {
	if a != nil {
		return *a
	}
	return time.Time{}
}

// ReadSeekCloser wraps a io.Reader returning a ReaderSeakerCloser
func ReadSeekCloser(r io.Reader) ReaderSeekerCloser {
	return ReaderSeekerCloser{r}
}

// ReaderSeekerCloser represents a reader that can also delegate io.Seeker and
// io.Closer interfaces to the underlying object if they are available.
type ReaderSeekerCloser struct {
	r io.Reader
}

// Read reads from the reader up to size of p. The number of bytes read, and
// error if it occurred will be returned.
//
// If the reader is not an io.Reader zero bytes read, and nil error will be returned.
//
// Performs the same functionality as io.Reader Read
func (r ReaderSeekerCloser) Read(p []byte) (int, error) {
	switch t := r.r.(type) {
	case io.Reader:
		return t.Read(p)
	}
	return 0, nil
}

// Seek sets the offset for the next Read to offset, interpreted according to
// whence: 0 means relative to the origin of the file, 1 means relative to the
// current offset, and 2 means relative to the end. Seek returns the new offset
// and an error, if any.
//
// If the ReaderSeekerCloser is not an io.Seeker nothing will be done.
func (r ReaderSeekerCloser) Seek(offset int64, whence int) (int64, error) {
	switch t := r.r.(type) {
	case io.Seeker:
		return t.Seek(offset, whence)
	}
	return int64(0), nil
}

// Close closes the ReaderSeekerCloser.
//
// If the ReaderSeekerCloser is not an io.Closer nothing will be done.
func (r ReaderSeekerCloser) Close() error {
	switch t := r.r.(type) {
	case io.Closer:
		return t.Close()
	}
	return nil
}

// A setState provides accessor methods for the set vs unset state.
type setState bool

// IsSet returns the set state.
func (s setState) IsSet() bool {
	return bool(s)
}

// Reset resets the set state to not set.
func (s *setState) Reset() {
	*s = false
}

// String returns the string value of the setState.
func (s setState) String() string {
	return fmt.Sprintf("%t", s)
}

// A String represents the wrapped Go string and also will maintain the state
// if the value was set or not.
type String struct {
	setState
	value string
}

// NewString returns a String initialized to the passed in value, and its state set.
func NewString(value string) String {
	return String{value: value, setState: true}
}

// Set sets the string value and updates the wrapper to be set.
func (s *String) Set(value string) {
	s.value = value
	s.setState = true
}

// Get returns the string value, or empty string if no value was set.
func (s String) Get() string {
	if !s.setState {
		return ""
	}
	return s.value
}

// String returns the string value. See String.Get()
func (s String) String() string {
	return s.Get()
}

// A Bool represents the wrapped Go bool and also will maintain the state
// if the value was set or not.
type Bool struct {
	setState
	value bool
}

// NewBool returns a Bool initialized to the passed in value, and its state set.
func NewBool(value bool) Bool {
	return Bool{value: value, setState: true}
}

// Set sets the bool value and updates the wrapper to be set.
func (b *Bool) Set(value bool) {
	b.value = value
	b.setState = true
}

// Get returns the bool value, or false if no value was set.
func (b Bool) Get() bool {
	if !b.setState {
		return false
	}
	return b.value
}

// String returns the string form of the bool value. See Bool.Get()
func (b Bool) String() string {
	return fmt.Sprintf("%t", b.Get())
}

// An Int represents the wrapped Go int and also will maintain the state
// if the value was set or not.
type Int struct {
	setState
	value int
}

// NewInt returns a Int initialized to the passed in value, and its state set.
func NewInt(value int) Int {
	return Int{value: value, setState: true}
}

// Set sets the int value and updates the wrapper to be set.
func (i *Int) Set(value int) {
	i.value = value
	i.setState = true
}

// Get returns the int value, or zero if no value was set.
func (i Int) Get() int {
	if !i.setState {
		return 0
	}
	return i.value
}

// String returns the string form of the int value. See Int.Get()
func (i Int) String() string {
	return fmt.Sprintf("%d", i.Get())
}

// An Int64 represents the wrapped Go int64 and also will maintain the state
// if the value was set or not.
type Int64 struct {
	setState
	value int64
}

// NewInt64 returns a Int64 initialized to the passed in value, and its state set.
func NewInt64(value int64) Int64 {
	return Int64{value: value, setState: true}
}

// Set sets the int64 value and updates the wrapper to be set.
func (i *Int64) Set(value int64) {
	i.value = value
	i.setState = true
}

// Get returns the int64 value, or zero if no value was set.
func (i Int64) Get() int64 {
	if !i.setState {
		return 0
	}
	return i.value
}

// String returns the string form of the int64 value. See Int64.Get()
func (i Int64) String() string {
	return fmt.Sprintf("%d", i.Get())
}

// A Float64 represents the wrapped Go float64 and also will maintain the state
// if the value was set or not.
type Float64 struct {
	setState
	value float64
}

// NewFloat64 returns a Float64 initialized to the passed in value, and its state set.
func NewFloat64(value float64) Float64 {
	return Float64{value: value, setState: true}
}

// Set sets the float64 value and updates the wrapper to be set.
func (f *Float64) Set(value float64) {
	f.value = value
	f.setState = true
}

// Get returns the float64 value, or zero if no value was set.
func (f Float64) Get() float64 {
	if !f.setState {
		return 0
	}
	return f.value
}

// String returns the string form of the float64 value. See Float64.Get()
func (f Float64) String() string {
	return fmt.Sprintf("%f", f.Get())
}
