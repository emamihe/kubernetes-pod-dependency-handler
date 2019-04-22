package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (ep *Endpoint) Call() ([]byte, error) {
	if ep.SelfSigned {
		return invalidHttpsCall(ep)
	}else {
		return validHttpsCall(ep)
	}
}

func validHttpsCall(ep *Endpoint) ([]byte, error) {
	token, err := ioutil.ReadFile(ep.Token)
	if err != nil {
		return nil, err
	}
	tokenStr:=string(token)

	client := http.Client{}

	req, err := http.NewRequest("GET",ep.Url, nil)
	if err != nil {
		return nil, err
	}


	req.Header.Add("Authorization", "Bearer "+tokenStr)
	res, err := client.Do(req)
	defer req.Body.Close()

	if err != nil {
		return nil, err
	}

	buffer, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("https request resulted in %v code" , res.StatusCode))
	}

	return buffer, nil
}

func invalidHttpsCall(ep *Endpoint) ([]byte, error) {
	caCert, err := ioutil.ReadFile(ep.Cacert)
	if err != nil {
		return nil, err
	}

	token, err := ioutil.ReadFile((*ep).Token)
	if err != nil {
		return nil, err
	}
	tokenStr:=string(token)

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
			},
		},
	}

	req, err := http.NewRequest("GET", (*ep).Url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+tokenStr)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	buffer, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("https request resulted in %v code" , res.StatusCode))
	}

	return buffer, nil
}