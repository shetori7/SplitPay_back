package usecase

import (
	"SplitPay_back/internal/domain"
	"errors"
)

type PaymentInteractor struct {
	WariPaymentRepository WariPaymentRepository
	WariLoanRepository    WariLoanRepository
}

func (interactor *PaymentInteractor) Add(groupId int, payerId int, amount float64, participantIds []int) error {
	wp := domain.Wari_payment{
		PayerGroupId: groupId,
		PayerUserId:  payerId,
		Amount:       amount,
	}
	//支払者の情報をDBに登録する
	interactor.WariPaymentRepository.Store(&wp)
	if wp.PaymentId == 0 {
		return errors.New("failed to store payment information")
	}
	//立替してもらった人の情報をDBに登録する
	for _, participantId := range participantIds {
		wl := domain.Wari_loan{
			PaymentId: wp.PaymentId,
			ToUserId:  participantId,
		}
		interactor.WariLoanRepository.Store(&wl)
	}
	//TODO:清算額の計算を行う
	return nil

}

func (interactor *PaymentInteractor) CalcuratePayment(amount float64, members int) float64 {
	if members == 0 {
		return 0
	}
	return amount / float64(members)
}

func (interactor *PaymentInteractor) ReCalcFinalPayment(groupId int) {
	//グループIDを元に、グループ内の全ての支払い情報を取得
	payments := interactor.WariPaymentRepository.SelectByGroupId(groupId)

	return
}
