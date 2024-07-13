package client

import (
	"encoding/json"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
)

type Config struct {
	BaseUrl string            `json:"baseurl"`
	UInfo   map[string]string `json:"user,omitempty"`
}

type User struct {
	C         *resty.Client
	Conf      *Config
	ConfPath  string
	ErrLogger *log.Logger
}

func Init(confPath string) (*User, error) {
	u := &User{
		C:         resty.New(),
		Conf:      &Config{},
		ConfPath:  confPath,
		ErrLogger: log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	err := u.readConfigFile()
	if err != nil {
		u.ErrLogger.Fatal(err)
	}

	u.C.SetBaseURL(u.Conf.BaseUrl)
	return u, nil
}

func (u *User) readConfigFile() error {
	d, err := os.ReadFile(u.ConfPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(d, &u.Conf)
	if err != nil {
		return err
	}

	return nil
}
