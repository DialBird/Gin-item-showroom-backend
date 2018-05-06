package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"s3_upload_api/item"
	"strconv"
)

var cl *mgo.Collection

// SetItemRouter is
func SetItemRouter(r *gin.Engine, session *mgo.Session) {
	cl = session.DB("test").C("items")

	r.GET("/", getRoot)
	r.GET("/items", getItems)
	r.POST("/items", postItem)
	r.PUT("/items", putItem)
	r.POST("/delete_item", deleteItem)
}

func getRoot(c *gin.Context) {
	c.String(200, "it worked?!!")
}

func getItems(c *gin.Context) {
	items := []item.Item{}
	err := cl.Find(bson.M{}).All(&items)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, items)
}

func postItem(c *gin.Context) {
	item := item.Item{}
	item.Name = c.PostForm("Name")
	item.ImageURL = c.PostForm("ImageURL")
	item.Price, _ = strconv.Atoi(c.PostForm("Price"))
	item.Description = c.PostForm("Description")

	err := cl.Insert(&item)
	if err != nil {
		if mgo.IsDup(err) {
			log.Fatal("Duplicate key error")
		} else {
			log.Fatal(err)
		}
	}
}

func putItem(c *gin.Context) {
	if name, newName := c.PostForm("Name"), c.PostForm("NewName"); name != "" && newName != "" {
		selector := bson.M{"name": name}
		update := bson.M{"$set": bson.M{"name": newName}}
		err := cl.Update(selector, update)
		if err != nil {
			if mgo.IsDup(err) {
				log.Fatal("Duplicate key error")
			} else {
				log.Fatal(err)
			}
		}
	}
}

func deleteItem(c *gin.Context) {
	if name := c.PostForm("Name"); name != "" {
		selector := bson.M{"name": name}
		err := cl.Remove(selector)
		if err != nil {
			if mgo.IsDup(err) {
				log.Fatal("Duplicate key error")
			} else {
				log.Fatal(err)
			}
		}
	}
}
