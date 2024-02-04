package function

import (
	"strings"
	"time"
)

/*
{
  "Records": [
    {
      "eventVersion": "2.1",
      "eventSource": "aws:s3",
      "awsRegion": "us-east-2",
      "eventTime": "2019-09-03T19:37:27.192Z",
      "eventName": "ObjectCreated:Put",
      "userIdentity": {
        "principalId": "AWS:AIDAINPONIXQXHT3IKHL2"
      },
      "requestParameters": {
        "sourceIPAddress": "205.255.255.255"
      },
      "responseElements": {
        "x-amz-request-id": "D82B88E5F771F645",
        "x-amz-id-2": "vlR7PnpV2Ce81l0PRw6jlUpck7Jo5ZsQjryTjKlc5aLWGVHPZLj5NeC6qMa0emYBDXOo6QBU0Wo="
      },
      "s3": {
        "s3SchemaVersion": "1.0",
        "configurationId": "828aa6fc-f7b5-4305-8584-487c791949c1",
        "bucket": {
          "name": "DOC-EXAMPLE-BUCKET",
          "ownerIdentity": {
            "principalId": "A3I5XTEXAMAI3E"
          },
          "arn": "arn:aws:s3:::lambda-artifacts-deafc19498e3f2df"
        },
        "object": {
          "key": "b21b84d653bb07b05b1e6b33684dc11b",
          "size": 1305107,
          "eTag": "b21b84d653bb07b05b1e6b33684dc11b",
          "sequencer": "0C0F6F405D6ED209E1"
        }
      }
    }
  ]
}
*/

type Record struct {
	EventVersion      string           `json:"eventVersion"`
	EventSource       string           `json:"eventSource"`
	AWSRegion         string           `json:"awsRegion"`
	EventTime         time.Time        `json:"eventTime"`
	EventName         string           `json:"eventName"`
	UserIdentity      UserIdentity     `json:"userIdentity"`
	RequestParameters RequestParams    `json:"requestParameters"`
	ResponseElements  ResponseElements `json:"responseElements"`
	S3                S3Model          `json:"s3"`
}

type UserIdentity struct {
	PrincipalID string `json:"principalId"`
}

type RequestParams struct {
	SourceIPAddress string `json:"sourceIPAddress"`
}

type ResponseElements struct {
	XAmzRequestID string `json:"x-amz-request-id"`
	XAmzID2       string `json:"x-amz-id-2"`
}

type S3Model struct {
	S3SchemaVersion string   `json:"s3SchemaVersion"`
	ConfigurationID string   `json:"configurationId"`
	Bucket          Bucket   `json:"bucket"`
	Object          S3Object `json:"object"`
}

type Bucket struct {
	Name          string        `json:"name"`
	OwnerIdentity OwnerIdentity `json:"ownerIdentity"`
	ARN           string        `json:"arn"`
}

type OwnerIdentity struct {
	PrincipalID string `json:"principalId"`
}

type S3Object struct {
	Key       string `json:"key"`
	Size      int    `json:"size"`
	ETag      string `json:"eTag"`
	Sequencer string `json:"sequencer"`
}

type Event struct {
	Records []Record `json:"Records"`
}

func (e Event) hasRecords() bool {
	return len(e.Records) > 0
}

func (e Event) BucketName() string {
	if !e.hasRecords() {
		return ""
	}

	return e.Records[0].S3.Bucket.Name
}

func (e Event) ObjectKey() string {
	if !e.hasRecords() {
		return ""
	}

	return e.Records[0].S3.Object.Key
}

func (e Event) FileName() string {
  objectKey := e.ObjectKey()

  parts := strings.Split(objectKey, "/")

  return parts[len(parts)-1]
}

func (e Event) FileNameWithoutExtension() string {
  fileName := e.FileName()

  parts := strings.Split(fileName, ".")

  return parts[0]
}
