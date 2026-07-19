package processor

import (
	"log"
	"sync"

	"github.com/google/uuid"
)

type EventService interface {
	Process(id uuid.UUID) error
}

type EventProcessor struct {
	service EventService
	jobs    chan uuid.UUID
	wg      sync.WaitGroup
}

func NewEventProcessor(
	service EventService,
	queueSize int,
) *EventProcessor {
	return &EventProcessor{
		service: service,
		jobs:    make(chan uuid.UUID, queueSize),
	}
}

func (p *EventProcessor) Enqueue(id uuid.UUID) {
	p.jobs <- id
}

func (p *EventProcessor) Start(workerCount int) {
	for workerID := 1; workerID <= workerCount; workerID++ {
		p.wg.Add(1)

		go p.worker(workerID)
	}
}

func (p *EventProcessor) worker(workerID int) {
	defer p.wg.Done()

	for eventID := range p.jobs {
		log.Printf("worker %d processing event %s", workerID, eventID)

		if err := p.service.Process(eventID); err != nil {
			log.Printf(
				"worker %d failed to process event %s: %v",
				workerID,
				eventID,
				err,
			)
		}
	}
}

func (p *EventProcessor) Stop() {
	close(p.jobs)
	p.wg.Wait()
}
