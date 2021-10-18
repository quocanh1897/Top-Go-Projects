package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
	"unicode"
)

// Repo describes a Github repository with additional field, last commit date
type Repo struct {
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	DefaultBranch  string    `json:"default_branch"`
	Stars          int       `json:"stargazers_count"`
	Forks          int       `json:"forks_count"`
	Issues         int       `json:"open_issues_count"`
	Created        time.Time `json:"created_at"`
	Updated        time.Time `json:"updated_at"`
	URL            string    `json:"html_url"`
	LastCommitDate time.Time `json:"-"`
}

const (
	head = `# TOP Go Projects

- [Command Line](#command-line)
- [Console UI](#console-ui)
	- [Console UI Engine/Library](#console-ui-enginelibrary)
- [Web Frameworks](#web-frameworks)
- [Game](#game)
	- [Game engine](#game-engine)
- [GUI Application](#gui-application)
- [Logging](#logging)
- [Object-Relational Mapping](#object-relational-mapping)
- [Go E-Books](#go-e-books)
`
	headerTable = `
| Repo | Stars  |  Forks  |  Description |
| ---- | :----: | :-----: | ------------ |
`
	footer = "\n*Last Update: %v*\n"
)

var (
	repos []Repo
)

func main() {
	accessToken := getAccessToken()

	if accessToken == "" {
		fmt.Errorf("Please provide valid access token")
		os.Exit(1)
	}

	// bar := progressbar.New(141)
	// bar.RenderBlank()
	writeTitle()

	byteContents, err := ioutil.ReadFile("list.repo")
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(byteContents), "\n")
	for _, url := range lines {
		var repo Repo

		if strings.HasPrefix(url, "##") {
			header := Repo{
				Name:        "",
				Description: url,
			}
			repos = append(repos, header)
		}

		idx := strings.Index(url, "https://github.com/")
		if idx != -1 {
			fmt.Println(url[idx:])

			url := fmt.Sprintf("https://api.github.com/repos/%s", url[idx+19:])

			client := &http.Client{
				Timeout: time.Second * 30,
			}
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				fmt.Println(err)
			}
			token := "token " + accessToken

			req.Header.Set("Authorization", token)
			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
			}
			if res.StatusCode != 200 {
				fmt.Println(res.Status)
			}

			decoder := json.NewDecoder(res.Body)
			if err = decoder.Decode(&repo); err != nil {
				fmt.Println(err)
			}
			repos = append(repos, repo)

			defer res.Body.Close()
		}

		if len(url) <= 1 {
			sort.Slice(repos, func(i, j int) bool {
				return repos[i].Stars > repos[j].Stars
			})
			saveRanking(repos)
			repos = nil
		}
	}

	writeFooter()
}

func trimSpaceAndSlash(r rune) bool {
	return unicode.IsSpace(r) || (r == rune('/'))
}

func getAccessToken() string {
	tokenBytes, err := ioutil.ReadFile("access-token.tok")
	if err != nil {
		fmt.Println("Try get accessToken from ENV OAuth ...")
		token := os.Getenv("OAuth")
		return token
	}
	return strings.TrimSpace(string(tokenBytes))
}

func writeTitle() {
	readme, err := os.OpenFile("README2.md", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
	}

	readme.WriteString(head)
	readme.Close()
}

func saveRanking(repos []Repo) {
	readme, err := os.OpenFile("README2.md", os.O_RDWR|os.O_APPEND, 0666)
	defer readme.Close()
	if err != nil {
		fmt.Println(err)
	}
	for _, repo := range repos {
		if len(repo.Name) == 0 {
			readme.WriteString(fmt.Sprintf("\n%s\n%s", repo.Description, headerTable))
		} else {
			readme.WriteString(fmt.Sprintf("| [%s](%s) | **%d** | **%d**  | %s |\n", repo.Name, repo.URL, repo.Stars, repo.Forks, repo.Description))
		}
	}

}

func writeFooter() {
	readme, err := os.OpenFile("README2.md", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}

	readme.WriteString(fmt.Sprintf(footer, time.Now().Format(time.RFC3339)))
	readme.Close()
}
