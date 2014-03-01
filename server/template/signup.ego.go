package template

//line signup.ego:3
import (
	"fmt"
	"io"
)

//line signup.ego:1
func (t *Template) Signup(w io.Writer) error {
//line signup.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line signup.ego:9
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line signup.ego:10
	if _, err := fmt.Fprintf(w, "<!DOCTYPE html>\n"); err != nil {
		return err
	}
//line signup.ego:11
	if _, err := fmt.Fprintf(w, "<html lang=\"en\">\n  "); err != nil {
		return err
	}
//line signup.ego:12
	t.Head(w, "")
//line signup.ego:13
	if _, err := fmt.Fprintf(w, "\n\n  "); err != nil {
		return err
	}
//line signup.ego:14
	if _, err := fmt.Fprintf(w, "<body id=\"signup\">\n    "); err != nil {
		return err
	}
//line signup.ego:15
	if _, err := fmt.Fprintf(w, "<div class=\"container\">\n      "); err != nil {
		return err
	}
//line signup.ego:16
	t.Nav(w)
//line signup.ego:17
	if _, err := fmt.Fprintf(w, "\n      "); err != nil {
		return err
	}
//line signup.ego:17
	t.Flash(w)
//line signup.ego:18
	if _, err := fmt.Fprintf(w, "\n\n      "); err != nil {
		return err
	}
//line signup.ego:19
	if _, err := fmt.Fprintf(w, "<div class=\"panel panel-default\">\n        "); err != nil {
		return err
	}
//line signup.ego:20
	if _, err := fmt.Fprintf(w, "<div class=\"panel-heading\">\n          "); err != nil {
		return err
	}
//line signup.ego:21
	if _, err := fmt.Fprintf(w, "<h3 class=\"panel-title\">Sign Up"); err != nil {
		return err
	}
//line signup.ego:21
	if _, err := fmt.Fprintf(w, "</h3>\n        "); err != nil {
		return err
	}
//line signup.ego:22
	if _, err := fmt.Fprintf(w, "</div>\n        "); err != nil {
		return err
	}
//line signup.ego:23
	if _, err := fmt.Fprintf(w, "<div class=\"panel-body\">\n          "); err != nil {
		return err
	}
//line signup.ego:24
	if _, err := fmt.Fprintf(w, "<form role=\"form\" action=\"/signup\" method=\"POST\">\n            "); err != nil {
		return err
	}
//line signup.ego:25
	if _, err := fmt.Fprintf(w, "<div class=\"form-group\">\n              "); err != nil {
		return err
	}
//line signup.ego:26
	if _, err := fmt.Fprintf(w, "<label for=\"email\">E-mail address"); err != nil {
		return err
	}
//line signup.ego:26
	if _, err := fmt.Fprintf(w, "</label>\n              "); err != nil {
		return err
	}
//line signup.ego:27
	if _, err := fmt.Fprintf(w, "<input type=\"text\" class=\"form-control\" id=\"email\" name=\"email\" placeholder=\"Enter your e-mail address\">\n            "); err != nil {
		return err
	}
//line signup.ego:28
	if _, err := fmt.Fprintf(w, "</div>\n            "); err != nil {
		return err
	}
//line signup.ego:29
	if _, err := fmt.Fprintf(w, "<div class=\"form-group\">\n              "); err != nil {
		return err
	}
//line signup.ego:30
	if _, err := fmt.Fprintf(w, "<label for=\"password\">Password"); err != nil {
		return err
	}
//line signup.ego:30
	if _, err := fmt.Fprintf(w, "</label>\n              "); err != nil {
		return err
	}
//line signup.ego:31
	if _, err := fmt.Fprintf(w, "<input type=\"password\" class=\"form-control\" id=\"password\" name=\"password\" placeholder=\"Choose a password\">\n            "); err != nil {
		return err
	}
//line signup.ego:32
	if _, err := fmt.Fprintf(w, "</div>\n            "); err != nil {
		return err
	}
//line signup.ego:33
	if _, err := fmt.Fprintf(w, "<button class=\"btn btn-lg btn-success btn-block\" type=\"submit\">Sign up"); err != nil {
		return err
	}
//line signup.ego:33
	if _, err := fmt.Fprintf(w, "</button>\n          "); err != nil {
		return err
	}
//line signup.ego:34
	if _, err := fmt.Fprintf(w, "</div>\n        "); err != nil {
		return err
	}
//line signup.ego:35
	if _, err := fmt.Fprintf(w, "</div>\n      "); err != nil {
		return err
	}
//line signup.ego:36
	if _, err := fmt.Fprintf(w, "</div>\n\n    "); err != nil {
		return err
	}
//line signup.ego:38
	if _, err := fmt.Fprintf(w, "</div> "); err != nil {
		return err
	}
//line signup.ego:38
	if _, err := fmt.Fprintf(w, "<!-- /container -->\n  "); err != nil {
		return err
	}
//line signup.ego:39
	if _, err := fmt.Fprintf(w, "</body>\n"); err != nil {
		return err
	}
//line signup.ego:40
	if _, err := fmt.Fprintf(w, "</html>\n\n"); err != nil {
		return err
	}
	return nil
}
