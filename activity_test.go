package twitterunfollow

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"fmt"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}
/*
func TestTwitterUnfollow_Success(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("consumerKey", "6f2ogTcNPK0Mf6OO6DSL0GxNB")
	tc.SetInput("consumerSecret", "Hb8gDvSnfIkLuO6L6NQWakf8w2OHBnfrH1z4Cjh8IrCsGa4BDM")
	tc.SetInput("accessToken", "1026895242-VTblWlyH3J24lsQbJcO7KBxkbd7ByGyUXIkAFTk")
	tc.SetInput("accessTokenSecret", "h41iwGi3oHwgP01tmVYwY1ceaqaly7RA14zY6oEqigD8b")
	tc.SetInput("twitterHandle", "LeonStigter")

	act.Eval(tc)

	//check result attr

	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "200")


}

func TestTwitterUnfollow_TwitterHandleInvalid(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("consumerKey", "6f2ogTcNPK0Mf6OO6DSL0GxNB")
	tc.SetInput("consumerSecret", "Hb8gDvSnfIkLuO6L6NQWakf8w2OHBnfrH1z4Cjh8IrCsGa4BDM")
	tc.SetInput("accessToken", "1026895242-VTblWlyH3J24lsQbJcO7KBxkbd7ByGyUXIkAFTk")
	tc.SetInput("accessTokenSecret", "h41iwGi3oHwgP01tmVYwY1ceaqaly7RA14zY6oEqigD8b")
	tc.SetInput("twitterHandle", "Leosigter")

	act.Eval(tc)

	//check result attr

	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "404")


}


func TestTwitterUnfollow_TwitterHandleBlank(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("consumerKey", "6f2ogTcNPK0Mf6OO6DSL0GxNB")
	tc.SetInput("consumerSecret", "Hb8gDvSnfIkLuO6L6NQWakf8w2OHBnfrH1z4Cjh8IrCsGa4BDM")
	tc.SetInput("accessToken", "1026895242-VTblWlyH3J24lsQbJcO7KBxkbd7ByGyUXIkAFTk")
	tc.SetInput("accessTokenSecret", "h41iwGi3oHwgP01tmVYwY1ceaqaly7RA14zY6oEqigD8b")
	tc.SetInput("twitterHandle", "")

	act.Eval(tc)

	//check result attr

	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "105")


} */


func TestTwitterUnfollow_InvalidToken(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("consumerKey", "6f2ogTcNPK0Mf6OO6DSL0GxN")
	tc.SetInput("consumerSecret", "Hb8gDvSnfIkLuO6L6NQWakf8w2OHBnfrH1z4Cjh8IrCsGa4BDM")
	tc.SetInput("accessToken", "1026895242-VTblWlyH3J24lsQbJcO7KBxkbd7ByGyUXIkAFTk")
	tc.SetInput("accessTokenSecret", "h41iwGi3oHwgP01tmVYwY1ceaqaly7RA14zY6oEqigD8b")
	tc.SetInput("twitterHandle", "themichaelowen")

	act.Eval(tc)

	//check result attr

	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "401")


}
