package main

import "testing"

func TestCurrentAccount_Deposit(t *testing.T) {
	type args struct {
		amount float64
	}
	tests := []struct {
		name    string
		c       *CurrentAccount
		args    args
		wantErr bool
	}{
		{"Test current account positive deposit", &CurrentAccount{}, args{100}, false},
		{"Test current account negative deposit", &CurrentAccount{}, args{-100}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Deposit(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("CurrentAccount.Deposit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSavingsAccount_Deposit(t *testing.T) {
	type args struct {
		amount float64
	}
	tests := []struct {
		name    string
		c       *SavingsAccount
		args    args
		wantErr bool
	}{
		{"Test savings account positive deposit", &SavingsAccount{}, args{100}, false},
		{"Test savings account negative deposit", &SavingsAccount{}, args{-100}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Deposit(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("SavingsAccount.Deposit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCurrentAccount_Withdraw(t *testing.T) {
	type args struct {
		amount float64
	}
	tests := []struct {
		name    string
		c       *CurrentAccount
		args    args
		wantErr bool
	}{
		{"Test current account positive amount withdraw", &CurrentAccount{500}, args{200}, false},
		{"Test current account negative amount withdraw", &CurrentAccount{500}, args{-200}, true},
		{"Test current account withdraw (not enough balance)", &CurrentAccount{500}, args{600}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Withdraw(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("CurrentAccount.Withdraw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSavingsAccount_Withdraw(t *testing.T) {
	type args struct {
		amount float64
	}
	tests := []struct {
		name    string
		c       *SavingsAccount
		args    args
		wantErr bool
	}{
		{"Test savings account positive amount withdraw", &SavingsAccount{500}, args{200}, false},
		{"Test savings account negative amount withdraw", &SavingsAccount{500}, args{-200}, true},
		{"Test savings account withdraw (not enough balance)", &SavingsAccount{500}, args{600}, true},
		{"Test savings account withdraw (balance less than 500)", &SavingsAccount{499}, args{100}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Withdraw(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("SavingsAccount.Withdraw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCurrentAccount_Balance(t *testing.T) {
	tests := []struct {
		name string
		c    *CurrentAccount
		want float64
	}{
		{"Test current account balance", &CurrentAccount{100}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Balance(); got != tt.want {
				t.Errorf("CurrentAccount.Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSavingsAccount_Balance(t *testing.T) {
	tests := []struct {
		name string
		c    *SavingsAccount
		want float64
	}{
		{"Test savings account balance", &SavingsAccount{100}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Balance(); got != tt.want {
				t.Errorf("SavingsAccount.Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessAccount(t *testing.T) {
	type args struct {
		a Accounter
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test proccess current account", args{&CurrentAccount{}}},
		{"Test proccess savings account", args{&SavingsAccount{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ProcessAccount(tt.args.a)
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
