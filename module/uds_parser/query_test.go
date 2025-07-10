package uds_parser

import (
	"context"
	"database/sql"
	uds "explain_uds/common"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"testing"
)

var testCtx context.Context

func TestMain(m *testing.M) {
	// Setup only 1 test database and re-use in all tests
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
	CREATE TABLE services (
		sid TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT,
		has_subfunction BOOLEAN NOT NULL DEFAULT 0,
		positive_response INTEGER
	);

	CREATE TABLE sub_functions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		sid TEXT NOT NULL,
		value TEXT NOT NULL,
		name TEXT NOT NULL,
		description TEXT,
		FOREIGN KEY (sid) REFERENCES services(sid)
	);

	INSERT INTO services (sid, name, description, has_subfunction, positive_response)
	VALUES 
		("0x10", "DiagnosticSession", "Start diagnostic session", 1, 80);

	INSERT INTO sub_functions (sid, value, name, description)
	VALUES
		("0x10", "0x01", "DefaultSession", "Default session desc."),
		("0x10", "0x02", "ProgrammingSession", "Programming desc.");
	`)
	if err != nil {
		log.Fatal(err)
	}

	testCtx = context.WithValue(context.Background(), uds.CtxKeyDB{}, db)
	code := m.Run() // execute all tests
	os.Exit(code)
}
func TestGetServiceById(t *testing.T) {
	t.Run("Valid SID", func(t *testing.T) {
		dto, err := GetServiceByID(testCtx, "0x10")
		if err != nil {
			t.Fatalf("GetServiceByID failed: %v", err)
		}
		if dto.SID != "0x10" {
			t.Errorf("Expected SID '0x10', got '%s'", dto.SID)
		}
		if dto.Name != "DiagnosticSession" {
			t.Errorf("Expected Name 'DiagnosticSession', got '%s'", dto.Name)
		}
		if dto.HasSubFunction != true {
			t.Errorf("Expected HasSubFunction true, got %v", dto.HasSubFunction)
		}
	})
	t.Run("Invalid SID", func(t *testing.T) {
		dto, err := GetServiceByID(testCtx, "0x99")
		if err != nil {
			t.Fatalf("GetServiceByID failed: %v", err)
		}
		if dto != nil {
			t.Errorf("Expected empty DTO for invalid SID, got %v", dto)
		}
	})

}

func TestGetSubfunctionById(t *testing.T) {}
