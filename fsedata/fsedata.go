// +build darwin

package fsedata

const FSENotifierVersion = "0.1"

//go:generate stringer -type=Errors

// Errors is type for error events
type Errors int32

// Reported errors types
const (
	NoError Errors = iota
	ErrorNoWatchPoint
	ErrorWrongWatchPoint
	ErrorNoDirWatchPoint
)

// FSEvent struct contain single file system event information it co-responding to fsevents.Event
type FSEvent struct {
	Path  string
	Flags EventFlags
}

// FSEvents is struct of events going to be Marshled to json
type FSEvents struct {
	Fserror Errors
	Events  []FSEvent
}

//go:generate stringer -type=EventFlags

// EventFlags co-responding fsevents.EventFlags
type EventFlags uint32

const (
	// MustScanSubDirs indicates that events were coalesced hierarchically.
	MustScanSubDirs EventFlags = 1 << iota
	// UserDropped or KernelDropped is set alongside MustScanSubDirs
	// to help diagnose the problem.
	UserDropped
	KernelDropped

	// EventIDsWrapped indicates the 64-bit event ID counter wrapped around.
	EventIDsWrapped

	// HistoryDone is a sentinel event when retrieving events sinceWhen.
	HistoryDone

	// RootChanged indicates a change to a directory along the path being watched.
	RootChanged

	// Mount for a volume mounted underneath the path being monitored.
	Mount
	// Unmount event occurs after a volume is unmounted.
	Unmount

	// The following flags are only set when using FileEvents.

	ItemCreated
	ItemRemoved
	ItemInodeMetaMod
	ItemRenamed
	ItemModified
	ItemFinderInfoMod
	ItemChangeOwner
	ItemXattrMod
	ItemIsFile
	ItemIsDir
	ItemIsSymlink
)
