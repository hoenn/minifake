package mimic

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGolden(t *testing.T) {
	entries, err := os.ReadDir("testdata")
	require.NoError(t, err)

	var inputs []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".go") && !strings.HasSuffix(e.Name(), ".go.golden") {
			inputs = append(inputs, e.Name())
		}
	}

	for _, input := range inputs {
		fileName := strings.TrimSuffix(input, ".go")
		t.Run(fileName, func(t *testing.T) {
			inputPath := filepath.Join("testdata", fileName)
			goldenPath := inputPath + "_fake.go.golden"

			src, err := os.ReadFile(inputPath + ".go")
			require.NoError(t, err)

			interfaceNames := extractInterfaceNames(t, string(src))

			got, err := ParseAndStubFromSrc(interfaceNames, string(src), true)
			require.NoError(t, err)

			expected, err := os.ReadFile(goldenPath)
			require.NoError(t, err, "golden file missing: %s", goldenPath)
			require.Equal(t, string(expected), string(got))
		})
	}
}
func extractInterfaceNames(t *testing.T, src string) []string {
	// Expected format: //go:generate go run github.com/hoenn/mimic ./file.go Name1,Name2
	t.Helper()
	for _, line := range strings.Split(src, "\n") {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "//go:generate go run github.com/hoenn/mimic") {
			continue
		}
		fields := strings.Fields(line)
		// //go:generate go run github.com/hoenn/mimic ./file.go Name1,Name2
		// fields[0]       [1]  [2]  [3]                          [4]      [5]
		require.GreaterOrEqual(t, len(fields), 6,
			"go:generate directive must include file path and interface names")
		names := strings.Split(fields[5], ",")
		for i := range names {
			names[i] = strings.TrimSpace(names[i])
		}
		return names
	}
	t.Fatal("no //go:generate go run github.com/hoenn/mimic directive found")
	return nil
}

func TestCrossPackageEmbedError(t *testing.T) {
	src := `package test

import "io"

type Outer interface {
	io.Reader
}
`
	_, err := ParseAndStubFromSrc([]string{"Outer"}, src, true)
	require.Error(t, err)
	require.Contains(t, err.Error(), "cross-package embed")
}
