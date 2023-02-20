package main

import (
	"crypto/tls"
	"flag"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/myzhan/boomer"
)

var client *http.Client
var address string
var verbose bool
var timeout int
var disableCompression bool
var disableKeepalive bool
var taskType string
var task *boomer.Task

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
	url := "/douyin/user/register/?username=&password="
	method := "POST"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func login() {
	url := "/douyin/user/login/?username=&password="
	method := "POST"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func user() {
	url := "/douyin/user/?user_id=&token="
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
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func publish_list() {
	url := "/douyin/publish/list/?token=&user_id="
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func like_action() {
	url := "/douyin/favorite/action/?token=&video_id=&action_type="
	method := "POST"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func like_list() {
	url := "/douyin/favorite/list/?user_id=&token="
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}
func comment_action() {
	url := "/douyin/comment/action/?token=&video_id=&action_type="
	method := "POST"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func comment_list() {
	url := "/douyin/comment/list/?token=&video_id="
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func relation_action() {
	url := "/douyin/relation/action/?token=&to_user_id=&action_type="
	method := "POST"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func follow_list() {
	url := "/douyin/relation/follow/list/?user_id=&token="
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func follower_list() {
	url := "/douyin/relation/follower/list/?user_id=&token="
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func friend_list() {
	url := "/douyin/relation/friend/list/?user_id=&token="
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func message_action() {
	url := "/douyin/message/action/?token=&to_user_id=&action_type=&content="
	method := "POST"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func chat() {
	url := "/douyin/message/chat/?token=&to_user_id="
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	worker_do(request)
}

func main() {
	flag.StringVar(&taskType, "task-type", "", "the task you want to test")
	flag.Parse()
	address = "43.139.147.169"
	timeout = 10
	disableCompression = false
	disableKeepalive = false
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

	//登录，获取token

	boomer.Run(task)
}
