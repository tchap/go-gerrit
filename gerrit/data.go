// Copyright (c) 2014 The go-gerrit AUTHORS
//
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

package gerrit

type Change struct {
	Project         string
	Branch          string
	Topic           string
	Id              string
	Subject         string
	Owner           *Account
	URL             string
	CommitMessage   string
	CreatedOn       uint64
	LastUpdated     uint64
	SortKey         string
	Open            bool
	Status          string
	Comments        []*PatchsetComment
	TrackingIds     []string
	CurrentPatchSet *PatchSet
	DependsOn       *Dependency
	NeededBy        *Dependency
	SubmitRecords   []*SubmitRecord
	AllReviewers    []*Account
}

type TrackingId struct {
	System string
	Id     string
}

type Account struct {
	Name     string
	Email    string
	Username string
}

type PatchSet struct {
	Number         int
	Revision       string
	Parents        []string
	Ref            string
	Uploader       *Account
	Author         *Account
	CreatedOn      uint64
	IsDraft        bool
	Approvals      []*Approval
	Comments       []*PatchsetComment
	Files          []*File
	SizeInsertions int
	SizeDeletions  int
}

type Approval struct {
	Type        string
	Description string
	Value       string
	GrantedOn   uint64
	By          *Account
}

type RefUpdate struct {
	OldRev  string
	NewRev  string
	RefName string
	Project string
}

type SubmitRecord struct {
	Status string
	Labels []*Label
}

type Label struct {
	Label  string
	Status string
	By     *Account
}

type Dependency struct {
	Id                string
	Number            int
	Revision          string
	Ref               string
	IsCurrentPatchSet bool
}

type Message struct {
	Timestamp uint64
	Reviewer  *Account
	Message   string
}

type PatchsetComment struct {
	File     string
	Line     int
	Reviewer *Account
	Message  string
}

type File struct {
	File       string
	FileOld    string
	Type       string
	Insertions int
	Deletions  int
}
