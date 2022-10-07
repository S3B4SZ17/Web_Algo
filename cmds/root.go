package cmds

import (
	"fmt"
	"log"
	"os"

	"github.com/S3B4SZ17/Web_Algo/app"
	"github.com/spf13/cobra"
)
var rootCmd = &cobra.Command{
  Use:   "kube_checks",
  Short: "A CLI app that automates and helps to debug issues in REST API calls env",
  Long: `A CLI app that automates and helps to debug issues in Kubernetes clusters env.
                With the help of a GUI we can easily detect issues in the cluster.
                Complete documentation is available at http://github.com/S3B4SZ17/kube_checks`,
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
    file, _ := cmd.Flags().GetString("file")
    config, err := app.ReadYaml(file); if err != nil {
      cmd.Usage()
      log.Fatal(err.Error())
    }
    
    app.StartApp(config)

  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
    var values string
    rootCmd.Flags().StringVarP(&values, "file", "f", "", "(Optional) A yaml file that defines custom values")
}