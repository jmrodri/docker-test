package main

import (
	"fmt"
	"os"

	docker "github.com/fsouza/go-dockerclient"
)

type Param map[string]string
type MyParam map[string]string

func t(myp MyParam) Param {
	return Param(myp)
}

func main() {

	/*
		myp := MyParam{"foo": "bar"}
		p := Param{"foo1": "bar1"}
		fmt.Println(myp)
		fmt.Println(p)
		os.Exit(1)
	*/

	fmt.Println("docker-test-client")
	if len(os.Args) < 2 {
		fmt.Println("Supply a number")
		os.Exit(1)
	}
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		fmt.Println(err)
	}

	arg := os.Args[1]

	if arg == "l" {
		imgs, err := client.ListImages(docker.ListImagesOptions{All: false})
		if err != nil {
			fmt.Println(err)
		}

		for _, img := range imgs {
			fmt.Println("ID: ", img.ID)
			fmt.Println("RepoTags: ", img.RepoTags)
		}
	}

	if arg == "p" {
		fmt.Println("docker pull")
		client.PullImage(docker.PullImageOptions{Repository: "hello-world", OutputStream: os.Stdout}, docker.AuthConfiguration{})
	}

	if arg == "r" {
		fmt.Println("docker run")
	}
}
