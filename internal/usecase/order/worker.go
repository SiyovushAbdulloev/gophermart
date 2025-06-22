package order

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/SiyovushAbdulloev/gophermart/internal/client/accrual"
	"github.com/SiyovushAbdulloev/gophermart/internal/entity/order"
	"github.com/SiyovushAbdulloev/gophermart/internal/repository"
)

type Worker struct {
	accrualClient *accrual.Client
	repo          repository.OrderRepository
	balanceRepo   repository.BalanceRepository
}

func NewWorker(client *accrual.Client, orderRepo repository.OrderRepository, balanceRepo repository.BalanceRepository) *Worker {
	return &Worker{
		accrualClient: client,
		repo:          orderRepo,
		balanceRepo:   balanceRepo,
	}
}

func (w *Worker) Process(ctx context.Context, o *order.Order) {
	err := w.accrualClient.SendOrder(accrual.OrderRequest{
		Order: strconv.Itoa(o.ID),
		Goods: []accrual.Good{
			{
				Description: "default item", // можно расширить
				Price:       1000,           // условная сумма
			},
		},
	})
	if err != nil {
		log.Printf("❌ failed to send order to accrual: %v", err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(5 * time.Second)

			res, err := w.accrualClient.GetOrder(strconv.Itoa(o.ID))
			if err != nil {
				log.Printf("❌ error polling accrual system: %v", err)
				continue
			}

			if res == nil {
				continue
			}

			log.Printf("ℹ️ order %s status: %s", res.Order, res.Status)
			log.Println("ORDER:", res)

			switch res.Status {
			case "INVALID":
				_ = w.repo.UpdateStatus(o.ID, "INVALID", float64(0))
				return
			case "PROCESSED":
				_ = w.repo.UpdateStatus(o.ID, "PROCESSED", res.Accrual)
				if res.Accrual > 0 {
					err = w.balanceRepo.AddPoints(o.UserID, res.Accrual)
					if err != nil {
						fmt.Println("Error of PROCESSING:", err)
					}
				}
				return
			case "PROCESSING":
				// ждём
			case "REGISTERED":
				// ждём
			}
		}
	}
}
