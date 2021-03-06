package cluster

import (
	"github.com/appcelerator/amp/cli"
	"github.com/appcelerator/amp/cmd/amplifier/server/configuration"
	"github.com/spf13/cobra"
)

type clusterOpts struct {
	managers      int
	workers       int
	provider      string
	name          string
	tag           string
	registration  string
	notifications bool
}

var (
	opts = &clusterOpts{
		managers:      3,
		workers:       2,
		provider:      "local",
		name:          "",
		tag:           "latest",
		registration:  configuration.RegistrationDefault,
		notifications: true,
	}
)

// NewClusterCommand returns a new instance of the cluster command.
func NewClusterCommand(c cli.Interface) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cluster",
		Short:   "Cluster management operations",
		PreRunE: cli.NoArgs,
		RunE:    c.ShowHelp,
	}
	cmd.AddCommand(NewCreateCommand(c))
	cmd.AddCommand(NewListCommand(c))
	cmd.AddCommand(NewRemoveCommand(c))
	cmd.AddCommand(NewStatusCommand(c))
	cmd.AddCommand(NewUpdateCommand(c))
	cmd.AddCommand(NewNodeCommand(c))
	return cmd
}

func queryCluster(c cli.Interface, args []string, env map[string]string) error {
	if err := check(opts.provider); err != nil {
		return err
	}
	err := Run(c, args, env)
	return err
}

// Map cli cluster flags to target bootstrap cluster command flags,
// append to and return args array
func reflag(cmd *cobra.Command, flags map[string]string, args []string) []string {
	// transform src flags to target flags and add flag and value to cargs
	for s, t := range flags {
		if cmd.Flag(s).Changed {
			args = append(args, t, cmd.Flag(s).Value.String())
		}
	}
	return args
}
