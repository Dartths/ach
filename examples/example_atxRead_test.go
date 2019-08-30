package examples

import (
	"fmt"
	"github.com/moov-io/ach"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func Example_atxRead() {
	f, err := os.Open(filepath.Join("testdata", "atx-read.ach"))
	if err != nil {
		log.Fatal(err)
	}
	r := ach.NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		fmt.Printf("Issue reading file: %+v \n", err)
	}
	// ensure we have a validated file structure
	if achFile.Validate(); err != nil {
		fmt.Printf("Could not validate entire read file: %v", err)
	}
	// If you trust the file but it's formatting is off building will probably resolve the malformed file.
	if err := achFile.Create(); err != nil {
		fmt.Printf("Could not create file with read properties: %v", err)
	}

	fmt.Printf("Total Amount Debit: %s", strconv.Itoa(achFile.Control.TotalDebitEntryDollarAmountInFile)+"\n")
	fmt.Printf("Total Amount Credit: %s", strconv.Itoa(achFile.Control.TotalCreditEntryDollarAmountInFile)+"\n")
	fmt.Printf("SEC Code: %s", achFile.Batches[0].GetHeader().StandardEntryClassCode+"\n")
	fmt.Printf("Total Amount: %s", strconv.Itoa(achFile.Batches[0].GetEntries()[0].Amount)+"\n")
	fmt.Printf("Original Trace Number: %s", achFile.Batches[0].GetEntries()[0].OriginalTraceNumberField()+"\n")
	fmt.Printf("Addenda1: %s", achFile.Batches[0].GetEntries()[0].Addenda05[0].String()+"\n")
	fmt.Printf("Addenda2: %s", achFile.Batches[0].GetEntries()[0].Addenda05[1].String()+"\n")
	fmt.Printf("Total Amount: %s", strconv.Itoa(achFile.Batches[0].GetEntries()[1].Amount)+"\n")
	fmt.Printf("Original Trace Number: %s", achFile.Batches[0].GetEntries()[1].OriginalTraceNumberField()+"\n")
	fmt.Printf("Addenda1: %s", achFile.Batches[0].GetEntries()[1].Addenda05[0].String()+"\n")
	fmt.Printf("Addenda2: %s", achFile.Batches[0].GetEntries()[1].Addenda05[1].String()+"\n")

	// Output:
	// Total Amount Debit: 0
	// Total Amount Credit: 0
	// SEC Code: ATX
	// Total Amount: 0
	// Original Trace Number: 031300010000001
	// Addenda1: 705Credit account 1 for service                                                    00010000001
	// Addenda2: 705Credit account 2 for service                                                    00020000001
	// Total Amount: 0
	// Original Trace Number: 031300010000002
	// Addenda1: 705Credit account 1 for leadership                                                 00010000002
	// Addenda2: 705Credit account 2 for leadership                                                 00020000002
}