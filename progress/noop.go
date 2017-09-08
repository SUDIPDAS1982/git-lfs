package progress

import "github.com/git-lfs/git-lfs/git/githistory/log"

func Noop() Meter {
	return &nonMeter{}
}

type nonMeter struct{}

func (m *nonMeter) Start()                                                               {}
func (m *nonMeter) Pause()                                                               {}
func (m *nonMeter) Add(size int64)                                                       {}
func (m *nonMeter) Skip(size int64)                                                      {}
func (m *nonMeter) StartTransfer(name string)                                            {}
func (m *nonMeter) TransferBytes(direction, name string, read, total int64, current int) {}
func (m *nonMeter) FinishTransfer(name string)                                           {}
func (m *nonMeter) Finish()                                                              {}

func (m *nonMeter) Updates() <-chan *log.Update { return nil }
func (m *nonMeter) Throttled() bool             { return false }
