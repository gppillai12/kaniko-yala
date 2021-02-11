package cmd

import (
	"bufio"
	"bytes"
	"github.com/fatih/color"
	docker "github.com/fsouza/go-dockerclient"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	harborAPI "yala/pkg/api/harbor"
	"yala/pkg/command/runner"
	"yala/pkg/config"
	"yala/pkg/util"
)

const bundleDir string = "bundle"
const dockerImageDir string = "docker-images"
const dockerCmd string = "docker"
const harborLabelScopeGlobal string = "g"
const harborLabelScopeProject string = "p"
const dockerPublic string = "docker.io"

var isSelfHosted bool = false

var bundleCmd = &cobra.Command{
	Use:   "bundle",
	Short: "Bundle all images as tar file.",
	Long:  "Bundle all images as tar file.",
	Run: func(cmd *cobra.Command, args []string) {

		clusterConfig, err := config.NewFromFile(globalOptions.ClusterConfigFile)
		if err != nil {
			log.WithError(err).WithField("filename", globalOptions.ClusterConfigFile).Fatal("Failed to create  config object from config file.")
		}
		var bundlePath string
		var bundleName string
		//Creates directory if doesn't exist
		if clusterConfig != nil {
			bundleVersion := clusterConfig.Version
			bundlePath = globalOptions.HomeDir + "/" + clusterConfig.ClusterName + "/" + bundleDir + "/" + dockerImageDir
			bundleName = bundleDir + "/" + dockerImageDir + "-" + bundleVersion + ".tar.gz"
			if clusterConfig.Docker.Bundle == nil {
				log.Warn("Bundle path not provided, using the default path!")
				log.Info("Using bundlePath ", bundlePath, " For creating the docker images bundle!")
			} else {
				if clusterConfig.Docker.Bundle.Path != "" {
					bundlePath = clusterConfig.Docker.Bundle.Path + "/" + dockerImageDir
				}
				if clusterConfig.Docker.Bundle.Name != "" {
					bundleName = clusterConfig.Docker.Bundle.Name
				}
			}
			if err := util.Mkdir(bundlePath, 0750); err != nil {
				log.Fatal("Error creating bundle directory ", err)
			}
		} else {
			log.Fatal("Nil clusterConfig provided!")
		}

		log.Info("Preparing Bundle.")

		//Create imageList via Labels
		var imagesByLabelsList []string
		projects := clusterConfig.Docker.Projects
		labelMap := make(map[string]string)
		harborHost := clusterConfig.Docker.Harbor.Host
		harborAuth := clusterConfig.Docker.Harbor.Auth
		for _, project := range projects {
			labels := project.Label
			for _, label := range labels {
				switch label.Scope {
				case "project":
					labelMap[label.Name] = harborLabelScopeProject
				case "global":
					labelMap[label.Name] = harborLabelScopeGlobal
				default:
					log.Fatal("Invalid scope selected!\n scope can be either global or project!")
				}
			}
			imagesByLabelsList = append(imagesByLabelsList, harborAPI.GetImagesByLabels(harborAuth, harborHost, project.Name, labelMap)...)
		}

		//Stage-1: Pull all images by labels
		if clusterConfig.Docker.Harbor.Host != dockerPublic {
			isSelfHosted = true
		}
		if clusterConfig.Docker.Harbor.Auth != nil {
			pullImages(&docker.AuthConfiguration{
				Username: clusterConfig.Docker.Harbor.Auth.Username,
				Password: clusterConfig.Docker.Harbor.Auth.Password,
			}, clusterConfig.Docker.Harbor.Host, imagesByLabelsList, isSelfHosted)
		} else {
			log.Info("Proceeding without docker credentials\n! Pulling images by labels list!")
			pullImages(&docker.AuthConfiguration{
				Username: "",
				Password: "",
			}, clusterConfig.Docker.Harbor.Host, imagesByLabelsList, isSelfHosted)
		}
		log.Info("Pulling images by labels list complete!")

		//Stage-2: Pull images from pull list
		var pullImageList []string
		pullObjectList := clusterConfig.Docker.Pull
		au := &docker.AuthConfiguration{
			Username: "",
			Password: "",
		}
		for i, pullObject := range pullObjectList {
			if pullObject.Registry != dockerPublic {
				isSelfHosted = true
			}
			color.Yellow("****** PULLING FROM REGISTRY %s %s%s%s %s ", pullObject.Registry, strconv.Itoa(i+1), "/", strconv.Itoa(len(pullObjectList)), " ******")
			pullImageList = pullObject.Images
			if pullObject.Auth == nil {
				log.Info("Proceeding without docker credentials!")
			} else {
				au = &docker.AuthConfiguration{
					Username: pullObject.Auth.Username,
					Password: pullObject.Auth.Password,
				}
			}
			pullImages(au, pullObject.Registry, pullImageList, isSelfHosted)
		}

		//Stage-3: Build images from build list
		var buildImageList, finalImageList, dockerfileList []string
		buildImageList = getBuildList(clusterConfig.Docker.Build)
		if clusterConfig.Docker.Build != nil {
			dockerfileList = clusterConfig.Docker.Build
			buildImages(dockerfileList)
			finalImageList = append(pullImageList, buildImageList...)
			finalImageList = append(finalImageList, imagesByLabelsList...)
			finalImageList = util.RemoveDuplicateStrings(finalImageList)
		} else {
			finalImageList = append(pullImageList, imagesByLabelsList...)
			finalImageList = util.RemoveDuplicateStrings(finalImageList)
		}
		util.WriteToFile(finalImageList, bundlePath+"/"+dockerImgListFile)
		var imageTarList []string
		for _, imageId := range finalImageList {
			imageTarList = append(imageTarList, tarImage(imageId, bundlePath))
		}
		createBundle(imageTarList, clusterConfig, bundleName, bundleDir)
		color.Green("****** Successfully created the bundle! ******")
	},
}

func getBuildList(dockerfileDirList []string) []string {
	var imgId, file string
	var imgList []string
	for _, dockerDir := range dockerfileDirList {
		file, _ = filepath.Abs(dockerDir)
		imgId = dockerfileParser(file)
		imgList = append(imgList, imgId)
	}
	return imgList
}

func pullImages(auth *docker.AuthConfiguration, registry string, images []string, isSelfHosted bool) {
	dockerClient, _ := docker.NewClientFromEnv()
	for _, imageId := range images {
		imageName, imageTag := util.SplitBeforeAfter(imageId, ":")

		if isSelfHosted {
			dockerPullSelfHosted(auth, imageName, imageTag, dockerClient)
		} else {
			dockerPull(auth, registry, imageName, imageTag, dockerClient)
		}
	}
}

func buildImages(dockerfiles []string) {
	for _, dockerfile := range dockerfiles {
		file, err := filepath.Abs(dockerfile)
		if err != nil {
			log.Fatal("Absolute filepath error for dockerfile ", dockerfile, "\n", err)
		}
		dockerBuildWithCLI(file, dockerfileParser(file))
	}
}

func createBundle(imageTarList []string, clusterConfig *config.ClusterConfig, bundleTarName, bundleDirName string) {
	bundleTarPath := globalOptions.HomeDir + "/" + clusterConfig.ClusterName + "/"
	// Archive to single file
	out, err := os.Create(bundleTarPath + "/" + bundleTarName)
	if err != nil {
		log.Fatalln("Error creating bundle stream :", err)
	}
	defer out.Close()
	if err := util.CreateArchive(imageTarList, out, bundleDirName); err != nil {
		log.Fatal("Error creating final bundle ", err)
	}
}

func tarImage(imageId, bundlePath string) string {
	dockerClient, _ := docker.NewClientFromEnv()
	imageName, imageTag := util.SplitBeforeAfter(imageId, ":")
	if strings.Contains(imageName, "/") {
		ss := strings.Split(imageName, "/")
		imageName = ss[len(ss)-1]
	}
	imageTarName := bundlePath + "/" + imageName + "-" + imageTag + ".tar"
	f, err := os.Create(imageTarName)
	if err != nil {
		log.Fatal("Error creating tar files ", err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	log.Info("Exporting docker image ", imageId)
	opts := docker.ExportImagesOptions{Names: []string{imageId}, OutputStream: w}
	if err := dockerClient.ExportImages(opts); err != nil {
		log.Fatal("Error exporting image ", err)
	}
	w.Flush()
	return imageTarName
}
func dockerPull(auth *docker.AuthConfiguration, registry, imageName, imageTag string, dockerClient *docker.Client) {
	pullOptions := &docker.PullImageOptions{
		Registry:   registry,
		Repository: imageName,
		Tag:        imageTag,
	}
	log.Info("Pulling docker Image ", registry+"/"+imageName+":"+imageTag)
	if err := dockerClient.PullImage(*pullOptions, *auth); err != nil {
		log.Fatal("Error pulling image", err)
	}
}

// Please refer this issue
// https://github.com/fsouza/go-dockerclient/issues/498
func dockerPullSelfHosted(auth *docker.AuthConfiguration, imageName, imageTag string, dockerClient *docker.Client) {
	pullOptions := &docker.PullImageOptions{
		Registry:   "",
		Repository: imageName,
		Tag:        imageTag,
	}
	log.Info("Pulling docker Image ", imageName+":"+imageTag, " from registry ", pullOptions.Registry)
	if err := dockerClient.PullImage(*pullOptions, *auth); err != nil {
		log.Fatal("Error pulling image", err)
	}
}

func dockerBuildWithCLI(dockerfile, imageId string) {
	initialDir, _ := os.Getwd()
	dockerfileDir := strings.Replace(dockerfile, "Dockerfile", "", -1)
	logger := log.StandardLogger()

	if err := os.Chdir(dockerfileDir); err != nil {
		log.Fatal("error changing directory ", err)
	}
	newDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current directory info ", err)
	}
	cmdRunner, err := runner.New(logger.Writer(), logger.Writer(), dockerfileDir)
	cmdRunner.SetDirectory(newDir)
	if err != nil {
		log.Fatal("Error creating new runner", err)
	}

	prov, err := cmdRunner.BinaryExists(dockerCmd)
	if err != nil {
		log.Fatal("error at docker setup ", err)
	}
	if !prov {
		log.Fatal("please install Docker before running the command")
	}
	err = cmdRunner.Run(dockerCmd, "build", "-f", dockerfile, ".", "-t", imageId)
	if err != nil {
		log.Fatal(err, "docker build command failed")
	}
	if err := os.Chdir(initialDir); err != nil {
		log.Fatal("error changing to initial directory ", err)
	}
}

func dockerfileParser(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open ", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()
	var imageId string
	for _, line := range text {
		if strings.Contains(line, "FROM") {
			imageId = strings.TrimSpace(strings.Trim(line, "FROM"))
		}
	}
	return imageId
}

func dockerBuild(dockerFile string, dockerClient *docker.Client) {
	log.Info("Building dockerfile ", dockerFile)
	var buf bytes.Buffer
	dockerAuth := &docker.AuthConfiguration{
		Username: "userName",
		Password: "pass",
	}
	buildOptions := &docker.BuildImageOptions{
		Auth:         *dockerAuth,
		Dockerfile:   "../dockerfiles/dex/Dockerfile",
		OutputStream: &buf,
		ContextDir:   "../dockerfiles/dex/",
	}
	if err := dockerClient.BuildImage(*buildOptions); err != nil {
		log.Fatal("Error Building image ", err)
	}
}
