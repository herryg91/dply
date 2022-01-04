package cli_image

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"

	image_usecase "github.com/herryg91/dply/dply/app/usecase/image"
	"github.com/herryg91/dply/dply/entity"
	"github.com/herryg91/dply/dply/pkg/serviceYaml"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/spf13/cobra"
)

type CmdImageCreate struct {
	*cobra.Command
	image_uc image_usecase.UseCase

	name        string
	description string
}

func newCmdImageCreate(image_uc image_usecase.UseCase) *CmdImageCreate {
	c := &CmdImageCreate{image_uc: image_uc}
	c.Command = &cobra.Command{
		Use:   "create",
		Short: "Create image",
		Long:  "Create image",
	}
	c.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "service (repository) name of image")
	c.Command.Flags().StringVarP(&c.description, "desc", "d", "", "image description")
	return c
}

func (c *CmdImageCreate) runCommand(cmd *cobra.Command, args []string) error {
	if c.image_uc == nil {
		return errors.New("You haven't setup the configuration. command: `dply config edit` then set the `dply_server_host``")
	} else if c.description == "" {
		log.Println("dasndjsanjdkasn")
		return errors.New("`--desc / -d` is required")
	} else if c.name == "" {
		data, err := serviceYaml.GetServiceYAML("service.yaml")
		if err != nil || data.Name == "" {
			return errors.New("`--name / -n` is required")
		}
		c.name = data.Name
	}

	cfg := entity.Config{}.FromFile()
	err := c.image_uc.Create(c.name, cfg.RegistryTagPrefix, c.description)
	if err != nil {
		return err
	}

	// cc := &client.Client{}
	// client.FromEnv(cc)
	// log.Println("client.FromEnv", cc, data.Name)

	// //// filess
	// certFile := "/Users/hg/go/src/github.com/herryg91/getwedding/certs3/client/cert.pem"
	// keyFile := "/Users/hg/go/src/github.com/herryg91/getwedding/certs3/client/key.pem"
	// caFile := "/Users/hg/go/src/github.com/herryg91/getwedding/certs3/client/ca.pem"
	// cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Load CA cert
	// caCert, err := ioutil.ReadFile(caFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// caCertPool := x509.NewCertPool()
	// caCertPool.AppendCertsFromPEM(caCert)

	// customTransport := &(*http.DefaultTransport.(*http.Transport)) // make shallow copy
	// customTransport.TLSClientConfig = &tls.Config{
	// 	Certificates: []tls.Certificate{cert},
	// 	RootCAs:      caCertPool,
	// }
	// httpclient := &http.Client{Transport: customTransport}

	// dockerHost := "tcp://localhost:12376"
	// dockerHost = "tcp://dind.dply.svc.cluster.local:2376"
	// dockerClient, err := client.NewClient(dockerHost, "1.41", httpclient, map[string]string{})
	// if err != nil {
	// 	return err
	// }

	// is, _ := dockerClient.ImageList(context.Background(), types.ImageListOptions{
	// 	All:     false,
	// 	Filters: filters.NewArgs(filters.Arg("dangling", "false"), filters.Arg("reference", "gcr.io/emplogy/backend/test-api:latest")),
	// })
	// log.Println(is)

	// // log.Println(dockerClient.ImageSearch(context.Background(), "gcr.io/emplogy/backend/test-api", types.ImageSearchOptions{
	// // 	Limit: 1,
	// // }))
	// return nil
	// baseFolder, _ := os.Getwd()
	// tar, err := archive.TarWithOptions(baseFolder+"/", &archive.TarOptions{})
	// if err != nil {
	// 	return err
	// }

	// opts := types.ImageBuildOptions{
	// 	Dockerfile: "Dockerfile",
	// 	Tags:       []string{"gcr.io/emplogy/backend/" + data.Name},
	// 	Remove:     true,
	// }

	// res, err := dockerClient.ImageBuild(context.Background(), tar, opts)
	// if err != nil {
	// 	return err
	// }

	// defer res.Body.Close()

	// err = print(res.Body)
	// if err != nil {
	// 	return err
	// }

	// err = imagePush(dockerClient)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func imagePush(dockerClient *client.Client) error {
	// jsonloc := "/Users/hg/go/src/github.com/herryg91/dply/dply/emplogy-860a466edaa5.json"
	// password, _ := ioutil.ReadFile(jsonloc)
	var authConfig = types.AuthConfig{
		// Username: "oauth2accesstoken",
		// // https://cloud.google.com/container-registry/docs/advanced-authentication | gcloud auth print-access-token
		// Password:      "ya29.a0ARrdaM8_V1Fj1EtPDnE1I_7AwM410XCBYwgS2WiThqLHHJRy0nLWZ56M0f7Qml8Ebs8rQQZmuQ-wdt3x6D_pcnj67SRiHog7KHnntUuZi6FGQG9Druk2IYG4eXlDfNi7iVG_3kB9wbT3p4XzIcaHIQDw22f2KNaCHfLBBxE",
		// ServerAddress: "https://gcr.io",
		Username:      "_json_key",
		Password:      `{"type":"service_account","project_id":"emplogy","private_key_id":"860a466edaa5b1f40a40fd9ab40d1e3d39b04b64","private_key":"-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDIHgwo3Lqs3vvU\nNh1KlZ4WoTWv5RzTjbs2Alw0s601kjHarv2v1q1BXfUsvzLluIiBq1ipiDQKNDsT\nZLW3W/bv7FRQ4XZTkA5B3Blc4r2jI5muonz1RAoLnikLYeGBcxPKIQD52mg4h02h\nlguw0eDz2VF2HxMO+mAaMAJ1v2WVyomRYzFgeMmROvIbwYtHmKdpLfKfP5hvti+g\nvBr7rnwnVTdZ8wKmTpUSOCn1HSBCcxT2xgTI43KuCLQM7wSUgDloiI9cWmKY/xsp\nzoC+Ttsv8Wemmxwsf2OeszZxJol3579CPPZnGHTW7+fyd1QZO1g5eyKrL8Me7nc/\n4ZgGU5CPAgMBAAECggEAYsz2+uuznKsA0U6gWpUQ9qJEE+I9v7MPlYRHyttU7oiN\n5aaU6I5IfufpJt0FP4bCmqaiwRzWeJ099362+t7ERcM8A6y1E1+hsmF9Ai+OKi/m\n7eIaaKtdfEvrfsumHxfWUurFhRYQc6xVpywh/Hw6oJoQTo3cBn6WGQfQBxtmh7OW\nH5UBGKhkDyGo6+zCOoVy0iR2yTCUoe6M3tZccflDfNhz9tPFnFmfw9IqcxIFovGK\nr2GvcO9bPqvrl4Lvc+mNghaM7bl1NmcKhUpNJA7w1dGMzSr+7rPxplCpbz3mLIiQ\nDHaZ0cGnxd3E6dRTqS/5Kb12VfKl+07KE7IkldCQvQKBgQDwQWye5JwPkn2GQJmC\nS2cpDcPnKuFuXzG0oeKKPw7ZAkg7ll4nP1CnhV0QLoOgWoYunqS6Xxjy1s5rS42z\nxCpo8sjgcC0H/TeQ4pinjjXfPycDujcv0Iyhjab2oHVmC+zS2mknNsXozn4AM2kl\nUKCkt8s+LwEoMlOi+SJHw9pg6wKBgQDVO0G/z0DXakayELLNzFtaPiFU1fg0o3yZ\nDLGbtz/xlP79J32Try+M2RVMVKc8p8tfTSQ5MjMwCZ3yoaQ5/kZ+XYvB9NFb/Hzv\nl0CGpwozoU4WEKbnfecqFvKX1dsSZmR8nF9vf/uxwa1QLJPMWyWLKY9To7SzKJor\nYtIgLPnF7QKBgQC5DbuLe4yVFgFnWfSjfk68OWUOdmHi8KHJfvOOBln6Xp6ifwSQ\neF04Wym+YAV0iqVV3U4GW19NFJUz4aMItuzvnymIbf7Ra4HUMCTi0k++X9c+ML13\nL8xSV1gmGJu0eTT1h9N8p9yyn/I/V1oCquLBXOvIPs5GVtVC72AvJLTc9wKBgQCu\nr58LvoTGdYB5PIjfZH2qjp/L2oc+yHi5Adc3VIcEKSZEyudr5+cyol16bRec73ID\nHzV/zgp1XkuRjK73+8JQn95xBVnG3DCWL/li1tHavlk0Zmv11gVdS/NuRHr2tf+4\nvnrI47aVR6/usLZcgoddXKzYvpK4+5hh1tGCHpZ5eQKBgAuRfrA0x4Roh0UHeBAt\n1Lzf+ematSlWd9yq6yE0A38UXBBxH4219EzBEVajeax3ZAZMTKH+sgR9IoxOS6VZ\n9GcBwAZAs2T8nuEkUVeGR4mYS4z8ovv1K8hEN1z2pbFMKpKci6SGOcKAKF2ZBUxS\n4aY7HRIhOr9SI7zdnLiQOoDQ\n-----END PRIVATE KEY-----\n","client_email":"gcr-agent@emplogy.iam.gserviceaccount.com","client_id":"114421699910684804407","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_x509_cert_url":"https://www.googleapis.com/robot/v1/metadata/x509/gcr-agent%40emplogy.iam.gserviceaccount.com"}`,
		ServerAddress: "https://gcr.io",
	}
	authConfigBytes, _ := json.Marshal(authConfig)
	authConfigEncoded := base64.URLEncoding.EncodeToString(authConfigBytes)

	optsPush := types.ImagePushOptions{RegistryAuth: authConfigEncoded}

	rd, err := dockerClient.ImagePush(context.Background(), "gcr.io/emplogy/backend/test-api", optsPush)
	if err != nil {
		return err
	}

	defer rd.Close()

	err = print(rd)
	if err != nil {
		return err
	}

	return nil
}

type ErrorLine struct {
	Error       string      `json:"error"`
	ErrorDetail ErrorDetail `json:"errorDetail"`
}

type ErrorDetail struct {
	Message string `json:"message"`
}

func print(rd io.Reader) error {
	var lastLine string

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		lastLine = scanner.Text()
		fmt.Println(scanner.Text())
	}

	errLine := &ErrorLine{}
	json.Unmarshal([]byte(lastLine), errLine)
	if errLine.Error != "" {
		return errors.New(errLine.Error)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
