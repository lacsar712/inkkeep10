package pathsafe

import (
	"errors"
	"path/filepath"
	"strings"
)

func JoinUnder(root, rel string) (string, error) {
	if strings.TrimSpace(rel) == "" {
		return "", errors.New("empty path")
	}
	if filepath.IsAbs(rel) {
		return "", errors.New("absolute path")
	}
	joined := filepath.Join(root, rel)
	// Reject traversal: the joined path must stay within root after
	// lexical cleaning. A relative path climbing above root via ".." escapes.
	relToRoot, err := filepath.Rel(root, joined)
	sep := string(filepath.Separator)
	if err != nil || filepath.IsAbs(relToRoot) || relToRoot == ".." ||
		strings.HasPrefix(relToRoot, ".."+sep) {
		return "", errors.New("path escape")
	}
	return joined, nil
}
