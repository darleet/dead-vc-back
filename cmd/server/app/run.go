package app

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/bootstrap"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/depgraph"
)

const (
	cmdConfigName      = "config"
	cmdConfigShorthand = "c"
	cmdConfigValue     = "dev.yaml"
	cmdConfigUsage     = ".yaml file path"
)

type RunArgs struct {
	EnvPath string
}

func InitRunCommand() (*cobra.Command, error) {
	args := RunArgs{}
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Starts server",
		RunE: func(cmd *cobra.Command, _ []string) error {
			dg := depgraph.NewDepGraph()
			logger, err := dg.GetLogger()
			if err != nil {
				return errors.Wrap(err, "get logger")
			}

			viper.SetConfigFile(args.EnvPath)

			var cfg bootstrap.Config

			err = viper.ReadInConfig()
			if err != nil {
				logger.Warn("Config file Not Found. Using cli arguments")
			} else {
				logger.Debug("Using config file")
				err = viper.Unmarshal(&cfg)
				if err != nil {
					return errors.Wrap(err, "unmarshal config")
				}
			}

			logger.Debugw(
				"Got config",
				"args", args,
			)

			return nil
		},
	}

	cmd.Flags().StringVarP(&args.EnvPath, cmdConfigName, cmdConfigShorthand, cmdConfigValue, cmdConfigUsage)

	return cmd, nil
}
