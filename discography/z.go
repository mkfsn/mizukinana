package discography

var (
	All DiscographyList
)

func init() {
	All = append(Singles, Albums...)
}
