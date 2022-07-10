package validator

import (
	"fmt"
	"regexp"
	"strings"
)

func ValidateAddress(crypto string, address string) (bool, error) {
	c := strings.ToLower(crypto)

	regex, err := getCryptoRegex(c)
	if err != nil {
		return false, fmt.Errorf("error to get regex %v", err)
	}

	re := regexp.MustCompile(regex)

	return re.MatchString(address), nil
}

func getCryptoRegex(crypto string) (string, error) {
	var cryptoRegexMap = map[string]string{
		"btc":   "^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$",
		"btg":   "^([GA])[a-zA-HJ-NP-Z0-9]{24,34}$",
		"dash":  "^([X7])[a-zA-Z0-9]{33}$",
		"dgb":   "^(D)[a-zA-Z0-9]{24,33}$",
		"eth":   "^(0x)[a-zA-Z0-9]{40}$",
		"smart": "^(S)[a-zA-Z0-9]{33}$",
		"xrp":   "^(r)[a-zA-Z0-9]{33}$",
		"zcr":   "^(Z)[a-zA-Z0-9]{33}$",
		"zec":   "^(t)[a-zA-Z0-9]{34}$",
	}

	regex, ok := cryptoRegexMap[crypto]
	if !ok {
		return "", fmt.Errorf("%s is not valid", crypto)
	}

	return regex, nil
}
