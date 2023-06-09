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
  return strings.Contains(s_.String(), pattn)
}

// -------------------------------------------------------------- //
// Contains
// ---------------------------------------------------------------//
func (s_ XString) Copy() XString {
  return XString(s_.String())
}

// -------------------------------------------------------------- //
// HasPrefix
// ---------------------------------------------------------------//
func (s_ XString) HasPrefix(pattn string) bool {
  return strings.HasPrefix(s_.String(), pattn)
}

// -------------------------------------------------------------- //
// HasSuffix
// ---------------------------------------------------------------//
func (s_ XString) HasSuffix(pattn string) bool {
  return strings.HasSuffix(s_.String(), pattn)
}

// -------------------------------------------------------------- //
// Length
// ---------------------------------------------------------------//
func (s_ XString) Length() int {
  return len(s_.String())
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
func (s_ *XString) Replace(old, new string, n int) XString {
  *s_ = XString(strings.Replace(s_.String(), old, new, n))
  return *s_
}

// -------------------------------------------------------------- //
// ReplaceAll
// ---------------------------------------------------------------//
func (s_ *XString) ReplaceAll(old, new string) XString {
  *s_ = XString(strings.ReplaceAll(s_.String(), old, new))
  return *s_
}

// -------------------------------------------------------------- //
// Set
// ---------------------------------------------------------------//
func (s_ *XString) Set(value string) XString {
  *s_ = XString(value)
  return *s_
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
func (s_ XString) SplitNKeepOne(delim string, n, keep int) XString {
  s := string(s_)
  x := strings.SplitN(s, delim, n)
  if (len(x)-1) < keep {
    return XString("")
  }
  return XString(x[keep])
}

// -------------------------------------------------------------- //
// SplitInThree
// ---------------------------------------------------------------//
func (s_ XString) SplitInThree(delim string) (XString, XString, XString) {
  s := string(s_)
  x := strings.SplitN(s, delim, 3)
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
  s := string(s_)
  x := strings.SplitN(s, delim, 4)
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
  return strconv.Atoi(s_.String())
}

// -------------------------------------------------------------- //
// Trim
// ---------------------------------------------------------------//
func (s_ XString) Trim() string {
  return strings.TrimSpace(s_.String())
}

// -------------------------------------------------------------- //
// XTrim
// ---------------------------------------------------------------//
func (s_ XString) XTrim() XString {
  return XString(strings.TrimSpace(s_.String()))
}

// -------------------------------------------------------------- //
// Ptr
// ---------------------------------------------------------------//
func (s_ XString) Ptr() *XString {
  return &s_
}
