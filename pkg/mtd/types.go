// Copyright 2019 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mtd

// Flasher defines the interface to flash drivers.
type Flasher interface {
	// ReadAt implements io.ReadAt for a flash device.
	ReadAt([]byte, int64) (int, error)
	// QueueWrite queues a sequence of writes into a flash device.
	QueueWrite([]byte, int64) (int, error)
	// SyncWrite assembles the queued writes and figures out a reasonable
	// plan for actually writing the part.
	SyncWrite() error
	// Close implements io.Close for a flash device.
	Close() error
}

// VendorName is the manufacturers name
type VendorName string

// VendorID is the integer associated with a VendorName
// It began as 8 bits, and never stopped growing.
type VendorID uint64

// ChipName is the device name
type ChipName string

// ChipID is the integer associated with a ChipName
// It began as 8 bits, and never stopped growing.
type ChipID uint64

// ChipSize is the size in bytes of the chip.
type ChipSize uint

// Vendor defines operations on vendor data.
type Vendor interface {
	// Chip returns a Chip, given a DeviceID
	ChipInfo(ChipID) (ChipInfo, error)
	// ID Returns the VendorID
	ID() VendorID
	// Name() returns the canonical name
	Name() VendorName
	// Synonyms returns all the names
	Synonyms() []VendorName
}

// ChipInfo provides information about a chip, possibly derived from /dev/mtdx,
// possibly from vendor and device ID read from a programmer, possibly from
// vendor and device ID provided by a user.
type ChipInfo interface {
	// Name returns the chip name
	Name() ChipName
	// ID returns the chip ID
	ID() ChipID
	// Size returns the chip size.
	Size() ChipSize
	// Synonyms returns all the alternate names for a chip
	Synonyms() []ChipName
	// String returns as much information as one can stand about a chip.
	String() string
}

// These struct are not designed to be efficient; rather, they are
// designed to compress efficiently into firmware. Several experiments
// show that this is about the best way to go, absent encoding it as a
// string and unpacking it. We leave dead vendors in for reference
// but comment them out.
type vendor struct {
	names []VendorName
	id    VendorID
}

// ChipDevice has information about a chip as found in a table,
// including Vendor, Device, sizes, and so on; and a reference to
// common properties.  As in Vendors, there are several names for a
// chip.
type ChipDevice struct {
	vendor   VendorName
	devices  []ChipName
	remarks  string
	id       ChipID
	pageSize int
	numPages int
}
