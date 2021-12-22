package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/saint-yellow/go-pl/toys/distributed-system/exception"
)

const ServerPort = ":3000"
const ServicesURL = "http://localhost" + ServerPort + "/services"

type registry struct {
	registrations []Registration
	mutex *sync.RWMutex
}

func (r *registry) add(reg Registration) error {
	r.mutex.Lock()
	r.registrations = append(r.registrations, reg)
	r.mutex.Unlock()
	err := r.sendRequiredServices(reg)
	r.notify(patch{
		Added: []patchEntry{
			{
				Name: reg.ServiceName,
				URL: reg.ServiceURL,
			},
		},
	})
	return err
}

func (r registry) sendRequiredServices(reg Registration) error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var p patch
	for _, serviceReg := range r.registrations {
		for _, reqService := range reg.RequiredServices {
			if serviceReg.ServiceName == reqService {
				p.Added = append(p.Added, patchEntry{
					Name: serviceReg.ServiceName,
					URL: serviceReg.ServiceURL,
				})
			}
		}
	}
	err := r.sendPatch(p, reg.ServiceUpdateURL)
	if err != nil {
		return err
	}
	return nil
}

func (r registry) notify(fullPatch patch) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, reg := range r.registrations {
		go func(reg Registration) {
			for _, reqService := range reg.RequiredServices {
				p := patch{
					Added: []patchEntry{},
					Removed: []patchEntry{},
				}
				sendUpdate := false
				for _, added := range fullPatch.Added {
					if added.Name == reqService {
						p.Added = append(p.Added, added)
						sendUpdate = true
					}
				}
				for _, removed := range fullPatch.Removed {
					if removed.Name == reqService {
						p.Removed = append(p.Removed, removed)
						sendUpdate =true
					}
				}
				if sendUpdate {
					err := r.sendPatch(p, reg.ServiceUpdateURL)
					if err != nil {
						log.Println(err)
						return
					}
				}
			}
		}(reg)
	}
}

func (r registry) sendPatch(p patch, url string) error {
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}

	_, err = http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	return nil
}

func (r *registry) remove(url string) error {
	for i := range reg.registrations {
		if reg.registrations[i].ServiceURL == url {
			r.notify(patch{
				Removed: []patchEntry{
					{
						Name: r.registrations[i].ServiceName,
						URL: r.registrations[i].ServiceURL,
					},
				},
			})
			r.mutex.Lock()
			reg.registrations = append(reg.registrations[:i], reg.registrations[:i+1]...)
			r.mutex.Unlock()
			return nil
		}
	}
	return fmt.Errorf("service at URL %s not found", url)
}

func (r *registry) heartbeat(frequency time.Duration) {
	for {
		var wg sync.WaitGroup
		for _, reg := range r.registrations {
			wg.Add(1)
			go func(reg Registration) {
				defer wg.Done()

				success := true
				for attempts := 0; attempts < 3; attempts++ {
					res, err := http.Get(reg.HeartbeatURL)
					if err != nil {
						log.Println(err)
					} else if res.StatusCode == http.StatusOK {
						log.Printf("heartbeat check passed for %v", reg.ServiceName)
						if !success {
							r.add(reg)
						}
						break;
					}

					log.Printf("heartbeat check failed for %v", reg.ServiceName)
					if success {
						success = false
						r.remove(reg.ServiceURL)
					}
					time.Sleep(1 * time.Second)
				}
			}(reg)

			wg.Wait()
			time.Sleep(frequency)
		}
	}
}

var reg = registry{
	registrations: make([]Registration, 0),
	mutex: new(sync.RWMutex),
}

type RegistryService struct {

}

func (s RegistryService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("request received")
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		var r Registration
		err := dec.Decode(&r)
		exception.Record(err, w, http.StatusBadRequest)
		log.Printf("adding service: %v with URL: %v \n", r.ServiceURL, r.ServiceName)
		err = reg.add(r)
		exception.Record(err, w, http.StatusBadRequest)
	case http.MethodDelete:
		payload, err := ioutil.ReadAll(r.Body)
		exception.Record(err, w, http.StatusInternalServerError)
		url := string(payload)
		log.Printf("removing service at URL %s", url)
		err = reg.remove(url)
		exception.Record(err, w, http.StatusInternalServerError)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

var once sync.Once

func SetupRegistryService() {
	once.Do(func() {
		go reg.heartbeat(3 * time.Second)
	})
}
