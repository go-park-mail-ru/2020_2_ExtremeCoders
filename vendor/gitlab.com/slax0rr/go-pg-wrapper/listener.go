/*
Copyright © 2020 Tomaz Lovrec <tomaz.lovrec@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package pgwrapper

import (
	"time"

	"github.com/go-pg/pg/v9"
)

type Listener interface {
	Channel() <-chan *pg.Notification
	ChannelSize(size int) <-chan *pg.Notification
	Close() error
	Listen(channels ...string) error
	Receive() (channel string, payload string, err error)
	ReceiveTimeout(timeout time.Duration) (channel, payload string, err error)
	String() string
}

type ListenerWrap struct {
	l Listener
}

func NewListener(l Listener) *ListenerWrap {
	return &ListenerWrap{l}
}

func (l *ListenerWrap) Channel() <-chan *pg.Notification {
	return l.Channel()
}

func (l *ListenerWrap) ChannelSize(size int) <-chan *pg.Notification {
	return l.ChannelSize(size)
}

func (l *ListenerWrap) Close() error {
	return l.Close()
}

func (l *ListenerWrap) Listen(channels ...string) error {
	return l.Listen(channels...)
}

func (l *ListenerWrap) Receive() (channel string, payload string, err error) {
	return l.Receive()
}

func (l *ListenerWrap) ReceiveTimeout(timeout time.Duration) (channel, payload string, err error) {
	return l.ReceiveTimeout(timeout)
}

func (l *ListenerWrap) String() string {
	return l.String()
}
