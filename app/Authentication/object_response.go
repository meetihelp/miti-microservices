package Authentication

type PreferenceContent struct{
	Code int `json:"Code"`
	Message string `json:"Message"`
	Preference int `json:"Preference"`
}