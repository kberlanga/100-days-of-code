package bookstore3

import (
	"testing"
)

func TestBook_NetPrice(t *testing.T) {
	type fields struct {
		Title           string
		Author          string
		PriceCents      int
		DiscountPercent int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "25 off",
			fields: fields{Title: "Title", Author: "Author", PriceCents: 100, DiscountPercent: 25},
			want:   75,
		},
		{
			name:   "without discount",
			fields: fields{Title: "Title", Author: "Author", PriceCents: 100, DiscountPercent: 0},
			want:   100,
		},
		{
			name:   "free",
			fields: fields{Title: "Title", Author: "Author", PriceCents: 100, DiscountPercent: 100},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Book{
				Title:           tt.fields.Title,
				Author:          tt.fields.Author,
				PriceCents:      tt.fields.PriceCents,
				DiscountPercent: tt.fields.DiscountPercent,
			}
			if got := b.NetPrice(); got != tt.want {
				t.Errorf("Book.NetPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSalePriceCents(t *testing.T) {
	b := Book{
		Title:      "A Clockwork Orange Soda",
		PriceCents: 500,
	}
	want := 250
	got := b.SalePriceCents()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestCustomer_MailingLabel(t *testing.T) {
	type fields struct {
		Title   string
		Name    string
		Address string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Normal customer",
			fields: fields{Title: "Internal Intranet Supervisor", Name: "Anita Jenkins", Address: "460 Barton Flats"},
			want:   "Customer: Internal Intranet Supervisor\nName: Anita Jenkins\nAddress: 460 Barton Flats",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Customer{
				Title:   tt.fields.Title,
				Name:    tt.fields.Name,
				Address: tt.fields.Address,
			}
			if got := b.MailingLabel(); got != tt.want {
				t.Errorf("Customer.MailingLabel() = %v, want %v", got, tt.want)
			}
		})
	}
}
