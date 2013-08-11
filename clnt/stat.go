// Copyright 2009 The Go9p Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clnt

import "code.google.com/p/go9p"
import "syscall"

// Returns the metadata for the file associated with the Fid, or an Error.
func (clnt *Clnt) Stat(fid *Fid) (*go9p.Dir, error) {
	tc := clnt.NewFcall()
	err := go9p.PackTstat(tc, fid.Fid)
	if err != nil {
		return nil, err
	}

	rc, err := clnt.Rpc(tc)
	if err != nil {
		return nil, err
	}
	if rc.Type == go9p.Rerror {
		return nil, &go9p.Error{rc.Error, syscall.Errno(rc.Errornum)}
	}

	return &rc.Dir, nil
}

// Returns the metadata for a named file, or an Error.
func (clnt *Clnt) FStat(path string) (*go9p.Dir, error) {
	fid, err := clnt.FWalk(path)
	if err != nil {
		return nil, err
	}

	d, err := clnt.Stat(fid)
	clnt.Clunk(fid)
	return d, err
}

// Modifies the data of the file associated with the Fid, or an Error.
func (clnt *Clnt) Wstat(fid *Fid, dir *go9p.Dir) error {
	tc := clnt.NewFcall()
	err := go9p.PackTwstat(tc, fid.Fid, dir, clnt.Dotu)
	if err != nil {
		return err
	}

	rc, err := clnt.Rpc(tc)
	if err != nil {
		return err
	}
	if rc.Type == go9p.Rerror {
		return &go9p.Error{rc.Error, syscall.Errno(rc.Errornum)}
	}

	return nil
}
