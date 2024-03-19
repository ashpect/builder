package builder

import (
	"context"
	"github.com/metacall/builder/pkg/staging"
	"github.com/moby/buildkit/client/llb"
	"github.com/spf13/cobra"
)

type DevOptions struct {
	DevImageFlags DevImageFlags
}

func NewDevOptions() *DevOptions {
	return &DevOptions{}
}

func NewDevCmd(o *DevOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dev",
		Short: "Build development base image for MetaCall",
		RunE: func(cmd *cobra.Command, args []string) error {
			base := cmd.Context().Value(baseKey{}).(llb.State)
			depsBase := staging.DevBase(base, o.DevImageFlags.Branch, args)
			depsBase, err := o.Run(depsBase)
			if err != nil {
				return err
			}
			// set final state
			cmd.SetContext(context.WithValue(cmd.Context(), finalKey{}, depsBase))
			return nil

		},
		Example: `builder dev -b develop nodejs typescript go rust wasm java c cobol`,
	}
	o.DevImageFlags.Set(cmd)

	return cmd
}

func (do *DevOptions) Run(devBase llb.State) (llb.State, error) {
	return devBase, nil
}
