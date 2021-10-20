package common

type Pagination struct {
	Page int
	Size int
}

func (p *Pagination) GetOffset() int {
	if p.Page > 0 && p.Size > 0 {
		return (p.Page - 1) * p.Size
	}
	return 0
}

func (p *Pagination) GetLimit() int {
	if p.Size > 0 {
		return p.Size
	} else {
		return 12
	}
}
