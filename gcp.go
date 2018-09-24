package main

import (
	"context"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
	pb "google.golang.org/genproto/googleapis/cloud/vision/v1"
)

func getTextAnnotation() (*pb.TextAnnotation, error) {

	f, err := os.Open("doc.png")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	c, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := c.DetectDocumentText(
		ctx,
		image,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
