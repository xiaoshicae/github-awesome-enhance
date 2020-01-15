package main

import (
	"encoding/json"
	"fmt"
	"github-awesome-enhance/spider"
	"github-awesome-enhance/util/file"
	"github-awesome-enhance/util/shell"
	"net/http"
	"strings"
	"time"
)

// gitHubPageRepoDir
var GitHubPageRepoDir = "/root/web/github-io-blog"

// github awesome repos
var AwesomeRepos = []string{
	//"sindresorhus/awesome",
	"avelino/awesome-go",
	"vinta/awesome-python",
	"vuejs/awesome-vue",
	"akullpp/awesome-java",
	"enaqx/awesome-react",
	"justjavac/awesome-wechat-weapp",
}

// AwesomeRepoReadmeMap
var AwesomeRepoReadmeMap = map[string]string{}

// LoadAwesomeContent
func LoadAwesomeContent() {
	// 加载markdown内容

	go func() {
		for _, userRepo := range AwesomeRepos {
			s := strings.Split(userRepo, "/")
			user, repo := s[0], s[1]
			filePath := fmt.Sprintf("./awesome/readme/%s_README.md", user+"_"+repo)

			// 1. 从本地加载markdown内容
			markdownContent, ok := file.ReadFile(filePath)

			if ok {
				_ = fmt.Sprintf("load content form %s", filePath)
			} else {
				// 2. 本地文件过期,从远程下载,并写入本地
				fileContent := spider.FetchGithubFileContent(userRepo, "README.md")

				markdownContent = spider.ParseFileContent(fileContent)

				file.WriteFile(filePath, markdownContent)

				// 同时写入github page 仓库中
				githubPageFilePath := fmt.Sprintf("%s/public/awsome_readme/%s_README.md", GitHubPageRepoDir, user+"_"+repo)
				file.WriteFile(githubPageFilePath, markdownContent)
			}

			if markdownContent != "" {
				AwesomeRepoReadmeMap[userRepo] = markdownContent
			}
		}

		// 发布github page
		deployGitHubPages()
	}()
}

func deployGitHubPages() {
	cmd := fmt.Sprintf("cd %s && sh %s/deploy.sh", GitHubPageRepoDir, GitHubPageRepoDir)
	fmt.Printf("cmd:  %s \n", cmd)
	resp := shell.RunCmd(cmd)
	fmt.Println(resp)
}

// RunTimedTask
func RunTimedTask() {
	fmt.Println("Starting timed  task ...")

	//定时任务
	ticker := time.NewTicker(time.Hour * 12)
	for range ticker.C {
		fmt.Println(fmt.Sprintf("Task begin at %s", time.Now().Format("2006-01-02 15:04:05")))
		LoadAwesomeContent()
	}
}

// RunHttpServer
func RunHttpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		js, err := json.Marshal(AwesomeRepoReadmeMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Set("content-type", "application/json")             //返回数据格式是json
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(js)
	})

	fmt.Println("Starting http server ...")

	_ = http.ListenAndServe(":5000", nil)
}

func main() {
	// 加载awesome内存
	LoadAwesomeContent()

	// 启动定时任务刷新,内存内容
	RunTimedTask()

	// 启动http服务对外提供接口
	//RunHttpServer()
}
