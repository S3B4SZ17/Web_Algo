package cmds

import (
	"fmt"
	"os"

	"github.com/S3B4SZ17/Web_Algo/app"
	"github.com/S3B4SZ17/Web_Algo/config"
	"github.com/S3B4SZ17/Web_Algo/management"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var rootCmd = &cobra.Command{
	Use:   "Web algorithms",
	Short: "A microservice backend that resolves different algorithms",
	Long: `A microservice backend that resolves different algorithms.
                Using the gRPC benefits of efficient serialization and easily interface interactions we can resolve different algorithms.  .
                Complete documentation is available at http://github.com/S3B4SZ17/Web_Algo`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		cfgFile, err := cmd.Flags().GetString("config")

		if err != nil {
			cmd.Usage()
			management.Log.Error(err.Error())
		} else {
			cfgObj, _ := config.LoadConfig(cfgFile)
			management.InitializeZapCustomLogger()
			management.Log.Info("Loaded config initial configuration", zap.String("configFile", cfgFile))
			management.Log.Info("Email_service configuration", zap.String("email", cfgObj.Smtp_server.Email_from), zap.Int("port", cfgObj.Smtp_server.Port), zap.String("redirect_uri", viper.GetString("google.redirect_uri")))
			app.StartApp(&cfgObj)
		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	var cfgFile string
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./app.env)")
}
