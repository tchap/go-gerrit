// Copyright (c) 2014 The go-gerrit AUTHORS
//
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

package gerrit

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"code.google.com/p/go.crypto/ssh"
)

// dialSSH establishes an SSH connection using the given username and identity file.
func dialSSH(host string, port uint16, user string, identityFile string) (*ssh.ClientConn, error) {
	auth, err := newKeyring(identityFile)
	if err != nil {
		return nil, err
	}

	portString := strconv.FormatUint(uint64(port), 10)
	return ssh.Dial("tcp", host+":"+portString, &ssh.ClientConfig{
		User: user,
		Auth: []ssh.ClientAuth{auth},
	})
}

// keyring implements ssh.ClientKeyring
type keyring struct {
	key *rsa.PrivateKey
}

func (k *keyring) Key(i int) (key ssh.PublicKey, err error) {
	if i != 0 {
		return nil, nil
	}
	return ssh.NewPublicKey(&k.key.PublicKey)
}

func (k *keyring) Sign(i int, rand io.Reader, data []byte) (sig []byte, err error) {
	hashFunc := crypto.SHA1
	hash := hashFunc.New()
	hash.Write(data)
	digest := hash.Sum(nil)
	return rsa.SignPKCS1v15(rand, k.key, hashFunc, digest)
}

func newKeyring(identityFile string) (ssh.ClientAuth, error) {
	// Check the private key access permissions.
	for _, filename := range [...]string{identityFile, filepath.Dir(identityFile)} {
		info, err := os.Stat(filename)
		if err != nil {
			return nil, err
		}

		if info.Mode()&0077 != 0 {
			return nil, fmt.Errorf(
				"Permissions %o for '%s' are too open", info.Mode()&os.ModePerm)
		}
	}

	// Read and decode the SSH key.
	p, err := ioutil.ReadFile(identityFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(p)
	if block == nil {
		return nil, fmt.Errorf("File '%s' does not contain a PEM-encoded key")
	}

	rsaKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// Return the keyring.
	return ssh.ClientAuthKeyring(&keyring{rsaKey}), nil
}
