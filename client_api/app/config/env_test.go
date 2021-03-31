package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEnv_loanEnvFile(t *testing.T) {
	tEnv := newEnv()

	// when specified env file exist
	// *****************************************************************
	err := tEnv.loadEnvFile(".env.example")
	if err != nil {
		t.Fatalf(".env.example always present, err must be nil: %v", err)
	}

	// when specified env file does not exist
	// *****************************************************************
	err = tEnv.loadEnvFile(".something_that_not_exist")
	if err == nil {
		t.Fatalf("err must not be nil when passing a filename that doesn't exist")
	}

	// when empty string passed as filename
	// *****************************************************************
	err = tEnv.loadEnvFile("")
	if err == nil {
		t.Fatalf("should error when empty filename passed to function")
		return
	}
	assert.Equal(
		t,
		"empty filename, specify the file that holds the environment variables",
		err.Error(),
	)
}

func TestEnv_loadValues(t *testing.T) {
	tEnv := newEnv()

	err := os.Setenv("GOLANG_ENV", "valueOne")
	err = os.Setenv("PORT", "3001")

	if err != nil {
		t.Fatalf("error loading env variables: %v", err)
	}

	tEnv.loadValues(&tEnv)

	assert.Equal(t, "valueOne", tEnv.GolangEnv)
	assert.Equal(t, 3001, tEnv.Port)
}

func TestEnv_Load(t *testing.T) {
	tEnv := newEnv()
	err := tEnv.load()
	if err != nil {
		t.Fatalf("should not error at this point, %v", err)
	}

	// when GOLANG_ENV not set or different from 'development' or 'tests'
	// *****************************************************************
	err = os.Setenv("GOLANG_ENV", "")
	if err != nil {
		t.Fatal(err)
	}
	var calledLoadEnvFile bool
	var calledWithFilename string
	tEnv.loadEnvFile = func(filename string) error {
		calledLoadEnvFile = true
		calledWithFilename = filename
		return nil
	}

	err = tEnv.load()
	if err != nil {
		t.Fatalf("should not error at this point, %v", err)
	}
	assert.Equal(t, false, calledLoadEnvFile)
	assert.Equal(t, "", calledWithFilename)

	// when GOLANG_ENV set to development
	// *****************************************************************
	err = os.Setenv("GOLANG_ENV", "development")
	if err != nil {
		t.Fatal(err)
	}
	err = tEnv.load()
	if err != nil {
		t.Fatalf("should not error at this point, %v", err)
	}
	assert.Equal(t, true, calledLoadEnvFile)
	assert.Equal(t, ".env", calledWithFilename)

	// when GOLANG_ENV set to development
	// *****************************************************************
	err = os.Setenv("GOLANG_ENV", "test")
	if err != nil {
		t.Fatal(err)
	}
	err = tEnv.load()
	if err != nil {
		t.Fatalf("should not error at this point, %v", err)
	}
	assert.Equal(t, true, calledLoadEnvFile)
	assert.Equal(t, ".env.test", calledWithFilename)

	//when loadEnvFile errors
	// *****************************************************************
	errorMessage := "erroring for testing"
	tEnv.loadEnvFile = func(filename string) error {
		return fmt.Errorf(errorMessage)
	}
	err = tEnv.load()
	if err == nil {
		t.Fatal("should error at this point")
	}

	assert.Equal(t, errorMessage, err.Error())
}
