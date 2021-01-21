// Copyright 2019 Liam White
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package command

import (
	"fmt"
	"os"
	"time"

	"github.com/liamawhite/licenser/pkg/file"
	"github.com/liamawhite/licenser/pkg/license"
	"github.com/spf13/cobra"
)

var (
	recurseDirectories bool
	licenseType        string
	template           license.Handler
	customStyle        []string
)

var rootCmd = &cobra.Command{
	Use:   "licenser",
	Short: "Applies and detects the absence of licenses in your repository",

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if licenseType == "Apache20" {
			template = license.NewApache20(time.Now().Year(), "")
		} else if licenseType == "MIT" {
			template = license.NewMIT(time.Now().Year(), "")
		}
		if template == nil {
			fmt.Fprintf(os.Stderr, "Unkown license type: %s\n", licenseType)
			os.Exit(1)
		}
	},
}

// Execute is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&recurseDirectories, "recurse", "r", false, "recurse from the passed directory")
	rootCmd.PersistentFlags().StringVarP(&licenseType, "license", "l", "Apache20", "Apache20 or MIT is supported")
	rootCmd.PersistentFlags().StringSliceVarP(&customStyle, "style", "s", []string{}, "<extension>:<style> available style is "+file.AllStyles())
}
