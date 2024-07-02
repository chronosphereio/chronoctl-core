package auth

import (
	"cmp"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/browser"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/chronosphereio/chronoctl-core/src/cmd"
	"github.com/chronosphereio/chronoctl-core/src/cmd/pkg/client"
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
	store, err := token.NewFileStore(chronoctlCacheDir)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return store, nil
}

// NewCommand returns a command for Chronosphere authentication
func NewCommand(store *token.Store) *cobra.Command {
	root := &cobra.Command{
		Use:   "auth",
		Short: "All commands related to authenticating the chronoctl session.",
	}

	root.AddGroup(groups.Commands)
	root.AddCommand(
		newAuthLoginCmd(store, browser.OpenURL),
		newSetDefaultOrgCmd(store),
		newPrintAccessTokenCmd(store),
		newListCmd(store),
	)
	return root
}

type loginOpts struct {
	orgName           string
	loginPath         string
	skipSetDefaultOrg bool
}

func (o *loginOpts) registerFlags(cmd *cobra.Command) {
	const loginPathFlag = "login-path"
	cmd.Flags().StringVar(&o.orgName, "org-name", "", "The name of your team's Chronosphere organization. Defaults to "+client.ChronosphereOrgNameKey+" environment variable.")
	cmd.Flags().StringVar(&o.loginPath, loginPathFlag, defaultLoginPath, "the path to the org's chronoctl login page")
	cmd.Flags().MarkHidden(loginPathFlag) //nolint:errcheck
	cmd.Flags().BoolVar(&o.skipSetDefaultOrg, "skip-set-default-org", false, "When true, will not the current org as default for future commands")
}

// browserOpenFunc is used to set the function that will be run to open the browser. Useful for testing with a fake browser.
type browserOpenFunc func(URL string) error

func newAuthLoginCmd(store *token.Store, openFunc browserOpenFunc) *cobra.Command {
	opts := &loginOpts{}
	cmd := &cobra.Command{
		Use:     "login",
		GroupID: groups.Commands.ID,
		Short:   "Authenticate the chronoctl session as your Chronosphere user.",
		RunE: func(command *cobra.Command, strings []string) error {
			ls, err := cmd.newLoginServer(store, openFunc, opts)
			if err != nil {
				return errors.WithStack(err)
			}
			return ls.login(command.Context(), command.InOrStdin(), command.OutOrStdout(), command.ErrOrStderr())
		},
	}

	opts.registerFlags(cmd)
	return cmd
}

func newSetDefaultOrgCmd(store *token.Store) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "set-default-org ORG_NAME",
		GroupID: groups.Commands.ID,
		Short:   "Sets the default Chronosphere organization for future commands.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.WithStack(errMustIncludeOrgName)
			}
			return setDefaultOrg(store, args[0])
		},
	}
	return cmd
}

func newPrintAccessTokenCmd(store *token.Store) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "print-session-id ORG_NAME",
		GroupID: groups.Commands.ID,
		Short:   "Print the stored session id for a Chronosphere organization.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.WithStack(errMustIncludeOrgName)
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

func newListCmd(store *token.Store) *cobra.Command {
	outputFlags := output.NewFlags()
	cmd := &cobra.Command{
		Use:     "list",
		GroupID: groups.Commands.ID,
		Short:   "List all authenticated Chronosphere organizations.",
		RunE: func(cmd *cobra.Command, args []string) error {
			defer outputFlags.Close(cmd.OutOrStdout()) //nolint:errcheck
			orgs, err := store.List()
			if err != nil {
				return errors.WithStack(err)
			}
			defaultOrg, err := getDefaultOrg(store)
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

func getDefaultOrg(store *token.Store) (string, error) {
	org, err := store.Get(defaultOrgPath)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(org.Value), nil
}

// trySetOrgSessionEnv sets Chronosphere environment variables if they are currently unset.
// This only lasts for the duration of the command execution and does not persist.
// TODO(sc-84204) This is a temporary fix while we're developing this feature outside of chronoctl-core.
// Remove this when chronoctl-core sets transport config based on token store
func trySetOrgSessionEnv(store *token.Store) error {
	// Only try to set the API token if the environment variable is not already set to avoid interfering with the
	// user's desired configuration.
	if os.Getenv(client.ChronosphereAPITokenKey) != "" {
		return nil
	}

	// Check if org is unset, otherwise use default org
	org := cmp.Or(
		os.Getenv(client.ChronosphereOrgNameKey),
		os.Getenv(client.ChronosphereOrgKey),
	)
	if org == "" {
		// The org is unset in environment variables, set the default org so we can use it in downstream commands if necessary
		var err error
		org, err = getDefaultOrg(store)
		if err != nil {
			// Swallow the error if the default org name expired or doesn't exist yet - this will happen during normal operation
			if errors.Is(err, token.ErrTokenExpired) || errors.Is(err, token.ErrNotExist) {
				return nil
			}
			return err
		}
		if org == "" {
			// return early as there won't be a valid session id for an empty org
			return errors.New("default org is unset")
		}
		if err := os.Setenv(client.ChronosphereOrgNameKey, org); err != nil {
			// Best-effort attempt to reset the org name env var if we fail to set the environment variable
			os.Setenv(client.ChronosphereOrgNameKey, "") //nolint:errcheck
			return errors.WithStack(err)
		}
	}

	sessionID, err := store.Get(org)
	if err != nil {
		// Swallow the error if the token expired or the token doesn't exist yet - these will happen during normal operation
		if errors.Is(err, token.ErrTokenExpired) || errors.Is(err, token.ErrNotExist) {
			return nil
		}
		return errors.WithStack(err)
	}

	if err := os.Setenv(client.ChronosphereAPITokenKey, string(sessionID.Value)); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
