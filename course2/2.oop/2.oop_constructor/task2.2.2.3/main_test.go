package main

import (
	"reflect"
	"sync"
	"testing"
)

func TestNewCustomer(t *testing.T) {
	type args struct {
		ID   int
		cops []CustomerOption
	}
	tests := []struct {
		name string
		args args
		want *Customer
	}{
		{"Test new customer", args{1, []CustomerOption{}}, &Customer{ID: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomer(tt.args.ID, tt.args.cops...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomer() = %v, want %v", got, tt.want)
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
		{"Test savings account positive amount withdraw", &SavingsAccount{1000, sync.Mutex{}}, args{200}, false},
		{"Test savings account negative amount withdraw", &SavingsAccount{1000, sync.Mutex{}}, args{-200}, true},
		{"Test savings account withdraw (not enough balance)", &SavingsAccount{1000, sync.Mutex{}}, args{1001}, true},
		{"Test savings account withdraw (balance less than 1000)", &SavingsAccount{999, sync.Mutex{}}, args{100}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Withdraw(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("SavingsAccount.Withdraw() error = %v, wantErr %v", err, tt.wantErr)
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
		{"Test savings account balance", &SavingsAccount{1000, sync.Mutex{}}, 1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Balance(); got != tt.want {
				t.Errorf("SavingsAccount.Balance() = %v, want %v", got, tt.want)
			}
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
