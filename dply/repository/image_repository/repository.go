package image_repository

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	grst_errors "github.com/herryg91/cdd/grst/errors"
	repository_intf "github.com/herryg91/dply/dply/app/repository"
	pbImage "github.com/herryg91/dply/dply/clients/grst/image"
	"github.com/herryg91/dply/dply/entity"
	"google.golang.org/grpc/metadata"
)

type repository struct {
	cli               pbImage.ImageApiClient
	docker_cli        *client.Client
	registry_username string
	registry_password string
	registry_host     string
}

func New(cli pbImage.ImageApiClient, cfg *entity.Config) (repository_intf.ImageRepository, error) {
	certFile := cfg.DockerCertificatesPath + "/cert.pem"
	keyFile := cfg.DockerCertificatesPath + "/key.pem"
	caFile := cfg.DockerCertificatesPath + "/ca.pem"

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	customTransport := &(*http.DefaultTransport.(*http.Transport)) // make shallow copy
	customTransport.TLSClientConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}

	dockerClient, err := client.NewClient(cfg.DockerHost, cfg.DockerVersion, &http.Client{Transport: customTransport}, map[string]string{})
	if err != nil {
		return nil, err
	}
	return &repository{
		cli:        cli,
		docker_cli: dockerClient,
		registry_host: cfg.
			RegistryHost, registry_username: cfg.
			RegistryUsername,
		registry_password: cfg.RegistryPassword}, nil
}

func (r *repository) Add(repoName, image, description string) error {
	u := entity.User{}.FromFile()
	if u == nil {
		return fmt.Errorf("%w: %s", repository_intf.ErrUserUnauthorized, "You are not login")
	}
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	_, err := r.cli.Add(ctx, &pbImage.AddReq{
		Image:       image,
		Repository:  repoName,
		Description: description,
	})
	if err != nil {
		grsterr, errparse := grst_errors.NewFromError(err)
		if errparse != nil {
			return err
		}
		return errors.New(grsterr.Message + ". " + fmt.Sprintf("%v", grsterr.OtherErrors))
	}
	return nil
}

func (r *repository) Remove(repoName, digest string) error {
	// Notes: unimplementated for now
	return nil
}

func (r *repository) Get(repoName string, page, size int) ([]entity.ContainerImage, error) {
	u := entity.User{}.FromFile()
	if u == nil {
		return []entity.ContainerImage{}, fmt.Errorf("%w: %s", repository_intf.ErrUserUnauthorized, "You are not login")
	}
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"Authorization": u.Token}))

	datas, err := r.cli.Get(ctx, &pbImage.GetReq{
		Repository: repoName,
		Page:       int32(page),
		Size:       int32(size),
	})
	if err != nil {
		grsterr, errparse := grst_errors.NewFromError(err)
		if errparse != nil {
			return []entity.ContainerImage{}, err
		}
		if grsterr.HTTPStatus == http.StatusForbidden {
			return []entity.ContainerImage{}, fmt.Errorf("%w: %s", repository_intf.ErrUnauthorized, grsterr.Message)
		}
		return []entity.ContainerImage{}, errors.New(grsterr.Message + ". " + fmt.Sprintf("%v", grsterr.OtherErrors))
	}

	resp := []entity.ContainerImage{}

	for _, data := range datas.Images {
		createdAt := data.CreatedAt.AsTime()
		resp = append(resp, entity.ContainerImage{
			Id:             int(data.Id),
			Digest:         data.Digest,
			Image:          data.Image,
			RepositoryName: data.Repository,
			Description:    data.Description,
			CreatedBy:      int(data.CreatedBy),
			CreatedAt:      &createdAt,
		})
	}
	return resp, nil
}

func (r *repository) BuildImage(repo_full_name string, src string) ([]string, error) {
	u := entity.User{}.FromFile()
	if u == nil {
		return []string{}, fmt.Errorf("%w: %s", repository_intf.ErrUserUnauthorized, "You are not login")
	}

	tar, err := archive.TarWithOptions(src+"/", &archive.TarOptions{})
	if err != nil {
		return []string{}, err
	}

	opts := types.ImageBuildOptions{
		Dockerfile: "Dockerfile",
		Tags:       []string{repo_full_name},
		// Tags:       []string{"gcr.io/emplogy/backend/" + data.Name},
		Remove: true,
	}

	res, err := r.docker_cli.ImageBuild(context.Background(), tar, opts)
	if err != nil {
		return []string{}, err
	}

	defer res.Body.Close()

	image_ids, err := print_build_docker_image(res.Body)
	if err != nil {
		return []string{}, err
	}

	return image_ids, nil
}

func (r *repository) PushImage(image_tag_name string) (digest string, err error) {
	u := entity.User{}.FromFile()
	if u == nil {
		return "", fmt.Errorf("%w: %s", repository_intf.ErrUserUnauthorized, "You are not login")
	}

	// https://cloud.google.com/container-registry/docs/advanced-authentication
	var authConfig = types.AuthConfig{
		Username:      r.registry_username,
		Password:      r.registry_password, //`{"type":"service_account","project_id":"emplogy","private_key_id":"860a466edaa5b1f40a40fd9ab40d1e3d39b04b64","private_key":"-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDIHgwo3Lqs3vvU\nNh1KlZ4WoTWv5RzTjbs2Alw0s601kjHarv2v1q1BXfUsvzLluIiBq1ipiDQKNDsT\nZLW3W/bv7FRQ4XZTkA5B3Blc4r2jI5muonz1RAoLnikLYeGBcxPKIQD52mg4h02h\nlguw0eDz2VF2HxMO+mAaMAJ1v2WVyomRYzFgeMmROvIbwYtHmKdpLfKfP5hvti+g\nvBr7rnwnVTdZ8wKmTpUSOCn1HSBCcxT2xgTI43KuCLQM7wSUgDloiI9cWmKY/xsp\nzoC+Ttsv8Wemmxwsf2OeszZxJol3579CPPZnGHTW7+fyd1QZO1g5eyKrL8Me7nc/\n4ZgGU5CPAgMBAAECggEAYsz2+uuznKsA0U6gWpUQ9qJEE+I9v7MPlYRHyttU7oiN\n5aaU6I5IfufpJt0FP4bCmqaiwRzWeJ099362+t7ERcM8A6y1E1+hsmF9Ai+OKi/m\n7eIaaKtdfEvrfsumHxfWUurFhRYQc6xVpywh/Hw6oJoQTo3cBn6WGQfQBxtmh7OW\nH5UBGKhkDyGo6+zCOoVy0iR2yTCUoe6M3tZccflDfNhz9tPFnFmfw9IqcxIFovGK\nr2GvcO9bPqvrl4Lvc+mNghaM7bl1NmcKhUpNJA7w1dGMzSr+7rPxplCpbz3mLIiQ\nDHaZ0cGnxd3E6dRTqS/5Kb12VfKl+07KE7IkldCQvQKBgQDwQWye5JwPkn2GQJmC\nS2cpDcPnKuFuXzG0oeKKPw7ZAkg7ll4nP1CnhV0QLoOgWoYunqS6Xxjy1s5rS42z\nxCpo8sjgcC0H/TeQ4pinjjXfPycDujcv0Iyhjab2oHVmC+zS2mknNsXozn4AM2kl\nUKCkt8s+LwEoMlOi+SJHw9pg6wKBgQDVO0G/z0DXakayELLNzFtaPiFU1fg0o3yZ\nDLGbtz/xlP79J32Try+M2RVMVKc8p8tfTSQ5MjMwCZ3yoaQ5/kZ+XYvB9NFb/Hzv\nl0CGpwozoU4WEKbnfecqFvKX1dsSZmR8nF9vf/uxwa1QLJPMWyWLKY9To7SzKJor\nYtIgLPnF7QKBgQC5DbuLe4yVFgFnWfSjfk68OWUOdmHi8KHJfvOOBln6Xp6ifwSQ\neF04Wym+YAV0iqVV3U4GW19NFJUz4aMItuzvnymIbf7Ra4HUMCTi0k++X9c+ML13\nL8xSV1gmGJu0eTT1h9N8p9yyn/I/V1oCquLBXOvIPs5GVtVC72AvJLTc9wKBgQCu\nr58LvoTGdYB5PIjfZH2qjp/L2oc+yHi5Adc3VIcEKSZEyudr5+cyol16bRec73ID\nHzV/zgp1XkuRjK73+8JQn95xBVnG3DCWL/li1tHavlk0Zmv11gVdS/NuRHr2tf+4\nvnrI47aVR6/usLZcgoddXKzYvpK4+5hh1tGCHpZ5eQKBgAuRfrA0x4Roh0UHeBAt\n1Lzf+ematSlWd9yq6yE0A38UXBBxH4219EzBEVajeax3ZAZMTKH+sgR9IoxOS6VZ\n9GcBwAZAs2T8nuEkUVeGR4mYS4z8ovv1K8hEN1z2pbFMKpKci6SGOcKAKF2ZBUxS\n4aY7HRIhOr9SI7zdnLiQOoDQ\n-----END PRIVATE KEY-----\n","client_email":"gcr-agent@emplogy.iam.gserviceaccount.com","client_id":"114421699910684804407","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_x509_cert_url":"https://www.googleapis.com/robot/v1/metadata/x509/gcr-agent%40emplogy.iam.gserviceaccount.com"}`,
		ServerAddress: r.registry_host,     //"https://gcr.io",
	}
	authConfigBytes, _ := json.Marshal(authConfig)
	authConfigEncoded := base64.URLEncoding.EncodeToString(authConfigBytes)

	optsPush := types.ImagePushOptions{RegistryAuth: authConfigEncoded}

	rd, err := r.docker_cli.ImagePush(context.Background(), image_tag_name, optsPush)
	if err != nil {
		return "", err
	}

	defer rd.Close()

	digest, err = print_push_image(rd)
	if err != nil {
		return "", err
	}

	return digest, nil
}

func (r *repository) DeleteImage(image_id string) error {
	u := entity.User{}.FromFile()
	if u == nil {
		return fmt.Errorf("%w: %s", repository_intf.ErrUserUnauthorized, "You are not login")
	}

	_, err := r.docker_cli.ImageRemove(context.Background(), image_id, types.ImageRemoveOptions{PruneChildren: true})
	if err != nil {
		return err
	}
	return nil
}
