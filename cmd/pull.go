/*
Copyright © 2020 Bo Yi <etng2004@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/etng/SpeeDocker/pkg/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("pull called")
		image := args[0]
		fmt.Printf("pulling image %s\n", image)
		baseUrl := cfg.GetString("client.server_url")
		apiUrl := fmt.Sprintf("%s/pull?image=%s", baseUrl, image)
		if resp, e := http.Get(apiUrl); e == nil {
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Printf("http response %s\n", body)
			data := map[string]string{}
			json.Unmarshal(body, &data)
			newImage := data["newImage"]
			utils.DockerPull(newImage)
			utils.DockerTag(newImage, image)
			utils.DockerRmi(newImage)
			fmt.Println("done")
		}

	},
	Args: cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
