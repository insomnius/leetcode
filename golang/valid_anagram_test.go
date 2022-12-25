package leetcode

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s, t string
		want bool
	}{
		{
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			s:    "aacc",
			t:    "ccac",
			want: false,
		},
		{
			s:    "a",
			t:    "ab",
			want: false,
		},
		{
			s:    "a",
			t:    "a",
			want: true,
		},
		{
			s:    "aa",
			t:    "a",
			want: false,
		},
		{
			s:    "a",
			t:    "aa",
			want: false,
		},
		{
			s:    "a",
			t:    "a",
			want: true,
		},
		{
			s:    "ab",
			t:    "a",
			want: false,
		},
		{
			s:    "a",
			t:    "ab",
			want: false,
		},
		{
			s:    "a",
			t:    "a",
			want: true,
		},
		{
			s:    "ab",
			t:    "a",
			want: false,
		},
		{
			s:    "a",
			t:    "ab",
			want: false,
		},
		{
			s:    "a",
			t:    "a",
			want: true,
		},
		{
			s:    "ab",
			t:    "a",
			want: false,
		},
		{
			s:    "a",
			t:    "ab",
			want: false,
		},
		{
			s:    "a",
			t:    "a",
			want: true,
		},
		{
			s:    "ab",
			t:    "a",
			want: false,
		},
		{
			s:    "a",
			t:    "ab",
			want: false,
		},
		{
			s:    "aa",
			t:    "bb",
			want: false,
		},
	}

	for _, test := range tests {
		if got := isAnagram(test.s, test.t); got != test.want {
			t.Errorf("isAnagram(%q, %q) = %v; want %v", test.s, test.t, got, test.want)
		}
	}
}

func TestIsAnagramAlt(t *testing.T) {
	tests := []struct {
		s, t string
		want bool
	}{
		{
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			s:    "aacc",
			t:    "ccac",
			want: false,
		},
		{
			s:    "a",
			t:    "ab",
			want: false,
		},
		{
			s:    "a",
			t:    "a",
			want: true,
		},
		{
			s:    "aa",
			t:    "a",
			want: false,
		},
		{
			s:    "a",
			t:    "aa",
			want: false,
		},
		{
			s:    "a",
			t:    "a",
			want: true,
		},
		{
			s:    "ab",
			t:    "a",
			want: false,
		},
		{
			s:    "a",
			t:    "ab",
			want: false,
		},
		{
			s:    "a",
			t:    "a",
			want: true,
		},
		{
			s:    "ab",
			t:    "a",
			want: false,
		},
		{
			s:    "a",
			t:    "ab",
			want: false,
		},
		{
			s:    "a",
			t:    "a",
			want: true,
		},
		{
			s:    "ab",
			t:    "a",
			want: false,
		},
		{
			s:    "a",
			t:    "ab",
			want: false,
		},
		{
			s:    "a",
			t:    "a",
			want: true,
		},
		{
			s:    "ab",
			t:    "a",
			want: false,
		},
		{
			s:    "a",
			t:    "ab",
			want: false,
		},
		{
			s:    "aa",
			t:    "bb",
			want: false,
		},
	}

	for _, test := range tests {
		if got := isAnagramAlt1(test.s, test.t); got != test.want {
			t.Errorf("isAnagramAlt1(%q, %q) = %v; want %v", test.s, test.t, got, test.want)
		}
	}
}
