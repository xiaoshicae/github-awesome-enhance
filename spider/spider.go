package spider

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github-awesome-enhance/awesome"
	"github-awesome-enhance/util/requests"
	"log"
	"os"
	"regexp"
	"strings"
)

// FetchGithubFileContent
func FetchGithubFileContent(repo string, file string) string {
	url := fmt.Sprintf("https://api.github.com/repos/%s/contents/%s", repo, file)

	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	if accessToken == "" {
		panic("miss env GITHUB_ACCESS_TOKEN")
	}
	params := map[string]string{
		"access_token": accessToken, // access_token, 增大接口调用频率限制
	}
	resp := requests.GET(url, params, nil, 0)

	gitHubFileInfo := awesome.GitHubFileInfo{}
	err := json.Unmarshal(resp, &gitHubFileInfo)
	if err != nil {
		panic(err)
	}

	// decode by base64
	decodeBytes, err := base64.StdEncoding.DecodeString(gitHubFileInfo.Content)
	if err != nil {
		log.Fatalln(err)
	}

	return string(decodeBytes)
}

// FetchGitHubRepoInfo
func FetchGitHubRepoInfo(repo string) awesome.GitHubRepoInfo {
	url := fmt.Sprintf("https://api.github.com/repos/%s", repo)

	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	if accessToken == "" {
		panic("miss env GITHUB_ACCESS_TOKEN")
	}
	params := map[string]string{
		"access_token": accessToken,
	}
	resp := requests.GET(url, params, nil, 0)

	gitHubRepoInfo := awesome.GitHubRepoInfo{}
	err := json.Unmarshal(resp, &gitHubRepoInfo)
	if err != nil {
		panic(err)
	}

	return gitHubRepoInfo
}

func ParseFileContent(fileContent string) string {
	markdownContent := ""
	re := regexp.MustCompile(`\(https://github\.com/(.*?)/(.*?)\)`)

	lines := strings.Split(fileContent, "\n")
	for _, line := range lines {
		matched := re.FindAllStringSubmatch(line, -1)
		if len(matched) > 0 {
			// 可能存在多匹配,只取第一个
			for _, match := range matched[:1] {
				username, repo := match[1], match[2]
				repoLength := len(repo)
				if repoLength > 0 && repo[repoLength-1] == '/' {
					repo = repo[:repoLength-2]
				}

				if !strings.Contains(repo, "/") {
					fmt.Printf("user is: %s, repo is: %s\n", username, repo)
					repoInfo := FetchGitHubRepoInfo(username + "/" + repo)
					line = fmt.Sprintf("%s 【🌟Star %d】\n", line, repoInfo.Star)
				}
			}
		}

		markdownContent += fmt.Sprintf("%s\n", line)
		fmt.Println(line)
	}

	return markdownContent
}
