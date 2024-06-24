package uranai

import "context"

type Batch struct {
	t *FortuneTeller
	p *Publisher
}

func NewBatch(t *FortuneTeller, p *Publisher) *Batch {
	return &Batch{t: t, p: p}
}

func (b *Batch) Run(ctx context.Context) error {
	resultSet, err := b.t.Listen(ctx)
	if err != nil {
		return err
	}
	err = b.p.Publish(ctx, resultSet)
	if err != nil {
		return err
	}
	return nil
}
