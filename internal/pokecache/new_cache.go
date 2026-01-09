package pokecache

import "sync"

type Cache struct {
	entry map[string]cacheEntry

}