package config

import "os"

type Config struct {
	EncryptionKey               string
	EncryptionSubstitution      string
	EncryptionSubstitutionStart string
	LastFmApiKey                string
	LastFmSharedSecret          string
}

func CreateConfig() *Config {
	return &Config{
		// 32 bytes
		EncryptionKey:               os.Getenv("ENCRYPTION_KEY"),
		EncryptionSubstitution:      os.Getenv("ENCRYPTION_SUBSTITUTION"),
		EncryptionSubstitutionStart: os.Getenv("ENCRYPTION_SUBSTITUTION_START"),
		LastFmApiKey:                os.Getenv("LAST_FM_API_KEY"),
		LastFmSharedSecret:          os.Getenv("LAST_FM_SHARED_SECRET"),
	}
}
