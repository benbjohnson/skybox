package template

//line login.ego:3
import (
	"fmt"
	"io"
)

//line login.ego:1
func (t *Template) Login(w io.Writer) error {
//line login.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line login.ego:9
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line login.ego:10
	if _, err := fmt.Fprintf(w, "<!DOCTYPE html>\n"); err != nil {
		return err
	}
//line login.ego:11
	if _, err := fmt.Fprintf(w, "<html lang=\"en\">\n  "); err != nil {
		return err
	}
//line login.ego:12
	t.Head(w, "")
//line login.ego:13
	if _, err := fmt.Fprintf(w, "\n\n  "); err != nil {
		return err
	}
//line login.ego:14
	if _, err := fmt.Fprintf(w, "<body id=\"login\">\n    "); err != nil {
		return err
	}
//line login.ego:15
	if _, err := fmt.Fprintf(w, "<div class=\"container\">\n      "); err != nil {
		return err
	}
//line login.ego:16
	if _, err := fmt.Fprintf(w, "<div class=\"panel panel-default\">\n        "); err != nil {
		return err
	}
//line login.ego:17
	if _, err := fmt.Fprintf(w, "<div class=\"panel-heading\">\n          "); err != nil {
		return err
	}
//line login.ego:18
	if _, err := fmt.Fprintf(w, "<h3 class=\"panel-title\">Sign In"); err != nil {
		return err
	}
//line login.ego:18
	if _, err := fmt.Fprintf(w, "</h3>\n        "); err != nil {
		return err
	}
//line login.ego:19
	if _, err := fmt.Fprintf(w, "</div>\n        "); err != nil {
		return err
	}
//line login.ego:20
	if _, err := fmt.Fprintf(w, "<div class=\"panel-body\">\n          "); err != nil {
		return err
	}
//line login.ego:21
	if _, err := fmt.Fprintf(w, "<form role=\"form\">\n            "); err != nil {
		return err
	}
//line login.ego:22
	if _, err := fmt.Fprintf(w, "<div class=\"form-group\">\n              "); err != nil {
		return err
	}
//line login.ego:23
	if _, err := fmt.Fprintf(w, "<label for=\"email\">E-mail address"); err != nil {
		return err
	}
//line login.ego:23
	if _, err := fmt.Fprintf(w, "</label>\n              "); err != nil {
		return err
	}
//line login.ego:24
	if _, err := fmt.Fprintf(w, "<input type=\"email\" class=\"form-control\" id=\"email\" placeholder=\"Enter email\">\n            "); err != nil {
		return err
	}
//line login.ego:25
	if _, err := fmt.Fprintf(w, "</div>\n            "); err != nil {
		return err
	}
//line login.ego:26
	if _, err := fmt.Fprintf(w, "<div class=\"form-group\">\n              "); err != nil {
		return err
	}
//line login.ego:27
	if _, err := fmt.Fprintf(w, "<label for=\"password\">Password"); err != nil {
		return err
	}
//line login.ego:27
	if _, err := fmt.Fprintf(w, "</label>\n              "); err != nil {
		return err
	}
//line login.ego:28
	if _, err := fmt.Fprintf(w, "<input type=\"password\" class=\"form-control\" id=\"password\" placeholder=\"Password\">\n            "); err != nil {
		return err
	}
//line login.ego:29
	if _, err := fmt.Fprintf(w, "</div>\n            "); err != nil {
		return err
	}
//line login.ego:30
	if _, err := fmt.Fprintf(w, "<button class=\"btn btn-lg btn-primary btn-block\" type=\"submit\">Sign in"); err != nil {
		return err
	}
//line login.ego:30
	if _, err := fmt.Fprintf(w, "</button>\n          "); err != nil {
		return err
	}
//line login.ego:31
	if _, err := fmt.Fprintf(w, "</div>\n        "); err != nil {
		return err
	}
//line login.ego:32
	if _, err := fmt.Fprintf(w, "</div>\n      "); err != nil {
		return err
	}
//line login.ego:33
	if _, err := fmt.Fprintf(w, "</div>\n\n    "); err != nil {
		return err
	}
//line login.ego:35
	if _, err := fmt.Fprintf(w, "</div> "); err != nil {
		return err
	}
//line login.ego:35
	if _, err := fmt.Fprintf(w, "<!-- /container -->\n  "); err != nil {
		return err
	}
//line login.ego:36
	if _, err := fmt.Fprintf(w, "</body>\n"); err != nil {
		return err
	}
//line login.ego:37
	if _, err := fmt.Fprintf(w, "</html>\n\n"); err != nil {
		return err
	}
	return nil
}
