package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/nlm/lucca-api/cmd/lucca/config"
	"github.com/nlm/lucca-api/internal/chrome-cookie-decrypt/cookies"
	"github.com/nlm/lucca-api/internal/chrome-cookie-decrypt/database"
	"github.com/nlm/lucca-api/internal/chrome-cookie-decrypt/keychain"
)

const (
	cookiesPathSuffix = "Library/Application Support/Google/Chrome/Default/Cookies"
)

func ImportAuthCookie(conf *config.Config, write bool) error {
	// open cookie db
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}
	cookiesPath := filepath.Join(home, cookiesPathSuffix)

	db, err := database.InitDB(cookiesPath)
	if err != nil {
		return fmt.Errorf("failed to open cookies sqlite DB: %w", err)
	}
	defer db.Close()

	// decrypot
	encryptionKey, err := keychain.GetEncryptionKey()
	if err != nil {
		return fmt.Errorf("failed to get encryption key: %w", err)
	}

	chromeCookies := []*cookies.ChromeCookie{}
	query, err := db.Preparex("SELECT * FROM cookies WHERE host_key = $1 AND name = $2")
	if err != nil {
		return fmt.Errorf("failed to prepare select query: %w", err)
	}
	if err := query.Select(&chromeCookies, conf.Host, "authToken"); err != nil {
		return fmt.Errorf("failed to query cookies sqlite DB: %w", err)
	}

	updated := false
	for _, c := range chromeCookies {
		if err := c.Decrypt(encryptionKey); err != nil {
			return fmt.Errorf("failed to decrypt cookie: %w (name=%s, host_key=%s)", err, c.Name, c.HostKey)
		}
		updated = true
		conf.AuthCookie = c.Value
	}

	if write {
		if updated {
			b := bytes.NewBuffer(nil)
			enc := toml.NewEncoder(b)
			err = enc.Encode(conf)
			if err != nil {
				return fmt.Errorf("failed encoding config: %w", err)
			}

			err = conf.Save()
			if err != nil {
				return fmt.Errorf("failed writing config to file: %w", err)
			}

			fmt.Println("config file updated")
		} else {
			return fmt.Errorf("lucca authToken cookie not found")
		}
	}

	return nil
}
