package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type IPFSService struct {
	IPFSGateway string
}

func NewIPFSService(gateway string) *IPFSService {
	return &IPFSService{IPFSGateway: gateway}
}

func (s *IPFSService) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	buffer := bytes.NewBuffer(nil)
	_, err := buffer.ReadFrom(file)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", s.IPFSGateway+"/api/v0/add", buffer)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "multipart/form-data")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to upload file to IPFS")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var ipfsResp map[string]interface{}
	if err := json.Unmarshal(body, &ipfsResp); err != nil {
		return "", err
	}

	cid, ok := ipfsResp["Hash"].(string)
	if !ok {
		return "", errors.New("invalid response from IPFS")
	}

	return cid, nil
}
