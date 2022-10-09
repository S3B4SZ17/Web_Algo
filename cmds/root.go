package cmds

import (
	"fmt"
	"log"
	"os"

	"github.com/S3B4SZ17/Web_Algo/app"
	"github.com/spf13/cobra"
)
var rootCmd = &cobra.Command{
  Use:   "Web algorithms",
  Short: "A microservice backend that resolves different algorithms",
  Long: `A microservice backend that resolves different algorithms.
                Using the gRPC benefits of efficient serialization and easily interface interactions we can resolve different algorithms.  .
                Complete documentation is available at http://github.com/S3B4SZ17/Web_Algo`,
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
    file, _ := cmd.Flags().GetString("file")
    config, err := app.ReadYaml(checkFile(file)); if err != nil {
      cmd.Usage()
      log.Fatal(err.Error())
    }
    
    app.StartApp(config)

  },
}

func checkFile(file string) string {
  if file == "" {
    config_yaml := "config.yml"
    return config_yaml
  }else{
    return file
  }
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