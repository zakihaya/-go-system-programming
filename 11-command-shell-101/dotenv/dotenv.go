package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	// go run dotenv.go printenv
	// を実行すると、.envに書いた内容がENVに取り込まれていることが分かる
	filename := flag.String("e", ".env", ".env file name to read")
	flag.Parse()
	cmdName := flag.Arg(0)
	args := flag.Args()[1:]
	flag.Args()

	cmd := exec.Command(cmdName, args...)

	envs := os.Environ()
	dotenvs, _ := godotenv.Read(*filename)
	for key, value := range dotenvs {
		envs = append(envs, key+"="+value)
	}
	cmd.Env = envs
	o, err := cmd.CombinedOutput()
	fmt.Println(string(o))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
