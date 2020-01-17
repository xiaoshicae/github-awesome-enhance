package main

import (
	"fmt"
	"github-awesome-enhance/spider"
	"github-awesome-enhance/util/file"
	"github-awesome-enhance/util/shell"
	"strings"
	"time"
)

// gitHubPageRepoDir
var GitHubPageRepoDir = "/root/web/github-io-blog"

//var GitHubPageRepoDir = "/Users/zhuangshui/VscodeProjects/FontendProjects/GitHubBlogFontend"

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

// LoadAwesomeContent
func LoadAwesomeContent() {
	// 加载markdown内容
	fmt.Println("load awesome content ...")

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
			}

			if markdownContent != "" {
				// 写入 github page 仓库中
				githubPageFilePath := fmt.Sprintf("%s/public/awsome_readme/%s_README.md", GitHubPageRepoDir, user+"_"+repo)
				file.WriteFile(githubPageFilePath, markdownContent)
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
	LoadAwesomeContent()

	//定时任务
	ticker := time.NewTicker(time.Hour * 8)
	for range ticker.C {
		fmt.Println(fmt.Sprintf("Task begin at %s", time.Now().Format("2006-01-02 15:04:05")))
		LoadAwesomeContent()
	}
}

func main() {
	// 启动定时任务刷新,内存内容
	RunTimedTask()
}
