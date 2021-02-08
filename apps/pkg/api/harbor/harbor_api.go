package harbor

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
	"yala/pkg/config"
)

type LabelResponse []struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Color        string    `json:"color"`
	Scope        string    `json:"scope"`
	ProjectID    int       `json:"project_id"`
	CreationTime time.Time `json:"creation_time"`
	UpdateTime   time.Time `json:"update_time"`
	Deleted      bool      `json:"deleted"`
}

type ProjectsResponse []struct {
	ProjectID          int           `json:"project_id"`
	OwnerID            int           `json:"owner_id"`
	Name               string        `json:"name"`
	CreationTime       time.Time     `json:"creation_time"`
	UpdateTime         time.Time     `json:"update_time"`
	Deleted            bool          `json:"deleted"`
	OwnerName          string        `json:"owner_name"`
	CurrentUserRoleID  int           `json:"current_user_role_id"`
	CurrentUserRoleIds []interface{} `json:"current_user_role_ids"`
	RepoCount          int           `json:"repo_count"`
	ChartCount         int           `json:"chart_count"`
	Metadata           struct {
		AutoScan             string `json:"auto_scan"`
		EnableContentTrust   string `json:"enable_content_trust"`
		PreventVul           string `json:"prevent_vul"`
		Public               string `json:"public"`
		ReuseSysCveWhitelist string `json:"reuse_sys_cve_whitelist"`
		Severity             string `json:"severity"`
	} `json:"metadata"`
	CveWhitelist struct {
		ID           int         `json:"id"`
		ProjectID    int         `json:"project_id"`
		Items        interface{} `json:"items"`
		CreationTime time.Time   `json:"creation_time"`
		UpdateTime   time.Time   `json:"update_time"`
	} `json:"cve_whitelist"`
}

type RepositoriesResponse []struct {
	ArtifactCount int       `json:"artifact_count"`
	CreationTime  time.Time `json:"creation_time"`
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	ProjectID     int       `json:"project_id"`
	PullCount     int       `json:"pull_count,omitempty"`
	UpdateTime    time.Time `json:"update_time"`
}

type ImageResponse []struct {
	AdditionLinks struct {
		BuildHistory struct {
			Absolute bool   `json:"absolute"`
			Href     string `json:"href"`
		} `json:"build_history"`
		Vulnerabilities struct {
			Absolute bool   `json:"absolute"`
			Href     string `json:"href"`
		} `json:"vulnerabilities"`
	} `json:"addition_links"`
	Digest     string `json:"digest"`
	ExtraAttrs struct {
		Architecture string      `json:"architecture"`
		Author       interface{} `json:"author"`
		Created      time.Time   `json:"created"`
		Os           string      `json:"os"`
	} `json:"extra_attrs"`
	ID     int `json:"id"`
	Labels []struct {
		CreationTime time.Time `json:"creation_time"`
		Description  string    `json:"description,omitempty"`
		ID           int       `json:"id"`
		Name         string    `json:"name"`
		ProjectID    int       `json:"project_id,omitempty"`
		Scope        string    `json:"scope"`
		UpdateTime   time.Time `json:"update_time"`
		Color        string    `json:"color,omitempty"`
	} `json:"labels"`
	ManifestMediaType string      `json:"manifest_media_type"`
	MediaType         string      `json:"media_type"`
	ProjectID         int         `json:"project_id"`
	PullTime          time.Time   `json:"pull_time"`
	PushTime          time.Time   `json:"push_time"`
	References        interface{} `json:"references"`
	RepositoryID      int         `json:"repository_id"`
	ScanOverview      struct {
		ApplicationVndScannerAdapterVulnReportHarborJSONVersion11 struct {
			CompletePercent int       `json:"complete_percent"`
			Duration        int       `json:"duration"`
			EndTime         time.Time `json:"end_time"`
			ReportID        string    `json:"report_id"`
			ScanStatus      string    `json:"scan_status"`
			Severity        string    `json:"severity"`
			StartTime       time.Time `json:"start_time"`
			Summary         struct {
			} `json:"summary"`
		} `json:"application/vnd.scanner.adapter.vuln.report.harbor+json; version=1.1"`
	} `json:"scan_overview"`
	Size int `json:"size"`
	Tags []struct {
		ArtifactID   int       `json:"artifact_id"`
		ID           int       `json:"id"`
		Immutable    bool      `json:"immutable"`
		Name         string    `json:"name"`
		PullTime     time.Time `json:"pull_time"`
		PushTime     time.Time `json:"push_time"`
		RepositoryID int       `json:"repository_id"`
		Signed       bool      `json:"signed"`
	} `json:"tags"`
	Type string `json:"type"`
}

func GetImagesByLabels(harborAuth *config.HarborAuth, harborHost, projectName string, labelMap map[string]string) []string {
	baseURL := "https://" + harborHost + "/api/v2.0"
	log.Info("Fetching all images for the project ", projectName, " with labels ", labelMap)
	var imageList []string
	for labelName, labelScope := range labelMap {
		log.Info("Fetching images for project ", projectName, " with label ", labelName)
		imageList = filterImagesByLabels(harborAuth, harborHost, baseURL, projectName, labelName, labelScope)
		if len(imageList) == 0 {
			log.Info("No images found for project ", projectName, " with label ", labelName, " found under scope ", labelScope)
		}
		for _, image := range imageList {
			log.Info("Images with the given label ", harborHost+"/"+projectName+"/"+image)
		}
	}
	return imageList
}

func filterImagesByLabels(harborAuth *config.HarborAuth, harborHost, baseURL, projectName, labelName, labelScope string) []string {
	projectId := getProjectIdByProjectName(harborAuth, baseURL+"/projects?name="+projectName)
	if projectId == 0 {
		log.Fatal("Project ID not found for Project ", projectName)
	}
	labelId := getLabelIdByLabelName(harborAuth, baseURL+"/labels?name="+labelName+"&scope="+labelScope+"&project_id="+strconv.Itoa(projectId))
	if labelId == 0 {
		return []string{}
	}
	labelIdStr := strconv.Itoa(labelId)
	repoList := getAllReposByProjectName(harborAuth, baseURL+"/projects/"+projectName+"/repositories")
	var images []string
	for _, repo := range repoList {
		repo = strings.Replace(repo, projectName+"/", "", -1)
		images = append(images, getimagesByLabelId(harborAuth, harborHost, baseURL+"/projects/"+projectName+"/repositories/"+repo+"/artifacts?with_tag=true&with_scan_overview=true&with_label=true&q=labels=("+labelIdStr+")", repo, projectName)...)
	}
	return images
}

func getAllReposByProjectName(harborAuth *config.HarborAuth, API string) []string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", API, nil)
	req.SetBasicAuth(harborAuth.Username, harborAuth.Password)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal("Error calling API getAll repos by project name ", err)
	}
	responseData, _ := ioutil.ReadAll(response.Body)
	if string(responseData) == "null" {
		log.Fatal("Null response, please check your credentials!")
	}
	var responseObject RepositoriesResponse
	json.Unmarshal(responseData, &responseObject)
	var repoList []string
	for _, respObj := range responseObject {
		repoList = append(repoList, respObj.Name)
	}
	return repoList
}

func getLabelIdByLabelName(harborAuth *config.HarborAuth, API string) int {
	client := &http.Client{}
	req, err := http.NewRequest("GET", API, nil)
	req.SetBasicAuth(harborAuth.Username, harborAuth.Password)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal("Error calling API getAll repos by project name ", err)
	}
	responseData, _ := ioutil.ReadAll(response.Body)
	if string(responseData) == "null" {
		log.Fatal("Null response, please check your credentials!")
	}
	var responseObject LabelResponse
	json.Unmarshal(responseData, &responseObject)
	if len(responseObject) == 0 {
		return 0
	}
	labelId := responseObject[0].ID
	return labelId
}

func getimagesByLabelId(harborAuth *config.HarborAuth, harborHost, API, repo, projectName string) []string {
	var imageList []string
	client := &http.Client{}
	req, err := http.NewRequest("GET", API, nil)
	req.SetBasicAuth(harborAuth.Username, harborAuth.Password)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal("Error calling API getAll repos by project name ", err)
	}
	responseData, _ := ioutil.ReadAll(response.Body)
	if string(responseData) == "null" {
		log.Fatal("Null response, please check your credentials!")
	}
	var responseObject ImageResponse
	json.Unmarshal(responseData, &responseObject)
	for _, respObj := range responseObject {
		tags := respObj.Tags
		for _, tag := range tags {
			imageList = append(imageList, harborHost+"/"+projectName+"/"+repo+":"+tag.Name)
		}
	}
	return imageList
}

func getProjectIdByProjectName(harborAuth *config.HarborAuth, API string) int {
	client := &http.Client{}
	req, err := http.NewRequest("GET", API, nil)
	req.SetBasicAuth(harborAuth.Username, harborAuth.Password)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	responseData, _ := ioutil.ReadAll(response.Body)
	if string(responseData) == "null" {
		log.Fatal("Null response, please check your credentials!")
	}
	var responseObject ProjectsResponse
	json.Unmarshal(responseData, &responseObject)
	if len(responseObject) == 0 {
		return 0
	}
	return responseObject[0].ProjectID
}
