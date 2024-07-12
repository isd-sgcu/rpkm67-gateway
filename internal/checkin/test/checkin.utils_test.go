package test

import (
	"github.com/bxcodec/faker/v4"
	checkinProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/checkin/checkin/v1"
)

func MockCheckInProto() *checkinProto.CheckIn {
	return &checkinProto.CheckIn{
		Id:     faker.UUIDDigit(),
		UserId: faker.UUIDDigit(),
		Email:  faker.Email(),
	}
}

func MockCheckInsProto() []*checkinProto.CheckIn {
	var checkIns []*checkinProto.CheckIn
	for i := 0; i < 10; i++ {
		checkIn := &checkinProto.CheckIn{
			Id:     faker.UUIDDigit(),
			UserId: faker.UUIDDigit(),
			Email:  faker.Email(),
		}
		checkIns = append(checkIns, checkIn)
	}
	return checkIns
}
