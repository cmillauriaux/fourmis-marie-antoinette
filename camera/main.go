package main

import (
	"log"
	"os"

	"golang.org/x/net/context"

	"strconv"
	"time"

	"io/ioutil"

	"cloud.google.com/go/datastore"
	"cloud.google.com/go/storage"
)

const picturesKind = "Pictures"
const projectID = "prototype-149014"

type Picture struct {
	FileName string
	Link     string
	DateTime int64
	CameraID int64
}

func main() {
	// Parse Arguments
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 2 {
		log.Fatalf("Not enough arguments")
	}

	cameraID, err := strconv.Atoi(argsWithoutProg[0])
	if err != nil {
		log.Fatalf("CameraID not valid")
	}

	// Prepares the new entity
	filePath := argsWithoutProg[1]
	dateTime := time.Now().Unix()
	fileName := strconv.Itoa(cameraID) + strconv.FormatInt(dateTime, 10)
	picture := Picture{CameraID: int64(cameraID), FileName: fileName, DateTime: dateTime}

	picture.Link = saveToFile(filePath, picture)
	saveToDB(picture)
}

func saveToFile(filePath string, picture Picture) string {
	ctx := context.Background()

	// Creates a client
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Prepares a new bucket
	bucket := client.Bucket("ants-photos")

	// Make Object in Data Store
	obj := bucket.Object(picture.FileName)

	w := obj.NewWriter(ctx)

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Cannot read file: %v", err)
	}
	w.Write(file)
	w.Close()

	attrs, err := obj.Attrs(ctx)
	if err != nil {
		log.Fatalf("Cannot retrieve object attributes: %v", err)
	}

	return attrs.MediaLink
}

func saveToDB(picture Picture) {
	// Make Google Cloud Context
	ctx := context.Background()

	// Connect to Google Cloud DataStore
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	pictureKey := datastore.NameKey(picturesKind, picture.FileName, nil)

	// Saves the entity
	if _, err := client.Put(ctx, pictureKey, &picture); err != nil {
		log.Fatalf("Failed to save task: %v", err)
	}
}
