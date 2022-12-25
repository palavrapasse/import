package database

import (
	"reflect"
	"testing"

	"github.com/palavrapasse/import/internal/entity"
)

func TestValuesReturnsSchemaValuesIfRecordIsBadActor(t *testing.T) {
	r := entity.BadActor{BaId: 1, Identifier: "l33t"}

	expectedValues := []Value{r.BaId, r.Identifier}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsCredentials(t *testing.T) {
	r := entity.Credentials{CredId: 1, Password: "my.password"}

	expectedValues := []Value{r.CredId, r.Password}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsHashCredentials(t *testing.T) {
	r := entity.HashCredentials{CredId: 1, HSHA256: entity.NewHSHA256("my.password")}

	expectedValues := []Value{r.CredId, r.HSHA256}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsHashUser(t *testing.T) {
	r := entity.HashUser{UserId: 1, HSHA256: entity.NewHSHA256("my.email@gmail.com")}

	expectedValues := []Value{r.UserId, r.HSHA256}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsLeakBadActor(t *testing.T) {
	r := entity.LeakBadActor{BaId: 1, LeakId: 2}

	expectedValues := []Value{r.BaId, r.LeakId}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsLeakCredentials(t *testing.T) {
	r := entity.LeakCredentials{CredId: 1, LeakId: 2}

	expectedValues := []Value{r.CredId, r.LeakId}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsLeakPlatform(t *testing.T) {
	r := entity.LeakPlatform{PlatId: 1, LeakId: 2}

	expectedValues := []Value{r.PlatId, r.LeakId}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsLeak(t *testing.T) {
	r := entity.Leak{LeakId: 1, ShareDateSC: entity.DateInSeconds(2), Context: "twitter breach"}

	expectedValues := []Value{r.LeakId, r.ShareDateSC, r.Context}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsLeakUser(t *testing.T) {
	r := entity.LeakUser{UserId: 1, LeakId: 2}

	expectedValues := []Value{r.UserId, r.LeakId}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsPlatform(t *testing.T) {
	r := entity.Platform{PlatId: 1, Name: "twitter"}

	expectedValues := []Value{r.PlatId, r.Name}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsUserCredentials(t *testing.T) {
	r := entity.UserCredentials{CredId: 1, UserId: 2}

	expectedValues := []Value{r.CredId, r.UserId}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestValuesReturnsSchemaValuesIfRecordIsUser(t *testing.T) {
	r := entity.User{UserId: 1, Email: "my.email@gmail.com"}

	expectedValues := []Value{r.UserId, r.Email}

	values := Values(r)

	if !reflect.DeepEqual(values, expectedValues) {
		t.Fatalf("Values should have return slice with values as defined in schema, but got: %v", values)
	}
}

func TestCopyWithNewKeyReturnsRecordWithAutoGenKeySetIfRecordIsBadActor(t *testing.T) {
	r := entity.BadActor{}
	k := entity.AutoGenKey(500)
	expectedRecord := r.Copy(k)

	copyRecord := CopyWithNewKey(r, k)

	if copyRecord != expectedRecord {
		t.Fatalf("CopyWithNewKey should have set auto gen key via Copy method, but match failed: %v\n", copyRecord)
	}
}

func TestCopyWithNewKeyReturnsRecordWithAutoGenKeySetIfRecordIsCredentials(t *testing.T) {
	r := entity.Credentials{}
	k := entity.AutoGenKey(500)
	expectedRecord := r.Copy(k)

	copyRecord := CopyWithNewKey(r, k)

	if copyRecord != expectedRecord {
		t.Fatalf("CopyWithNewKey should have set auto gen key via Copy method, but match failed: %v\n", copyRecord)
	}
}

func TestCopyWithNewKeyReturnsRecordWithAutoGenKeySetIfRecordIsLeak(t *testing.T) {
	r := entity.Leak{}
	k := entity.AutoGenKey(500)
	expectedRecord := r.Copy(k)

	copyRecord := CopyWithNewKey(r, k)

	if copyRecord != expectedRecord {
		t.Fatalf("CopyWithNewKey should have set auto gen key via Copy method, but match failed: %v\n", copyRecord)
	}
}

func TestCopyWithNewKeyReturnsRecordWithAutoGenKeySetIfRecordIsPlatform(t *testing.T) {
	r := entity.Platform{}
	k := entity.AutoGenKey(500)
	expectedRecord := r.Copy(k)

	copyRecord := CopyWithNewKey(r, k)

	if copyRecord != expectedRecord {
		t.Fatalf("CopyWithNewKey should have set auto gen key via Copy method, but match failed: %v\n", copyRecord)
	}
}

func TestCopyWithNewKeyReturnsRecordWithAutoGenKeySetIfRecordIsUser(t *testing.T) {
	r := entity.User{}
	k := entity.AutoGenKey(500)
	expectedRecord := r.Copy(k)

	copyRecord := CopyWithNewKey(r, k)

	if copyRecord != expectedRecord {
		t.Fatalf("CopyWithNewKey should have set auto gen key via Copy method, but match failed: %v\n", copyRecord)
	}
}

func TestCopyWithNewKeyReturnsSameRecordIfRecordIsNotPrimary(t *testing.T) {
	r := entity.LeakCredentials{}
	k := entity.AutoGenKey(500)
	expectedRecord := r

	copyRecord := CopyWithNewKey(r, k)

	if copyRecord != expectedRecord {
		t.Fatalf("CopyWithNewKey should have not set auto gen key via Copy method, but match failed: %v\n", copyRecord)
	}
}

// func TestBadActorTableNameReturnsBadActor(t *testing.T) {
// 	tb := BadActorTable{Records: Records{entity.BadActor{}}}
// 	expectedTableName := "BadActor"

// 	name := tb.Name()

// 	if name != expectedTableName {
// 		t.Fatalf("BadActor table name in database is %s, but Name() returned: %s", expectedTableName, name)
// 	}
// }
