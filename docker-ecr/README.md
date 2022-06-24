# Push Images from Docker Hub to ECR

This script is used to pull public images from docker hub and then push it to AWS Elastic Container Registry if an image with that specific SHA256 does not exist.


### Prerequisite to run locally on MAC

```
brew install docker
```

```
brew install cue-lang/tap/cue
```
```
brew install awscli
```
Install GO

https://go.dev/doc/install version >= 1.18

Set up following environment variables:

        # URI of the Registry
        export ECR_REGISTRY=XXXXXX
        # Name of the Registry
        export ECR_REGISTRY_NAME=XXXXX
        # Region Where it is running
        export ECR_REGION=xxx
        # Schema file for the linux config
        export Schemafile=schema-linux.json


Login to the AWS account from the CLI and make sure you have access to the ECR this will generate the token that will be used later to list and push images in ECR.


```
go build docker-push.go
```

```
chmod +x docker-push
```

Create your Schema-linux.json file with the imagedigst and image tag in the following manner.

```JSON
{
    "imageIds": [
        {
            "imageDigest": "sha256:af6de82006b8c1296ec539f3cb628956006d0df6ac6b93cb484c1c030a148257",
            "imageTag": "golang:1.17-buster"
        }
}
```

If an image with the specific digest already exists it will ignore the push.

Run the Script
```
./docker-push
```

## GitHub Actions
There are three workflows that use github actions to build the tool and then run the workflow based on the schema based on the operating system for the image(LINUX/WINDOWS)

1. release.yaml - This workflow complies and builds the releases based on the operating system, currently this supports windows and linux of architecture amd64. 

      - It uses https://goreleaser.com/ to build the artifacts

      - .goreleaser.yaml defines what needs to built.

2. go-linux.yaml- Here we configure the AWS Creds , ECR and use the build from the previous step to trigger and push the linux based images into ECR.

3. go-win.yaml- Here we configure the AWS Creds , ECR and use the build from the previous step to trigger and push the windows based images into ECR.


### Future work

1. Separete the build process and store the artifacts into a central repo such as S3 or artifactory.
2. Create tests for the script using go testing.
3. use the pusblished artifacts to run the workflow to push the images.