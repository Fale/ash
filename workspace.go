package ash

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type workspaceConfig struct {
	Repositories []struct {
		URL      string
		Location string
	}
}

type Workspace struct {
	location     string
	repositories []Repository
}

func NewWorkspace(location string) (*Workspace, error) {
	w := &Workspace{
		location: location,
	}
	if _, err := os.Stat(location); !os.IsNotExist(err) {
		if err := w.loadConfig(); err != nil {
			return nil, err
		}
		return w, nil
	}
	if err := os.MkdirAll(location, os.ModePerm); err != nil {
		return nil, err
	}
	if err := w.saveConfig(); err != nil {
		return nil, err
	}
	return w, nil
}

func (w *Workspace) loadConfig() error {
	configFile := fmt.Sprintf("%v/.workspace.yaml", w.location)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return fmt.Errorf("config not found: %v\n", err)
	}
	buf, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	wcs := &workspaceConfig{}
	err = yaml.Unmarshal(buf, wcs)
	if err != nil {
		return fmt.Errorf("in file %q: %v", configFile, err)
	}
	for _, r := range wcs.Repositories {
		u, err := url.Parse(r.URL)
		if err != nil {
			return err
		}
		nr, err := NewRepository(*u, r.Location)
		if err != nil {
			return err
		}
		w.repositories = append(w.repositories, *nr)
	}
	return nil
}

func (w *Workspace) saveConfig() error {
	var wcs workspaceConfig
	for _, r := range w.repositories {
		nr := struct {
			URL      string
			Location string
		}{
			URL:      r.URL.String(),
			Location: r.Location,
		}
		wcs.Repositories = append(wcs.Repositories, nr)
	}
	configs, err := yaml.Marshal(wcs)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("%v/.workspace.yaml", w.location), configs, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (w *Workspace) AddRepositories(us []url.URL) error {
	for _, u := range us {
		for _, repo := range w.repositories {
			if repo.URL.String() == u.String() {
				return fmt.Errorf("the repository %v is already present", u)
			}
		}
		location := path.Join(w.location, u.Host, u.Path)
		if err := os.MkdirAll(path.Dir(location), os.ModePerm); err != nil {
			return err
		}
		nr, err := NewRepository(u, location)
		if err != nil {
			return err
		}
		w.repositories = append(w.repositories, *nr)
	}
	return w.saveConfig()
}

func (w *Workspace) ListRepositories() []Repository {
	return w.repositories
}

func (w *Workspace) DeleteRepository(url string) error { return nil }
