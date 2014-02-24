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

}

type DraftPublishedEvent struct {

}

type ChangeAbandonedEvent struct {

}

type ChangeRestoredEvent struct {

}

type ChangeMergedEvent struct {

}

type MergeFailedEvent struct {

}

type CommentAddedEvent struct {

}

type RefUpdatedEvent struct {

}

type ReviewerAddedEvent struct {

}

type TopicChangedEvent struct {

}
