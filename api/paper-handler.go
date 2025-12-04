package api

import (
	"buku-kotak-api/helper"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type EnvData struct {
	CLOUDINARY_API_NAME,
	CLOUDINARY_API_KEY,
	CLOUDINARY_API_SECRET,
	CLOUDINARY_URL,
	UPSTASH_REDIS_REST_TCP string
}

type JsonData struct {
	Img_base64 string `json:"img_base64"`
	Public_id  string `json:"public_id"`
}

func envData() *EnvData {
	return &EnvData{
		CLOUDINARY_API_NAME:    os.Getenv("CLOUDINARY_API_NAME"),
		CLOUDINARY_API_KEY:     os.Getenv("CLOUDINARY_API_KEY"),
		CLOUDINARY_API_SECRET:  os.Getenv("CLOUDINARY_API_SECRET"),
		CLOUDINARY_URL:         os.Getenv("CLOUDINARY_URL"),
		UPSTASH_REDIS_REST_TCP: os.Getenv("UPSTASH_REDIS_REST_TCP"),
	}
}

func recoverResponse(c *gin.Context) {
	// recover == catch error,
	// must place in defer function
	message := recover()
	if message != nil {
		errorMessage := fmt.Sprint(message)
		fmt.Println(fmt.Printf("error: %s", errorMessage))
		c.JSON(400, helper.HandlerResponse(400, nil, errors.New(errorMessage)))
	}
}

func PaperHistory(c *gin.Context) {
	defer recoverResponse(c)
	envData := envData()

	// initialize redis
	rdbOpt, _ := redis.ParseURL(envData.UPSTASH_REDIS_REST_TCP)
	rdb := redis.NewClient(rdbOpt)
	getHistory, getHistoryErr := rdb.Get(c, "buku_kotak_history").Result()
	if getHistoryErr == redis.Nil || getHistoryErr != nil {
		// key not exist, return empty array
		c.JSON(http.StatusOK, helper.HandlerResponse(200, []any{}, nil))
	} else {
		// Return JSON response
		c.JSON(http.StatusOK, helper.HandlerResponse(200, strings.Split(getHistory, ";"), nil))
	}
}

func PaperUpload(c *gin.Context) {
	defer recoverResponse(c)

	envData := envData()
	var jsonData JsonData
	// return if body empty
	if jsonErr := c.BindJSON(&jsonData); jsonErr != nil {
		panic("body cannot be empty")
	}

	// set cloudinary
	cld, _ := cloudinary.NewFromParams(envData.CLOUDINARY_API_NAME, envData.CLOUDINARY_API_KEY, envData.CLOUDINARY_API_SECRET)
	// upload image
	uploadRes, uploadErr := cld.Upload.Upload(c,
		jsonData.Img_base64,
		uploader.UploadParams{
			PublicID:       jsonData.Public_id,
			Folder:         "buku-kotak",
			UploadPreset:   "ml_default",
			AllowedFormats: []string{"jpg", "jpeg", "png"},
		})
	// check upload error
	if uploadErr != nil {
		panic("cloudinary, " + uploadErr.Error())
	}
	imgFormatErr, _ := regexp.MatchString(`file format .* not allowed`, uploadRes.Error.Message)
	if imgFormatErr {
		panic("cloudinary, " + uploadRes.Error.Message)
	}

	// initialize redis
	rdbOpt, _ := redis.ParseURL(envData.UPSTASH_REDIS_REST_TCP)
	rdb := redis.NewClient(rdbOpt)

	// save uploaded image to redis
	getHistory, getHistoryErr := rdb.Get(c, "buku_kotak_history").Result()
	if getHistoryErr == redis.Nil || getHistoryErr != nil {
		// if EOF, set new value to redis
		rdb.Set(c, "buku_kotak_history", uploadRes.SecureURL, 0)
	} else {
		// else, push new value to redis
		concatHistory := getHistory + ";" + uploadRes.SecureURL
		rdb.Set(c, "buku_kotak_history", concatHistory, 0)
	}

	// Return JSON response
	c.JSON(http.StatusOK, helper.HandlerResponse(200, []any{uploadRes.PublicID}, nil))
}
