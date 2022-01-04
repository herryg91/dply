package image_repository

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/gosuri/uilive"
)

type ErrorLine struct {
	Error       string      `json:"error"`
	ErrorDetail ErrorDetail `json:"errorDetail"`
}

type ErrorDetail struct {
	Message string `json:"message"`
}

type BuildDockerImageMsg struct {
	Stream string              `json:"stream"`
	Aux    BuildDockerImageAux `json:"aux"`
}

type BuildDockerImageAux struct {
	Id string `json:"ID"`
}

func print_build_docker_image(rd io.Reader) ([]string, error) {
	var last_line string
	image_ids := []string{}

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		last_line = scanner.Text()

		var msg BuildDockerImageMsg
		json.Unmarshal([]byte(scanner.Text()), &msg)
		if msg.Stream != "" {
			fmt.Print(msg.Stream)
		} else {
			fmt.Println(scanner.Text())
		}
		if msg.Aux.Id != "" {
			image_ids = append(image_ids, msg.Aux.Id)
		}
	}

	errLine := &ErrorLine{}
	json.Unmarshal([]byte(last_line), errLine)
	if errLine.Error != "" {
		return []string{}, errors.New(errLine.Error)
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	return image_ids, nil
}

type PushImageMsg struct {
	Status         string                  `json:"status"`
	Aux            PushImageAux            `json:"aux"`
	ProgressDetail PushImageProgressDetail `json:"progressDetail"`
	Id             string                  `json:"id"`
	Progress       string                  `json:"progress"`
}
type PushImageAux struct {
	Tag    string `json:"Tag"`
	Digest string `json:"Digest"`
	Size   int    `json:"Size"`
}
type PushImageProgressDetail struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}

func print_push_image(rd io.Reader) (string, error) {
	var lastLine string
	digest := ""
	writer := uilive.New()

	for i := 0; i <= 100; i++ {

		time.Sleep(time.Millisecond * 5)
	}
	scanner := bufio.NewScanner(rd)
	first_push := false
	for scanner.Scan() {
		lastLine = scanner.Text()

		var msg PushImageMsg
		json.Unmarshal([]byte(scanner.Text()), &msg)
		if msg.Status == "Pushing" {
			if !first_push {
				writer.Start()
				first_push = true

				fmt.Fprintf(writer, "Pushing.. (%d/%d): %s\n", msg.ProgressDetail.Current, msg.ProgressDetail.Total, msg.Progress)
			}
		} else if msg.Aux.Digest != "" {
			digest = msg.Aux.Digest
			fmt.Println(scanner.Text())
		} else {
			fmt.Println(scanner.Text())
		}

	}
	writer.Stop() // flush and stop rendering

	errLine := &ErrorLine{}
	json.Unmarshal([]byte(lastLine), errLine)
	if errLine.Error != "" {
		return "", errors.New(errLine.Error)
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return digest, nil
}
