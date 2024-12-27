package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/nlm/lucca-api/api"
	"github.com/nlm/lucca-api/internal/chrome-cookie-decrypt/cookies"
	"github.com/nlm/lucca-api/internal/chrome-cookie-decrypt/database"
	"github.com/nlm/lucca-api/internal/chrome-cookie-decrypt/keychain"
)

func init() {
	RegisterCommand("import-auth-cookie", ImportAuthCookie)
}

const (
	cookiesPathSuffix = "Library/Application Support/Google/Chrome/Default/Cookies"
)

func ImportAuthCookie(ctx context.Context, client *api.Client, args []string) error {

	// read config
	var config Config
	md, err := toml.DecodeFile(*flagConfigFile, &config)
	if len(md.Undecoded()) > 0 {
		return fmt.Errorf("extra config keys: %v", md.Undecoded())
	}
	if err != nil {
		log.Fatal(err)
	}

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
	if err := query.Select(&chromeCookies, config.Host, "authToken"); err != nil {
		return fmt.Errorf("failed to query cookies sqlite DB: %w", err)
	}

	updated := false
	for _, c := range chromeCookies {
		if err := c.Decrypt(encryptionKey); err != nil {
			return fmt.Errorf("failed to decrypt cookie: %w (name=%s, host_key=%s)", err, c.Name, c.HostKey)
		}
		updated = true
		config.AuthCookie = c.Value
	}

	if updated {
		b := bytes.NewBuffer(nil)
		enc := toml.NewEncoder(b)
		err = enc.Encode(config)
		if err != nil {
			return fmt.Errorf("failed encoding config: %w", err)
		}

		f, err := os.Create(*flagConfigFile)
		if err != nil {
			return fmt.Errorf("failed opening config file for writing: %w", err)
		}
		defer f.Close()
		_, err = f.Write(b.Bytes())
		if err != nil {
			return fmt.Errorf("failed writing config to file: %w", err)
		}

		fmt.Println("config file updated")
	} else {
		return fmt.Errorf("lucca authToken cookie not found")
	}

	return nil
}
