package exception

type ItemListing struct{}

// Override error
func (e *ItemListing) Error() string {
	return "item listing failed"
}
