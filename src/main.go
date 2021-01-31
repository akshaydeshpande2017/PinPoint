package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/pinpoint"
	//"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/astaxie/beego"
)

type KafkaMsg struct {
	Type string			`json:"Type"`
	MessageBody string	`json:"MessageBody"`
	Mobile string		`json:"Mobile"`
}

type StatementEntry struct {
	Effect   string
	Action   []string
	Resource string
}

type PolicyDocument struct {
	Version   string
	Statement []StatementEntry
}

var policy = PolicyDocument{
	Version: "2012-10-17",
	Statement: []StatementEntry{
		StatementEntry{
			Effect: "Allow",
			// Allows for DeleteItem, GetItem, PutItem, Scan, and UpdateItem
			Action: []string{
				"mobiletargeting:SendMessages",
				//"mobiletargeting:GetApp",
				//"mobiletargeting:SendUsersMessages",
				//"mobiletargeting:GetApps",
				//"mobiletargeting:GetApplicationSettings",
			},
			Resource: "arn:aws:mobiletargeting:ap-south-1:123080666138:apps/*",
		},
		{
			Effect: "Allow",
			Action: []string{
				//"iam:GetPolicy",
				//"iam:ListPolicies",
				"iam:*",
			},
			// Resource: "arn:aws:iam::123080666138:policy/*",
			Resource: "*",
		},
	},
}

type mainController struct {
    beego.Controller
}


func main(){
    beego.Router("/:operation/:num1:int/:num2:int", &mainController{})
    beego.Run()
    //beego.Router("/:operation", &mainController{})
    //beego.Run()
}

func (c *mainController) Get() {
//func main(){
	var msg KafkaMsg
	msg.Type = "TRANSACTIONAL"
        msg.MessageBody = "Test msg"
        msg.Mobile = "+919820716498"
	c.TplName = "result.html"
	Send(msg)  
	//consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
	//	"bootstrap.servers":    "localhost:9092",
	//	"group.id":             "group1",
	//	//"auto.offset.reset":    "newest",
	//})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = consumer.SubscribeTopics([]string{"test"}, nil)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer consumer.Close()
	//var msg map[string]interface{}
	//var msg KafkaMsg
	//run := true
	//for run == true {
	//	ev := consumer.Poll(0)
	//	switch e := ev.(type) {
	//	case *kafka.Message:
	//		fmt.Printf("%% Message on %s:\n%s\n",
	//			e.TopicPartition, string(e.Value))
	//		err = json.Unmarshal(e.Value, &msg)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		fmt.Println(msg)
	//		//run = false
	//		Send(msg)
	//	case kafka.PartitionEOF:
	//		fmt.Println("EOF", e)
	//	case kafka.Error:
	//		fmt.Println("Error: ", e)
	//		run = false
	//	}
	//}
}

func Send(msg KafkaMsg) {

	// Creating session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
	})
	if err != nil {
		fmt.Println(err)
	}

	//iamSvc := iam.New(sess)
	// List IAM users
	//// List IAM users
	//r, err := iamSvc.ListUsers(&iam.ListUsersInput{
	//	MaxItems: aws.Int64(10),
	//})
	//
	//if err != nil {
	//	fmt.Println("Error", err)
	//	return
	//}
	//
	//for i, user := range r.Users {
	//	if user == nil {
	//		continue
	//	}
	//	if *user.UserName == "AkshayTest" {
	//		userExist = true
	//	}
	//	fmt.Printf("%d user %s created %v\n", i, *user.UserName, user.CreateDate)
	//}

	// Get or Create User from UserName
	//userName := "AkshayTest1"
	//_, err = iamSvc.GetUser(&iam.GetUserInput{
	//	UserName: aws.String(userName),
	//})
	//if awserr, ok := err.(awserr.Error); ok && awserr.Code() == iam.ErrCodeNoSuchEntityException {
	//	fmt.Println("Error raised user not found")
	//	_, err := iamSvc.CreateUser(&iam.CreateUserInput{
	//		UserName: aws.String(userName),
	//	})
	//	if err != nil {
	//		fmt.Println("CreateUser Error", err)
	//		return
	//	}
	//} else {
	//	fmt.Println("GetUser Error", err)
	//}


	// Create IAM policy for PinPoint if it does not exists.
	//arn := "arn:aws:iam::123080666138:policy/PinPointSendSMS"
	//result, err := iamSvc.GetPolicy(&iam.GetPolicyInput{
	//	PolicyArn: aws.String(arn),
	//})

	//if err != nil {
	//	// Creating New Policy
	//	b, err := json.Marshal(&policy)
	//	if err != nil {
	//		fmt.Println("Error marshaling policy", err)
	//		return
	//	}
	//	result, err := iamSvc.CreatePolicy(&iam.CreatePolicyInput{
	//		PolicyDocument: aws.String(string(b)),
	//		PolicyName:     aws.String("PinPointSendSMS"),
	//	})
	//	arn = *result.Policy.Arn
	//	if err != nil {
	//		fmt.Println("Error", err)
	//		return
	//	}
	//	fmt.Println("New policy", result)
	//} else {
	//	arn = *result.Policy.Arn
	//}


	// Attach Policy to the user
	//_, err = iamSvc.AttachUserPolicy(&iam.AttachUserPolicyInput{
	//	UserName: aws.String(userName),
	//	PolicyArn: aws.String(arn),
	//})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	// Creating pinpoint Config
	svc := pinpoint.New(sess)

	// Creating or Getting existing App
	appID := "0e17d5e0af6f4761bc64bfb5d14b195e"
	// Create App
	// asd := pinpoint.CreateApplicationRequest{Name: &appID}
	// zxc := pinpoint.CreateAppInput{CreateApplicationRequest: &asd}
	// app, err := svc.CreateApp(&zxc)

	// Get Existing App
	// qwe := pinpoint.GetAppInput{ApplicationId: &appID}
	// app, err := svc.GetApp(&qwe)
	// if err != nil {
	//	 fmt.Println(err)
	// }

	// Sending message through pinpoint.
	sms := "SMS"
	message := pinpoint.MessageRequest{
		Addresses: map[string]*pinpoint.AddressConfiguration{
			msg.Mobile: &pinpoint.AddressConfiguration{
				ChannelType: &sms,
			},
		},
		// Context: nil,
		// Endpoints: nil,
		MessageConfiguration: &pinpoint.DirectMessageConfiguration{
			SMSMessage: &pinpoint.SMSMessage{
				Body: &msg.MessageBody,
				// Keyword: nil,
				// MediaUrl: nil,
				MessageType: &msg.Type,
				// OriginationNumber: nil,
				// SenderId: nil,
				// Substitutions: nil
			},
		},
		// TraceId: nil,
		TemplateConfiguration: nil,
	}
	m := pinpoint.SendMessagesInput{ApplicationId: &appID, MessageRequest: &message}
	msgOut, err := svc.SendMessages(&m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(msgOut)
	fmt.Println(m)
}


