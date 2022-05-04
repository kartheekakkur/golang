package main

import (
	"encoding/json"
	"fmt"
	"os"
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

type ImageIds struct {
	ImageDigest string `json:"imageDigest"`
	ImageTag    string `json:"imageTag"`
}

type ImageName struct {
	ImageIds []ImageIds `json:"ImageIds"`
}


func main() {

	svc := ecr.New(session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")})))
    input := &ecr.ListImagesInput{
    RepositoryName: aws.String("docker-io-cache"),
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

// fmt.Println(result)


jsonresult := result.String()
var (
	c *cue.Context
	v cue.Value
)

c = cuecontext.New()


// compile some CUE into a Value
v = c.CompileString(jsonresult)

bs,_ := json.Marshal(jsonresult)
data := ImageName{}

_ = json.Unmarshal([]byte(bs), &data)

for i, v := range data.ImageIds {
   fmt.Println(v.ImageDigest[i])

}


outfile, err := os.Create("./ecr.cue")
if err != nil {
	panic(err)
}
defer outfile.Close()
fmt.Fprintf(outfile, "%v\n", v.Value())


}