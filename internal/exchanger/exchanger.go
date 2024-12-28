package exchanger

import (
	"context"
	"errors"
	"log"

	pb "github.com/MaximKlimenko/proto-exchange/exchange"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ExchangeServiceServer struct {
	pb.UnimplementedExchangeServiceServer
	DB *gorm.DB
}

// Реализация метода GetExchangeRateForCurrency
func (s *ExchangeServiceServer) GetExchangeRateForCurrency(ctx context.Context, req *pb.CurrencyRequest) (*pb.ExchangeRateResponse, error) {
	log.Printf("Получен запрос на курс %s к %s", req.FromCurrency, req.ToCurrency)

	// Проверка валидности входных данных
	if req.FromCurrency == "" || req.ToCurrency == "" {
		return nil, status.Errorf(codes.InvalidArgument, "параметры валют не могут быть пустыми")
	}

	var rate float64

	// Ищем курс в базе данных
	err := s.DB.Table("exchange_rates").
		Select("rate").
		Where("from_currency = ? AND to_currency = ?", req.FromCurrency, req.ToCurrency).
		Scan(&rate).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "курс %s к %s не найден", req.FromCurrency, req.ToCurrency)
		}
		log.Printf("Ошибка при запросе к базе данных: %v", err)
		return nil, status.Errorf(codes.Internal, "не удалось получить курс из базы данных")
	}

	// Возвращаем успешный ответ
	return &pb.ExchangeRateResponse{
		FromCurrency: req.FromCurrency,
		ToCurrency:   req.ToCurrency,
		Rate:         float32(rate),
	}, nil
}
