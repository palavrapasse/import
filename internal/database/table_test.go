package database

import (
	"reflect"
	"testing"

	. "github.com/palavrapasse/import/internal/entity"
)

func TestTableNameReturnsUnknownTableNameIfItDoesNotContainAnyRecord(t *testing.T) {
	tb := DatabaseTable{}

	name := tb.Name()

	if name != unknownTableName {
		t.Fatalf("Table does not contain any record, so it's name can't be inferred, but got: %s\n", name)
	}
}

func TestBadActorTableNameReturnsBadActor(t *testing.T) {
	tb := NewBadActorTable([]BadActor{{}})
	expectedTableName := "BadActor"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("BadActor table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestCredentialsTableNameReturnsCredentials(t *testing.T) {
	tb := NewCredentialsTable([]Credentials{{}})
	expectedTableName := "Credentials"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("Credentials table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestHashCredentialsTableNameReturnsHashCredentials(t *testing.T) {
	tb := NewHashCredentialsTable([]Credentials{{}})
	expectedTableName := "HashCredentials"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("HashCredentials table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestHashUserTableNameReturnsHashUser(t *testing.T) {
	tb := NewHashUserTable([]User{{}})
	expectedTableName := "HashUser"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("HashUser table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestLeakBadActorTableNameReturnsLeakBadActor(t *testing.T) {
	tb := NewLeakBadActorTable(map[Leak][]BadActor{{}: {BadActor{}}})
	expectedTableName := "LeakBadActor"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("LeakBadActor table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestLeakCredentialsTableNameReturnsLeakCredentials(t *testing.T) {
	tb := NewLeakCredentialsTable(map[Leak][]Credentials{{}: {Credentials{}}})
	expectedTableName := "LeakCredentials"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("LeakCredentials table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestLeakPlatformTableNameReturnsLeakPlatform(t *testing.T) {
	tb := NewLeakPlatformTable(map[Leak][]Platform{{}: {Platform{}}})
	expectedTableName := "LeakPlatform"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("LeakPlatform table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestLeakTableNameReturnsLeak(t *testing.T) {
	tb := NewLeakTable(Leak{})
	expectedTableName := "Leak"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("Leak table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestLeakUserTableNameReturnsLeakUser(t *testing.T) {
	tb := NewLeakUserTable(map[Leak][]User{{}: {User{}}})
	expectedTableName := "LeakUser"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("LeakUser table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestPlatformTableNameReturnsPlatform(t *testing.T) {
	tb := NewPlatformTable([]Platform{{}})
	expectedTableName := "Platform"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("Platform table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestUserCredentialsTableNameReturnsUserCredentials(t *testing.T) {
	tb := NewUserCredentialsTable(map[User]Credentials{{}: {}})
	expectedTableName := "UserCredentials"

	name := tb.Name()

	if name != expectedTableName {
		t.Fatalf("UserCredentials table name in database is %s, but Name() returned: %s", expectedTableName, name)
	}
}

func TestUserTableNameReturnsUser(t *testing.T) {
	tb := NewUserTable([]User{{}})
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
	tb := NewBadActorTable([]BadActor{{}})

	expectedInsertStatement := "INSERT OR IGNORE INTO BadActor (identifier) VALUES (?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestCredentialsTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := NewCredentialsTable([]Credentials{{}})

	expectedInsertStatement := "INSERT OR IGNORE INTO Credentials (password) VALUES (?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestHashCredentialsTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := NewHashCredentialsTable([]Credentials{{}})

	expectedInsertStatement := "INSERT OR IGNORE INTO HashCredentials (credid, hsha256) VALUES (?, ?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestHashUserTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := NewHashUserTable([]User{{}})

	expectedInsertStatement := "INSERT OR IGNORE INTO HashUser (userid, hsha256) VALUES (?, ?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestLeakBadActorTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := NewLeakBadActorTable(map[Leak][]BadActor{{}: {BadActor{}}})

	expectedInsertStatement := "INSERT OR IGNORE INTO LeakBadActor (baid, leakid) VALUES (?, ?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestLeakCredentialsTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := NewLeakCredentialsTable(map[Leak][]Credentials{{}: {Credentials{}}})

	expectedInsertStatement := "INSERT OR IGNORE INTO LeakCredentials (credid, leakid) VALUES (?, ?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestLeakPlatformTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := NewLeakPlatformTable(map[Leak][]Platform{{}: {Platform{}}})

	expectedInsertStatement := "INSERT OR IGNORE INTO LeakPlatform (platid, leakid) VALUES (?, ?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestLeakTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := NewLeakTable(Leak{})

	expectedInsertStatement := "INSERT OR IGNORE INTO Leak (sharedatesc, context) VALUES (?, ?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestLeakUserTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := NewLeakUserTable(map[Leak][]User{{}: {User{}}})

	expectedInsertStatement := "INSERT OR IGNORE INTO LeakUser (userid, leakid) VALUES (?, ?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestPlatformTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := NewPlatformTable([]Platform{{}})

	expectedInsertStatement := "INSERT OR IGNORE INTO Platform (name) VALUES (?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestUserCredentialsTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := NewUserCredentialsTable(map[User]Credentials{{}: {}})

	expectedInsertStatement := "INSERT OR IGNORE INTO UserCredentials (credid, userid) VALUES (?, ?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestUserTablePrepareInsertStatementReturnsSchemaInsertStatement(t *testing.T) {
	tb := NewUserTable([]User{{}})

	expectedInsertStatement := "INSERT OR IGNORE INTO User (email) VALUES (?)"

	insertStatement := tb.prepareInsertStatementString()

	if insertStatement != expectedInsertStatement {
		t.Fatalf("Prepared insert statement should be the same as defined in the schema, but got: %v", insertStatement)
	}
}

func TestPrimaryTableCopyReturnsDatabaseTableWithNewRecords(t *testing.T) {
	oldRecords := Records{User{UserId: 1}}
	newRecords := Records{User{UserId: 2}}

	oldTable := UserTable{Records: oldRecords}
	expectedNewTable := UserTable{Records: newRecords}

	copyTable := oldTable.Copy(newRecords)

	if !reflect.DeepEqual(copyTable.Records, expectedNewTable.Records) {
		t.Fatalf("Copy should have return a database table with new records, but got: %v", copyTable.Records)
	}
}

func TestForeignTableCopyReturnsDatabaseTableWithNewRecords(t *testing.T) {
	oldRecords := Records{LeakUser{UserId: 1}}
	newRecords := Records{LeakUser{UserId: 2}}

	oldTable := LeakUserTable{Records: oldRecords}
	expectedNewTable := LeakUserTable{Records: newRecords}

	copyTable := oldTable.Copy(newRecords)

	if !reflect.DeepEqual(copyTable.Records, expectedNewTable.Records) {
		t.Fatalf("Copy should have return a database table with new records, but got: %v", copyTable.Records)
	}
}

func TestToBadActorSliceReturnsTableBadActorRecords(t *testing.T) {
	expectedBadActorSlice := []BadActor{{BaId: 1}, {BaId: 2}}

	bat := NewBadActorTable(expectedBadActorSlice)
	bas := bat.ToBadActorSlice()

	if !reflect.DeepEqual(bas, expectedBadActorSlice) {
		t.Fatalf("ToBadActorSlice should have return a slice with all BadActor records, but got: %v", bas)
	}
}

func TestToCredentialsSliceReturnsTableCredentialsRecords(t *testing.T) {
	expectedCredentialsSlice := []Credentials{{CredId: 1}, {CredId: 2}}

	bat := NewCredentialsTable(expectedCredentialsSlice)
	bas := bat.ToCredentialsSlice()

	if !reflect.DeepEqual(bas, expectedCredentialsSlice) {
		t.Fatalf("ToCredentialsSlice should have return a slice with all Credentials records, but got: %v", bas)
	}
}

func TestToLeakSliceReturnsTableLeakRecords(t *testing.T) {
	expectedLeakSlice := []Leak{{LeakId: 1}, {LeakId: 2}}

	bat := NewLeakTable(expectedLeakSlice...)
	bas := bat.ToLeakSlice()

	if !reflect.DeepEqual(bas, expectedLeakSlice) {
		t.Fatalf("ToLeakSlice should have return a slice with all Leak records, but got: %v", bas)
	}
}

func TestToPlatformSliceReturnsTablePlatformRecords(t *testing.T) {
	expectedPlatformSlice := []Platform{{PlatId: 1}, {PlatId: 2}}

	bat := NewPlatformTable(expectedPlatformSlice)
	bas := bat.ToPlatformSlice()

	if !reflect.DeepEqual(bas, expectedPlatformSlice) {
		t.Fatalf("ToPlatformSlice should have return a slice with all Platform records, but got: %v", bas)
	}
}

func TestToUserSliceReturnsTableUserRecords(t *testing.T) {
	expectedUserSlice := []User{{UserId: 1}, {UserId: 2}}

	bat := NewUserTable(expectedUserSlice)
	bas := bat.ToUserSlice()

	if !reflect.DeepEqual(bas, expectedUserSlice) {
		t.Fatalf("ToUserSlice should have return a slice with all User records, but got: %v", bas)
	}
}
