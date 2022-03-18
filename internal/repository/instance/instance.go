package instance

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

func GetConnection() *dynamodb.DynamoDB {

	sess := sesssion.Must(session.NewSessionWithOptions(sesion.Options{
		SharedConfigState: sesssion.SharedConfigEnable,
	}))
}
