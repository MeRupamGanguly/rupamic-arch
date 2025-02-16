package usecase

type service struct{}

func NewOrderService() *service {
	return &service{}
}

func (svc *service) AddOrder() {
	// step1: Sent the Processed order to Exchange.
	// step2: DB
	// step3: MQ publish
	// step4: Wait for Subscriber to Send Notification to user.
	// Modify DB
}

func (svc *service) GetOrder() {

}
func (svc *service) ListOrder() {

}
