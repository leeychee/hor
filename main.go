package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

var path string

func init() {
	flag.StringVar(&path, "p", "./", "images's root path")
}

func main() {
	flag.Parse()
	store, err := newStore(path)
	if err != nil {
		log.Fatalf("Fail to new a db store, %s", err)
	}
	defer store.Close()

	r := gin.Default()

	// Rebuild db from the given path
	r.POST("/manage/_rebuild", func(c *gin.Context) {
		if err := store.Rebuild(); err != nil {
			log.Printf("fail to rebuild")
			c.JSON(500, gin.H{"status": "failed"})
		} else {
			c.JSON(200, gin.H{"status": "rebuilding"})
		}
	})
	r.GET("/image/_next", func(c *gin.Context) {
		img, err := store.NextImage()
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
		img := store.GetImage(uint64(id))
		c.JSON(200, img)
	})
	r.POST("/image/:id/_tag", func(c *gin.Context) {
		idstr := c.Params.ByName("id")
		id, _ := strconv.Atoi(idstr)
		img := store.GetImage(uint64(id))
		img.Identified = 1
		objs := &objects{}
		if err := c.BindJSON(objs); err != nil {
			c.JSON(500, gin.H{"status": err.Error()})
		} else {
			img.Objects = objs.Objs
			if err := store.UpdateImage(img); err != nil {
				c.JSON(500, gin.H{"status": err.Error()})
			} else {
				c.JSON(200, gin.H{"status": "ok"})
			}
		}
	})
	r.POST("/image/:id/_review", func(c *gin.Context) {
		idstr := c.Params.ByName("id")
		id, _ := strconv.Atoi(idstr)
		img := store.GetImage(uint64(id))
		img.Reviewed++
		objs := &objects{}
		if err := c.BindJSON(objs); err != nil {
			c.JSON(500, gin.H{"status": err.Error()})
		} else {
			img.Objects = objs.Objs
			if err := store.UpdateImage(img); err != nil {
				c.JSON(500, gin.H{"status": err.Error()})
			} else {
				c.JSON(200, gin.H{"status": "ok"})
			}
		}
	})
	r.Run(":9000")
}
