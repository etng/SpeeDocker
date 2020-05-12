package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"os"

	"io/ioutil"
	"os/exec"
)

func GetDockerPass() string{
	fmt.Printf("Now, please type in the password (mandatory): ")
	defer fmt.Println("")
	for{
		if rawPassword, e := terminal.ReadPassword(int(os.Stdin.Fd()));e==nil && len(rawPassword)>0{

			return string(rawPassword)
		}
	}
	return ""
}
func DockerLogin(username, password string){
	cmd := exec.Command(
		getDocker(),
		"login",
		fmt.Sprintf("-u%s", username),
		"--password-stdin",
	)
	stdin, _:=cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	stdin.Write([]byte(password))
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	failMsg, _ := ioutil.ReadAll(stderr)
	okMsg, _ := ioutil.ReadAll(stdout)
	log.Printf("Waiting for pull command to finish...")
	log.Printf("stderr:%s\n", okMsg)
	log.Printf("failMsg:%s\n", failMsg)
	if err := cmd.Wait(); err != nil {
		log.Printf("Command finished with error: %v", err)
	}
}
func getDocker() string {
	path, err := exec.LookPath("docker")
	if err != nil {
		log.Fatal("please install docker for your later use")
	}
	fmt.Printf("docker is available at %s\n", path)
	return path
}
func DockerExec(args ...string) {
	cmd := exec.Command(
		getDocker(),
		args...
	)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	failMsg, _ := ioutil.ReadAll(stderr)
	okMsg, _ := ioutil.ReadAll(stdout)
	log.Printf("Waiting for pull command to finish...")
	log.Printf("stderr:%s\n", okMsg)
	log.Printf("failMsg:%s\n", failMsg)
	if err := cmd.Wait(); err != nil {
		log.Printf("Command finished with error: %v", err)
	}
}
func DockerRmi(image string)  {
	log.Printf("removing image %s\n", image)
	DockerExec("rmi", image)
}
func DockerPush(image string)  {
	log.Printf("pushing image %s\n", image)
	DockerExec("push", image)
}
func DockerPull(image string)  {
	log.Printf("pulling image %s\n", image)
	DockerExec("pull", image)
}
func DockerTag(from, to string)  {
	log.Printf("tagging image %s %s\n", from, to)
	DockerExec("tag", from, to)
}