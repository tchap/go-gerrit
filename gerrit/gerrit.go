// Copyright (c) 2014 The go-gerrit AUTHORS
//
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

package gerrit

import (
	"os/user"
	"path/filepath"
	"strconv"

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

	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	if opts.Host == "" {
		opts.Host = DefaultHost
	}
	if opts.Port == 0 {
		opts.Port = DefaultPort
	}
	if opts.User == "" {
		opts.User = usr.Username
	}
	if opts.IdentityFile == "" {
		opts.IdentityFile = filepath.Join(usr.HomeDir, ".ssh", "id_rsa")
	}

	// Connect to Gerrit using given configuration.
	clientKeyring, err := newKeyring(opts.IdentityFile)
	if err != nil {
		return nil, err
	}

	clientConfig := &ssh.ClientConfig{
		User: opts.User,
		Auth: []ssh.ClientAuth{clientKeyring},
	}

	port := strconv.FormatUint(uint64(opts.Port), 10)
	conn, err := ssh.Dial("tcp", opts.Host+":"+port, clientConfig)
	if err != nil {
		return nil, err
	}

	return &Session{conn}, nil
}
