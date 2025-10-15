package token

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/optimuscrime/lastfm-on-this-day/pgk/config"
)

var (
	ErrInvalidConfigurationProvided  = errors.New("invalid configuration provided")
	ErrMissingEncryptionSubstitution = errors.New("missing encryption substitution")
	ErrInvalidToken                  = errors.New("invalid token")
)

type Service struct {
	config *config.Config
}

func New(c *config.Config) *Service {
	return &Service{
		config: c,
	}
}

func (s *Service) EncryptToken(accessToken string) (string, error) {
	fullToken, err := s.constructFullToken(accessToken)
	block, err := aes.NewCipher([]byte(s.config.EncryptionKey))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nonce, nonce, []byte(fullToken), nil)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (s *Service) ValidateToken(fullEncryptedToken string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(fullEncryptedToken)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(s.config.EncryptionKey))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	plainTextString := string(plainText)
	token, err := s.deconstructFullToken(plainTextString)

	return token, nil
}

func (s *Service) deconstructFullToken(token string) (string, error) {
	if len(token) != 128 {
		return "", ErrInvalidToken
	}

	prefix := token[:32]
	err := s.validatePrefix(prefix)
	if err != nil {
		return "", err
	}

	return token[32:64], nil
}

func (s *Service) validatePrefix(prefix string) error {
	substitutionString := s.config.EncryptionSubstitution
	if len(substitutionString) == 0 {
		return ErrInvalidConfigurationProvided
	}

	substitutionStartIndex, err := strconv.Atoi(s.config.EncryptionSubstitutionStart)
	if err != nil {
		return err
	}

	maxIndex := substitutionStartIndex + len(s.config.EncryptionSubstitution)
	if maxIndex > len(prefix) {
		return ErrMissingEncryptionSubstitution
	}

	substitutionCheck := prefix[substitutionStartIndex : substitutionStartIndex+len(s.config.EncryptionSubstitution)]

	if substitutionCheck != substitutionString {
		return ErrInvalidToken
	}

	return nil
}

func (s *Service) constructFullToken(token string) (string, error) {
	prefix, err := generateSecureString(32)
	if err != nil {
		return "", err
	}
	suffix1, err := generateSecureString(32)
	if err != nil {
		return "", err
	}
	suffix2, err := generateSecureString(32)
	if err != nil {
		return "", err
	}

	// In the prefix, replace parts of the string with the substitution, to ensure a valid token
	substitutionPrefix, err := s.replacePrefix(prefix)
	if err != nil {
		return "", err
	}

	return substitutionPrefix + token + suffix1 + suffix2, nil

}

func (s *Service) replacePrefix(prefix string) (string, error) {
	substitutionString := s.config.EncryptionSubstitution
	if len(substitutionString) == 0 {
		return "", ErrInvalidConfigurationProvided
	}

	substitutionStartIndex, err := strconv.Atoi(s.config.EncryptionSubstitutionStart)
	if err != nil {
		return "", err
	}

	maxIndex := substitutionStartIndex + len(s.config.EncryptionSubstitution)
	if maxIndex > len(prefix) {
		return "", ErrMissingEncryptionSubstitution
	}

	constructedPrefix := prefix[:substitutionStartIndex] + substitutionString + prefix[substitutionStartIndex+len(substitutionString):]
	return constructedPrefix, nil
}

func generateSecureString(length int) (string, error) {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
