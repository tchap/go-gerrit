// Copyright (c) 2014 The go-gerrit AUTHORS
//
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

package gerrit

import (
	"os/user"
	"path/filepath"

	"code.google.com/p/go.crypto/ssh"
)

const (
	DefaultHost = "127.0.0.1"
	DefaultPort = 29418
)

type Session struct {
	conn *ssh.ClientConn
}

type DialOptions struct {
	// Gerrit SSH endpoint host.
	// 127.0.0.1 is used by default.
	Host string

	// Gerrit SSH endpoint port.
	// 29418 is used by default.
	Port uint16

	// Gerrit username.
	// The current user's username is used by default.
	User string

	// The file from which the SSH private key is read.
	// ~/.ssh/id_rsa is used by default.
	IdentityFile string
}

func Dial(opts *DialOptions) (*Session, error) {
	// Check and potentially fill in the options.
	if opts == nil {
		opts = &DialOptions{}
	}

	if opts.Host == "" {
		opts.Host = DefaultHost
	}
	if opts.Port == 0 {
		opts.Port = DefaultPort
	}
	if opts.User == "" {
		usr, err := user.Current()
		if err != nil {
			return nil, err
		}
		opts.User = usr.Username
	}
	if opts.IdentityFile == "" {
		usr, err := user.Current()
		if err != nil {
			return nil, err
		}
		opts.IdentityFile = filepath.Join(usr.HomeDir, ".ssh", "id_rsa")
	}

	// Connect to Gerrit using given configuration.
	conn, err := dialSSH(opts.Host, opts.Port, opts.User, opts.IdentityFile)
	if err != nil {
		return nil, err
	}

	return &Session{conn}, nil
}

func (session *Session) NewEventStream() (*EventStream, error) {
	s, err := session.conn.NewSession()
	if err != nil {
		return nil, err
	}

	return newEventStream(s)
}

func (session *Session) Close() error {
	return session.conn.Close()
}
