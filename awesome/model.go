package awesome

// GitHubFileLinks
type GitHubFileLinks struct {
	Self string `json:"self"`
	Git  string `json:"git"`
	Html string `json:"html"`
}

// GitHubFileInfo, github content api json response model
type GitHubFileInfo struct {
	Name            string          `json:"name"`
	Path            string          `json:"path"`
	Sha             string          `json:"sha"`
	Size            int             `json:"size"`
	Url             string          `json:"url"`
	HtmlUrl         string          `json:"html_url"`
	DownloadUrl     string          `json:"download_url"`
	Type            string          `json:"type"`
	Content         string          `json:"content"`
	Encoding        string          `json:"encoding"`
	Links           GitHubFileLinks `json:"_links"`
	MarkdownContent string          `json:"markdown_content"` // content 解码,增加星数后的markdown内容
}

// GitHubRepoInfo
type GitHubRepoInfo struct {
	Description string `json:"description"`
	Star        int    `json:"stargazers_count"`
	Archived    bool   `json:"archived"`
}
