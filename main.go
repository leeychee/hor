package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

var path string

func init() {
	flag.StringVar(&path, "p", "./images", "images's root path")
}

func main() {
	flag.Parse()
	store, err := newStore(path)
	if err != nil {
		log.Fatalf("Fail to new a db store, %s", err)
	}
	defer store.Close()

	r := gin.Default()
	r.Static("f", path)

	// Rebuild db from the given path
	r.POST("/manage/_rebuild", func(c *gin.Context) {
		if err := store.Rebuild(); err != nil {
			log.Printf("fail to rebuild")
			c.JSON(500, gin.H{"status": "failed"})
		} else {
			c.JSON(200, gin.H{"status": "rebuilding"})
		}
	})
	// Get next image
	// type=tag, return a untaged image
	// type=review, return a taged but unreviewed image
	r.GET("/images/_next", func(c *gin.Context) {
		_type := c.DefaultQuery("type", "tag")
		img, err := store.NextImage(_type)
		if err != nil {
			c.JSON(500, gin.H{"status": err})
		} else {
			c.JSON(200, img)
		}
	})
	// Get image by id
	r.GET("/image/:id", func(c *gin.Context) {
		idstr := c.Params.ByName("id")
		id, _ := strconv.Atoi(idstr)
		img, err := store.GetImage(uint64(id))
		if err != nil {
			c.JSON(500, gin.H{"status": err.Error()})
		} else {
			c.JSON(200, img)
		}
	})
	// Delete image by id
	r.DELETE("/image/:id", func(c *gin.Context) {
		idstr := c.Params.ByName("id")
		id, _ := strconv.Atoi(idstr)
		err := store.RemoveImage(uint64(id))
		if err != nil {
			c.JSON(500, gin.H{"status": err.Error()})
		} else {
			c.JSON(200, gin.H{"status": "delete successed."})
		}
	})
	r.POST("/image/:id/_tag", func(c *gin.Context) {
		idstr := c.Params.ByName("id")
		id, _ := strconv.Atoi(idstr)
		objs := &objects{}
		if err := c.BindJSON(objs); err != nil {
			c.JSON(500, gin.H{"status": err.Error()})
		} else {
			img, err := store.GetImage(uint64(id))
			if err != nil {
				c.JSON(500, gin.H{"status": err.Error()})
			} else {
				img.Identified = 1
				img.Objects = objs.Objs
				if err := store.UpdateImage(img); err != nil {
					c.JSON(500, gin.H{"status": err.Error()})
				} else {
					c.JSON(200, gin.H{"status": "ok"})
				}
			}
		}
	})
	r.POST("/image/:id/_review", func(c *gin.Context) {
		idstr := c.Params.ByName("id")
		id, _ := strconv.Atoi(idstr)
		objs := &objects{}
		if err := c.BindJSON(objs); err != nil {
			c.JSON(500, gin.H{"status": err.Error()})
		} else {
			img, err := store.GetImage(uint64(id))
			if err != nil {
				c.JSON(500, gin.H{"status": err.Error()})
			} else {
				img.Reviewed++
				img.Objects = objs.Objs
				if err := store.UpdateImage(img); err != nil {
					c.JSON(500, gin.H{"status": err.Error()})
				} else {
					c.JSON(200, gin.H{"status": "ok"})
				}
			}
		}
	})
	r.Run(":9000")
}
