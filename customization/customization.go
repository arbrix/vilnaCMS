package customization

import "github.com/arbrix/uadmin"

// ApplySystemCustomization invokes functions for changing default language, rootURL, UI theme, etc.
func ApplySystemCustomization() {
	setDefaultTheme()
	setRootURL()
	setDefaultLang()
}

func setDefaultTheme() {
	customTheme := "vilna"
	// NOTE: This code works only if database does not exist yet.
	if uadmin.Theme != customTheme {
		uadmin.Theme = customTheme
	}

	// ----- IF YOU RUN YOUR APPLICATION AGAIN, DO THIS BELOW -----

	// "vilna" is the name of the folder inside the templates/uadmin path
	// that uAdmin will run when the user starts the server
	setting := uadmin.Setting{}

	uadmin.Get(&setting, "code = ?", "uAdmin.Theme")
	if setting.Value != customTheme {
		setting.ParseFormValue([]string{customTheme})
		setting.Save()
	}
}

func setRootURL() {
	adminURL := "/admin/"
	// NOTE: This code works only if database does not exist yet.
	if uadmin.RootURL != adminURL {
		uadmin.RootURL = adminURL
	}

	// ----- IF YOU RUN YOUR APPLICATION AGAIN, DO THIS BELOW -----

	// Assign the Root URL value to /admin/
	setting := uadmin.Setting{}
	uadmin.Get(&setting, "code = ?", "uAdmin.RootURL")
	if setting.Value != adminURL {
		setting.ParseFormValue([]string{adminURL})
		setting.Save()
	}

}

func setDefaultLang() {
	ukLang := uadmin.Language{
		Code:           "uk",
		Default:        true,
		Active:         true,
		AvailableInGui: true,
		EnglishName:    "Ukrainian",
		Name:           "Українська",
	}

	// NOTE: This code works only if database does not exist yet.
	if uadmin.DefaultLang.Code != ukLang.Code {
		uadmin.DefaultLang = ukLang
		for i := range uadmin.ActiveLangs {
			if uadmin.ActiveLangs[i].Code == ukLang.Code {
				uadmin.ActiveLangs[i].Default = true
			} else {
				uadmin.ActiveLangs[i].Default = false
			}
		}
	}

	// ----- IF YOU RUN YOUR APPLICATION AGAIN, DO THIS BELOW -----
	defaultLang := uadmin.Language{}
	uadmin.Get(&defaultLang, `"default" = ?`, true)
	if defaultLang.Code != ukLang.Code {
		enLang := uadmin.Language{
			Code:           "en",
			Default:        false,
			AvailableInGui: false,
		}

		uadmin.Update(&enLang, "default", enLang.Default, "code = ?", enLang.Code)

		uadmin.Update(&ukLang, "default", ukLang.Default, "code = ?", ukLang.Code)
		uadmin.Update(&ukLang, "active", ukLang.Active, "code = ?", ukLang.Code)
		uadmin.Update(&ukLang, "available_in_gui", ukLang.AvailableInGui, "code = ?", ukLang.Code)
	}
}
