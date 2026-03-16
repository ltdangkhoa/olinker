package core

// VendorDriver is the interface the EncodeService expects, preventing circular deps with vendors pkg.
type VendorDriver interface {
	EncodeCard(req EncodeRequest) (EncodeResult, error)
	CancelCard(cardID string) error
	ExtendCard(req ExtendRequest) error
	ReadCard() (CardInfo, error)
}

// EncodeService coordinates vendor actions through the single worker JobQueue
type EncodeService struct {
	queue  *JobQueue
	vendor VendorDriver
}

func NewEncodeService(queue *JobQueue, vendor VendorDriver) *EncodeService {
	return &EncodeService{
		queue:  queue,
		vendor: vendor,
	}
}

func (s *EncodeService) EncodeCard(req EncodeRequest) (EncodeResult, error) {
	res, err := s.queue.Enqueue(func() (interface{}, error) {
		return s.vendor.EncodeCard(req)
	})
	if err != nil {
		return EncodeResult{}, err
	}
	return res.(EncodeResult), nil
}

func (s *EncodeService) CancelCard(req CancelRequest) error {
	_, err := s.queue.Enqueue(func() (interface{}, error) {
		return nil, s.vendor.CancelCard(req.CardID)
	})
	return err
}

func (s *EncodeService) ExtendCard(req ExtendRequest) error {
	_, err := s.queue.Enqueue(func() (interface{}, error) {
		return nil, s.vendor.ExtendCard(req)
	})
	return err
}

func (s *EncodeService) ReadCard() (CardInfo, error) {
	res, err := s.queue.Enqueue(func() (interface{}, error) {
		return s.vendor.ReadCard()
	})
	if err != nil {
		return CardInfo{}, err
	}
	return res.(CardInfo), nil
}
