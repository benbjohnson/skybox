package template

//line flash.ego:3
import "fmt"

//line flash.ego:4
import "io"

//line flash.ego:5
import "unicode"

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
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
//line flash.ego:6
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line flash.ego:7
	for _, flash := range t.Flashes {
//line flash.ego:8
		if _, err := fmt.Fprintf(w, "\n  "); err != nil {
			return err
		}
//line flash.ego:8
		if len(flash) > 0 {
//line flash.ego:9
			if _, err := fmt.Fprintf(w, "\n    "); err != nil {
				return err
			}
//line flash.ego:9
			flash = fmt.Sprintf("%c%s", unicode.ToUpper(rune(flash[0])), flash[1:])
//line flash.ego:10
			if _, err := fmt.Fprintf(w, "\n    "); err != nil {
				return err
			}
//line flash.ego:10
			if _, err := fmt.Fprintf(w, "<p class=\"alert bg-danger\">\n      "); err != nil {
				return err
			}
//line flash.ego:11
			if _, err := fmt.Fprintf(w, "%v", flash); err != nil {
				return err
			}
//line flash.ego:12
			if _, err := fmt.Fprintf(w, "\n    "); err != nil {
				return err
			}
//line flash.ego:12
			if _, err := fmt.Fprintf(w, "</p>\n  "); err != nil {
				return err
			}
//line flash.ego:13
		}
//line flash.ego:14
		if _, err := fmt.Fprintf(w, "\n"); err != nil {
			return err
		}
//line flash.ego:14
	}
//line flash.ego:15
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
	return nil
}
