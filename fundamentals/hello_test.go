package main

// go test

import "testing"

func TestHello(t *testing.T) {
	got := Hello("James")
	want := "Hello, James"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

/**
Writing a test is just like writing a function, with a few rules

It needs to be in a file with a name like xxx_test.go

The test function must start with the word Test

The test function takes one argument only t *testing.T

To use the *testing.T type, you need to import "testing", like we did with fmt in the other file

For now, it's enough to know that your t of type *testing.T is your "hook" into the testing framework so you can do things like t.Fail() when you want to fail.
*/

func TestCleverHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := CleverHello("James")
		want := "Hello, James"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := CleverHello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
