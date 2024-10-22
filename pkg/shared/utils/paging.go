package utils

func CalculateTotalPages(totalItems, itemsPerPage int) int {
	// Handle edge cases
	if totalItems <= 0 || itemsPerPage <= 0 {
		return 0
	}

	// Calculate total pages, rounding up using integer division
	totalPages := (totalItems + itemsPerPage - 1) / itemsPerPage
	return totalPages
}
