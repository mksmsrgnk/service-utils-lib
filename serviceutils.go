package serviceutils

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

func statusCodeParser(URL string, statusCode int) error {
	if statusCode > 199 && statusCode < 300 {
		return nil
	}
	if statusCode > 499 && statusCode < 600 {
		return fmt.Errorf("service  is NOT OK, test URL: %s, status code %d", URL, statusCode)
	}
	return nil
}

func CheckURL(URL string) error {
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Get(URL)
	if err != nil {
		return err
	}
	return statusCodeParser(URL, resp.StatusCode)
}

func RestartService(serviceName string) error {
	stopCmd := exec.Command("sudo", "systemctl", "stop", serviceName)
	startCmd := exec.Command("sudo", "systemctl", "start", serviceName)

	if err := stopCmd.Run(); err == nil {
		time.Sleep(5 * time.Second)
		if err := startCmd.Run(); err == nil {
			return nil
		} else {
			return fmt.Errorf("can not start %s, error: %v", serviceName, err)
		}
	} else {
		return fmt.Errorf("can not stop %s, error: %v", serviceName, err)
	}
}
