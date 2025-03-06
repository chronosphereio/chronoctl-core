// Copyright 2024 Chronosphere Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package token provides utilities for storing and retrieving data from a local file store.
package token

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
)

const (
	// Directory that's user read-write-execute, but read-execute for all others.
	dirRWPublic os.FileMode = 0755
	// File that's user read-write, but no group/world perms.
	fileRWPrivate  os.FileMode = 0600
	defaultOrgPath             = "default-org"
)

var (
	// ErrTokenExpired is returned from Get when a token is past its expiration.
	ErrTokenExpired = errors.New("expired")

	// ErrNotExist is returned when a token doesn't exist.
	ErrNotExist = errors.New("does not exist")
)

// Token is a generic token stored with an expiration time.
type Token struct {
	Value  []byte    `json:"value"`
	Expiry time.Time `json:"expires"`
}

// Store can store and retrieve tokens from the local file system.
type Store struct {
	root string
}

// NewFileStore returns a token store that stores tokens on disk relative to a
// given root. All files will have 0600 permissions. That is, they are only
// readable by the user putting them on disk.
//
// Paths passed to storage methods must be valid filesystem paths. Paths also
// must be file paths, not directories.
func NewFileStore(root string) *Store {
	return &Store{
		root: root,
	}
}

// NewChronoctlStore creates a new token store in the user's local cache directory to store short-lived chronoctl credentials
func NewChronoctlStore() (*Store, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	chronoctlCacheDir := filepath.Join(cacheDir, "chronoctl")
	return NewFileStore(chronoctlCacheDir), nil
}

// Get retrieves a token from the local file system.
func (s *Store) Get(path string) (Token, error) {
	full := filepath.Join(s.root, path)
	f, err := os.Open(filepath.Clean(full))
	if err != nil {
		if os.IsNotExist(err) {
			return Token{}, errors.WithStack(ErrNotExist)
		}

		return Token{}, errors.WithMessagef(err, "error opening %q to decode", full)
	}
	defer f.Close() // nolint:errcheck

	var token Token
	if err := json.NewDecoder(f).Decode(&token); err != nil {
		return Token{}, errors.WithMessagef(err, "error decoding from %q", full)
	}

	if time.Now().After(token.Expiry) {
		return Token{}, errors.WithStack(ErrTokenExpired)
	}

	return token, nil
}

// Put stores a token on the local file system.
func (s *Store) Put(path string, token Token) (putErr error) {
	// Path cannot exist as a directory.
	full := filepath.Join(s.root, path)
	stat, err := os.Stat(full)
	if err == nil && stat.IsDir() {
		return errors.Errorf("path %q is a directory", full)
	}

	dir := filepath.Dir(full)
	// Dir can be 0755: the file contents are sensitive, not the tree structure.
	if err := os.MkdirAll(dir, dirRWPublic); err != nil {
		return errors.WithMessagef(err, "error ensuring path %q", dir)
	}

	// We write the contents of the new file to a temp file and then atomically
	// move it to the target. This prevents two concurrent writers from clobbering
	// the same file.
	//
	// The temporary file is in the destination folder, because rename(3p) works
	// only on the same device. There are owners that have devices (hi!) that keep
	// /tmp and ~/.config on different file systems.
	tmp, err := os.CreateTemp(dir, "droid-token")
	if err != nil {
		return errors.WithMessage(err, "could not create tmp file")
	}
	defer func() {
		if putErr != nil {
			// best-effort delete to avoid leaving temporary files
			_ = os.Remove(tmp.Name()) // nolint:errcheck
		}
	}()

	if err := os.Chmod(tmp.Name(), fileRWPrivate); err != nil {
		return errors.WithMessagef(err, "could not make %q user-only read/write", tmp.Name())
	}

	if err := json.NewEncoder(tmp).Encode(token); err != nil {
		return errors.WithMessage(err, "error encoding")
	}

	if err := tmp.Close(); err != nil {
		return errors.WithMessagef(err, "could not close %q", tmp.Name())
	}

	if err := os.Rename(tmp.Name(), full); err != nil {
		return errors.WithMessagef(err, "atomic move error %q -> %q", tmp.Name(), full)
	}

	return nil
}

// Entry provides metadata about a token stored in the token store
type Entry struct {
	Name  string
	Valid bool
}

// List returns the list of valid tokens stored in the store's root
func (s *Store) List() ([]Entry, error) {
	// Get the files in the root directory. Each file likely corresponds to a token.
	files, err := os.ReadDir(s.root)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	entries := make([]Entry, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		_, err := s.Get(file.Name())
		expired := errors.Is(err, ErrTokenExpired)
		// If the error isn't ErrTokenExpired, the file we checked is likely not a token.
		if err != nil && !expired {
			continue
		}
		// Don't list the default-org token, as it doesn't contain a session id
		if file.Name() == defaultOrgPath {
			continue
		}
		entries = append(entries, Entry{
			Name:  file.Name(),
			Valid: !expired,
		})
	}
	return entries, nil
}

// GetDefaultOrg returns the default org within the store
func (s *Store) GetDefaultOrg() (string, error) {
	org, err := s.Get(defaultOrgPath)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(org.Value), nil
}

func (s *Store) SetDefaultOrg(org string) error {
	return errors.WithStack(s.Put(defaultOrgPath, Token{
		Value: []byte(org),
		// Expire the default org if not updated for 1 year
		Expiry: time.Now().Add(time.Hour * 24 * 365),
	}))
}
