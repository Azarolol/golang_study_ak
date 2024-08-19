package main

import "testing"

func TestCreditCard_Pay(t *testing.T) {
	type args struct {
		amount float64
	}
	tests := []struct {
		name    string
		c       *CreditCard
		args    args
		wantErr bool
	}{
		{"Test credit card pay", &CreditCard{500.00}, args{200.00}, false},
		{"Test credit card pay (negative amount)", &CreditCard{500.00}, args{-100.00}, true},
		{"Test credit card pay (not enough balance)", &CreditCard{500.00}, args{600.00}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Pay(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("CreditCard.Pay() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBitcoin_Pay(t *testing.T) {
	type args struct {
		amount float64
	}
	tests := []struct {
		name    string
		c       *Bitcoin
		args    args
		wantErr bool
	}{
		{"Test bitcoin pay", &Bitcoin{2.00}, args{1.00}, false},
		{"Test bitcoin pay (negative amount)", &Bitcoin{2.00}, args{-1.00}, true},
		{"Test bitcoin pay (not enough balance)", &Bitcoin{2.00}, args{2.01}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Pay(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("Bitcoin.Pay() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProccessPayment(t *testing.T) {
	type args struct {
		p      PaymentMethod
		amount float64
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test credit card proccess payment 1", args{&CreditCard{500.00}, 200.00}},
		{"Test credit card proccess payment 2", args{&CreditCard{500.00}, -100.00}},
		{"Test credit card proccess payment 3", args{&CreditCard{500.00}, 600.00}},
		{"Test bitcoin proccess payment 1", args{&Bitcoin{2.00}, 1.00}},
		{"Test bitcoin proccess payment 2", args{&CreditCard{2.00}, -1.00}},
		{"Test bitcoin card proccess payment 3", args{&CreditCard{2.00}, 2.01}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ProccessPayment(tt.args.p, tt.args.amount)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test main"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
