package main

import (
	"context"
	"encoding/json"
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/alanchchen/go-project-skeleton/pkg/api/user"
	"github.com/alanchchen/go-project-skeleton/pkg/app"
)

func init() {
	rootCmd.AddCommand(findUserCommand)

	findUserCommand.Flags().String("name", "", "the user name")
	findUserCommand.Flags().Int64("id", -1, "the user id")
}

var findUserCommand = &cobra.Command{
	Use:   "find",
	Short: "finds an new user by ID or name",
	Long:  "finds an new user by ID or name",
	RunE: func(cmd *cobra.Command, args []string) error {
		runner := app.NewRunner()
		if err := runner.BindCobraCommand(cmd, args...); err != nil {
			return err
		}

		initializers := []interface{}{
			NewConnection,
			NewClient,
		}

		return runner.RunCustom(func(client user.ServiceClient, cfg *viper.Viper) error {
			var resp *user.Users
			var err error

			if name := cfg.GetString("name"); name != "" {
				resp, err = client.FindUserByName(context.Background(), &user.FindUserByNameRequest{
					Name: name,
				})
			}
			if id := cfg.GetInt64("id"); id >= 0 {
				resp, err = client.FindUserById(context.Background(), &user.FindUserByIdRequest{
					Id: id,
				})
			}

			if err != nil {
				return err
			}

			rawData, err := json.MarshalIndent(resp, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(rawData))

			return nil
		}, initializers...)
	},
}
