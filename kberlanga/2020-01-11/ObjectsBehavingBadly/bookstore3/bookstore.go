package bookstore3

/* structs ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++*/

/* represent information about a book */
type Book struct {
	Title           string
	Author          string
	PriceCents      int
	DiscountPercent int
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ structs */

/*################################################################################################*/
/*=== methods ===*/
/*################################################################################################*/
/* NetPrice XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX*/
/**
 * metodo para traer el precio del libro con el descuento correspondiente
 */
func (b *Book) NetPrice() int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving
}
