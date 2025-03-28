package certificateparser

import (
	"crypto/x509"
	"fmt"
)

// Parse extracts the certificate's serial number from ASN.1 DER data.
// Note: this is an example function to demonstrate unit test execution.
func Parse(der []byte) (string, error) {
	certificate, err := x509.ParseCertificate(der)
	if err != nil {
		return "", fmt.Errorf("failed to parse certificate: %w", err)
	}

	return certificate.Subject.SerialNumber, nil
}
