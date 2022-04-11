package cmd

import (
	"os"
	"path"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"https://github.com/fsmile2/ckm8/common"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize ckm8 node configuration.",
	Long:  ``,
	Run:   runInit,
}

func init() {
	RootCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, args []string) {
	if _, err := os.Stat(cfgPath); !os.IsNotExist(err) {
		log.WithFields(log.Fields{"err": err, "path": cfgPath}).Fatal("Folder already exists!")
	}

	if err := os.Mkdir(cfgPath, 0700); err != nil {
		log.WithFields(log.Fields{"err": err, "path": cfgPath}).Fatal("Failed to create config folder")
	}

	if err := common.WriteInitialConfig(path.Join(cfgPath, "config.yaml")); err != nil {
		log.WithFields(log.Fields{"err": err, "path": cfgPath}).Fatal("Failed to write config")
	}
}
