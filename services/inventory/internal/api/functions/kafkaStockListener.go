package functions

import (
	"github.com/IBM/sarama"
	"github.com/akulsharma1/distributed-analytics-platform/services/inventory/internal/api/models"
	"github.com/akulsharma1/distributed-analytics-platform/services/inventory/internal/db"
	"github.com/akulsharma1/distributed-analytics-platform/services/inventory/internal/kafka"
	"gorm.io/gorm"
)

func ChangeStock(message kafka.OrderCreationMessage) {

	for _, orderItem := range message.Order.OrderItems {
		err := db.DATABASE.Model(&models.Variant{}).Where("id = ?", orderItem.VariantID).Update("quantity", gorm.Expr("quantity - ?", orderItem.Quantity)).Error

		if err != nil {
			_, _, err = kafka.Producer.SendMessage(&sarama.ProducerMessage{
				Topic: kafka.GET_ORDER_STATUS,
				Key: sarama.StringEncoder(message.MessageID),
			})
		}
	}
}