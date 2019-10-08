package environment

import (
	"os"
	"reflect"
	"testing"
)

const TestEnvVar = "ENV_VAR_NAME"

func TestGetVar(t *testing.T) {
	stringVar := "test"

	setTestEnv(stringVar, t)
	defer unsetTestEnv(t)

	stringEnvVar := GetVar(TestEnvVar, "")
	if stringEnvVar != stringVar {
		t.Fatalf("expected %s got %s", stringVar, stringEnvVar)
	}
}

func TestGetVarAsInt(t *testing.T) {
	setTestEnv("999", t)
	defer unsetTestEnv(t)

	intEnvVar := GetVarAsInt(TestEnvVar, 0)
	if intEnvVar != 999 {
		t.Fatalf("expected %d got %d", 999, intEnvVar)
	}
}

func TestGetVarAsBool_true(t *testing.T) {
	setTestEnv("true", t)
	defer unsetTestEnv(t)

	boolEnvVar := GetVarAsBool(TestEnvVar, false)
	if boolEnvVar != true {
		t.Fatalf("expected %t got %t", true, boolEnvVar)
	}
}

func TestGetVarAsBool_false(t *testing.T) {
	setTestEnv("false", t)
	defer unsetTestEnv(t)

	boolEnvVar := GetVarAsBool(TestEnvVar, true)
	if boolEnvVar != false {
		t.Fatalf("expected %t got %t", false, boolEnvVar)
	}
}

func TestGetVarAsSlice(t *testing.T) {
	setTestEnv("one,two,three", t)
	defer unsetTestEnv(t)

	sliceEnvVar := GetVarAsSlice(TestEnvVar, make([]string, 0), ",")
	if len(sliceEnvVar) != 3 {
		t.Fatal("We expected there to be 3 elements in the env var slice")
	}
	expected := []string{"one", "two", "three"}
	if !reflect.DeepEqual(expected, sliceEnvVar) {
		t.Fatalf("expected %v got %v", expected, sliceEnvVar)
	}
}

func setTestEnv(v string, t *testing.T) {
	err := os.Setenv(TestEnvVar, v)
	if err != nil {
		t.Error("Failed to set env var for test")
	}
}

func unsetTestEnv(t *testing.T) {
	err := os.Unsetenv(TestEnvVar)
	if err != nil {
		t.Errorf("Failed to unset env var %s after test", TestEnvVar)
	}
}
