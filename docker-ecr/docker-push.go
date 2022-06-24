package main

import (
	"context"
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// Global Variables used in the script
var ecrRegistry string
var ecrRegistryName string
var ecrRegistryToken string
var osSchema string //defines the schemafile based on the Operating System (linux or windows)
var ecrRegion string

//Struct for ImageIds
type ImageIds struct {
	ImageDigest string `json:"imageDigest"`
	ImageTag    string `json:"imageTag"`
}

//Struct for ImageName
type ImageName struct {
	ImageIds []ImageIds `json:"imageIds"`
}

// Init function to intialize the values
func init() {
	ecrRegistry = os.Getenv("ECR_REGISTRY")
	ecrRegistryName = os.Getenv("ECR_REGISTRY_NAME")
	ecrRegion = os.Getenv("ECR_REGION")
	//ECR_GetAuthorizationToken() gets key value pair of the token and AWS user in the format [AWS: eyyy] the below split will just get the token value
	ecrRegistryToken = strings.Split(ECR_GetAuthorizationToken(), ":")[1]
	osSchema = os.Getenv("Schemafile")
}

func main() {

	// Below func will list all the images from ECR into ECR.json file
	ListECRImages()

	// Reads the Docker Images need to be updated schema json File
	file := readJsonFile(osSchema)

	// Reads ECR JSON File
	ecrfile := readJsonFile("ecr.json")

	// Defining the Struct for Images that need to pushed
	data := ImageName{}

	// Defining the Struct for ECR-Data
	ecrdata := ImageName{}

	// unmarshal the JSON files
	_ = json.Unmarshal([]byte(file), &data)
	_ = json.Unmarshal([]byte(ecrfile), &ecrdata)

	// Adding the digest to slice of string
	ecrdigest := retrundigests(ecrdata)

	// Digests not present in the registry
	var digestnotpresent []string

	for _, v := range data.ImageIds {
		if !contains(ecrdigest, v.ImageDigest) {
			digestnotpresent = append(digestnotpresent, v.ImageDigest)
		}

	}

	// An Empty list will determine no new images need to be pushed
	if len(digestnotpresent) == 0 {

		fmt.Println("No new images to be pushed to the registry")
	}

	//Push Images that are not present
	for _, v := range digestnotpresent {

		fmt.Println("Pushing the Image with SHA", v)

		for j := 0; j < len(data.ImageIds); j++ {

			if v == data.ImageIds[j].ImageDigest {

				imagePull(data.ImageIds[j].ImageTag + "@" + data.ImageIds[j].ImageDigest)
				imageTag(data.ImageIds[j].ImageTag+"@"+data.ImageIds[j].ImageDigest, ecrRegistry+":"+strings.ReplaceAll(data.ImageIds[j].ImageTag, ":", "-"))
				imagePush(ecrRegistry + ":" + strings.ReplaceAll(data.ImageIds[j].ImageTag, ":", "-"))
			}

		}

	}

	rm := exec.Command("rm", "ecr.json")
	rm.Run()

}

// Function to verify if a digest is present in the registry
func contains(ecrdigest []string, digest string) bool {

	for _, v := range ecrdigest {
		if v == digest {
			return true
		}
	}
	return false
}

// Function that returns a slice of string digests from a give struct
func retrundigests(d ImageName) []string {
	var ds []string
	for j := 0; j < len(d.ImageIds); j++ {
		ds = append(ds, d.ImageIds[j].ImageDigest)
	}

	return ds

}

//Function to pull Image from the Docker Registry
func imagePull(s string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	out, err := cli.ImagePull(ctx, s, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer out.Close()

	io.Copy(os.Stdout, out)
}

//Function to push an image to the Docker Registry
func imagePush(s string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	authConfig := types.AuthConfig{
		Username: "AWS",
		Password: ecrRegistryToken,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		panic(err)
	}

	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	out, err := cli.ImagePush(ctx, s, types.ImagePushOptions{RegistryAuth: authStr})
	if err != nil {
		panic(err)
	}

	defer out.Close()

	io.Copy(os.Stdout, out)
}

// Function to tag an image
func imageTag(s string, t string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	err = cli.ImageTag(ctx, s, t)
	if err != nil {
		panic(err)
	}
}

// Func to authenticate to ECR and Pull list of images
func ListECRImages() {

	svc := ecr.New(session.Must(session.NewSession(&aws.Config{Region: aws.String(ecrRegion)})))
	input := &ecr.ListImagesInput{
		RepositoryName: aws.String(ecrRegistryName),
	}

	result, err := svc.ListImages(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecr.ErrCodeServerException:
				fmt.Println(ecr.ErrCodeServerException, aerr.Error())
			case ecr.ErrCodeInvalidParameterException:
				fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
			case ecr.ErrCodeRepositoryNotFoundException:
				fmt.Println(ecr.ErrCodeRepositoryNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	strResult := result.String()
	var (
		c *cue.Context
		v cue.Value
	)

	c = cuecontext.New()

	v = c.CompileString(strResult)

	outfile, err := os.Create("./ecr.cue")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	fmt.Fprintf(outfile, "%v\n", v.Value())

	ecr := exec.Command("cue", "export", "ecr.cue", "-o", "ecr.json")

	err = ecr.Run()
	if err != nil {
		fmt.Println("Cannot generate the ECR json file")
		os.Exit(1)
	}

}

//Function to return the byte slice for the read json
func readJsonFile(filename string) []byte {

	bytefile, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error")
	}

	return bytefile

}

//Function to get the Auth Token
func ECR_GetAuthorizationToken() (r string) {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := ecr.New(sess)

	params := &ecr.GetAuthorizationTokenInput{}
	resp, err := svc.GetAuthorizationToken(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	str1, _ := base64.StdEncoding.DecodeString(*resp.AuthorizationData[0].AuthorizationToken)

	return string(str1)

}
