package models

type User struct {
	Username string
	Password string
}

type MultiAuthResult struct {
	Successed      int             `json:"successed"`
	Failed         int             `json:"failed"`
	FailedMessages []FailedMessage `json:"failed_messages"`
}

type FailedMessage struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type MultiSearchUserResult struct {
	Users          []LADP_RESULT   `json:"users"`
	Successed      int             `json:"successed"`
	Failed         int             `json:"failed"`
	FailedMessages []FailedMessage `json:"failed_messages"`
}

func SingleAuth(lc *LDAP_CONFIG, username, password string) (success bool, err error) {
	err = lc.Connect()
	defer lc.Close()
	if err != nil {
		return
	}

	success, err = lc.Auth(username, password)
	return
}

func MultiAuth(lc *LDAP_CONFIG, userList []User) (result MultiAuthResult, err error) {
	err = lc.Connect()
	defer lc.Close()
	if err != nil {
		return
	}

	var failedMsg FailedMessage
	for _, user := range userList {
		success, err := lc.Auth(user.Username, user.Password)
		if !success {
			result.Failed++
			failedMsg = FailedMessage{
				Username: user.Username,
				Message:  err.Error(),
			}
			result.FailedMessages = append(result.FailedMessages, failedMsg)
		} else {
			result.Successed++
		}
	}
	return
}

func HealthCheck(lc *LDAP_CONFIG) (success bool, err error) {
	err = lc.Connect()
	defer lc.Close()

	if err != nil {
		return
	}
	success = true
	return
}

func SingleSearchUser(lc *LDAP_CONFIG, username string) (user LADP_RESULT, err error) {
	err = lc.Connect()
	defer lc.Close()
	if err != nil {
		return
	}

	user, err = lc.SearchUser(username)
	return
}

func MultiSearchUser(lc *LDAP_CONFIG, userList []string) (result MultiSearchUserResult, err error) {
	err = lc.Connect()
	defer lc.Close()
	if err != nil {
		return
	}

	var failedMsg FailedMessage
	for _, username := range userList {
		user, err := lc.SearchUser(username)
		if err != nil {
			result.Failed++
			failedMsg = FailedMessage{
				Username: username,
				Message:  err.Error(),
			}
			result.FailedMessages = append(result.FailedMessages, failedMsg)
		} else {
			result.Users = append(result.Users, user)
			result.Successed++
		}
	}
	return
}

func SingleSearch(lc *LDAP_CONFIG, searchFilter string) (results []LADP_RESULT, err error) {
	err = lc.Connect()
	defer lc.Close()
	if err != nil {
		return
	}

	results, err = lc.Search(searchFilter)
	return
}
