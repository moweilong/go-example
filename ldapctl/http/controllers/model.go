package controllers

import "ldapctl/models"

type MsgResult struct {
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
}

type AuthResult struct {
	Success bool                   `json:"success"`
	Result  models.MultiAuthResult `json:"result"`
}

type SearchResult struct {
	Results []models.LADP_RESULT `json:"results"`
	Success bool                 `json:"success"`
}
