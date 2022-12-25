package database

import (
	"reflect"
	"testing"

	"github.com/palavrapasse/import/internal/entity"
)

func TestTableNameReturnsUnknownTableNameIfItDoesNotContainAnyRecord(t *testing.T) {
	tb := DatabaseTable{}

	name := tb.Name()

	if name != unknownTableName {
		t.Fatalf("Table does not contain any record, so it's name can't be inferred, but got: %s\n", name)
	}
}

func TestBadActorTableNameReturnsBadActor(t *testing.T) {
	tb := BadActorTable{Records: Records{entity.BadActor{}}}
	expectedTableName := "BadActor"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("BadActor table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestCredentialsTableNameReturnsCredentials(t *testing.T) {
	tb := CredentialsTable{Records: Records{entity.Credentials{}}}
	expectedTableName := "Credentials"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("Credentials table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestHashCredentialsTableNameReturnsHashCredentials(t *testing.T) {
	tb := HashCredentialsTable{Records: Records{entity.HashCredentials{}}}
	expectedTableName := "HashCredentials"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("HashCredentials table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestHashUserTableNameReturnsHashUser(t *testing.T) {
	tb := HashUserTable{Records: Records{entity.HashUser{}}}
	expectedTableName := "HashUser"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("HashUser table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestLeakBadActorTableNameReturnsLeakBadActor(t *testing.T) {
	tb := LeakBadActorTable{Records: Records{entity.LeakBadActor{}}}
	expectedTableName := "LeakBadActor"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("LeakBadActor table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestLeakCredentialsTableNameReturnsLeakCredentials(t *testing.T) {
	tb := LeakCredentialsTable{Records: Records{entity.LeakCredentials{}}}
	expectedTableName := "LeakCredentials"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("LeakCredentials table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestLeakPlatformTableNameReturnsLeakPlatform(t *testing.T) {
	tb := LeakPlatformTable{Records: Records{entity.LeakPlatform{}}}
	expectedTableName := "LeakPlatform"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("LeakPlatform table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestLeakTableNameReturnsLeak(t *testing.T) {
	tb := LeakTable{Records: Records{entity.Leak{}}}
	expectedTableName := "Leak"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("Leak table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestLeakUserTableNameReturnsLeakUser(t *testing.T) {
	tb := LeakUserTable{Records: Records{entity.LeakUser{}}}
	expectedTableName := "LeakUser"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("LeakUser table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestPlatformTableNameReturnsPlatform(t *testing.T) {
	tb := PlatformTable{Records: Records{entity.Platform{}}}
	expectedTableName := "Platform"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("Platform table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestUserCredentialsTableNameReturnsUserCredentials(t *testing.T) {
	tb := UserCredentialsTable{Records: Records{entity.UserCredentials{}}}
	expectedTableName := "UserCredentials"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("UserCredentials table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestUserTableNameReturnsUser(t *testing.T) {
	tb := UserTable{Records: Records{entity.User{}}}
	expectedTableName := "User"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("User table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestTableFieldsReturnsUnknownTableIfItDoesNotContainAnyRecord(t *testing.T) {
	tb := DatabaseTable{}

	fields := tb.Fields()

	if !reflect.DeepEqual(fields, unknownTableFields) {
		t.Fatalf("Table does not contain any record, so it's fields can't be inferred, but got: %v\n", fields)
	}
}

func TestBadActorTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := BadActorTable{Records: Records{entity.BadActor{}}}

	expectedInsertStatement := "INSERT OR IGNORE INTO BadActor (identifier) VALUES (?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestCredentialsTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := CredentialsTable{Records: Records{entity.Credentials{}}}

	expectedInsertStatement := "INSERT OR IGNORE INTO Credentials (password) VALUES (?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestHashCredentialsTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := HashCredentialsTable{Records: Records{entity.HashCredentials{}}}

	expectedInsertStatement := "INSERT OR IGNORE INTO HashCredentials (credid, hsha256) VALUES (?, ?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestHashUserTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := HashUserTable{Records: Records{entity.HashUser{}}}

	expectedInsertStatement := "INSERT OR IGNORE INTO HashUser (userid, hsha256) VALUES (?, ?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestLeakBadActorTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := LeakBadActorTable{Records: Records{entity.LeakBadActor{}}}

	expectedInsertStatement := "INSERT OR IGNORE INTO LeakBadActor (baid, leakid) VALUES (?, ?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestLeakCredentialsTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := LeakCredentialsTable{Records: Records{entity.LeakCredentials{}}}

	expectedInsertStatement := "INSERT OR IGNORE INTO LeakCredentials (credid, leakid) VALUES (?, ?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestLeakPlatformTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := LeakPlatformTable{Records: Records{entity.LeakPlatform{}}}

	expectedInsertStatement := "INSERT OR IGNORE INTO LeakPlatform (platid, leakid) VALUES (?, ?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestLeakTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := LeakTable{Records: Records{entity.Leak{}}}

	expectedInsertStatement := "INSERT OR IGNORE INTO Leak (sharedatesc, context) VALUES (?, ?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestLeakUserTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := LeakUserTable{Records: Records{entity.LeakUser{}}}

	expectedInsertStatement := "INSERT OR IGNORE INTO LeakUser (userid, leakid) VALUES (?, ?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestPlatformTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := PlatformTable{Records: Records{entity.Platform{}}}

	expectedInsertStatement := "INSERT OR IGNORE INTO Platform (name) VALUES (?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestUserCredentialsTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := UserCredentialsTable{Records: Records{entity.UserCredentials{}}}

	expectedInsertStatement := "INSERT OR IGNORE INTO UserCredentials (credid, userid) VALUES (?, ?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestUserTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := UserTable{Records: Records{entity.User{}}}

	expectedInsertStatement := "INSERT OR IGNORE INTO User (email) VALUES (?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestCopyReturnsDatabaseTableWithNewRecords(t *testing.T) {
	oldRecords := Records{entity.User{UserId: 1}}
	newRecords := Records{entity.User{UserId: 2}}

	oldTable := UserTable{Records: oldRecords}
	expectedNewTable := UserTable{Records: newRecords}

	copyTable := oldTable.Copy(newRecords)

	if !reflect.DeepEqual(copyTable.Records, expectedNewTable.Records) {
		t.Fatalf("Copy should have return a database table with new records, but got: %v", copyTable.Records)
	}
}
