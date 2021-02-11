package IOrderBiz

import (
	"context"
	md "github.com/goclub/project-seckilling/internal/memory_data"
	"github.com/goclub/project-seckilling/internal/persistence_data"
)

type Interface interface {
	ConsumerPlaceOrder(ctx context.Context, data ConsumerPlaceOrder) (ticketID md.IDTicket, reject error)
	ConsumerQueryTicket(ctx context.Context, data ConsumerQueryTicket) (reply ConsumerQueryTicketReply, reject error)
	ConsumerOrderList(ctx context.Context, data ConsumerOrderList)(orderList ConsumerOrderListReply, reject error)
	JobCancelUnpaidOrder(ctx context.Context) (reject error)
}

type ConsumerPlaceOrder struct {
	ConsumerID pd.IDConsumer
	GoodsID pd.IDGoods
}
type ConsumerQueryTicket struct {
	TicketID md.IDTicket
}
type ConsumerQueryTicketReply struct {
	Status QueryTicketReplyStatus
	Message string
}

type QueryTicketReplyStatus string
func (QueryTicketReplyStatus) Enum() (e struct {
	Processing string
	Success string
	Fail string
}) {
	e.Processing = "processing"
	e.Success = "success"
	e.Fail = "fail"
	return
}

type ConsumerOrderList struct {
	Page uint
	PerPage uint
}
type ConsumerOrderListReply struct {
	Items []OrderListReplyItem
	Total uint
}
type OrderListReplyItem struct {
	OrderID pd.IDOrder
	ConsumerID pd.IDConsumer
	GoodsID pd.IDGoods
	OrderStatus OrderListReplyItemStatus
}
type OrderListReplyItemStatus string

func (OrderListReplyItemStatus) Enum() (e struct {
	NotPaid string
	Paid string
	Canceled string
}) {
	e.NotPaid = "notPaid"
	e.Paid = "paid"
	e.Canceled = "canceled"
	return
}