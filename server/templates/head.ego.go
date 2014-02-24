package templates

//line head.ego:3
import (
	"fmt"
	"io"
)

//line head.ego:1
func Head(w io.Writer, title string) error {
//line head.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line head.ego:9
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line head.ego:11
	if title == "" {
		title = "Skybox"
	}

//line head.ego:16
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line head.ego:17
	if _, err := fmt.Fprintf(w, "<head>\n  "); err != nil {
		return err
	}
//line head.ego:18
	if _, err := fmt.Fprintf(w, "<meta charset=\"utf-8\">\n  "); err != nil {
		return err
	}
//line head.ego:19
	if _, err := fmt.Fprintf(w, "<meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n  "); err != nil {
		return err
	}
//line head.ego:20
	if _, err := fmt.Fprintf(w, "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n  "); err != nil {
		return err
	}
//line head.ego:21
	if _, err := fmt.Fprintf(w, "<title>"); err != nil {
		return err
	}
//line head.ego:21
	if _, err := fmt.Fprintf(w, title); err != nil {
		return err
	}
//line head.ego:21
	if _, err := fmt.Fprintf(w, "</title>\n  "); err != nil {
		return err
	}
//line head.ego:22
	if _, err := fmt.Fprintf(w, "<link href=\"/assets/bootstrap.min.css\" rel=\"stylesheet\">\n  "); err != nil {
		return err
	}
//line head.ego:23
	if _, err := fmt.Fprintf(w, "<link href=\"/assets/bootstrap-theme.min.css\" rel=\"stylesheet\">\n  "); err != nil {
		return err
	}
//line head.ego:24
	if _, err := fmt.Fprintf(w, "<link href=\"/assets/application.css\" rel=\"stylesheet\">\n  "); err != nil {
		return err
	}
//line head.ego:25
	if _, err := fmt.Fprintf(w, "<script src=\"/assets/jquery-2.1.0.min.js\">"); err != nil {
		return err
	}
//line head.ego:25
	if _, err := fmt.Fprintf(w, "</script>\n  "); err != nil {
		return err
	}
//line head.ego:26
	if _, err := fmt.Fprintf(w, "<script src=\"/assets/bootstrap.min.js\">"); err != nil {
		return err
	}
//line head.ego:26
	if _, err := fmt.Fprintf(w, "</script>\n"); err != nil {
		return err
	}
//line head.ego:27
	if _, err := fmt.Fprintf(w, "</head>\n"); err != nil {
		return err
	}
	return nil
}
