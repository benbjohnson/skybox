package template

//line index.ego:3
import "fmt"

//line index.ego:4
import "io"

//line index.ego:1
func (t *Template) Index(w io.Writer) error {
//line index.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line index.ego:4
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
//line index.ego:5
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line index.ego:6
	if _, err := fmt.Fprintf(w, "<!DOCTYPE html>\n"); err != nil {
		return err
	}
//line index.ego:7
	if _, err := fmt.Fprintf(w, "<html lang=\"en\">\n  "); err != nil {
		return err
	}
//line index.ego:8
	t.Head(w, "")
//line index.ego:9
	if _, err := fmt.Fprintf(w, "\n\n  "); err != nil {
		return err
	}
//line index.ego:10
	if _, err := fmt.Fprintf(w, "<body id=\"index\">\n    "); err != nil {
		return err
	}
//line index.ego:11
	if _, err := fmt.Fprintf(w, "<div class=\"container\">\n      "); err != nil {
		return err
	}
//line index.ego:12
	t.Nav(w)
//line index.ego:13
	if _, err := fmt.Fprintf(w, "\n\n    "); err != nil {
		return err
	}
//line index.ego:14
	if _, err := fmt.Fprintf(w, "</div> "); err != nil {
		return err
	}
//line index.ego:14
	if _, err := fmt.Fprintf(w, "<!-- /container -->\n  "); err != nil {
		return err
	}
//line index.ego:15
	if _, err := fmt.Fprintf(w, "</body>\n"); err != nil {
		return err
	}
//line index.ego:16
	if _, err := fmt.Fprintf(w, "</html>\n\n"); err != nil {
		return err
	}
	return nil
}
