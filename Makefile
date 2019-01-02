.PHONY: deps clean build package deploy

deps:
	go get -u github.com/aws/aws-lambda-go/...
	go get -u github.com/nlopes/slack

clean: 
	rm -rf ./presence/presence
	
build:
	GOOS=linux GOARCH=amd64 go build -o presence/presence ./presence

package:
	sam package --template-file template.yaml --output-template-file packaged.yaml --s3-bucket presence-function $(filter-out $@,$(MAKECMDGOALS))

deploy:
	sam deploy --template-file packaged.yaml --stack-name presence --capabilities CAPABILITY_IAM --parameter-overrides PresenceSlackToken=$(filter-out $@,$(MAKECMDGOALS))