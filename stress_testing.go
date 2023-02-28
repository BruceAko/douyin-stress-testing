package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
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
var task1 *boomer.Task
var task2 *boomer.Task
var task3 *boomer.Task
var task4 *boomer.Task
var task5 *boomer.Task
var task6 *boomer.Task
var task7 *boomer.Task
var task8 *boomer.Task
var task9 *boomer.Task
var task10 *boomer.Task
var task11 *boomer.Task
var task12 *boomer.Task
var task13 *boomer.Task
var task14 *boomer.Task
var task15 *boomer.Task
var task16 *boomer.Task

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

func preLogin() {
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
	postBody, err := os.ReadFile("dummy.py")
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	request, err := http.NewRequest(method, address+url, bytes.NewBuffer(postBody))
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
	url := "/douyin/message/action/?token=" + token + "&to_user_id=" + user_id + "&action_type=" + action_type + "&content="
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

func initTask() {
	task1 = &boomer.Task{
		Name:   "feed",
		Weight: 10,
		Fn:     feed,
	}
	task2 = &boomer.Task{
		Name:   "register",
		Weight: 10,
		Fn:     register,
	}
	task3 = &boomer.Task{
		Name:   "login",
		Weight: 10,
		Fn:     login,
	}
	task4 = &boomer.Task{
		Name:   "user",
		Weight: 10,
		Fn:     user,
	}
	task5 = &boomer.Task{
		Name:   "publish_action",
		Weight: 10,
		Fn:     publish_action,
	}
	task6 = &boomer.Task{
		Name:   "publish_list",
		Weight: 10,
		Fn:     publish_list,
	}
	task7 = &boomer.Task{
		Name:   "like_action",
		Weight: 10,
		Fn:     like_action,
	}
	task8 = &boomer.Task{
		Name:   "like_listr",
		Weight: 10,
		Fn:     like_list,
	}
	task9 = &boomer.Task{
		Name:   "comment_action",
		Weight: 10,
		Fn:     comment_action,
	}
	task10 = &boomer.Task{
		Name:   "comment_list",
		Weight: 10,
		Fn:     comment_list,
	}
	task11 = &boomer.Task{
		Name:   "relation_action",
		Weight: 10,
		Fn:     relation_action,
	}
	task12 = &boomer.Task{
		Name:   "follow_list",
		Weight: 10,
		Fn:     follow_list,
	}
	task13 = &boomer.Task{
		Name:   "follower_list",
		Weight: 10,
		Fn:     follower_list,
	}
	task14 = &boomer.Task{
		Name:   "friend_list",
		Weight: 10,
		Fn:     friend_list,
	}
	task15 = &boomer.Task{
		Name:   "message_action",
		Weight: 10,
		Fn:     message_action,
	}
	task16 = &boomer.Task{
		Name:   "chat",
		Weight: 10,
		Fn:     chat,
	}
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
	initTask()
	switch taskType {
	case "feed":
		task = task1
	case "register":
		task = task2
	case "login":
		task = task3
	case "user":
		task = task4
	case "publish_action":
		task = task5
	case "publish_list":
		task = task6
	case "like_action":
		task = task7
	case "like_list":
		task = task8
	case "comment_action":
		task = task9
	case "comment_list":
		task = task10
	case "relation_action":
		task = task11
	case "follow_list":
		task = task12
	case "follower_list":
		task = task13
	case "friend_list":
		task = task14
	case "message_action":
		task = task15
	case "chat":
		task = task16
	case "mix":
		preLogin()
		boomer.Run(task1, task2, task3, task4, task5, task6, task7, task8,
			task9, task10, task11, task12, task13, task14, task15, task16)
	default:
		log.Fatalln("Wrong task type.")
	}
	preLogin()
	boomer.Run(task)
}
