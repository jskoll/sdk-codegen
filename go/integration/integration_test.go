package integration

import (
	"testing"

	"github.com/looker-open-source/sdk-codegen/go/rtl"
	v4 "github.com/looker-open-source/sdk-codegen/go/sdk/v4"
)

func TestCRUDuser(t *testing.T) {
	cfg, err := rtl.NewSettingsFromEnv()

	if err != nil {
		t.Errorf("TestCRUDuser() error getting settings from env error = %v", err)
	}

	sdk := v4.NewLookerSDK(rtl.NewAuthSession(cfg))

	t.Run("Create user", func(t *testing.T) {
		firstName := "John"
		lastName := "Doe"
		isDisabled := false
		locale := "fr"

		user, err := sdk.CreateUser(v4.WriteUser{
			FirstName: &firstName,
			LastName: &lastName,
			IsDisabled: &isDisabled,
			Locale: &locale,
		}, "", nil)

		if err != nil {
			t.Errorf("CreateUser() failed. error=%v", err)
		}
		if *user.FirstName != firstName {
			t.Errorf("Create user FirstName not the same. got=%v want=%v", *user.FirstName, firstName )
		}
		if *user.LastName != lastName {
			t.Errorf("Create user LastName not the same. got=%v want=%v", *user.LastName, lastName )
		}
		if *user.IsDisabled != isDisabled {
			t.Errorf("Create user LastName not the same. got=%v want=%v", *user.IsDisabled, isDisabled )
		}
		if *user.Locale != locale {
			t.Errorf("Create user Locale not the same. got=%v want=%v", *user.Locale, locale )
		}
	})
}