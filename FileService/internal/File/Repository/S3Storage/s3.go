package S3Storage

import (
	repo "Mailer/FileService/internal/File/Repository"
	fileProto "Mailer/FileService/proto"
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"Mailer/config"
	"io/ioutil"
	"os"
	"strconv"
)


type S3Conf struct {
	ctx context.Context
	sess *s3.S3
	AccessKey  string
	SecretKey  string
	BucketName string
	BucketID   string
	Password   string
	Token      string
}


func New() repo.Interface {
	rep:=S3Conf{
		AccessKey:  config.AccessKey,
		SecretKey:  config.SecretKey,
		BucketName: config.BucketName,
		BucketID:   config.BucketID,
		Password:   config.Password,
		Token:      config.Token,
	}
	return rep
}

func (storage *S3Conf)Connect() error{
	storage.ctx = context.Background()
	creds:=credentials.NewStaticCredentials(storage.AccessKey, storage.SecretKey, storage.Token)
	_, err:=creds.Get()
	if err!=nil{
		fmt.Printf("error cred: %v\n", err)
		return nil
	}
	cfg:=aws.NewConfig().WithRegion(
		"ru-msk",
	).WithEndpoint(
		"http://hb.bizmrg.com",
	).WithCredentials(
		creds,
	)
	sess:=session.Must(session.NewSession(cfg))
	storage.sess=s3.New(sess)
	return nil
}

func (storage S3Conf)SaveFiles(file *fileProto.Files) error{
	err:=storage.Connect()
	if err!=nil{
		return err
	}
	for pos, f:=range file.Files{
		_, err := storage.sess.PutObjectWithContext(storage.ctx, &s3.PutObjectInput{
			Bucket: aws.String(storage.BucketName),
			Key:    aws.String(strconv.Itoa(int(file.LetterId))+"_"+strconv.Itoa(pos)),
			Body:   bytes.NewReader(f.Content),
		})
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
				fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
			} else {
				fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
			}
			return err
		}
	}
	return nil
}
func (storage S3Conf)GetFiles(lid *fileProto.LetterId) (*fileProto.Files, error){
	err:=storage.Connect()
	if err!=nil{
		return nil, err
	}
	var container fileProto.Files
	var i=0
	for ;;{
		resp, err := storage.sess.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(storage.BucketName),
			Key:    aws.String(strconv.Itoa(int(lid.Id))+"_"+strconv.Itoa(i)),
		})
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case s3.ErrCodeNoSuchBucket:
					var err error
					return nil, errors.Wrapf(err, "bucket  does not exist")
				case s3.ErrCodeNoSuchKey:
					break
				}
			}
			break
		}
		defer resp.Body.Close()
		var f fileProto.File
		body, _ := ioutil.ReadAll(resp.Body)
		f.Content=body
		container.Files=append(container.Files, &f)
	}
	return &container, nil
}
func (storage S3Conf)SaveAvatar(file *fileProto.Avatar) error{
	err:=storage.Connect()
	if err!=nil{
		return err
	}
	storage.ctx = context.Background()
	creds:=credentials.NewStaticCredentials(storage.AccessKey, storage.SecretKey, storage.Token)
	_, err=creds.Get()
	if err!=nil{
		fmt.Print("error cred: %v\n", err)
		return nil
	}
	cfg:=aws.NewConfig().WithRegion(
		"ru-msk",
	).WithEndpoint(
		"http://hb.bizmrg.com",
	).WithCredentials(
		creds,
	)
	sess:=session.Must(session.NewSession(cfg))
	storage.sess=s3.New(sess)
	_, err = storage.sess.PutObjectWithContext(storage.ctx, &s3.PutObjectInput{
		Bucket: aws.String(storage.BucketName),
		Key:    aws.String(file.Email),
		Body:   bytes.NewReader(file.Content),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
		}
		return err
	}
	return nil
}
func (storage S3Conf)GetAvatar(user *fileProto.User) (*fileProto.Avatar, error){
	err:=storage.Connect()
	if err!=nil{
		return nil, err
	}
	resp, err := storage.sess.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(storage.BucketName),
		Key:    aws.String(user.Email),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				_, _ = storage.GetDefaultAvatar()
				return nil, nil
			case s3.ErrCodeNoSuchKey:
				_, _ = storage.GetDefaultAvatar()
				return nil, nil
			}
		}
		_, _ = storage.GetDefaultAvatar()
		return nil, nil
	}
	defer resp.Body.Close()
	var f fileProto.Avatar
	body, _ := ioutil.ReadAll(resp.Body)
	f.Content=body
	return &f, nil
}
func (storage S3Conf)GetDefaultAvatar() (*fileProto.Avatar, error){
	err:=storage.Connect()
	if err!=nil{
		return nil, err
	}
	resp, err := storage.sess.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(storage.BucketName),
		Key:    aws.String("default.jpeg"),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				return nil, nil
			case s3.ErrCodeNoSuchKey:
				return nil, nil
			}
		}
		return nil, nil
	}
	defer resp.Body.Close()
	var f fileProto.Avatar
	body, err := ioutil.ReadAll(resp.Body)
	f.Content=body
	return &f, nil
}
