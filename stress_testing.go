package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/myzhan/boomer"
)

var client *http.Client
var address string
var verbose bool
var taskType string
var task *boomer.Task
var username string
var password string
var user_id string
var token string
var action_type string
var video_id string

func worker_do(request *http.Request) {
	startTime := time.Now()
	response, err := client.Do(request)
	elapsed := time.Since(startTime)
	if err != nil {
		if verbose {
			log.Printf("%v\n", err)
		}
		boomer.RecordFailure("http", "error", 0.0, err.Error())
	} else {
		boomer.RecordSuccess("http", strconv.Itoa(response.StatusCode),
			elapsed.Nanoseconds()/int64(time.Millisecond), response.ContentLength)
		if verbose {
			body, err := io.ReadAll(response.Body)
			if err != nil {
				log.Printf("%v\n", err)
			} else {
				log.Printf("Status Code: %d\n", response.StatusCode)
				log.Println(string(body))
			}
		} else {
			io.Copy(io.Discard, response.Body)
		}
		response.Body.Close()
		if err != nil {
			log.Printf("%v\n", err)
		}
	}
}

func feed() {
	url := "/douyin/feed/"
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func register() {
	url := "/douyin/user/register/?username=" + username + "&password=" + password
	method := "POST"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func login() {
	url := "/douyin/user/login/?username=" + username + "&password=" + password
	method := "POST"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func user() {
	url := "/douyin/user/?user_id=" + user_id + "&token=" + token
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func publish_action() {
	url := "/douyin/publish/action/"
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open("dummy.py")
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("data", filepath.Base(""))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}
	_ = writer.WriteField("token", "")
	_ = writer.WriteField("title", "")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	request, err := http.NewRequest(method, address+url, payload)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func publish_list() {
	url := "/douyin/publish/list/?token=" + token + "&user_id=" + user_id
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func like_action() {
	url := "/douyin/favorite/action/?token=" + token + "&video_id=" + video_id + "&action_type=" + action_type
	method := "POST"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func like_list() {
	url := "/douyin/favorite/list/?user_id=" + user_id + "&token=" + token
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}
func comment_action() {
	url := "/douyin/comment/action/?token=" + token + "&video_id=" + video_id + "&action_type=" + action_type
	method := "POST"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func comment_list() {
	url := "/douyin/comment/list/?token=" + token + "&video_id=" + video_id
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func relation_action() {
	url := "/douyin/relation/action/?token=" + token + "&to_user_id=" + user_id + "&action_type=" + action_type
	method := "POST"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func follow_list() {
	url := "/douyin/relation/follow/list/?user_id=" + user_id + "&token=" + token
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func follower_list() {
	url := "/douyin/relation/follower/list/?user_id=" + user_id + "&token=" + token
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func friend_list() {
	url := "/douyin/relation/friend/list/?user_id=" + user_id + "&token=" + token
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func message_action() {
	url := "/douyin/message/action/?token=&to_user_id=" + "2" + "&action_type=" + action_type + "&content=test"
	method := "POST"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func chat() {
	url := "/douyin/message/chat/?token=" + token + "&to_user_id=" + user_id
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func main() {
	flag.StringVar(&taskType, "task", "", "the task you want to test")
	flag.Parse()
	address = "http://43.139.147.169:8070"
	username = "stress_testing"
	password = "stress_testing"
	action_type = "1"
	video_id = "1"
	timeout := 10
	disableCompression := false
	disableKeepalive := false
	verbose = false
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 2000
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		MaxIdleConnsPerHost: 2000,
		DisableCompression:  disableCompression,
		DisableKeepAlives:   disableKeepalive,
	}
	client = &http.Client{
		Transport: tr,
		Timeout:   time.Duration(timeout) * time.Second,
	}
	switch taskType {
	case "feed":
		task = &boomer.Task{
			Name:   "feed",
			Weight: 10,
			Fn:     feed,
		}
	case "register":
		task = &boomer.Task{
			Name:   "register",
			Weight: 10,
			Fn:     register,
		}
	case "login":
		task = &boomer.Task{
			Name:   "login",
			Weight: 10,
			Fn:     login,
		}
	case "user":
		task = &boomer.Task{
			Name:   "user",
			Weight: 10,
			Fn:     user,
		}
	case "publish_action":
		task = &boomer.Task{
			Name:   "publish_action",
			Weight: 10,
			Fn:     publish_action,
		}
	case "publish_list":
		task = &boomer.Task{
			Name:   "publish_list",
			Weight: 10,
			Fn:     publish_list,
		}
	case "like_action":
		task = &boomer.Task{
			Name:   "like_action",
			Weight: 10,
			Fn:     like_action,
		}
	case "like_list":
		task = &boomer.Task{
			Name:   "like_listr",
			Weight: 10,
			Fn:     like_list,
		}
	case "comment_action":
		task = &boomer.Task{
			Name:   "comment_action",
			Weight: 10,
			Fn:     comment_action,
		}
	case "comment_list":
		task = &boomer.Task{
			Name:   "comment_list",
			Weight: 10,
			Fn:     comment_list,
		}
	case "relation_action":
		task = &boomer.Task{
			Name:   "relation_action",
			Weight: 10,
			Fn:     relation_action,
		}
	case "follow_list":
		task = &boomer.Task{
			Name:   "follow_list",
			Weight: 10,
			Fn:     follow_list,
		}
	case "follower_list":
		task = &boomer.Task{
			Name:   "follower_list",
			Weight: 10,
			Fn:     follower_list,
		}
	case "friend_list":
		task = &boomer.Task{
			Name:   "friend_list",
			Weight: 10,
			Fn:     friend_list,
		}
	case "message_action":
		task = &boomer.Task{
			Name:   "message_action",
			Weight: 10,
			Fn:     message_action,
		}
	case "chat":
		task = &boomer.Task{
			Name:   "chat",
			Weight: 10,
			Fn:     chat,
		}
	default:
		log.Fatalln("Wrong task type.")
	}

	//登录，获取token和user_id
	url := "/douyin/user/login/?username=" + username + "&password=" + password
	method := "POST"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Printf("%v\n", err)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("%v\n", err)
	}
	type loginResponse struct {
		StatusCode int    `json:"status_code"`
		StatusMsg  string `json:"status_msg"`
		UserId     int    `json:"user_id"`
		Token      string `json:"token"`
	}
	var login_response loginResponse
	err = json.Unmarshal(body, &login_response)
	if err != nil {
		log.Printf("%v\n", err)
	}
	user_id = strconv.Itoa(login_response.UserId)
	token = login_response.Token
	err = response.Body.Close()
	if err != nil {
		log.Printf("%v\n", err)
	}
	boomer.Run(task)
}
