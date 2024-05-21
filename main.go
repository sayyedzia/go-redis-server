package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	r := gin.Default()

	// routes
	r.GET("/v1/api/cache", GetCaches)
	r.GET("/v1/api/cache/:key", GetCacheByKey)
	r.POST("/v1/api/cache", CreateCache)
	r.PUT("/v1/api/cache", UpdateCache)
	r.DELETE("/v1/api/cache/:key", DeleteCache)

	r.Run(":8000")
}

// get all Caches
func GetCaches(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "In Progress"})
}

// get cache by key
func GetCacheByKey(c *gin.Context) {
	key := c.Param("key")

	// Connect to Redis server
	ctx := context.Background()
	opt, conn_err := redis.ParseURL("redis://:@localhost:6379/0")
	if conn_err != nil {
		panic(conn_err)
	}
	rdb := redis.NewClient(opt)
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Cache not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{key: val})

}

// create cache
func CreateCache(c *gin.Context) {

	// Connect to Redis server
	ctx := context.Background()
	opt, conn_err := redis.ParseURL("redis://:@localhost:6379/0")
	if conn_err != nil {
		panic(conn_err)
	}
	rdb := redis.NewClient(opt)

	var cache Cache

	if err := c.BindJSON(&cache); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err1 := rdb.Set(ctx, cache.Key, cache.Value, 0).Err()
	if err1 != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Unable to cache data"})
		return
	}

	c.JSON(http.StatusCreated, cache)
}

// update cache
func UpdateCache(c *gin.Context) {
	// Connect to Redis server
	ctx := context.Background()
	opt, conn_err := redis.ParseURL("redis://:@localhost:6379/0")
	if conn_err != nil {
		panic(conn_err)
	}
	rdb := redis.NewClient(opt)

	var cache Cache

	if err := c.BindJSON(&cache); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err1 := rdb.Set(ctx, cache.Key, cache.Value, 0).Err()
	if err1 != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Unable to cache data"})
		return
	}

	c.JSON(http.StatusCreated, cache)
}

// delete cache
func DeleteCache(c *gin.Context) {
	key := c.Param("key")

	// Connect to Redis server
	ctx := context.Background()
	opt, conn_err := redis.ParseURL("redis://:@localhost:6379/0")
	if conn_err != nil {
		panic(conn_err)
	}
	rdb := redis.NewClient(opt)
	_, err := rdb.Del(ctx, key).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Cache not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cache deleted successfully"})
}
