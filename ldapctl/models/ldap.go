package models

import (
	"crypto/tls"
	"errors"
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

type LDAP_CONFIG struct {
	Addr       string     `json:"addr"`
	BaseDn     string     `json:"base_dn"`
	BindDn     string     `json:"bind_dn"`
	BindPass   string     `json:"bind_pass"`
	AuthFilter string     `json:"auth_filter"`
	Attributes []string   `json:"attributes"`
	TLS        bool       `json:"tls"`
	StartTLS   bool       `json:"start_tls"`
	Conn       *ldap.Conn `json:"-"`
}

type LADP_RESULT struct {
	DN         string              `json:"dn"`
	Attributes map[string][]string `json:"attributes"`
}

func (lc *LDAP_CONFIG) SearchUser(username string) (user LADP_RESULT, err error) {
	searchRequest := ldap.NewSearchRequest(
		lc.BaseDn,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(lc.AuthFilter, username),
		lc.Attributes,
		nil,
	)

	sr, err := lc.Conn.Search(searchRequest)
	if err != nil {
		return
	}
	if len(sr.Entries) == 0 {
		err = errors.New("could not find any user")
		return
	}
	if len(sr.Entries) > 1 {
		err = errors.New("found multiple users")
		return
	}

	attributes := make(map[string][]string)
	for _, attr := range sr.Entries[0].Attributes {
		attributes[attr.Name] = attr.Values
	}
	user.DN = sr.Entries[0].DN
	user.Attributes = attributes

	return
}

func (lc *LDAP_CONFIG) Search(searchFilter string) (results []LADP_RESULT, err error) {
	searchRequest := ldap.NewSearchRequest(
		lc.BaseDn,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		searchFilter,
		lc.Attributes,
		nil,
	)

	sr, err := lc.Conn.Search(searchRequest)
	if err != nil {
		return
	}
	if len(sr.Entries) == 0 {
		err = errors.New("could not find any user")
		return
	}

	var result LADP_RESULT
	results = make([]LADP_RESULT, 0)
	for _, entry := range sr.Entries {
		attributes := make(map[string][]string)
		for _, attr := range entry.Attributes {
			attributes[attr.Name] = attr.Values
		}
		result.DN = entry.DN
		result.Attributes = attributes
		results = append(results, result)
	}

	return
}

func (lc *LDAP_CONFIG) Auth(username, password string) (success bool, err error) {
	searchRequest := ldap.NewSearchRequest(
		lc.BaseDn,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(lc.AuthFilter, username),
		lc.Attributes,
		nil,
	)
	sr, err := lc.Conn.Search(searchRequest)
	if err != nil {
		return
	}
	if len(sr.Entries) == 0 {
		err = errors.New("could not find any user")
		return
	}
	if len(sr.Entries) > 1 {
		err = errors.New("found multiple users")
		return
	}

	err = lc.Conn.Bind(sr.Entries[0].DN, password)
	if err != nil {
		return
	}

	//Rebind as the search user for any further queries
	err = lc.Conn.Bind(lc.BindDn, lc.BindPass)
	if err != nil {
		return
	}

	success = true
	return
}

func (lc *LDAP_CONFIG) Connect() (err error) {
	if lc.TLS {
		lc.Conn, err = ldap.DialURL(lc.Addr, ldap.DialWithTLSConfig(&tls.Config{InsecureSkipVerify: true}))
	} else {
		lc.Conn, err = ldap.DialURL(lc.Addr)
	}
	if err != nil {
		return err
	}

	if !lc.TLS && lc.StartTLS {
		err = lc.Conn.StartTLS(&tls.Config{InsecureSkipVerify: true})
		if err != nil {
			lc.Conn.Close()
			return err
		}
	}

	err = lc.Conn.Bind(lc.BindDn, lc.BindPass)
	if err != nil {
		lc.Conn.Close()
		return err
	}
	return nil
}

func (lc *LDAP_CONFIG) Close() {
	if lc.Conn != nil {
		lc.Conn.Close()
		lc.Conn = nil
	}
}
