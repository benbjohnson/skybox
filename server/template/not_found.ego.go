package template

//line not_found.ego:3
import (
	"fmt"
	"io"
)

//line not_found.ego:1
func (t *Template) NotFound(w io.Writer) error {
//line not_found.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line not_found.ego:9
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line not_found.ego:10
	if _, err := fmt.Fprintf(w, "<!DOCTYPE html>\n"); err != nil {
		return err
	}
//line not_found.ego:11
	if _, err := fmt.Fprintf(w, "<html lang=\"en\">\n  "); err != nil {
		return err
	}
//line not_found.ego:12
	t.Head(w, "")
//line not_found.ego:13
	if _, err := fmt.Fprintf(w, "\n\n  "); err != nil {
		return err
	}
//line not_found.ego:14
	if _, err := fmt.Fprintf(w, "<body id=\"index\">\n    "); err != nil {
		return err
	}
//line not_found.ego:15
	if _, err := fmt.Fprintf(w, "<div class=\"container\">\n      "); err != nil {
		return err
	}
//line not_found.ego:16
	t.Nav(w)
//line not_found.ego:17
	if _, err := fmt.Fprintf(w, "\n\n      "); err != nil {
		return err
	}
//line not_found.ego:18
	if _, err := fmt.Fprintf(w, "<div class=\"page-header\">\n        "); err != nil {
		return err
	}
//line not_found.ego:19
	if _, err := fmt.Fprintf(w, "<h3>Not Found"); err != nil {
		return err
	}
//line not_found.ego:19
	if _, err := fmt.Fprintf(w, "</h3>\n      "); err != nil {
		return err
	}
//line not_found.ego:20
	if _, err := fmt.Fprintf(w, "</div>\n\n      "); err != nil {
		return err
	}
//line not_found.ego:22
	if _, err := fmt.Fprintf(w, "<p>The page you requested could not be found."); err != nil {
		return err
	}
//line not_found.ego:22
	if _, err := fmt.Fprintf(w, "</p>\n\n      "); err != nil {
		return err
	}
//line not_found.ego:24
	if _, err := fmt.Fprintf(w, "<p>\n        "); err != nil {
		return err
	}
//line not_found.ego:25
	if _, err := fmt.Fprintf(w, "<button class=\"btn btn-default\" onclick=\"window.history.back()\">Back"); err != nil {
		return err
	}
//line not_found.ego:25
	if _, err := fmt.Fprintf(w, "</button>\n      "); err != nil {
		return err
	}
//line not_found.ego:26
	if _, err := fmt.Fprintf(w, "</p>\n    "); err != nil {
		return err
	}
//line not_found.ego:27
	if _, err := fmt.Fprintf(w, "</div> "); err != nil {
		return err
	}
//line not_found.ego:27
	if _, err := fmt.Fprintf(w, "<!-- /container -->\n  "); err != nil {
		return err
	}
//line not_found.ego:28
	if _, err := fmt.Fprintf(w, "</body>\n"); err != nil {
		return err
	}
//line not_found.ego:29
	if _, err := fmt.Fprintf(w, "</html>\n\n"); err != nil {
		return err
	}
	return nil
}
