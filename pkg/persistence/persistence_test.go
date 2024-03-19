package persistence

import (
	"testing"
)

func TestGetDb(t *testing.T) {
	db := GetDb(":memory:")
	defer db.Close()

	if db.dbHandle == nil {
		t.Error("Expected non-nil dbHandle, got nil")
	}
}

func TestStoreAndGetAll(t *testing.T) {
	db := GetDb(":memory:")
	db.Initialise()
	defer db.Close()

	fqdn := "example.com"
	err := db.Store(fqdn)
	if err != nil {
		t.Errorf("Failed to store FQDN: %v", err)
	}

	fqdns, err := db.GetAllFqdns()

	if err != nil {
		t.Errorf("Failed to get FQDNs: %v", err)
	}

	if len(fqdns) != 1 {
		t.Errorf("Expected 1 FQDN, got %d", len(fqdns))
	}

	if fqdns[0] != fqdn {
		t.Errorf("Expected %s, got %s", fqdn, fqdns[0])
	}
}
