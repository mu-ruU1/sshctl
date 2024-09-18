package get

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mu-ruU1/sshctl/cmd"
	"github.com/spf13/cobra"
)

type SSHFileInfo struct {
	Dir   string
	Files []string
	Path  string
}

func NewSSHFileInfo() (*[]SSHFileInfo, error) {
	var dirFilesMap = make(map[string][]string)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	configDir := filepath.Join(homeDir, ".ssh", "conf.d")
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		return nil, err
	}

	err = filepath.Walk(configDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			path := filepath.Dir(path)
			fileName := info.Name()

			dirFilesMap[path] = append(dirFilesMap[path], fileName)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	var sshFileInfo []SSHFileInfo
	for path, file := range dirFilesMap {
		dir := strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
		sshFileInfo = append(sshFileInfo, SSHFileInfo{
			Dir:   dir,
			Files: file,
			Path:  path,
		})
	}

	return &sshFileInfo, nil
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(getCmd)
}
