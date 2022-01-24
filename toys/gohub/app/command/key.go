package command

import (
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/console"
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/helpers"
	"github.com/spf13/cobra"
)

var KeyCommand = &cobra.Command{
	Use: "key",
	Short: "Generate an App Key, will print the generated Key",
	Run: runKeyGenerate,
	Args: cobra.NoArgs, // 不允许传参
}

func runKeyGenerate(cd *cobra.Command, args []string) {
	console.Success("---")
    console.Success("App Key:")
    console.Success(helpers.RandomString(32))
    console.Success("---")
    console.Warning("please go to .env file to change the APP_KEY option")
}