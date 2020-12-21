package S3Storage

import (
	repo "Mailer/FileService/internal/File/Repository"
	fileProto "Mailer/FileService/proto"
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/pkg/errors"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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
		AccessKey:  "vUEv3F69WEeN1D85oiiFgt",
		SecretKey:  "c5yvQ6ANBnxvU2txz6dQwY7rJjDvMmVxVEakNjgJfH4X",
		BucketName: "maila",
		BucketID:   "mcs6132821991",
		Password:   "CherDan985fy1aasdf681553",
		Token:      "",
	}
	var timeout time.Duration
	timeout=0

	rep.ctx = context.Background()
	var cancelFn func()
	if timeout > 0 {
		rep.ctx, cancelFn = context.WithTimeout(rep.ctx, timeout)
	}

	defer cancelFn()


	creds:=credentials.NewStaticCredentials(rep.AccessKey, rep.SecretKey, rep.Token)
	_, err:=creds.Get()
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
	rep.sess=s3.New(sess)
	return rep
}

func (storage S3Conf)SaveFiles(file *fileProto.Files) error{
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
					return nil, errors.Wrapf(err, "bucket %s does not exist")
				case s3.ErrCodeNoSuchKey:
					break
				}
			}
			break
		}
		defer resp.Body.Close()
		var f fileProto.File
		body, err := ioutil.ReadAll(resp.Body)
		f.Content=body
		container.Files=append(container.Files, &f)
	}
	return &container, nil
}
func (storage S3Conf)SaveAvatar(file *fileProto.Avatar) error{
	_, err := storage.sess.PutObjectWithContext(storage.ctx, &s3.PutObjectInput{
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
	resp, err := storage.sess.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(storage.BucketName),
		Key:    aws.String(user.Email),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				storage.GetDefaultAvatar()
				return nil, nil
			case s3.ErrCodeNoSuchKey:
				storage.GetDefaultAvatar()
				return nil, nil
			}
		}
		storage.GetDefaultAvatar()
		return nil, nil
	}
	defer resp.Body.Close()
	var f fileProto.Avatar
	body, err := ioutil.ReadAll(resp.Body)
	f.Content=body
	return &f, nil
}
func (storage S3Conf)GetDefaultAvatar() (*fileProto.Avatar, error){
	file, err := os.Open("default.jpeg")
	if err != nil {
		return &fileProto.Avatar{}, repo.GetAvatarError
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return &fileProto.Avatar{}, repo.GetAvatarError
	}

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		return &fileProto.Avatar{}, repo.GetAvatarError
	}

	return &fileProto.Avatar{
		FileName: "default.jpeg",
		Content:  buffer.Bytes(),
	}, nil
}
