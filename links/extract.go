package links

import (
	"fmt"
	"html"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

var hrefPattern = regexp.MustCompile(`(?is)<a\b[^>]*?\bhref\s*=\s*(?:"([^"]*)"|'([^']*)'|([^\s"'<>` + "`" + `]+))`)

// Extract fetches a page and returns all anchor href values as absolute URLs
// when they can be resolved against the source page.
func Extract(rawURL string) ([]string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch %s: %s", rawURL, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	base, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	matches := hrefPattern.FindAllStringSubmatch(string(body), -1)
	seen := make(map[string]struct{}, len(matches))
	links := make([]string, 0, len(matches))

	for _, match := range matches {
		href := firstNonEmpty(match[1], match[2], match[3])
		if href == "" {
			continue
		}

		href = html.UnescapeString(href)
		parsed, err := url.Parse(href)
		if err != nil {
			continue
		}

		resolved := base.ResolveReference(parsed).String()
		if _, ok := seen[resolved]; ok {
			continue
		}
		seen[resolved] = struct{}{}
		links = append(links, resolved)
	}

	return links, nil
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}
