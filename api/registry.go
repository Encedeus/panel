package api

import (
	"encoding/json"
	"fmt"
	"github.com/Encedeus/panel/services"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"
)

type getPluginResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	OwnerName string `json:"ownerName"`
	Source    struct {
		RepoUri string `json:"repoUri"`
	} `json:"source"`
	Releases []struct {
		Name             string    `json:"name"`
		GithubReleaseTag string    `json:"githubReleaseTag"`
		PublishedAt      time.Time `json:"publishedAt"`
		DownloadURI      string    `json:"DownloadURI"`
	} `json:"releases"`
}

func GetLatestReleaseDownloadURI(pluginId string) (name string, uri string, err error) {
	url := fmt.Sprintf("http://127.0.0.1:3001/plugin/%s", pluginId)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Error(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error(err)
	}
	defer res.Body.Close()

	var body getPluginResponse
	json.NewDecoder(res.Body).Decode(&body)

	if res.StatusCode != 200 {
		return "", "", services.ErrApiFailure
	}

	if len(body.Releases) == 0 {
		return "", "", services.ErrModuleHasNoReleases
	}

	return body.Name, body.Releases[0].DownloadURI, nil
}
