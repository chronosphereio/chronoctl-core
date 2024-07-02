package token

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileStore(t *testing.T) {
	t.Parallel()
	root, err := os.MkdirTemp("", "chronoctltest")
	require.NoError(t, err)
	defer os.RemoveAll(root) // nolint:errcheck,gosec

	store, err := NewFileStore(root)
	assert.NoError(t, err)

	_, err = store.Get("foo/bar")
	assert.True(t, errors.Is(err, ErrNotExist))

	expiry := time.Date(2020, 11, 05, 14, 30, 24, 0, time.UTC)
	err = store.Put("foo/bar", Token{
		Value:  []byte("hello"),
		Expiry: expiry,
	})

	assert.NoError(t, err)
	path := filepath.Join(root, "foo", "bar")
	stat, err := os.Stat(path)
	assert.NoError(t, err)
	assert.Equal(t, os.FileMode(0600), stat.Mode().Perm())

	f, err := os.Open(path) // nolint:gosec
	assert.NoError(t, err)
	data, err := io.ReadAll(f)
	assert.NoError(t, err)

	assert.Equal(t, `{"value":"aGVsbG8=","expires":"2020-11-05T14:30:24Z"}`, strings.TrimSpace(string(data)))

	// Fail to get expired token
	_, err = store.Get("foo/bar")
	assert.ErrorIs(t, err, ErrTokenExpired)

	// Get valid token
	tokenVal := []byte("hello")
	err = store.Put("foo/baz", Token{
		Value:  tokenVal,
		Expiry: time.Now().Add(time.Hour),
	})
	require.NoError(t, err)
	token, err := store.Get("foo/baz")
	require.NoError(t, err)
	require.Equal(t, tokenVal, token.Value)
}

func TestList(t *testing.T) {
	tests := []struct {
		name        string
		tokens      map[string]Token
		wantEntries []Entry
	}{
		{
			name: "valid token in store",
			tokens: map[string]Token{
				"token1": {
					Value:  []byte("my-token"),
					Expiry: time.Now().Add(time.Hour),
				},
			},
			wantEntries: []Entry{
				{
					Name:  "token1",
					Valid: true,
				},
			},
		},
		{
			name: "invalid token in store",
			tokens: map[string]Token{
				"token1": {
					Value:  []byte("my-token"),
					Expiry: time.Now().Add(-time.Hour),
				},
			},
			wantEntries: []Entry{
				{
					Name:  "token1",
					Valid: false,
				},
			},
		},
		{
			name: "multiple tokens in store",
			tokens: map[string]Token{
				"token1": {
					Value:  []byte("my-token"),
					Expiry: time.Now().Add(-time.Hour),
				},
				"token2": {
					Value:  []byte("my-token"),
					Expiry: time.Now().Add(time.Hour),
				},
			},
			wantEntries: []Entry{
				{
					Name:  "token1",
					Valid: false,
				},
				{
					Name:  "token2",
					Valid: true,
				},
			},
		},
		{
			name:        "no tokens in store",
			tokens:      map[string]Token{},
			wantEntries: []Entry{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store, err := NewFileStore(t.TempDir())
			require.NoError(t, err)

			for name, token := range tt.tokens {
				require.NoError(t, store.Put(name, token))
			}

			gotEntries, err := store.List()
			require.NoError(t, err)
			require.EqualValues(t, tt.wantEntries, gotEntries)
		})
	}
}

func TestList_NonTokenFileNotReturned(t *testing.T) {
	tempDir := t.TempDir()
	store, err := NewFileStore(tempDir)
	require.NoError(t, err)

	// Create a file in the same directory that the token store uses and ensure it's not returned as a token entry
	fd, err := os.Create(path.Join(tempDir, "randomFile"))
	require.NoError(t, err)
	defer fd.Close() //nolint:errcheck
	_, err = fd.Write([]byte("not a token"))
	require.NoError(t, err)

	require.NoError(t, store.Put("realToken", Token{
		Value:  []byte("my-token"),
		Expiry: time.Now().Add(time.Hour),
	}))

	gotEntries, err := store.List()
	require.NoError(t, err)
	require.EqualValues(t, []Entry{{Name: "realToken", Valid: true}}, gotEntries)
}
