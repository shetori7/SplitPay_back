package usecase

import (
	"SplitPay_back/internal/domain"
	"errors"
	"sort"
)

type PaymentInteractor struct {
	WariPaymentRepository      WariPaymentRepository
	WariLoanRepository         WariLoanRepository
	WariFinalPaymentRepository WariFinalPaymentRepository
}

func (interactor *PaymentInteractor) Add(groupUuid string, payerId int, amount float64, participantIds []int) error {
	wp := domain.Wari_payment{
		PayerGroupId: groupUuid,
		PayerUserId:  payerId,
		PayerAmount:  amount,
	}
	//支払者の情報をDBに登録する
	interactor.WariPaymentRepository.Store(&wp)
	if wp.PaymentId == 0 {
		return errors.New("failed to store payment information")
	}
	//精算額を計算する
	PayeeAmount := interactor.CalcuratePayment(amount, len(participantIds))

	//立替してもらった人の情報をDBに登録する
	for _, participantId := range participantIds {
		wl := domain.Wari_loan{
			PaymentId:   wp.PaymentId,
			PayeeUserId: participantId,
			PayeeAmount: PayeeAmount,
		}
		interactor.WariLoanRepository.Store(&wl)
	}
	//Messageの取得
	//TODO:清算額の計算を行う
	return nil

}

func (interactor *PaymentInteractor) CalcuratePayment(amount float64, members int) float64 {
	if members == 0 {
		return 0
	}
	return amount / float64(members)
}

func (interactor *PaymentInteractor) ReCalcFinalPayment(groupUuid string) error {
	//精算額を計算するための構造体
	type PersonBalance struct {
		userId int
		Amount float64
	}

	var debtors []PersonBalance
	var creditors []PersonBalance
	var balances map[int]float64
	//構造体の初期化
	balances = make(map[int]float64)
	//既存の最終支払テーブルのレコードを削除する
	if err := interactor.WariPaymentRepository.DeleteByUuid(groupUuid); err != nil {
		return err
	}

	//グループIDを元に、グループ内の全ての支払い情報を取得
	reCalcFinalPaymentDtos := interactor.WariPaymentRepository.SelectPaymentAndLoanByGroupId(groupUuid)
	if len(reCalcFinalPaymentDtos) == 0 {
		return errors.New("no payment information found")
	}

	for _, dto := range reCalcFinalPaymentDtos {
		balances[dto.PayerUserId] += dto.PayeeAmount
		balances[dto.PayeeUserId] -= dto.PayeeAmount
	}

	// 分類：正の残高 → もらう、負の残高 → 払う
	for name, amount := range balances {
		if amount < 0 {
			debtors = append(debtors, PersonBalance{name, amount})
		} else if amount > 0 {
			creditors = append(creditors, PersonBalance{name, amount})
		}
	}

	// 並べ替え：残高が小さい順（支払う側）、大きい順（受け取る側）
	sort.Slice(debtors, func(i, j int) bool {
		return debtors[i].Amount < debtors[j].Amount
	})
	sort.Slice(creditors, func(i, j int) bool {
		return creditors[i].Amount > creditors[j].Amount
	})

	// 結果を格納するスライスを宣言

	var finalPayments []*domain.WariFinalPayment

	i, j := 0, 0
	for i < len(debtors) && j < len(creditors) {
		debtor := &debtors[i]
		creditor := &creditors[j]

		payment := min(-debtor.Amount, creditor.Amount)

		// 結果を構造体に保持する
		finalPayment := domain.WariFinalPayment{
			GroupUuid:  groupUuid,
			FromUserId: debtor.userId,
			ToUserId:   creditor.userId,
			Amount:     payment,
		}

		finalPayments = append(finalPayments, &finalPayment)
		debtor.Amount += payment
		creditor.Amount -= payment

		if debtor.Amount == 0 {
			i++
		}
		if creditor.Amount == 0 {
			j++
		}
	}

	interactor.WariFinalPaymentRepository.StoreAll(finalPayments)
	return nil
}
