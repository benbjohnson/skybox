package template

//line flash.ego:3
import "fmt"

//line flash.ego:4
import "io"

//line flash.ego:1
func (t *Template) Flash(w io.Writer) error {
//line flash.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line flash.ego:4
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
//line flash.ego:5
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line flash.ego:6
	if t.Session != nil {
//line flash.ego:7
		if _, err := fmt.Fprintf(w, "\n  "); err != nil {
			return err
		}
//line flash.ego:7
		for _, flash := range t.Session.Flashes() {
//line flash.ego:8
			if _, err := fmt.Fprintf(w, "\n    "); err != nil {
				return err
			}
//line flash.ego:8
			if _, err := fmt.Fprintf(w, "<p class=\"bg-danger\">"); err != nil {
				return err
			}
//line flash.ego:8
			if _, err := fmt.Fprintf(w, fmt.Sprintf("%v", flash)); err != nil {
				return err
			}
//line flash.ego:8
			if _, err := fmt.Fprintf(w, "</p>\n  "); err != nil {
				return err
			}
//line flash.ego:9
		}
//line flash.ego:10
		if _, err := fmt.Fprintf(w, "\n"); err != nil {
			return err
		}
//line flash.ego:10
	}
//line flash.ego:11
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
	return nil
}
