// Copyright (c) 2014 The go-gerrit AUTHORS
//
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

package gerrit

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"code.google.com/p/go.crypto/ssh"
)

const (
	EventPatchsetCreated = "patchset-created"
	EventDraftPublished  = "draft-published"
	EventChangeAbandoned = "change-abandoned"
	EventChangeRestored  = "change-restored"
	EventChangeMerged    = "change-merged"
	EventMergeFailed     = "merge-failed"
	EventCommentAdded    = "comment-added"
	EventRefUpdated      = "ref-updated"
	EventReviewerAdded   = "reviewer-added"
	EventTopicChanged    = "topic-changed"
)

type EventStream struct {
	session  *ssh.Session
	eventCh  chan interface{}
	closedCh chan struct{}
	err      error
}

func newEventStream(session *ssh.Session) (*EventStream, error) {
	pipe, err := session.StdoutPipe()
	if err != nil {
		session.Close()
		return nil, err
	}

	if err := session.Start("gerrit stream-events"); err != nil {
		session.Close()
		return nil, err
	}

	stream := &EventStream{
		session:  session,
		eventCh:  make(chan interface{}),
		closedCh: make(chan struct{}),
	}

	go stream.readEvents(pipe)
	return stream, nil
}

func (stream *EventStream) readEvents(pipe io.Reader) {
	reader := bufio.NewReader(pipe)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			stream.err = err
			close(stream.eventCh)
			close(stream.closedCh)
			return
		}

		stream.eventCh <- parseEvent(line)
	}
}

func (stream *EventStream) Chan() <-chan interface{} {
	return stream.eventCh
}

func (stream *EventStream) Close() error {
	if err := stream.session.Close(); err != nil {
		return err
	}
	<-stream.closedCh
	return stream.err
}

type event struct {
	Type string
}

func parseEvent(line []byte) interface{} {
	var evt event
	if err := json.Unmarshal(line, &evt); err != nil {
		return err
	}

	switch evt.Type {
	case EventPatchsetCreated:
		var event PatchsetCreatedEvent
		if err := json.Unmarshal(line, &event); err != nil {
			return err
		}
		return &event

	case EventDraftPublished:
		var event DraftPublishedEvent
		if err := json.Unmarshal(line, &event); err != nil {
			return err
		}
		return &event

	case EventChangeAbandoned:
		var event ChangeAbandonedEvent
		if err := json.Unmarshal(line, &event); err != nil {
			return err
		}
		return &event

	case EventChangeRestored:
		var event ChangeRestoredEvent
		if err := json.Unmarshal(line, &event); err != nil {
			return err
		}
		return &event

	case EventChangeMerged:
		var event ChangeMergedEvent
		if err := json.Unmarshal(line, &event); err != nil {
			return err
		}
		return &event

	case EventMergeFailed:
		var event MergeFailedEvent
		if err := json.Unmarshal(line, &event); err != nil {
			return err
		}
		return &event

	case EventCommentAdded:
		var event PatchsetCreatedEvent
		if err := json.Unmarshal(line, &event); err != nil {
			return err
		}
		return &event

	case EventRefUpdated:
		var event RefUpdatedEvent
		if err := json.Unmarshal(line, &event); err != nil {
			return err
		}
		return &event

	case EventReviewerAdded:
		var event ReviewerAddedEvent
		if err := json.Unmarshal(line, &event); err != nil {
			return err
		}
		return &event

	case EventTopicChanged:
		var event TopicChangedEvent
		if err := json.Unmarshal(line, &event); err != nil {
			return err
		}
		return &event

	default:
		return fmt.Errorf("unknown event received: %s", evt.Type)
	}
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
