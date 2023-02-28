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
var verbose bool
var address string

var username string
var password string
var userID string
var token string
var actionType string
var videoID string

var task *boomer.Task
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

func workerDo(request *http.Request) {
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
			_, err := io.Copy(io.Discard, response.Body)
			if err != nil {
				log.Printf("%v\n", err)
			}
		}
		err := response.Body.Close()
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
	var responseData loginResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		log.Printf("%v\n", err)
	}
	userID = strconv.Itoa(responseData.UserId)
	token = responseData.Token
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
	workerDo(request)
}

func register() {
	url := "/douyin/user/register/?username=" + username + "&password=" + password
	method := "POST"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	workerDo(request)
}

func login() {
	url := "/douyin/user/login/?username=" + username + "&password=" + password
	method := "POST"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	workerDo(request)
}

func user() {
	url := "/douyin/user/?user_id=" + userID + "&token=" + token
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	workerDo(request)
}

func publishAction() {
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
	workerDo(request)
}

func publishList() {
	url := "/douyin/publish/list/?token=" + token + "&user_id=" + userID
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	workerDo(request)
}

func likeAction() {
	url := "/douyin/favorite/action/?token=" + token + "&video_id=" + videoID + "&action_type=" + actionType
	method := "POST"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	workerDo(request)
}

func likeList() {
	url := "/douyin/favorite/list/?user_id=" + userID + "&token=" + token
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	workerDo(request)
}
func commentAction() {
	url := "/douyin/comment/action/?token=" + token + "&video_id=" + videoID + "&action_type=" + actionType
	method := "POST"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	workerDo(request)
}

func commentList() {
	url := "/douyin/comment/list/?token=" + token + "&video_id=" + videoID
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	workerDo(request)
}

func relationAction() {
	url := "/douyin/relation/action/?token=" + token + "&to_user_id=" + userID + "&action_type=" + actionType
	method := "POST"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	workerDo(request)
}

func followList() {
	url := "/douyin/relation/follow/list/?user_id=" + userID + "&token=" + token
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	workerDo(request)
}

func followerList() {
	url := "/douyin/relation/follower/list/?user_id=" + userID + "&token=" + token
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	workerDo(request)
}

func friendList() {
	url := "/douyin/relation/friend/list/?user_id=" + userID + "&token=" + token
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	workerDo(request)
}

func messageAction() {
	url := "/douyin/message/action/?token=" + token + "&to_user_id=" + userID + "&action_type=" + actionType + "&content="
	method := "POST"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	workerDo(request)
}

func chat() {
	url := "/douyin/message/chat/?token=" + token + "&to_user_id=" + userID
	method := "GET"
	request, err := http.NewRequest(method, address+url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	workerDo(request)
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
		Fn:     publishAction,
	}
	task6 = &boomer.Task{
		Name:   "publish_list",
		Weight: 10,
		Fn:     publishList,
	}
	task7 = &boomer.Task{
		Name:   "like_action",
		Weight: 10,
		Fn:     likeAction,
	}
	task8 = &boomer.Task{
		Name:   "like_listr",
		Weight: 10,
		Fn:     likeList,
	}
	task9 = &boomer.Task{
		Name:   "comment_action",
		Weight: 10,
		Fn:     commentAction,
	}
	task10 = &boomer.Task{
		Name:   "comment_list",
		Weight: 10,
		Fn:     commentList,
	}
	task11 = &boomer.Task{
		Name:   "relation_action",
		Weight: 10,
		Fn:     relationAction,
	}
	task12 = &boomer.Task{
		Name:   "follow_list",
		Weight: 10,
		Fn:     followList,
	}
	task13 = &boomer.Task{
		Name:   "follower_list",
		Weight: 10,
		Fn:     followerList,
	}
	task14 = &boomer.Task{
		Name:   "friend_list",
		Weight: 10,
		Fn:     friendList,
	}
	task15 = &boomer.Task{
		Name:   "message_action",
		Weight: 10,
		Fn:     messageAction,
	}
	task16 = &boomer.Task{
		Name:   "chat",
		Weight: 10,
		Fn:     chat,
	}
}

func main() {
	var taskType string
	flag.StringVar(&taskType, "task", "", "the task you want to test")
	flag.Parse()

	address = "http://43.139.147.169:8060"
	username = "stress_testing"
	password = "stress_testing"
	actionType = "1"
	videoID = "1"

	timeout := 10   //超时时间
	verbose = false //调试开关

	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 2000
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		MaxIdleConnsPerHost: 2000,
		DisableCompression:  false, //disableCompression
		DisableKeepAlives:   false, //disableKeepalive
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
