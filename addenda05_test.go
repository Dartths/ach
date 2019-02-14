// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"

	"github.com/moov-io/base"
)

func mockAddenda05() *Addenda05 {
	addenda05 := NewAddenda05()
	addenda05.SequenceNumber = 1
	addenda05.PaymentRelatedInformation = "This is an Addenda05"
	addenda05.EntryDetailSequenceNumber = 0000001
	return addenda05
}

func TestMockAddenda05(t *testing.T) {
	addenda05 := mockAddenda05()
	if err := addenda05.Validate(); err != nil {
		t.Error("mockAddenda05 does not validate and will break other tests")
	}
	if addenda05.EntryDetailSequenceNumber != 0000001 {
		t.Error("EntryDetailSequenceNumber dependent default value has changed")
	}
}

// testParseAddenda05 parses an Addenda05 record for a PPD detail entry
func testParseAddenda05(t testing.TB) {
	addendaPPD := NewAddenda05()
	var line = "705PPD                                        DIEGO MAY                            00010000001"
	addendaPPD.Parse(line)

	r := NewReader(strings.NewReader(line))

	//Add a new BatchPPD
	r.addCurrentBatch(NewBatchPPD(mockBatchPPDHeader()))

	//Add a PPDEntryDetail
	entryDetail := mockPPDEntryDetail()

	//Add an addenda to the PPD EntryDetail
	entryDetail.AddAddenda05(addendaPPD)

	// add the PPD entry detail to the batch
	r.currentBatch.AddEntry(entryDetail)

	record := r.currentBatch.GetEntries()[0].Addenda05[0]

	if record.recordType != "7" {
		t.Errorf("RecordType Expected '7' got: %v", record.recordType)
	}
	if record.TypeCode != "05" {
		t.Errorf("TypeCode Expected 05 got: %v", record.TypeCode)
	}
	if record.PaymentRelatedInformationField() != "PPD                                        DIEGO MAY                            " {
		t.Errorf("PaymentRelatedInformation Expected 'PPD                                        DIEGO MAY                            ' got: %v", record.PaymentRelatedInformationField())
	}
	if record.SequenceNumberField() != "0001" {
		t.Errorf("SequenceNumber Expected '0001' got: %v", record.SequenceNumberField())
	}
	if record.EntryDetailSequenceNumberField() != "0000001" {
		t.Errorf("EntryDetailSequenceNumber Expected '0000001' got: %v", record.EntryDetailSequenceNumberField())
	}
}

// TestParseAddenda05 tests parsing an Addenda05 record for a PPD detail entry
func TestParseAddenda05(t *testing.T) {
	testParseAddenda05(t)
}

// BenchmarkParseAddenda05 benchmarks parsing an Addenda05 record for a PPD detail entry
func BenchmarkParseAddenda05(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testParseAddenda05(b)
	}
}

// testAddenda05String validates that a known parsed file can be return to a string of the same value
func testAddenda05String(t testing.TB) {
	addenda05 := NewAddenda05()
	var line = "705WEB                                        DIEGO MAY                            00010000001"
	addenda05.Parse(line)

	if addenda05.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestAddenda05 String tests validating that a known parsed file can be return to a string of the same value
func TestAddenda05String(t *testing.T) {
	testAddenda05String(t)
}

// BenchmarkAddenda05 String benchmarks validating that a known parsed file can be return to a string of the same value
func BenchmarkAddenda05String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda05String(b)
	}
}

func TestValidateAddenda05RecordType(t *testing.T) {
	addenda05 := mockAddenda05()
	addenda05.recordType = "63"
	err := addenda05.Validate()
	// TODO: are we not expecting any errors here?
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestValidateAddenda05TypeCode(t *testing.T) {
	addenda05 := mockAddenda05()
	addenda05.TypeCode = "23"
	err := addenda05.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddenda05FieldInclusion(t *testing.T) {
	addenda05 := mockAddenda05()
	addenda05.EntryDetailSequenceNumber = 0
	err := addenda05.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddenda05FieldInclusionSequenceNumber(t *testing.T) {
	addenda05 := mockAddenda05()
	addenda05.SequenceNumber = 0
	err := addenda05.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddenda05FieldInclusionRecordType(t *testing.T) {
	addenda05 := mockAddenda05()
	addenda05.recordType = ""
	err := addenda05.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddenda05PaymentRelatedInformationAlphaNumeric(t *testing.T) {
	addenda05 := mockAddenda05()
	addenda05.PaymentRelatedInformation = "®©"
	err := addenda05.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

func testAddenda05TypeCodeNil(t testing.TB) {
	addenda05 := mockAddenda05()
	addenda05.TypeCode = ""
	err := addenda05.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddenda05TypeCodeNil(t *testing.T) {
	testAddenda05TypeCodeNil(t)
}

func BenchmarkAddenda05TypeCodeNil(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda05TypeCodeNil(b)
	}
}

// testAddenda05TypeCode05 TypeCode is 05
func testAddenda05TypeCode05(t testing.TB) {
	addenda05 := mockAddenda05()
	addenda05.TypeCode = "99"
	err := addenda05.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda05TypeCode05 tests TypeCode is 05
func TestAddenda05TypeCode05(t *testing.T) {
	testAddenda05TypeCode05(t)
}

// BenchmarkAddenda05TypeCode05 benchmarks TypeCode is 05
func BenchmarkAddenda05TypeCode05(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda05TypeCode05(b)
	}
}
