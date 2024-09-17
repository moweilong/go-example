package cmd

import (
	"encoding/csv"
	"fmt"
	"ldapctl/g"
	"ldapctl/models"
	"os"
	"strings"
	"time"
)

func PrintStart(action string) {
	fmt.Printf("LDAP %s Start \n", action)
	fmt.Println("==================================")
	fmt.Println()
}

func PrintEnd(action string, startTime time.Time) {
	endTime := time.Now()
	fmt.Println()
	fmt.Println("==================================")
	fmt.Printf("LDAP %s Finished, Time Usage %s \n", action, endTime.Sub(startTime))
}

func PrintSearchResult(result models.LADP_RESULT) {
	fmt.Println()
	fmt.Printf("DN: %s \n", result.DN)
	fmt.Println("Attributes:")

	longestKeyLenth := getLongestKeyLen(result.Attributes)
	for k, v := range result.Attributes {
		kLen := len(k)
		if kLen < longestKeyLenth {
			k = addSpace(k, (longestKeyLenth - kLen))
		}
		vStr := strings.Join(v, ";")
		fmt.Printf(" -- %s : %s \n", k, vStr)
	}
	fmt.Println()
}

func addSpace(s string, l int) string {
	for i := 0; i < l; i++ {
		s += " "
	}
	return s
}

func getLongestKeyLen(m map[string][]string) int {
	l := 0
	for k, _ := range m {
		if len(k) > l {
			l = len(k)
		}
	}
	return l
}

func inArray(str string, array []string) bool {
	for _, v := range array {
		if v == str {
			return true
		}
	}
	return false
}

func WriteUsersToCsv(users []models.LADP_RESULT, filename string) (err error) {
	csvFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvFile.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	writer := csv.NewWriter(csvFile)
	title := make([]string, len(g.Config().Ldap.Attributes))
	title[0] = "DN"
	for i, v := range g.Config().Ldap.Attributes {
		title[i+1] = v
	}
	err = writer.Write(title)
	if err != nil {
		return err
	}

	for _, user := range users {
		s := make([]string, len(g.Config().Ldap.Attributes)+1)
		s[0] = user.DN
		for key, value := range user.Attributes {
			valueStr := strings.Join(value, ";")
			for i, v := range title {
				if v == key {
					s[i] = valueStr
				}
			}
		}
		err = writer.Write(s)
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

func WriteFailsToCsv(msgs []models.FailedMessage, filename string) (err error) {
	csvFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvFile.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	writer := csv.NewWriter(csvFile)
	title := make([]string, 2)
	title[0] = "username"
	title[1] = "message"
	err = writer.Write(title)
	if err != nil {
		return err
	}

	for _, msg := range msgs {
		s := make([]string, 2)
		s[0] = msg.Username
		s[1] = msg.Message
		err = writer.Write(s)
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return
}
