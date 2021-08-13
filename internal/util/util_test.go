package util

import "testing"

func TestFormatHex(t *testing.T) {
	testCases := []struct {
		input  []byte
		output string
	}{
		{[]byte{1, 2, 3, 4}, "01:02:03:04"},
		{[]byte{33, 79, 21, 35, 77, 72, 76, 97, 27, 80, 70, 67, 52, 73, 38, 9, 51, 42, 92, 1, 5, 93, 25, 62, 47, 43, 13, 66, 98, 30, 11, 53, 100, 44, 60, 16, 56, 85, 12, 82, 90, 57, 74, 55, 22, 59, 7, 15, 32, 86, 6, 96, 83, 39, 34, 65, 3, 31, 49}, "21:4F:15:23:4D:48:4C:61:1B:50:46:43:34:49:26:09:33:2A:5C:01:05:5D:19:3E:2F:2B:0D:42:62:1E:0B:35:64:2C:3C:10:38:55:0C:52:5A:39:4A:37:16:3B:07:0F:20:56:06:60:53:27:22:41:03:1F:31"},
	}
	for _, tc := range testCases {
		if out := FormatHex(tc.input); out != tc.output {
			t.Errorf("expected %s got %s", tc.output, out)
		}
	}
}
