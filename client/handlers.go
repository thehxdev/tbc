package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func (u *User) UseraddHandler(password string) error {
	if password == "" {
		u.ErrLogger.Fatal("password could not be empty")
	}

	c := u.C

	resp, err := c.R().
		SetBody(map[string]string{
			"password": password,
		}).
		Post("/useradd")

	if err != nil {
		return err
	}

	err = json.Unmarshal(resp.Body(), &u.Conf.UInfo)
	if err != nil {
		return err
	}

	jdata, err := json.MarshalIndent(u.Conf, "", "    ")
	if err != nil {
		return err
	}

	fmt.Println(string(jdata))
	err = os.WriteFile(u.ConfPath, jdata, 644)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) UserdelHandler() error {
	c := u.C
	uinfo := u.Conf.UInfo

	resp, err := c.R().
		SetHeader("Authorization", uinfo["authKey"]).
		Delete("/userdel")

	if err != nil {
		return err
	}

	if stat := resp.StatusCode(); stat != 200 {
		return fmt.Errorf(string(resp.Body()))
	}

	return nil
}

func (u *User) LsHandler() ([]byte, error) {
	c := u.C
	uinfo := u.Conf.UInfo

	resp, err := c.R().
		SetHeader("Authorization", uinfo["authKey"]).
		Get("/ls")

	if err != nil {
		return nil, err
	}

	if stat := resp.StatusCode(); stat != 200 {
		return nil, errors.New(string(resp.Body()))
	}

	return resp.Body(), nil
}

func (u *User) TeeHandler(name, path string) ([]byte, error) {
	if name == "" {
		return nil, errors.New("name could not be empty")
	}

	c := u.C
	uinfo := u.Conf.UInfo

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	resp, err := c.R().
		SetQueryParam("name", name).
		SetHeader("Authorization", uinfo["authKey"]).
		SetBody(data).
		Post("/tee")

	if err != nil {
		return nil, err
	}

	if stat := resp.StatusCode(); stat != 200 {
		return nil, errors.New(string(resp.Body()))
	}

	return resp.Body(), nil
}

func (u *User) RmHandler(txtid string) error {
	if txtid == "" {
		return errors.New("txt id could not be empty")
	}

	c := u.C
	uinfo := u.Conf.UInfo

	resp, err := c.R().
		SetQueryParam("txtid", txtid).
		SetHeader("Authorization", uinfo["authKey"]).
		Delete("/rm")

	if err != nil {
		return err
	}

	if stat := resp.StatusCode(); stat != 200 {
		return errors.New(string(resp.Body()))
	}

	return nil
}

func (u *User) MvHandler(txtid string) ([]byte, error) {
	if txtid == "" {
		return nil, errors.New("txt id could not be empty")
	}

	c := u.C
	uinfo := u.Conf.UInfo

	resp, err := c.R().
		SetQueryParam("txtid", txtid).
		SetHeader("Authorization", uinfo["authKey"]).
		Put("/mv")

	if err != nil {
		return nil, err
	}

	if stat := resp.StatusCode(); stat != 200 {
		return nil, errors.New(string(resp.Body()))
	}

	return resp.Body(), nil
}

func (u *User) ChtxtHandler(txtid, path string) error {
	if txtid == "" {
		return errors.New("txt id could not be empty")
	}

	c := u.C
	uinfo := u.Conf.UInfo

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	resp, err := c.R().
		SetQueryParam("txtid", txtid).
		SetHeader("Authorization", uinfo["authKey"]).
		SetBody(data).
		Put("/tee")

	if err != nil {
		return err
	}

	if stat := resp.StatusCode(); stat != 200 {
		return errors.New(string(resp.Body()))
	}

	return nil
}

func (u *User) RenameHandler(name, txtid string) error {
	if txtid == "" || name == "" {
		return errors.New("txt id or name could not be empty")
	}

	c := u.C
	uinfo := u.Conf.UInfo

	resp, err := c.R().
		SetQueryParam("txtid", txtid).
		SetBody(map[string]string{"name": name}).
		SetHeader("Authorization", uinfo["authKey"]).
		Put("/rename")

	if err != nil {
		return err
	}

	if stat := resp.StatusCode(); stat != 200 {
		return errors.New(string(resp.Body()))
	}

	return nil
}
