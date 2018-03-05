package cache

import (
	"errors"
	"time"

	"github.com/konojunya/HEW2018/model"
	gocache "github.com/patrickmn/go-cache"
)

type productCache struct {
	cache *gocache.Cache
}

// Product productsのcache
var Product = newProductCache()

func newProductCache() productCache {
	c := productCache{
		cache: gocache.New(5*time.Minute, 10*time.Minute),
	}
	c.load()

	return c
}

func (p *productCache) Reload() {
	p.load()
}

func (p *productCache) load() {
	products := make([]model.Product, 0)
	err := db.Find(&products).Error
	if err != nil {
		panic(err)
	}

	for i, product := range products {
		product.SetVote()
		products[i] = product
	}

	p.cache.Set("products", products, gocache.NoExpiration)
}

func (p *productCache) GetAll() ([]model.Product, error) {
	products, found := p.cache.Get("products")
	if found {
		return products.([]model.Product), nil
	}
	return nil, errors.New("products not loaded")
}
