package template

//line index.ego:3
import (
	"fmt"
	"io"
	"os"
)

//line index.ego:1
func (t *Template) Index(w io.Writer) error {
//line index.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line index.ego:10
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line index.ego:12
	hostname, _ := os.Hostname()
	if hostname == "" {
		hostname = "HOSTNAME"
	}

//line index.ego:18
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line index.ego:19
	if _, err := fmt.Fprintf(w, "<!DOCTYPE html>\n"); err != nil {
		return err
	}
//line index.ego:20
	if _, err := fmt.Fprintf(w, "<html lang=\"en\">\n  "); err != nil {
		return err
	}
//line index.ego:21
	t.Head(w, "")
//line index.ego:22
	if _, err := fmt.Fprintf(w, "\n\n  "); err != nil {
		return err
	}
//line index.ego:23
	if _, err := fmt.Fprintf(w, "<body id=\"index\">\n    "); err != nil {
		return err
	}
//line index.ego:24
	if _, err := fmt.Fprintf(w, "<div class=\"container\">\n      "); err != nil {
		return err
	}
//line index.ego:25
	t.Nav(w)
//line index.ego:26
	if _, err := fmt.Fprintf(w, "\n\n      "); err != nil {
		return err
	}
//line index.ego:27
	if _, err := fmt.Fprintf(w, "<!-- Main component for a primary marketing message or call to action -->\n      "); err != nil {
		return err
	}
//line index.ego:28
	if _, err := fmt.Fprintf(w, "<div class=\"jumbotron\">\n        "); err != nil {
		return err
	}
//line index.ego:29
	if _, err := fmt.Fprintf(w, "<h1>Welcome to Skybox"); err != nil {
		return err
	}
//line index.ego:29
	if _, err := fmt.Fprintf(w, "</h1>\n        "); err != nil {
		return err
	}
//line index.ego:30
	if _, err := fmt.Fprintf(w, "<p>\n          Skybox is a tool for tracking user events and querying them through funnel analysis.\n          Funnel analysis is a way to see how users in your application move through multi-step tasks (such as signing up or purchasing).\n        "); err != nil {
		return err
	}
//line index.ego:33
	if _, err := fmt.Fprintf(w, "</p>\n\n        "); err != nil {
		return err
	}
//line index.ego:35
	if _, err := fmt.Fprintf(w, "<p>\n          To add event tracking to your application, simple copy the following HTML snippet and paste it into your code:\n        "); err != nil {
		return err
	}
//line index.ego:37
	if _, err := fmt.Fprintf(w, "</p>\n\n        "); err != nil {
		return err
	}
//line index.ego:39
	if _, err := fmt.Fprintf(w, "<pre>\n&lt;script type=\"text/javascript\" src=\"//"); err != nil {
		return err
	}
//line index.ego:40
	if _, err := fmt.Fprintf(w, "%v", hostname); err != nil {
		return err
	}
//line index.ego:40
	if _, err := fmt.Fprintf(w, "/skybox.js\" async\n  data-api-key=\""); err != nil {
		return err
	}
//line index.ego:41
	if _, err := fmt.Fprintf(w, "%v", t.Account.APIKey); err != nil {
		return err
	}
//line index.ego:41
	if _, err := fmt.Fprintf(w, "\"&gt;\n&lt;/script&gt;\n        "); err != nil {
		return err
	}
//line index.ego:43
	if _, err := fmt.Fprintf(w, "</pre>\n        "); err != nil {
		return err
	}
//line index.ego:44
	if _, err := fmt.Fprintf(w, "<p>\n          "); err != nil {
		return err
	}
//line index.ego:45
	if _, err := fmt.Fprintf(w, "<a class=\"btn btn-lg btn-primary\" href=\"https://github.com/skybox/skybox\" target=\"_blank\" role=\"button\">Find us on Github &raquo;"); err != nil {
		return err
	}
//line index.ego:45
	if _, err := fmt.Fprintf(w, "</a>\n        "); err != nil {
		return err
	}
//line index.ego:46
	if _, err := fmt.Fprintf(w, "</p>\n      "); err != nil {
		return err
	}
//line index.ego:47
	if _, err := fmt.Fprintf(w, "</div>\n\n    "); err != nil {
		return err
	}
//line index.ego:49
	if _, err := fmt.Fprintf(w, "</div> "); err != nil {
		return err
	}
//line index.ego:49
	if _, err := fmt.Fprintf(w, "<!-- /container -->\n  "); err != nil {
		return err
	}
//line index.ego:50
	if _, err := fmt.Fprintf(w, "</body>\n"); err != nil {
		return err
	}
//line index.ego:51
	if _, err := fmt.Fprintf(w, "</html>\n\n"); err != nil {
		return err
	}
	return nil
}
