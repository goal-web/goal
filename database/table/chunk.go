package table

import "github.com/goal-web/contracts"

func (this *table) Chunk(size int, handler func(collection contracts.Collection, page int) error) (err error) {
	page := 1
	for err == nil {
		newCollection := this.SimplePaginate(int64(size), int64(page))
		err = handler(newCollection, page)
		page++
		if newCollection.Len() < size {
			break
		}
	}
	return
}

func (this *table) ChunkById(size int, handler func(collection contracts.Collection, page int) error) error {
	return this.OrderBy("id").Chunk(size, handler)
}
