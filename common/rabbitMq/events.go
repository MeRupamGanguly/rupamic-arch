package rabbitmq

type Event string

const (
	AddOrder     Event = "Add Order"
	EditOrder    Event = "Edit Order"
	LimitOrder   Event = "Limit Order"
	CancelOrder  Event = "Cancel Order"
	CloseOrder   Event = "Close Order"
	TriggerOrder Event = "Trigger Order"
)
