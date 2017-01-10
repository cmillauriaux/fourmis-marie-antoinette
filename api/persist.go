package ants

import (
	"log"

	"golang.org/x/net/context"

	"google.golang.org/appengine/datastore"
)

const picturesKind = "Pictures"

type Persistance struct {
}

func (p *Persistance) GetLastPicture(ctx context.Context, cameraID int) (*Picture, error) {
	q := datastore.NewQuery(picturesKind).Filter("CameraID =", cameraID).Order("-DateTime").Limit(1)
	pictures := make([]Picture, 0, 1)
	if _, err := q.GetAll(ctx, &pictures); err != nil {
		return nil, err
	}

	if len(pictures) == 0 {
		return nil, nil
	}

	return &pictures[0], nil
}

func (p *Persistance) GetPreviousPicture(ctx context.Context, cameraID int, dateTime int64) (*Picture, error) {
	q := datastore.NewQuery(picturesKind).Filter("DateTime <", dateTime).Filter("CameraID =", cameraID).Order("-DateTime").Limit(1)
	pictures := make([]Picture, 0, 1)
	if _, err := q.GetAll(ctx, &pictures); err != nil {
		return nil, err
	}

	if len(pictures) == 0 {
		return nil, nil
	}

	return &pictures[0], nil
}

func (p *Persistance) GetNextPicture(ctx context.Context, cameraID int, dateTime int64) (*Picture, error) {
	q := datastore.NewQuery(picturesKind).Filter("DateTime >", dateTime).Filter("CameraID =", cameraID).Order("-DateTime").Limit(1)
	pictures := make([]Picture, 0, 1)
	if _, err := q.GetAll(ctx, &pictures); err != nil {
		return nil, err
	}

	if len(pictures) == 0 {
		return nil, nil
	}

	return &pictures[0], nil
}

func (p *Persistance) GetAllPicture(ctx context.Context, cameraID int) ([]Picture, error) {
	q := datastore.NewQuery(picturesKind).Filter("CameraID =", cameraID).Order("DateTime")
	nbElements, err := q.Count(ctx)
	if err != nil {
		return nil, err
	}

	pictures := make([]Picture, 0, nbElements)

	if _, err := q.GetAll(ctx, &pictures); err != nil {
		return nil, err
	}

	if len(pictures) == 0 {
		return nil, nil
	}

	return pictures, nil
}

func (p *Persistance) AddPicture(ctx context.Context, picture *Picture) (*datastore.Key, error) {
	newKey := datastore.NewKey(ctx, picturesKind, "", picture.DateTime, nil)
	return datastore.Put(ctx, newKey, picture)
}

func (p *Persistance) PutDataTest(ctx context.Context) error {
	for i := 00; i < 50; i++ {
		picture := &Picture{CameraID: 1, DateTime: int64(201612011310 + i), FileName: "11482059315", Link: "https://storage.googleapis.com/ants-photos/11482059315"}
		key, err := p.AddPicture(ctx, picture)
		if err != nil {
			return err
		}
		log.Println(key.String())
	}
	return nil
}
