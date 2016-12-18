package ants

import (
	"bufio"
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"io"
	"log"

	"golang.org/x/net/context"

	"google.golang.org/appengine/urlfetch"
)

func makeGif(files []string, ctx context.Context) []byte {
	outGif := &gif.GIF{}
	for _, name := range files {

		// Download file
		resp, err := getFile(name, ctx)
		if err != nil {
			log.Println("Cannot download file : ", err)
			continue
		}

		// Get image
		img, err := readImage(resp)
		if err != nil {
			log.Println("Cannot get image : ", err)
			continue
		}

		// Convert to GIF
		imgGif, err := formatgif(img)
		if err != nil {
			log.Println("Cannot convert GIF : ", err)
			continue
		}

		inGif, err := gif.Decode(imgGif)
		if err != nil {
			log.Println("Cannot convert file : ", err)
			continue
		}
		log.Println(inGif)
		resp.Close()

		outGif.Image = append(outGif.Image, inGif.(*image.Paletted))
		outGif.Delay = append(outGif.Delay, 10)
	}

	// Convert to byte array
	buffer := bytes.NewBuffer(nil)
	bufferWriter := bufio.NewWriter(buffer)
	gif.EncodeAll(bufferWriter, outGif)

	return buffer.Bytes()
}

func getFile(url string, ctx context.Context) (io.ReadCloser, error) {
	log.Println("GET FILE")
	client := urlfetch.Client(ctx)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func readImage(image io.ReadCloser) (image.Image, error) {
	log.Println("READ IMAGE")
	img, err := jpeg.Decode(image)

	return img, err
}

func formatgif(img image.Image) (io.Reader, error) {
	log.Println("FORMAT GIF")
	r, w := io.Pipe()

	var opt gif.Options
	opt.NumColors = 256
	go func() {
		_ = gif.Encode(w, img, &opt)
		w.Close()
		log.Println("CLOSE GIF")
	}()

	return r, nil
}
