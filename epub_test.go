package epub

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/pirmd/verify"
)

const (
	testdataPath = "/data/ebooks/10/"
)

func TestGetPackageFromFile(t *testing.T) {
	testCases, err := filepath.Glob(filepath.Join(testdataPath, "*.epub"))
	if err != nil {
		t.Fatalf("cannot read test data in %s:%v", testdataPath, err)
	}

	out := []*PackageDocument{}
	for _, tc := range testCases {
		opf, err := GetPackageFromFile(tc)
		if err != nil {
			t.Errorf("Fail to get package for %s: %v", tc, err)
		}
		out = append(out, opf)
	}

	got, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		t.Fatalf("Fail to marshal test output to json: %v", err)
	}

	if failure := verify.MatchGolden(t.Name(), string(got)); failure != nil {
		t.Fatalf("Package Document is not as expected:\n%v", failure)
	}
}

func TestOpenItem(t *testing.T) {
	testCases, err := filepath.Glob(filepath.Join(testdataPath, "*.epub"))
	if err != nil {
		t.Fatalf("cannot read test data in %s:%v", testdataPath, err)
	}

	for _, tc := range testCases {
		opf, err := GetPackageFromFile(tc)
		if err != nil {
			t.Errorf("Fail to get package for %s: %v", tc, err)
		}

		items := opf.Manifest.Items
		found := false
		for i := range items {
			if items[i].ID == `cover-image` {
				fmt.Println(`book: `, opf.Metadata.Title[0].Value, `cover index : `, i, `len: `, len(items))
				found = true
				break
			}
		}

		if !found {
			fmt.Println(`book: `, opf.Metadata.Title[0].Value, `cover index : `, len(items), `len: `, len(items))

		}

		// out = append(out, opf)
	}

	// e, err := Open(testCases[2])
	// if err != nil {
	// 	t.Fatalf("cannot read test data in %s:%v", testdataPath, err)

	// }
	// defer e.Close()

	// pp, _ := e.Package()

	// items :=

	// 	println(pp != nil)
}
