package templates

//line index.ego:3
import (
	"fmt"
	"io"
)

//line index.ego:1
func Index(w io.Writer) error {
//line index.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line index.ego:9
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line index.ego:10
	if _, err := fmt.Fprintf(w, "<!DOCTYPE html>\n"); err != nil {
		return err
	}
//line index.ego:11
	if _, err := fmt.Fprintf(w, "<html lang=\"en\">\n  "); err != nil {
		return err
	}
//line index.ego:12
	Head(w, "")
//line index.ego:13
	if _, err := fmt.Fprintf(w, "\n\n  "); err != nil {
		return err
	}
//line index.ego:14
	if _, err := fmt.Fprintf(w, "<body>\n    "); err != nil {
		return err
	}
//line index.ego:15
	if _, err := fmt.Fprintf(w, "<div class=\"container\">\n      "); err != nil {
		return err
	}
//line index.ego:16
	Nav(w)
//line index.ego:17
	if _, err := fmt.Fprintf(w, "\n\n      "); err != nil {
		return err
	}
//line index.ego:18
	if _, err := fmt.Fprintf(w, "<!-- Main component for a primary marketing message or call to action -->\n      "); err != nil {
		return err
	}
//line index.ego:19
	if _, err := fmt.Fprintf(w, "<div class=\"jumbotron\">\n        "); err != nil {
		return err
	}
//line index.ego:20
	if _, err := fmt.Fprintf(w, "<h1>Navbar example"); err != nil {
		return err
	}
//line index.ego:20
	if _, err := fmt.Fprintf(w, "</h1>\n        "); err != nil {
		return err
	}
//line index.ego:21
	if _, err := fmt.Fprintf(w, "<p>This example is a quick exercise to illustrate how the default, static navbar and fixed to top navbar work. It includes the responsive CSS and HTML, so it also adapts to your viewport and device."); err != nil {
		return err
	}
//line index.ego:21
	if _, err := fmt.Fprintf(w, "</p>\n        "); err != nil {
		return err
	}
//line index.ego:22
	if _, err := fmt.Fprintf(w, "<p>\n          "); err != nil {
		return err
	}
//line index.ego:23
	if _, err := fmt.Fprintf(w, "<a class=\"btn btn-lg btn-primary\" href=\"../../components/#navbar\" role=\"button\">View navbar docs &raquo;"); err != nil {
		return err
	}
//line index.ego:23
	if _, err := fmt.Fprintf(w, "</a>\n        "); err != nil {
		return err
	}
//line index.ego:24
	if _, err := fmt.Fprintf(w, "</p>\n      "); err != nil {
		return err
	}
//line index.ego:25
	if _, err := fmt.Fprintf(w, "</div>\n\n    "); err != nil {
		return err
	}
//line index.ego:27
	if _, err := fmt.Fprintf(w, "</div> "); err != nil {
		return err
	}
//line index.ego:27
	if _, err := fmt.Fprintf(w, "<!-- /container -->\n  "); err != nil {
		return err
	}
//line index.ego:28
	if _, err := fmt.Fprintf(w, "</body>\n"); err != nil {
		return err
	}
//line index.ego:29
	if _, err := fmt.Fprintf(w, "</html>\n\n"); err != nil {
		return err
	}
	return nil
}
