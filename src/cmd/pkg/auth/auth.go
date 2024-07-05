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

// Package auth provides utilities for authenticating a chronoctl session.
package auth

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/browser"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/env"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/groups"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/output"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/token"
)

const (
	defaultLoginPath = "/login/cli"
	defaultOrgPath   = "default-org"
)

var (
	errMustIncludeOrgName = errors.New("A single organization name is required")
)

// NewChronoctlStore creates a new token store in the user's local cache directory to store short-lived chronoctl credentials
func NewChronoctlStore() (*token.Store, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	chronoctlCacheDir := filepath.Join(cacheDir, "chronoctl")
	store := token.NewFileStore(chronoctlCacheDir)
	return store, nil
}

// NewCommand returns a command for Chronosphere authentication
func NewCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "auth",
		Short: "All commands related to authenticating the chronoctl session.",
	}

	root.AddGroup(groups.Commands)
	c := subcommand{}
	root.AddCommand(
		c.newAuthLoginCmd(browser.OpenURL),
		c.newSetDefaultOrgCmd(),
		c.newPrintAccessTokenCmd(),
		c.newListCmd(),
	)
	return root
}

type subcommand struct {
	store *token.Store
}

// getStore returns a token store contained in the subcommand if one is defined, otherwise returns a NewChronoctlStore
func (c *subcommand) getStore() (*token.Store, error) {
	if c.store != nil {
		return c.store, nil
	}
	store, err := NewChronoctlStore()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return store, nil
}

type loginOpts struct {
	orgName           string
	loginPath         string
	skipSetDefaultOrg bool
}

func (o *loginOpts) addFlags(cmd *cobra.Command) {
	const loginPathFlag = "login-path"
	cmd.Flags().StringVar(&o.orgName, "org-name", "", "The name of your team's Chronosphere organization. Defaults to "+env.ChronosphereOrgNameKey+" environment variable.")
	cmd.Flags().StringVar(&o.loginPath, loginPathFlag, defaultLoginPath, "the path to the org's chronoctl login page")
	cmd.Flags().MarkHidden(loginPathFlag) //nolint:errcheck
	cmd.Flags().BoolVar(&o.skipSetDefaultOrg, "skip-set-default-org", false, "When true, will not the current org as default for future commands")
}

// browserOpenFunc is used to set the function that will be run to open the browser. Useful for testing with a fake browser.
type browserOpenFunc func(URL string) error

func (c *subcommand) newAuthLoginCmd(openFunc browserOpenFunc) *cobra.Command {
	opts := &loginOpts{}
	cmd := &cobra.Command{
		Use:     "login",
		GroupID: groups.Commands.ID,
		Short:   "Authenticate the chronoctl session as your Chronosphere user.",
		RunE: func(command *cobra.Command, strings []string) error {
			store, err := c.getStore()
			if err != nil {
				return errors.WithStack(err)
			}
			ls, err := newLoginServer(store, openFunc, opts)
			if err != nil {
				return errors.WithStack(err)
			}
			return ls.login(command.Context(), command.InOrStdin(), command.OutOrStdout(), command.ErrOrStderr())
		},
	}

	opts.addFlags(cmd)
	return cmd
}

func (c *subcommand) newSetDefaultOrgCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "set-default-org ORG_NAME",
		GroupID: groups.Commands.ID,
		Short:   "Sets the default Chronosphere organization for future commands.",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			store, err := c.getStore()
			if err != nil {
				return errors.WithStack(err)
			}
			return setDefaultOrg(store, args[0])
		},
	}
	return cmd
}

func (c *subcommand) newPrintAccessTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "print-session-id ORG_NAME",
		GroupID: groups.Commands.ID,
		Short:   "Print the stored session id for a Chronosphere organization.",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			store, err := c.getStore()
			if err != nil {
				return errors.WithStack(err)
			}
			sessionIDToken, err := store.Get(args[0])
			if err != nil {
				return errors.WithStack(err)
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), string(sessionIDToken.Value))
			return errors.WithStack(err)
		},
	}
	return cmd
}

type listEntry struct {
	Organization string
	Valid        bool
	Default      bool
}

func (c *subcommand) newListCmd() *cobra.Command {
	outputFlags := output.NewFlags()
	cmd := &cobra.Command{
		Use:     "list",
		GroupID: groups.Commands.ID,
		Short:   "List all authenticated Chronosphere organizations.",
		RunE: func(cmd *cobra.Command, args []string) error {
			defer outputFlags.Close(cmd.OutOrStdout()) //nolint:errcheck
			store, err := c.getStore()
			if err != nil {
				return errors.WithStack(err)
			}
			orgs, err := store.List()
			if err != nil {
				return errors.WithStack(err)
			}
			defaultOrg, err := GetDefaultOrg(store)
			if err != nil {
				// Proceed to list orgs even if we fail to get the default
				defaultOrg = ""
			}
			for _, org := range orgs {
				// Don't list the default-org token, as it doesn't contain a session id
				if org.Name == defaultOrgPath {
					continue
				}
				e := listEntry{
					Organization: org.Name,
					Valid:        org.Valid,
					Default:      org.Name == defaultOrg,
				}
				if err := outputFlags.WriteObject(e, cmd.OutOrStdout()); err != nil {
					return errors.WithStack(err)
				}
			}
			return nil
		},
	}
	outputFlags.AddFlags(cmd)
	return cmd
}

func setDefaultOrg(store *token.Store, org string) error {
	return errors.WithStack(store.Put(defaultOrgPath, token.Token{
		Value: []byte(org),
		// Expire the default org if not updated for 1 year
		Expiry: time.Now().Add(time.Hour * 24 * 365),
	}))
}

// GetDefaultOrg returns the default org within the store
func GetDefaultOrg(store *token.Store) (string, error) {
	org, err := store.Get(defaultOrgPath)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(org.Value), nil
}
