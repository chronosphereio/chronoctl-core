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
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/browser"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/client"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/env"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/groups"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/output"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/token"
)

const (
	defaultLoginPath = "/login/cli"
)

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
		c.newWhoAmICmd(),
		c.newLogoutCmd(),
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
	store, err := token.NewChronoctlStore()
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
			return store.SetDefaultOrg(args[0])
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
			defaultOrg, err := store.GetDefaultOrg()
			if err != nil {
				// Proceed to list orgs even if we fail to get the default
				defaultOrg = ""
			}
			for _, org := range orgs {
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

type whoAmIResponse struct {
	Email string `json:"email"`
}

func (c *subcommand) newWhoAmICmd() *cobra.Command {
	authFlags := client.NewClientFlags()
	cmd := &cobra.Command{
		Use:     "whoami",
		GroupID: groups.Commands.ID,
		Short:   "Print the currently authenticated user.",
		RunE: func(cmd *cobra.Command, args []string) error {
			req, err := authFlags.NewRequest(http.MethodGet, "/auth/whoami", nil /* body */)
			if err != nil {
				return errors.WithStack(err)
			}
			client := http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return errors.WithStack(err)
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return errors.WithStack(err)
			}
			if resp.StatusCode != http.StatusOK {
				return errors.Errorf("%d: %s", resp.StatusCode, string(body))
			}
			var whoAmI whoAmIResponse
			err = json.Unmarshal(body, &whoAmI)
			if err != nil {
				return errors.WithStack(err)
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), whoAmI.Email)
			return errors.WithStack(err)
		},
	}
	authFlags.AddFlags(cmd)
	return cmd
}

func (c *subcommand) newLogoutCmd() *cobra.Command {
	authFlags := client.NewClientFlags()
	cmd := &cobra.Command{
		Use:     "logout",
		GroupID: groups.Commands.ID,
		Short:   "Logs out an authenticated chronoctl session.",
		RunE: func(cmd *cobra.Command, args []string) error {
			req, err := authFlags.NewRequest(http.MethodGet, "/auth/logout", nil /* body */)
			if err != nil {
				return errors.WithStack(err)
			}
			client := http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return errors.WithStack(err)
			}
			defer resp.Body.Close() //nolint: errcheck
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return errors.WithStack(err)
			}
			if resp.StatusCode != http.StatusOK {
				return errors.Errorf("%d: %s", resp.StatusCode, string(body))
			}
			cmd.OutOrStdout().Write(body) //nolint: errcheck
			return nil
		},
	}
	authFlags.AddFlags(cmd)
	return cmd
}
