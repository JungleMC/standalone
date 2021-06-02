package nbt

import (
	"bufio"
	"bytes"
	"github.com/junglemc/JungleTree/nbt/test"
	"os"
	"path/filepath"
	"testing"
)

func TestMarshalCompoundMap(t *testing.T) {
	tests := []struct {
		name          string
		inputTagName  string
		input         map[string]interface{}
		expected      []byte
		expectedError bool
	}{
		{
			name:         "unnamed root compound tag",
			inputTagName: "",
			input: map[string]interface{}{
				"ByteTag":   byte(0xFF),
				"StringTag": "hello, world",
			},
			expected:      test.UnnamedRootCompoundBytes,
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := Marshal(tt.inputTagName, tt.input)

			path, _ := os.MkdirTemp("", "nbt")
			f, _ := os.Create(filepath.Join(path, "bigtest_go.nbt"))
			w := bufio.NewWriter(f)
			_, _ = w.Write(data)
			_ = w.Flush()

			if !bytes.Equal(data, tt.expected) {
				t.Errorf("got:\n[% 2x]\nwant:\n[% 2x]", data, tt.expected)
			}
		})
	}
}

func TestMarshalCompoundStruct(t *testing.T) {
	tests := []struct {
		name    string
		tagName string
		tag     interface{}
		want    []byte
		wantErr bool
	}{
		{
			name:    "unnamed root compound tag",
			tagName: "",
			tag: test.UnnamedRootCompound{
				ByteTag:   0xFF,
				StringTag: "hello, world",
			},
			want:    test.UnnamedRootCompoundBytes,
			wantErr: false,
		},
		{
			name:    "bananrama",
			tagName: "hello world",
			tag:     test.BananramaStruct,
			want:    test.BananramaBytes,
			wantErr: false,
		},
		{
			name:    "bigtest",
			tagName: "Level",
			tag: test.BigTest{
				LongTest:   9223372036854775807,
				ShortTest:  32767,
				StringTest: "HELLO WORLD THIS IS A TEST STRING \xc3\x85\xc3\x84\xc3\x96!",
				FloatTest:  0.49823147058486938,
				IntTest:    2147483647,
				NCT: test.BigTestNCT{
					Egg: test.BigTestNameAndFloat32{
						Name:  "Eggbert",
						Value: 0.5,
					},
					Ham: test.BigTestNameAndFloat32{
						Name:  "Hampus",
						Value: 0.75,
					},
				},
				ListTest: []int64{11, 12, 13, 14, 15},
				ListTest2: [2]test.BigTestCompound{
					{
						Name:      "Compound tag #0",
						CreatedOn: 1264099775885,
					},
					{
						Name:      "Compound tag #1",
						CreatedOn: 1264099775885,
					},
				},
				ByteTest:      127,
				ByteArrayTest: test.BigTestByteArray(),
				DoubleTest:    0.49312871321823148,
			},
			want:    test.BigTestBytes,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := Marshal(tt.tagName, tt.tag)

			path, _ := os.MkdirTemp("", "nbt")
			f, _ := os.Create(filepath.Join(path, "bigtest_go.nbt"))
			w := bufio.NewWriter(f)
			_, _ = w.Write(data)
			_ = w.Flush()

			if !bytes.Equal(data, tt.want) {
				t.Errorf("got:\n[% 2x]\nwant:\n[% 2x]", data, tt.want)
			}
		})
	}
}
