package ants

import (
	"errors"
	"log"

	uuid "github.com/satori/go.uuid"

	"golang.org/x/net/context"

	"time"

	"google.golang.org/appengine/datastore"
)

const picturesKind = "Pictures"
const articlesKind = "Articles"

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
	q := datastore.NewQuery(picturesKind).Filter("DateTime >", dateTime).Filter("CameraID =", cameraID).Order("DateTime").Limit(1)
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

func (p *Persistance) AddArticle(ctx context.Context, article *Article) (*datastore.Key, error) {
	if article.ID == "" {
		article.ID = uuid.NewV4().String()
	}

	if article.DateTime == 0 {
		article.DateTime = time.Now().Unix()
	}

	newKey := datastore.NewKey(ctx, articlesKind, article.ID, 0, nil)
	return datastore.Put(ctx, newKey, article)
}

func (p *Persistance) GetAllArticles(ctx context.Context, published bool, details bool) ([]*Article, error) {
	q := datastore.NewQuery(articlesKind).Order("-DateTime")
	if published {
		q = q.Filter("Published =", published)
	}

	nbElements, err := q.Count(ctx)
	if err != nil {
		return nil, err
	}

	articles := make([]*Article, 0, nbElements)

	if _, err := q.GetAll(ctx, &articles); err != nil {
		return nil, err
	}

	if len(articles) == 0 {
		return nil, nil
	}

	return articles, nil
}

func (p *Persistance) GetArticle(ctx context.Context, ID string) (*Article, error) {
	q := datastore.NewQuery(articlesKind).Filter("ID =", ID).Limit(1)
	articles := make([]Article, 0, 1)
	if _, err := q.GetAll(ctx, &articles); err != nil {
		return nil, err
	}

	if len(articles) == 0 {
		return nil, nil
	}

	return &articles[0], nil
}

func (p *Persistance) DeleteArticle(ctx context.Context, ID string) error {
	q := datastore.NewQuery(articlesKind).Filter("ID =", ID).Limit(1)
	articles := make([]Article, 0, 1)
	if keys, err := q.GetAll(ctx, &articles); err != nil {
		return err
	} else {
		if len(keys) == 0 {
			return errors.New("Cannot retrieve article")
		}

		err = datastore.Delete(ctx, keys[0])
		if err != nil {
			return err
		}

		return nil
	}
}
