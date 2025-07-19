package strutil

import "github.com/microcosm-cc/bluemonday"

// sanitizeHTML sanitizes an input HTML string by removing potentially unsafe or harmful content.
func sanitizeHTML(s string) string {
	p := bluemonday.UGCPolicy()
	return p.Sanitize(s)
}

// sanitizeHTMLCustom sanitizes the input HTML string by allowing only the specified elements in allowedElements.
func sanitizeHTMLCustom(s string, allowedElements []string) string {
	p := bluemonday.NewPolicy()
	//TODO: extend implementation to better address complex options
	p.AllowElements(allowedElements...)
	return p.Sanitize(s)
}
