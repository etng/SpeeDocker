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
	"fmt"
	"github.com/etng/SpeeDocker/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		port := cfg.GetInt("server.port")
		log.Printf("listening at port %d\n", port)
		dockerUser := cfg.GetString("server.docker.username")
		if len(dockerUser) == 0 {
			panic("should config docker user")
		}
		dockerPass := cfg.GetString("server.docker.password")
		if len(dockerPass) == 0 {
			dockerPass = utils.GetDockerPass()
		}
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), MakeHandler(dockerUser, dockerPass)))
	},
}

func MakeHandler(dockerUser, dockerPass string) *gin.Engine {
	engine := gin.New()
	engine.GET("/pull", func(context *gin.Context) {
		image := ""
		if _image, ok := context.GetQuery("image"); ok && len(_image) > 0 {
			image = _image
		} else {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "image is required",
			})
			return
		}
		newImage := SpeedUp(image, dockerUser, dockerPass)
		context.JSON(http.StatusOK, gin.H{
			"newImage": newImage,
		})
	})
	gin.SetMode(gin.DebugMode)
	return engine
}
func SpeedUp(image, dockerUser, dockerPass string) string {
	fmt.Printf("pulling image %s\n", image)
	newImage := fmt.Sprintf("%s/%s", dockerUser, strings.ReplaceAll(image, "/", "_"))
	//utils.DockerLogin(dockerUser, dockerPass)
	utils.DockerPull(image)
	utils.DockerTag(image, newImage)
	utils.DockerPush(newImage)
	return newImage
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
