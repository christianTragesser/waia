package system

import (
	"os"
	"testing"
)

type mockSystem struct {
	Hostname string
	Address  string
	Version  string
	BgColor  string
	TxtColor string
}

func (m *mockSystem) getIP() string {
	return m.Address
}

func (m *mockSystem) getColors() (string, string) {
	testSystem := System{}

	return testSystem.getColors()
}

func TestGetIP(t *testing.T) {
	t.Run("Get IP", func(t *testing.T) {
		expect := "2.2.2.2"

		testSystem := mockSystem{Address: expect}

		got := testSystem.getIP()

		if got != expect {
			t.Errorf("getIP unit test failure\ngot: '%v'\nwant: '%v", got, expect)
		}
	})

}

func TestGetColors(t *testing.T) {
	tests := []struct {
		name             string
		envVar           string
		expectBackground string
		expectText       string
	}{
		{
			name:             "No BG_COLOR set",
			envVar:           "",
			expectBackground: "white",
			expectText:       "black",
		},
		{
			name:             "BG_COLOR equals 'blue'",
			envVar:           "blue",
			expectBackground: "blue",
			expectText:       "white",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv("BG_COLOR", tc.envVar)

			testSystem := mockSystem{}

			gotBg, gotTxt := testSystem.getColors()

			if gotBg != tc.expectBackground || gotTxt != tc.expectText {
				t.Errorf("getColors unit test failure\ngot: '%v %v'\nwant: '%v %v'", gotBg, gotTxt, tc.expectBackground, tc.expectText)
			}
		})
	}
}
