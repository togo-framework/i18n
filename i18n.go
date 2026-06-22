// Package i18n is togo's default JSON-file translator provider. Blank-import (or
// `togo install togo-framework/i18n`) to register it with the kernel.
package i18n

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("i18n", togo.PriorityService, func(k *togo.Kernel) error {
		k.I18n = Load(k.Config.LocaleDir, k.Config.Locale)
		return nil
	})
}

// Bundle holds loaded locales and satisfies togo/i18n.Translator.
type Bundle struct {
	mu       sync.RWMutex
	locales  map[string]map[string]string
	fallback string
}

// Load reads every <locale>.json in dir (missing dir is fine).
func Load(dir, fallback string) togo.Translator {
	b := &Bundle{locales: map[string]map[string]string{}, fallback: fallback}
	matches, _ := filepath.Glob(filepath.Join(dir, "*.json"))
	for _, f := range matches {
		data, err := os.ReadFile(f)
		if err != nil {
			continue
		}
		var m map[string]string
		if json.Unmarshal(data, &m) != nil {
			continue
		}
		b.locales[strings.TrimSuffix(filepath.Base(f), ".json")] = m
	}
	return b
}

// T translates key for locale, falling back to the fallback locale then the key.
func (b *Bundle) T(locale, key string) string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if m, ok := b.locales[locale]; ok {
		if v, ok := m[key]; ok {
			return v
		}
	}
	if m, ok := b.locales[b.fallback]; ok {
		if v, ok := m[key]; ok {
			return v
		}
	}
	return key
}
