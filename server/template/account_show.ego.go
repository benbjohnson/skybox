package template

//line account_show.ego:3
import (
	"fmt"
	"io"
)

//line account_show.ego:1
func (t *AccountTemplate) Show(w io.Writer) error {
//line account_show.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line account_show.ego:9
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line account_show.ego:10
	if _, err := fmt.Fprintf(w, "<!DOCTYPE html>\n"); err != nil {
		return err
	}
//line account_show.ego:11
	if _, err := fmt.Fprintf(w, "<html lang=\"en\">\n  "); err != nil {
		return err
	}
//line account_show.ego:12
	t.Head(w, "")
//line account_show.ego:13
	if _, err := fmt.Fprintf(w, "\n\n  "); err != nil {
		return err
	}
//line account_show.ego:14
	if _, err := fmt.Fprintf(w, "<body id=\"index\">\n    "); err != nil {
		return err
	}
//line account_show.ego:15
	if _, err := fmt.Fprintf(w, "<div class=\"container\">\n      "); err != nil {
		return err
	}
//line account_show.ego:16
	t.Nav(w)
//line account_show.ego:17
	if _, err := fmt.Fprintf(w, "\n\n      "); err != nil {
		return err
	}
//line account_show.ego:18
	if _, err := fmt.Fprintf(w, "<div class=\"page-header\">\n        "); err != nil {
		return err
	}
//line account_show.ego:19
	if _, err := fmt.Fprintf(w, "<h3>My Account"); err != nil {
		return err
	}
//line account_show.ego:19
	if _, err := fmt.Fprintf(w, "</h3>\n      "); err != nil {
		return err
	}
//line account_show.ego:20
	if _, err := fmt.Fprintf(w, "</div>\n\n      "); err != nil {
		return err
	}
//line account_show.ego:22
	if _, err := fmt.Fprintf(w, "<div class=\"row\">\n        "); err != nil {
		return err
	}
//line account_show.ego:23
	if _, err := fmt.Fprintf(w, "<form role=\"form\" class=\"col-sm-6 col-md-5 col-lg-5\">\n          "); err != nil {
		return err
	}
//line account_show.ego:24
	if _, err := fmt.Fprintf(w, "<div class=\"form-group\">\n            "); err != nil {
		return err
	}
//line account_show.ego:25
	if _, err := fmt.Fprintf(w, "<label>API Key"); err != nil {
		return err
	}
//line account_show.ego:25
	if _, err := fmt.Fprintf(w, "</label>\n            "); err != nil {
		return err
	}
//line account_show.ego:26
	if _, err := fmt.Fprintf(w, "<input type=\"text\" class=\"form-control\" value=\""); err != nil {
		return err
	}
//line account_show.ego:26
	if _, err := fmt.Fprintf(w, "%v", t.Account.APIKey); err != nil {
		return err
	}
//line account_show.ego:26
	if _, err := fmt.Fprintf(w, "\" disabled/>\n          "); err != nil {
		return err
	}
//line account_show.ego:27
	if _, err := fmt.Fprintf(w, "</div>\n        "); err != nil {
		return err
	}
//line account_show.ego:28
	if _, err := fmt.Fprintf(w, "</form>\n      "); err != nil {
		return err
	}
//line account_show.ego:29
	if _, err := fmt.Fprintf(w, "</div>\n    "); err != nil {
		return err
	}
//line account_show.ego:30
	if _, err := fmt.Fprintf(w, "</div> "); err != nil {
		return err
	}
//line account_show.ego:30
	if _, err := fmt.Fprintf(w, "<!-- /container -->\n  "); err != nil {
		return err
	}
//line account_show.ego:31
	if _, err := fmt.Fprintf(w, "</body>\n"); err != nil {
		return err
	}
//line account_show.ego:32
	if _, err := fmt.Fprintf(w, "</html>\n\n"); err != nil {
		return err
	}
	return nil
}
