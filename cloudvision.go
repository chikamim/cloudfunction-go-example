package main

import (
	"fmt"
	"io"

	vision "cloud.google.com/go/vision/apiv1"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

// func main() {
//
// 	err := detectSafeSearch(os.Stdout, "gs://cloudvision-sandbox/pexels-photo-603560.jpeg")
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// detectSafeSearch gets image properties from the Vision API for an image at the given file path.
func detectSafeSearch(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx, option.WithServiceAccountFile("./assets/google-service-account.json"))
	if err != nil {
		return err
	}

	image := vision.NewImageFromURI(file)
	props, err := client.DetectSafeSearch(ctx, image, nil)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, "Safe Search properties:")
	fmt.Fprintln(w, "Adult:", props.Adult)
	fmt.Fprintln(w, "Medical:", props.Medical)
	fmt.Fprintln(w, "Racy:", props.Racy)
	fmt.Fprintln(w, "Spoofed:", props.Spoof)
	fmt.Fprintln(w, "Violence:", props.Violence)

	return nil
}
