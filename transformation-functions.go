package main

import (
	"crypto/sha1"
	"encoding/hex"
	"time"

	"github.com/xwb1989/sqlparser"
	"syreclabs.com/go/faker"
)

func generateUsername(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Internet().UserName()))
}

func generatePassword(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	// TODO encrypt this value
	return sqlparser.NewStrVal([]byte(faker.Internet().Password(8, 14)))
}

func generateEmail(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Internet().SafeEmail()))
}

func generateURL(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Internet().Url()))
}

func generateName(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Name().Name()))
}

func generateFirstName(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Name().FirstName()))
}

func generateLastName(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Name().LastName()))
}

func generateParagraph(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Lorem().Sentence(3)))
}

func generateIPv4(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Internet().IpV4Address()))
}

func generatePhone(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte("06" + faker.PhoneNumber().SubscriberNumber(8)))
}

func generateCompany(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Company().Name()))
}

func generateCity(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Address().City()))
}

func generateAddress(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Address().StreetAddress()))
}

func generatePostcode(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Address().Postcode()))
}

func generateJob(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Name().Title()))
}

func generateSHA1(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	// thanks to https://stackoverflow.com/questions/10701874/generating-the-sha-hash-of-a-string-using-golang
	now := time.Now()
	str := now.String()
	hasher := sha1.New()
	hasher.Write([]byte(str))
	sha1_hash := hex.EncodeToString(hasher.Sum(nil))
	return sqlparser.NewStrVal([]byte(sha1_hash))
}

func generatePrice(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Number().Decimal(6, 2)))
}
