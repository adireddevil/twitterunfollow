package twitterunfollow

import (
	s "strings"
	"github.com/DipeshTest/allstarsshared/Twitter"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

var log = logger.GetLogger("activity-twitterunfollow")

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	consumerKey := s.TrimSpace(context.GetInput("consumerKey").(string))
	consumerSecret := s.TrimSpace(context.GetInput("consumerSecret").(string))
	accessToken := s.TrimSpace(context.GetInput("accessToken").(string))
	accessTokenSecret := s.TrimSpace(context.GetInput("accessTokenSecret").(string))
	twitterHandle := s.TrimSpace(context.GetInput("twitterHandle").(string))

	if len(consumerKey) == 0 {

		context.SetOutput("statusCode", "101")

		context.SetOutput("message", "Consumer Key field is blank")
		//context.SetOutput("failedNumbers", to)

		//respond with this
	} else if len(consumerSecret) == 0 {

		context.SetOutput("statusCode", "102")

		context.SetOutput("message", "Consumer Key field is blank")

	} else if len(accessToken) == 0 {

		context.SetOutput("statusCode", "103")

		context.SetOutput("message", "Access Token field is blank")

	} else if len(accessTokenSecret) == 0 {

		context.SetOutput("statusCode", "104")

		context.SetOutput("message", "Access Token Secret field is blank")

	} else if len(twitterHandle) == 0 {

		context.SetOutput("statusCode", "105")

		context.SetOutput("message", "Tweet cannot be blank")

	} else {


		//code, msg := GDrive.CreateFile(accessToken, fileFullPath, emailAddr, role, sendNotification, timeout)

		code, msg := twitter.UnFollow(consumerKey, consumerSecret, accessToken, accessTokenSecret, twitterHandle)
		context.SetOutput("statusCode", code)

		context.SetOutput("message", msg)
	}

	return true, err
}