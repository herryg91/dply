package image_usecase

import (
	"errors"
	"strings"

	"github.com/herryg91/dply/dply-server/app/repository"
)

func (uc *usecase) imageToDigest(image string) (string, error) {
	splitImage := strings.Split(image, "@")
	if len(splitImage) != 2 {
		return "", ErrInvalidImageFormat
	}

	fullDigest := strings.Replace(splitImage[1], "sha256:", "", -1)
	return fullDigest, nil
}

func (uc *usecase) generateShortDigest(fullDigest string) (string, error) {
	digestLength := 12
	loop := 0
	maxLoop := 10
	for {
		if loop > maxLoop {
			break
		}
		// Create short digest
		shortDigest := fullDigest
		if len(fullDigest) > digestLength {
			shortDigest = shortDigest[:digestLength]
		}

		// Check if short digest allowed
		_, err := uc.repo.GetByDigest(shortDigest)
		if err != nil {
			if !errors.Is(err, repository.ErrImageNotFound) {
				return "", err
			}
		} else {
			// if there duplicate digest then next loop
			digestLength++
			loop++
			continue
		}
		return shortDigest, nil
	}
	return "", errors.New("Failed to generate digest. Too many retry")
}
