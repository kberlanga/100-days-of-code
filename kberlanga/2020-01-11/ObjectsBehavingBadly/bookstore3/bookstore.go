package bookstore3

/* structs ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++*/

/* represent information about a book */
type Book struct {
	Title           string
	Author          string
	PriceCents      int
	DiscountPercent int
}

/* represent information about a customer */
type Customer struct {
	Title   string
	Name    string
	Address string
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ structs */

/*################################################################################################*/
/*=== methods ===*/
/*################################################################################################*/
/* NetPrice XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX*/
/**
 * metodo that returns the discounted price of a book
 */
func (b *Book) NetPrice() int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving
}

/* SalePriceCents XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX*/
/**
 * method that returns 50 % off book price
 */
func (b *Book) SalePriceCents() int {
	return b.PriceCents - (b.PriceCents * 50 / 100)
}

func (b *Customer) MailingLabel() string {
	return "Customer: " + b.Title + "\nName: " + b.Name + "\nAddress: " + b.Address
}
