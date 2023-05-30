package wget

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"strings"
)

//Site represent data needed to build sitemap
type Site struct {
	rootLink     string
	visitedLinks map[string]struct{}
	directory    string
}

//NewSite creates instance if sitemap
func NewSite(rootLink, directory string) *Site {
	v := map[string]struct{}{
		rootLink: {},
	}

	return &Site{
		rootLink:     rootLink,
		visitedLinks: v,
		directory:    directory,
	}
}

// DownloadSite recursively visit links in queue and download them;
// If depth is greater than zero is number of recursive calls
func (s *Site) DownloadSite(queue []string, depth int) error {
	if depth == 0 {
		return nil
	}

	discoveredLinks := make([]string, 0)

	for _, value := range queue {
		resp, err := http.Get(value)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		mediaType, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			fmt.Printf("cant parse link type [%s]\n", err.Error())
		}

		ext, err := mime.ExtensionsByType(mediaType)
		if err != nil || len(ext) == 0 {
			ext = append(ext, "")
			fmt.Printf("cant parse link type [%s]\n", err.Error())
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		r := bytes.NewReader(body)
		fileName := path.Join(s.directory, path.Base(resp.Request.URL.Path)+ext[0])
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		size, err := io.Copy(file, r)
		if err != nil {
			return err
		}

		fmt.Printf("Download a file [%s] with size [%s] bytes\n", fileName, size)

		if _, err := r.Seek(0, 0); err != nil {
			return err
		}

		links, err := s.parseLinks(r)
		if err != nil {
			return err
		}
		discoveredLinks = append(discoveredLinks, links...)
	}
	if len(discoveredLinks) > 0 {
		depth--
		return s.DownloadSite(discoveredLinks, depth)
	}

	return nil
}

// parseLinks read html data from io.Reader and create array of Links
func (s *Site) parseLinks(r io.Reader) ([]string, error) {
	res, err := ParseHTML(r)
	if err != nil {
		return nil, err
	}

	links := make([]string, 0)
	for _, value := range res {
		href := value.href.Host + value.href.Path
		visited := true

		switch {
		case strings.HasPrefix(href, s.rootLink):
			visited = s.Visited(href)
		case strings.HasPrefix(href, "/"):
			href = s.rootLink + href
			visited = s.Visited(href)
		}

		if !visited {
			links = append(links, href)
		}
	}

	return links, nil
}

// Visited check is url visited
func (s *Site) Visited(href string) (visited bool) {
	if _, visited := s.visitedLinks[href]; !visited {
		s.visitedLinks[href] = struct{}{}
	}
	return visited
}