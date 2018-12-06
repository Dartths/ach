// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"log"
	"testing"
)

// mockBatchCTXHeader creates a BatchCTX BatchHeader
func mockBatchCTXHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.StandardEntryClassCode = CTX
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "ACH CTX"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockCTXEntryDetail creates a BatchCTX EntryDetail
func mockCTXEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.IdentificationNumber = "45689033"
	entry.SetCATXAddendaRecords(1)
	entry.SetCATXReceivingCompany("Receiver Company")
	entry.SetTraceNumber(mockBatchCTXHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.Category = CategoryForward
	return entry
}

// mockBatchCTX creates a BatchCTX
func mockBatchCTX() *BatchCTX {
	mockBatch := NewBatchCTX(mockBatchCTXHeader())
	mockBatch.AddEntry(mockCTXEntryDetail())
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	if err := mockBatch.Create(); err != nil {
		log.Fatal(err)
	}
	return mockBatch
}

// testBatchCTXHeader creates a BatchCTX BatchHeader
func testBatchCTXHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchCTXHeader())
	err, ok := batch.(*BatchCTX)
	if !ok {
		t.Errorf("Expecting BatchCTX got %T", err)
	}
}

// TestBatchCTXHeader tests validating BatchCTX BatchHeader
func TestBatchCTXHeader(t *testing.T) {
	testBatchCTXHeader(t)
}

// BenchmarkBatchCTXHeader benchmarks validating BatchCTX BatchHeader
func BenchmarkBatchCTXHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXHeader(b)
	}
}

// testBatchCTXCreate validates BatchCTX create
func testBatchCTXCreate(t testing.TB) {
	mockBatch := mockBatchCTX()
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCTXCreate tests validating BatchCTX create
func TestBatchCTXCreate(t *testing.T) {
	testBatchCTXCreate(t)
}

// BenchmarkBatchCTXCreate benchmarks validating BatchCTX create
func BenchmarkBatchCTXCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXCreate(b)
	}
}

// testBatchCTXStandardEntryClassCode validates BatchCTX create for an invalid StandardEntryClassCode
func testBatchCTXStandardEntryClassCode(t testing.TB) {
	mockBatch := mockBatchCTX()
	mockBatch.Header.StandardEntryClassCode = "WEB"
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "StandardEntryClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCTXStandardEntryClassCode tests validating BatchCTX create for an invalid StandardEntryClassCode
func TestBatchCTXStandardEntryClassCode(t *testing.T) {
	testBatchCTXStandardEntryClassCode(t)
}

// BenchmarkBatchCTXStandardEntryClassCode benchmarks validating BatchCTX create for an invalid StandardEntryClassCode
func BenchmarkBatchCTXStandardEntryClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXStandardEntryClassCode(b)
	}
}

// testBatchCTXServiceClassCodeEquality validates service class code equality
func testBatchCTXServiceClassCodeEquality(t testing.TB) {
	mockBatch := mockBatchCTX()
	mockBatch.GetControl().ServiceClassCode = MixedDebitsAndCredits
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ServiceClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCTXServiceClassCodeEquality tests validating service class code equality
func TestBatchCTXServiceClassCodeEquality(t *testing.T) {
	testBatchCTXServiceClassCodeEquality(t)
}

// BenchmarkBatchCTXServiceClassCodeEquality benchmarks validating service class code equality
func BenchmarkBatchCTXServiceClassCodeEquality(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXServiceClassCodeEquality(b)
	}
}

// testBatchCTXAddendaCount validates BatchCTX Addendum count of 2
func testBatchCTXAddendaCount(t testing.TB) {
	mockBatch := mockBatchCTX()
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addendum" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCTXAddendaCount tests validating BatchCTX Addendum count of 2
func TestBatchCTXAddendaCount(t *testing.T) {
	testBatchCTXAddendaCount(t)
}

// BenchmarkBatchCTXAddendaCount benchmarks validating BatchCTX Addendum count of 2
func BenchmarkBatchCTXAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXAddendaCount(b)
	}
}

// testBatchCTXAddendaCountZero validates Addendum count of 0
func testBatchCTXAddendaCountZero(t testing.TB) {
	mockBatch := NewBatchCTX(mockBatchCTXHeader())
	mockBatch.AddEntry(mockCTXEntryDetail())
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addendum" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCTXAddendaCountZero tests validating Addendum count of 0
func TestBatchCTXAddendaCountZero(t *testing.T) {
	testBatchCTXAddendaCountZero(t)
}

// BenchmarkBatchCTXAddendaCountZero benchmarks validating Addendum count of 0
func BenchmarkBatchCTXAddendaCountZero(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXAddendaCountZero(b)
	}
}

// testBatchCTXInvalidAddendum validates Addendum must be Addenda05
func testBatchCTXInvalidAddendum(t testing.TB) {
	mockBatch := NewBatchCTX(mockBatchCTXHeader())
	mockBatch.AddEntry(mockCTXEntryDetail())
	mockBatch.GetEntries()[0].Addenda02 = mockAddenda02()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addendum" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCTXInvalidAddendum tests validating Addendum must be Addenda05
func TestBatchCTXInvalidAddendum(t *testing.T) {
	testBatchCTXInvalidAddendum(t)
}

// BenchmarkBatchCTXInvalidAddendum benchmarks validating Addendum must be Addenda05
func BenchmarkBatchCTXInvalidAddendum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXInvalidAddendum(b)
	}
}

// testBatchCTXInvalidAddenda validates Addendum must be Addenda05 with record type 7
func testBatchCTXInvalidAddenda(t testing.TB) {
	mockBatch := NewBatchCTX(mockBatchCTXHeader())
	mockBatch.AddEntry(mockCTXEntryDetail())
	addenda05 := mockAddenda05()
	addenda05.recordType = "63"
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCTXInvalidAddenda tests validating Addendum must be Addenda05 with record type 7
func TestBatchCTXInvalidAddenda(t *testing.T) {
	testBatchCTXInvalidAddenda(t)
}

// BenchmarkBatchCTXInvalidAddenda benchmarks validating Addendum must be Addenda05 with record type 7
func BenchmarkBatchCTXInvalidAddenda(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXInvalidAddenda(b)
	}
}

// testBatchCTXInvalidBuild validates an invalid batch build
func testBatchCTXInvalidBuild(t testing.TB) {
	mockBatch := mockBatchCTX()
	mockBatch.GetHeader().recordType = "3"
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCTXInvalidBuild tests validating an invalid batch build
func TestBatchCTXInvalidBuild(t *testing.T) {
	testBatchCTXInvalidBuild(t)
}

// BenchmarkBatchCTXInvalidBuild benchmarks validating an invalid batch build
func BenchmarkBatchCTXInvalidBuild(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXInvalidBuild(b)
	}
}

// testBatchCTXAddenda10000 validates error for 10000 Addenda
func testBatchCTXAddenda10000(t testing.TB) {

	bh := NewBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.StandardEntryClassCode = CTX
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "ACH CTX"
	bh.ODFIIdentification = "12104288"

	entry := NewEntryDetail()
	entry.TransactionCode = CheckingCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.IdentificationNumber = "45689033"
	entry.SetCATXAddendaRecords(9999)
	entry.SetCATXReceivingCompany("Receiver Company")
	entry.SetTraceNumber(mockBatchCTXHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.Category = CategoryForward

	mockBatch := NewBatchCTX(bh)
	mockBatch.AddEntry(entry)

	for i := 0; i < 10000; i++ {
		mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	}
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "AddendaCount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCTXAddenda10000 tests validating error for 10000 Addenda
func TestBatchCTXAddenda10000(t *testing.T) {
	testBatchCTXAddenda10000(t)
}

// BenchmarkBatchCTXAddenda10000 benchmarks validating error for 10000 Addenda
func BenchmarkBatchCTXAddenda10000(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXAddenda10000(b)
	}
}

// testBatchCTXAddendaRecords validates error for AddendaRecords not equal to addendum
func testBatchCTXAddendaRecords(t testing.TB) {
	bh := NewBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.StandardEntryClassCode = CTX
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "ACH CTX"
	bh.ODFIIdentification = "12104288"

	entry := NewEntryDetail()
	entry.TransactionCode = CheckingCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.IdentificationNumber = "45689033"
	entry.SetCATXAddendaRecords(500)
	entry.SetCATXReceivingCompany("Receiver Company")
	entry.SetTraceNumber(mockBatchCTXHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.Category = CategoryForward

	mockBatch := NewBatchCTX(bh)
	mockBatch.AddEntry(entry)

	for i := 0; i < 565; i++ {
		mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	}
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addendum" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCTXAddendaRecords tests validating error for AddendaRecords not equal to addendum
func TestBatchCTXAddendaRecords(t *testing.T) {
	testBatchCTXAddendaRecords(t)
}

// BenchmarkBatchAddendaRecords benchmarks validating error for AddendaRecords not equal to addendum
func BenchmarkBatchCTXAddendaRecords(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXAddendaRecords(b)
	}
}

// testBatchCTXReceivingCompany validates CTXReceivingCompany
func testBatchCTXReceivingCompany(t testing.TB) {
	mockBatch := mockBatchCTX()
	//mockBatch.GetEntries()[0].SetCATXReceivingCompany("Receiver")

	if mockBatch.GetEntries()[0].CATXReceivingCompanyField() != "Receiver Company" {
		t.Errorf("expected %v got %v", "Receiver Company", mockBatch.GetEntries()[0].CATXReceivingCompanyField())
	}
}

// TestBatchCTXReceivingCompany tests validating CTXReceivingCompany
func TestBatchCTXReceivingCompany(t *testing.T) {
	testBatchCTXReceivingCompany(t)
}

// BenchmarkBatchCTXReceivingCompany benchmarks validating CTXReceivingCompany
func BenchmarkBatchCTXReceivingCompany(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXReceivingCompany(b)
	}
}

// testBatchCTXReserved validates CTXReservedField
func testBatchCTXReserved(t testing.TB) {
	mockBatch := mockBatchCTX()

	if mockBatch.GetEntries()[0].CATXReservedField() != "  " {
		t.Errorf("expected %v got %v", "  ", mockBatch.GetEntries()[0].CATXReservedField())
	}
}

// TestBatchCTXReserved tests validating CTXReservedField
func TestBatchCTXReserved(t *testing.T) {
	testBatchCTXReserved(t)
}

// BenchmarkBatchCTXReserved benchmarks validating CTXReservedField
func BenchmarkBatchCTXReserved(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXReserved(b)
	}
}

// testBatchCTXZeroAddendaRecords validates zero addenda records
func testBatchCTXZeroAddendaRecords(t testing.TB) {
	bh := NewBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.StandardEntryClassCode = CTX
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "ACH CTX"
	bh.ODFIIdentification = "12104288"

	entry := NewEntryDetail()
	entry.TransactionCode = CheckingCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.IdentificationNumber = "45689033"
	entry.SetCATXAddendaRecords(1)
	entry.SetCATXReceivingCompany("Receiver Company")
	entry.SetTraceNumber(mockBatchCTXHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.Category = CategoryForward

	mockBatch := NewBatchCTX(bh)
	mockBatch.AddEntry(entry)

	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addendum" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCTXZeroAddendaRecords tests validating zero addenda records
func TestBatchCTXZeroAddendaRecords(t *testing.T) {
	testBatchCTXZeroAddendaRecords(t)
}

// BenchmarkBatchZeroAddendaRecords benchmarks validating zero addenda records
func BenchmarkBatchCTXZeroAddendaRecords(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXZeroAddendaRecords(b)
	}
}

// testBatchCTXPrenoteAddendaRecords validates prenote addenda records
func testBatchCTXPrenoteAddendaRecords(t testing.TB) {
	bh := NewBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.StandardEntryClassCode = CTX
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "ACH CTX"
	bh.ODFIIdentification = "12104288"
	bh.OriginatorStatusCode = 2

	entry := NewEntryDetail()
	entry.TransactionCode = CheckingPrenoteCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.IdentificationNumber = "45689033"
	entry.SetCATXAddendaRecords(1)
	entry.SetCATXReceivingCompany("Receiver Company")
	entry.SetTraceNumber(mockBatchCTXHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.Category = CategoryForward

	mockBatch := NewBatchCTX(bh)
	mockBatch.AddEntry(entry)
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addendum" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCTXPrenoteAddendaRecords tests validating prenote addenda records
func TestBatchCTXPrenoteAddendaRecords(t *testing.T) {
	testBatchCTXPrenoteAddendaRecords(t)
}

// BenchmarkBatchPrenoteAddendaRecords benchmarks validating prenote addenda records
func BenchmarkBatchCTXPrenoteAddendaRecords(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXPrenoteAddendaRecords(b)
	}
}

// TestBatchCTXAddendum98 validates Addenda98 returns an error
func TestBatchCTXAddendum98(t *testing.T) {
	mockBatch := NewBatchCTX(mockBatchCTXHeader())
	mockBatch.AddEntry(mockCTXEntryDetail())
	mockAddenda98 := mockAddenda98()
	mockAddenda98.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.GetEntries()[0].Addenda98 = mockAddenda98
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCTXAddendum99 validates Addenda99 returns an error
func TestBatchCTXAddendum99(t *testing.T) {
	mockBatch := NewBatchCTX(mockBatchCTXHeader())
	mockBatch.AddEntry(mockCTXEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockAddenda99.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryReturn
	mockBatch.GetEntries()[0].Addenda99 = mockAddenda99
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCTXValidTranCodeForServiceClassCode validates a transactionCode based on ServiceClassCode
func TestBatchCTXValidTranCodeForServiceClassCode(t *testing.T) {
	mockBatch := mockBatchCTX()
	mockBatch.GetHeader().ServiceClassCode = DebitsOnly
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TransactionCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCTXAddenda02 validates BatchCTX cannot have Addenda02
func TestBatchCTXAddenda02(t *testing.T) {
	mockBatch := mockBatchCTX()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].Addenda02 = mockAddenda02()
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addenda02" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}
