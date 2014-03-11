package template

//line dashboard.ego:3
import (
	"fmt"
	"io"
	"os"
)

//line dashboard.ego:1
func (t *Template) Dashboard(w io.Writer) error {
//line dashboard.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line dashboard.ego:10
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line dashboard.ego:12
	hostname, _ := os.Hostname()
	if hostname == "" {
		hostname = "HOSTNAME"
	}

//line dashboard.ego:18
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line dashboard.ego:19
	if _, err := fmt.Fprintf(w, "<!DOCTYPE html>\n"); err != nil {
		return err
	}
//line dashboard.ego:20
	if _, err := fmt.Fprintf(w, "<html lang=\"en\">\n  "); err != nil {
		return err
	}
//line dashboard.ego:21
	t.Head(w, "")
//line dashboard.ego:22
	if _, err := fmt.Fprintf(w, "\n\n  "); err != nil {
		return err
	}
//line dashboard.ego:23
	if _, err := fmt.Fprintf(w, "<body id=\"index\">\n    "); err != nil {
		return err
	}
//line dashboard.ego:24
	if _, err := fmt.Fprintf(w, "<div class=\"container\">\n      "); err != nil {
		return err
	}
//line dashboard.ego:25
	t.Nav(w)
//line dashboard.ego:26
	if _, err := fmt.Fprintf(w, "\n\n      "); err != nil {
		return err
	}
//line dashboard.ego:27
	if _, err := fmt.Fprintf(w, "<div class=\"jumbotron\">\n        "); err != nil {
		return err
	}
//line dashboard.ego:28
	if _, err := fmt.Fprintf(w, "<h1>Welcome to Skybox"); err != nil {
		return err
	}
//line dashboard.ego:28
	if _, err := fmt.Fprintf(w, "</h1>\n        "); err != nil {
		return err
	}
//line dashboard.ego:29
	if _, err := fmt.Fprintf(w, "<p>\n          Skybox is a tool for tracking user events and querying them through funnel analysis.\n          Funnel analysis is a way to see how users in your application move through multi-step tasks (such as signing up or purchasing).\n        "); err != nil {
		return err
	}
//line dashboard.ego:32
	if _, err := fmt.Fprintf(w, "</p>\n\n        "); err != nil {
		return err
	}
//line dashboard.ego:34
	if _, err := fmt.Fprintf(w, "<p>\n          To add event tracking to your application, simple copy the following HTML snippet and paste it into your code:\n        "); err != nil {
		return err
	}
//line dashboard.ego:36
	if _, err := fmt.Fprintf(w, "</p>\n\n        "); err != nil {
		return err
	}
//line dashboard.ego:38
	if _, err := fmt.Fprintf(w, "<pre>\n&lt;script type=\"text/javascript\" src=\"//"); err != nil {
		return err
	}
//line dashboard.ego:39
	if _, err := fmt.Fprintf(w, "%v", hostname); err != nil {
		return err
	}
//line dashboard.ego:39
	if _, err := fmt.Fprintf(w, "/skybox.js\" async\n  data-api-key=\""); err != nil {
		return err
	}
//line dashboard.ego:40
	if _, err := fmt.Fprintf(w, "%v", t.Account.APIKey); err != nil {
		return err
	}
//line dashboard.ego:40
	if _, err := fmt.Fprintf(w, "\"&gt;\n&lt;/script&gt;\n        "); err != nil {
		return err
	}
//line dashboard.ego:42
	if _, err := fmt.Fprintf(w, "</pre>\n        "); err != nil {
		return err
	}
//line dashboard.ego:43
	if _, err := fmt.Fprintf(w, "<p>\n          "); err != nil {
		return err
	}
//line dashboard.ego:44
	if _, err := fmt.Fprintf(w, "<a class=\"btn btn-lg btn-primary\" href=\"https://github.com/skybox/skybox\" target=\"_blank\" role=\"button\">Find us on Github &raquo;"); err != nil {
		return err
	}
//line dashboard.ego:44
	if _, err := fmt.Fprintf(w, "</a>\n        "); err != nil {
		return err
	}
//line dashboard.ego:45
	if _, err := fmt.Fprintf(w, "</p>\n      "); err != nil {
		return err
	}
//line dashboard.ego:46
	if _, err := fmt.Fprintf(w, "</div>\n\n    "); err != nil {
		return err
	}
//line dashboard.ego:48
	if _, err := fmt.Fprintf(w, "</div> "); err != nil {
		return err
	}
//line dashboard.ego:48
	if _, err := fmt.Fprintf(w, "<!-- /container -->\n  "); err != nil {
		return err
	}
//line dashboard.ego:49
	if _, err := fmt.Fprintf(w, "</body>\n"); err != nil {
		return err
	}
//line dashboard.ego:50
	if _, err := fmt.Fprintf(w, "</html>\n\n"); err != nil {
		return err
	}
	return nil
}
