package dataloader

import (
	"time"

	. "github.com/matthew188/crex"
)

type DataLoader interface {
	Setup(start time.Time, end time.Time) error
	ReadOrderBooks() []*OrderBook
	ReadRecords(limit int) []*Record
	HasMoreData() bool
}
