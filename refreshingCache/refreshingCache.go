package util

import (
	"context"
	"time"

	"k8s.io/klog/v2"
)

// update example interface
type client interface {
	Get(context.Context) ([]string, error)
}

type RefreshingCache struct {
	output             chan map[string]string
	refreshedTimestamp time.Time
}

func (c *RefreshingCache) Get() map[string]string {
	return <-c.output
}

func (c *RefreshingCache) RefreshedTimestamp() time.Time {
	return c.refreshedTimestamp
}

func NewRefreshingCache(ctx context.Context, client client) (*RefreshingCache, error) {
	currentIdentities, err := client.Get(ctx)
	if err != nil {
		return nil, err
	}
	current := identitiesToMap(currentIdentities)
	output := make(chan map[string]string)
	updateChan := make(chan map[string]string)
	c := &RefreshingCache{
		output:             output,
		refreshedTimestamp: time.Now(),
	}
	// go routine to return the current cache
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case current = <-updateChan:
				// update current
				klog.Infof("successfully updated cache")
				c.refreshedTimestamp = time.Now()
			case output <- current:
				// nop, we just want to send the current data any time someone
				// wants it
			}
		}
	}()

	// go routine to update the cache
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Duration(updateInterval) * time.Second):
				nextIdentities, err := client.Get(context.Background())
				if err != nil {
					klog.Errorf("failed to call client.Get: %s", err)
					continue
				}
				// create map
				next := transform(nextIdentities)
				// update cache
				updateChan <- next
			}
		}
	}()

	return c, nil
}

func transform(arr []string) map[string]string {
	m := make(map[string]string, len(arr))
	for _, a := range arr {
		m[a] = a
	}
	return m
}