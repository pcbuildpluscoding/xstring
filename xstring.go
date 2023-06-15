package xstring

import (
	"strconv"
	"strings"
)

//=============================================================== //
// NewXString
//=============================================================== //
func NewXString(s string) *XString {
  x := XString(s)
  return &x
}

//=============================================================== //
// XString
//=============================================================== //
type XString string

// -------------------------------------------------------------- //
// Append
// ---------------------------------------------------------------//
func (s_ *XString) Append(text string) *XString {
  *s_ = XString(s_.String() + text)
  return s_
}

// -------------------------------------------------------------- //
// Contains
// ---------------------------------------------------------------//
func (s_ XString) Contains(pattn string) bool {
  return strings.Contains(string(s_), pattn)
}

// -------------------------------------------------------------- //
// Contains
// ---------------------------------------------------------------//
func (s_ XString) Copy() XString {
  return XString(string(s_))
}

// -------------------------------------------------------------- //
// HasPrefix
// ---------------------------------------------------------------//
func (s_ XString) HasPrefix(pattn string) bool {
  return strings.HasPrefix(string(s_), pattn)
}

// -------------------------------------------------------------- //
// HasSuffix
// ---------------------------------------------------------------//
func (s_ XString) HasSuffix(pattn string) bool {
  return strings.HasSuffix(string(s_), pattn)
}

// -------------------------------------------------------------- //
// Length
// ---------------------------------------------------------------//
func (s_ XString) Length() int {
  return len(string(s_))
}

// -------------------------------------------------------------- //
// Prepend
// ---------------------------------------------------------------//
func (s_ *XString) Prepend(text string) *XString {
  *s_ = XString(text + s_.String())
  return s_
}

// -------------------------------------------------------------- //
// Replace
// ---------------------------------------------------------------//
func (s_ *XString) Replace(old, new string, n int) *XString {
  *s_ = XString(strings.Replace(s_.String(), old, new, n))
  return s_
}

// -------------------------------------------------------------- //
// ReplaceAll
// ---------------------------------------------------------------//
func (s_ *XString) ReplaceAll(old, new string) *XString {
  *s_ = XString(strings.ReplaceAll(s_.String(), old, new))
  return s_
}

// -------------------------------------------------------------- //
// Set
// ---------------------------------------------------------------//
func (s_ *XString) Set(value string) *XString {
  *s_ = XString(value)
  return s_
}

// -------------------------------------------------------------- //
// SplitInTwo
// ---------------------------------------------------------------//
func (s_ XString) SplitInTwo(delim string) (XString, XString) {
  x := strings.SplitN(s_.String(), delim, 2)
  if len(x) == 1 {
    return XString(x[0]), ""
  }
  return XString(x[0]), XString(x[1])
}

// -------------------------------------------------------------- //
// SplitN
// ---------------------------------------------------------------//
func (s_ XString) SplitN(delim string, n int) []string {
  return strings.SplitN(s_.String(), delim, n)
}

// -------------------------------------------------------------- //
// SplitNKeepOne
// ---------------------------------------------------------------//
func (s_ XString) SplitNKeepOne(delim string, n, keep int) *XString {
  x := strings.SplitN(string(s_), delim, n)
  var y XString
  last := len(x)-1
  if last < keep {
    y = XString(x[last])
  } else {
    y = XString(x[keep])
  }
  return &y
}

// -------------------------------------------------------------- //
// SplitInThree
// ---------------------------------------------------------------//
func (s_ XString) SplitInThree(delim string) (XString, XString, XString) {
  x := strings.SplitN(string(s_), delim, 3)
  switch len(x) {
  case 1:
    return XString(x[0]), "", ""
  case 2:
    return XString(x[0]), XString(x[1]), ""
  }
  return XString(x[0]), XString(x[1]), XString(x[2])
}

// -------------------------------------------------------------- //
// SplitInFour
// ---------------------------------------------------------------//
func (s_ XString) SplitInFour(delim string) (XString, XString, XString, XString) {
  x := strings.SplitN(string(s_), delim, 4)
  switch len(x) {
  case 1:
    return XString(x[0]), "", "", ""
  case 2:
    return XString(x[0]), XString(x[1]), "", ""
  case 3:
    return XString(x[0]), XString(x[1]), XString(x[2]), ""
  }
  return XString(x[0]), XString(x[1]), XString(x[2]), XString(x[3])
}

// -------------------------------------------------------------- //
// String
// ---------------------------------------------------------------//
func (s_ XString) String() string {
  return string(s_)
}

// -------------------------------------------------------------- //
// ToInt
// ---------------------------------------------------------------//
func (s_ XString) ToInt() (int, error) {
  return strconv.Atoi(string(s_))
}

// -------------------------------------------------------------- //
// Trim
// ---------------------------------------------------------------//
func (s_ XString) Trim() string {
  return strings.TrimSpace(string(s_))
}

// -------------------------------------------------------------- //
// XTrim
// ---------------------------------------------------------------//
func (s_ XString) XTrim() *XString {
  y := XString(strings.TrimSpace(string(s_)))
  return &y
}

// -------------------------------------------------------------- //
// Ptr
// ---------------------------------------------------------------//
func (s_ XString) Ptr() *XString {
  return &s_
}
