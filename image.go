package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/boltdb/bolt"
)

type version struct {
	Version string `json:"version"`
}

type object struct {
	X    int    `json:"x"`
	Y    int    `json:"y"`
	W    int    `json:"w"`
	H    int    `json:"h"`
	Type string `json:"type"`
	Path string `json:"path"`
}

type objects struct {
	Objs []*object `json:"objects"`
}

type image struct {
	ID           uint64    `json:"id"`
	Path         string    `json:"path"`
	Identified   int       `json:"identified"`
	Reviewed     int       `json:"reviewed"`
	Objects      []*object `json:"objects"`
	CheckoutTime time.Time `json:"checkout_time"`
	DemarkedTime time.Time `json:"demarked_time"`
	ReviewedTime time.Time `json:"reviewed_time"`
}

type stat struct {
	All        int `json:"all"`
	Identified int `json:"identified"`
	Reviewed   int `josn:"reviewed"`
}

type store struct {
	db          *bolt.DB
	path        string
	dbpath      string
	initialized bool
}

var dbFilename = "hor.db"
var imgBucketName = []byte("image")

func newStore(path string) (*store, error) {
	dbPath := filepath.Join(path, dbFilename)
	dbFileExist := false
	if _, err := os.Stat(dbPath); err == nil {
		dbFileExist = true
	}
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, err
	}
	s := &store{
		db,
		path,
		dbPath,
		dbFileExist,
	}
	if !dbFileExist {
		log.Printf("initializing db...\n")
		if err := s.init(); err != nil {
			return nil, err
		}
	} else {
		log.Println("already has db.")
	}
	return s, nil
}

func (s *store) init() error {
	imgC := make(chan string)
	err := s.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket(imgBucketName)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	// Search path for images
	go func() {
		filepath.Walk(s.path, func(path string, f os.FileInfo, err error) error {
			if !f.IsDir() && strings.HasPrefix(f.Name(), "P_") {
				log.Printf("Find a image: %s\n", path)
				relativePath := path[len(s.path)-1:]
				imgC <- relativePath
			}
			return nil
		})
		close(imgC)
	}()
	// Write images data into db
	go func() {
		for {
			path, ok := <-imgC
			if !ok {
				break
			}
			img := &image{Path: path}
			if err := s.CreateImage(img); err != nil {
				log.Printf("fail to create image %s, %v\n", path, err)
			} else {
				log.Printf("save %d: %s", img.ID, img.Path)
			}
		}
		s.initialized = true
	}()
	return nil
}

func (s *store) Rebuild() error {
	if err := s.Close(); err != nil {
		return err
	}
	backpath := fmt.Sprintf("%s.%d", s.dbpath, time.Now().Unix())
	if err := os.Rename(s.dbpath, backpath); err != nil {
		return err
	}
	fmt.Printf("backup db to %s.", backpath)
	s1, err := newStore(s.path)
	if err != nil {
		return err
	}
	s.db = s1.db
	return nil
}

func (s *store) Close() error {
	return s.db.Close()
}

func (s *store) NextImage(t string) (*image, error) {
	var img *image
	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(imgBucketName)
		bc := b.Cursor()
		for k, v := bc.First(); k != nil; k, v = bc.Next() {
			img = &image{}
			if err := json.Unmarshal(v, img); err != nil {
				return err
			}
			if t == "tag" && img.Identified == 0 && time.Since(img.CheckoutTime) > time.Second*60 {
				img.CheckoutTime = time.Now()
				buf, _ := json.Marshal(img)
				b.Put(itob(img.ID), buf)
				return nil
			} else if t == "review" && img.Identified > 0 && img.Reviewed == 0 && time.Since(img.CheckoutTime) > time.Second*60 {
				img.CheckoutTime = time.Now()
				buf, _ := json.Marshal(img)
				b.Put(itob(img.ID), buf)
				return nil
			}
		}
		return errors.New("no more images")
	})
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (s *store) CreateImage(img *image) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(imgBucketName)
		s, err := b.NextSequence()
		if err != nil {
			return err
		}
		img.ID = s
		buf, _ := json.Marshal(img)
		b.Put(itob(s), buf)
		return nil
	})
}

func (s *store) UpdateImage(img *image) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(imgBucketName)
		img.CheckoutTime = time.Time{}
		buf, _ := json.Marshal(img)
		b.Put(itob(img.ID), buf)
		return nil
	})
}

func (s *store) GetImage(id uint64) (*image, error) {
	img := &image{}
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(imgBucketName)
		v := b.Get(itob(id))
		return json.Unmarshal(v, img)
	})
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (s *store) RemoveImage(id uint64) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(imgBucketName)
		return b.Delete(itob(id))
	})
}

func (s *store) Stat() (*stat, error) {
	st := &stat{}
	err := s.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(imgBucketName).Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			img := &image{}
			err := json.Unmarshal(v, img)
			if err != nil {
				return err
			}
			st.All++
			if img.Identified > 0 {
				st.Identified++
			}
			if img.Reviewed > 0 {
				st.Reviewed++
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return st, nil
}
