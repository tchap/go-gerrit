// Copyright (c) 2014 The go-gerrit AUTHORS
//
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

package gerrit

type EventStream struct {
	PatchsetCreated <-chan *PatchsetCreatedEvent
	DraftPublished  <-chan *DraftPublishedEvent
	ChangeAbandoned <-chan *ChangeAbandonedEvent
	ChangeRestored  <-chan *ChangeRestoredEvent
	ChangeMerged    <-chan *ChangeMergedEvent
	MergeFailed     <-chan *MergeFailedEvent
	CommentAdded    <-chan *CommentAddedEvent
	RefUpdated      <-chan *RefUpdatedEvent
	ReviewerAdded   <-chan *ReviewerAddedEvent
	TopicChanged    <-chan *TopicChangedEvent
}

type PatchsetCreatedEvent struct {
	Type     string
	Change   *Change
	PatchSet *PatchSet
	Uploader *Account
}

type DraftPublishedEvent struct {
	Type     string
	Change   *Change
	PatchSet *PatchSet
	Uploader *Account
}

type ChangeAbandonedEvent struct {
	Type      string
	Change    *Change
	PatchSet  *PatchSet
	Abandoner *Account
	Reason    string
}

type ChangeRestoredEvent struct {
	Type     string
	Change   *Change
	PatchSet *PatchSet
	Restorer *Account
	Reason   string
}

type ChangeMergedEvent struct {
	Type      string
	Change    *Change
	PatchSet  *PatchSet
	Submitter *Account
}

type MergeFailedEvent struct {
	Type      string
	Change    *Change
	PatchSet  *PatchSet
	Submitter *Account
	Reason    string
}

type CommentAddedEvent struct {
	Type      string
	Change    *Change
	PatchSet  *PatchSet
	Author    *Account
	Approvals []Approval
	Comment   string
}

type RefUpdatedEvent struct {
	Type      string
	Submitter *Account
	RefUpdate *RefUpdate
}

type ReviewerAddedEvent struct {
	Type     string
	Change   *Change
	PatchSet *PatchSet
	Reviewer *Account
}

type TopicChangedEvent struct {
	Type     string
	Change   *Change
	Changer  *Account
	OldTopic string
}
