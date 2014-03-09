package template

//line funnel_edit.ego:3
import "encoding/json"

//line funnel_edit.ego:4
import "fmt"

//line funnel_edit.ego:5
import "io"

//line funnel_edit.ego:1
func (t *FunnelTemplate) Edit(w io.Writer) error {
//line funnel_edit.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line funnel_edit.ego:4
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
//line funnel_edit.ego:5
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
//line funnel_edit.ego:6
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line funnel_edit.ego:7
	if _, err := fmt.Fprintf(w, "<!DOCTYPE html>\n"); err != nil {
		return err
	}
//line funnel_edit.ego:8
	if _, err := fmt.Fprintf(w, "<html lang=\"en\">\n  "); err != nil {
		return err
	}
//line funnel_edit.ego:9
	t.Head(w, "")
//line funnel_edit.ego:10
	if _, err := fmt.Fprintf(w, "\n\n  "); err != nil {
		return err
	}
//line funnel_edit.ego:11
	if _, err := fmt.Fprintf(w, "<body id=\"index\">\n    "); err != nil {
		return err
	}
//line funnel_edit.ego:12
	if _, err := fmt.Fprintf(w, "<div class=\"container\">\n      "); err != nil {
		return err
	}
//line funnel_edit.ego:13
	t.Nav(w)
//line funnel_edit.ego:14
	if _, err := fmt.Fprintf(w, "\n\n      "); err != nil {
		return err
	}
//line funnel_edit.ego:15
	if _, err := fmt.Fprintf(w, "<div class=\"page-header\">\n        "); err != nil {
		return err
	}
//line funnel_edit.ego:16
	if _, err := fmt.Fprintf(w, "<h3>\n          "); err != nil {
		return err
	}
//line funnel_edit.ego:17
	if t.Funnel.ID() == 0 {
//line funnel_edit.ego:18
		if _, err := fmt.Fprintf(w, "\n            New Funnel\n          "); err != nil {
			return err
		}
//line funnel_edit.ego:19
	} else {
//line funnel_edit.ego:20
		if _, err := fmt.Fprintf(w, "\n            Edit Funnel\n          "); err != nil {
			return err
		}
//line funnel_edit.ego:21
	}
//line funnel_edit.ego:22
	if _, err := fmt.Fprintf(w, "\n        "); err != nil {
		return err
	}
//line funnel_edit.ego:22
	if _, err := fmt.Fprintf(w, "</h3>\n      "); err != nil {
		return err
	}
//line funnel_edit.ego:23
	if _, err := fmt.Fprintf(w, "</div>\n\n      "); err != nil {
		return err
	}
//line funnel_edit.ego:25
	if _, err := fmt.Fprintf(w, "<div class=\"row\">\n        "); err != nil {
		return err
	}
//line funnel_edit.ego:26
	if _, err := fmt.Fprintf(w, "<form action=\"/funnels/"); err != nil {
		return err
	}
//line funnel_edit.ego:26
	if _, err := fmt.Fprintf(w, "%v", t.Funnel.ID()); err != nil {
		return err
	}
//line funnel_edit.ego:26
	if _, err := fmt.Fprintf(w, "\" method=\"POST\" role=\"form\" class=\"col-sm-6 col-md-5 col-lg-5\">\n          "); err != nil {
		return err
	}
//line funnel_edit.ego:27
	if _, err := fmt.Fprintf(w, "<input type=\"hidden\" name=\"id\" value=\""); err != nil {
		return err
	}
//line funnel_edit.ego:27
	if _, err := fmt.Fprintf(w, "%v", t.Funnel.ID()); err != nil {
		return err
	}
//line funnel_edit.ego:27
	if _, err := fmt.Fprintf(w, "\"/>\n\n          "); err != nil {
		return err
	}
//line funnel_edit.ego:29
	if _, err := fmt.Fprintf(w, "<div class=\"form-group\">\n            "); err != nil {
		return err
	}
//line funnel_edit.ego:30
	if _, err := fmt.Fprintf(w, "<label for=\"name\">Funnel Name"); err != nil {
		return err
	}
//line funnel_edit.ego:30
	if _, err := fmt.Fprintf(w, "</label>\n            "); err != nil {
		return err
	}
//line funnel_edit.ego:31
	if _, err := fmt.Fprintf(w, "<input type=\"text\" class=\"form-control\" id=\"name\" name=\"name\" value=\""); err != nil {
		return err
	}
//line funnel_edit.ego:31
	if _, err := fmt.Fprintf(w, "%v", t.Funnel.Name); err != nil {
		return err
	}
//line funnel_edit.ego:31
	if _, err := fmt.Fprintf(w, "\"/>\n          "); err != nil {
		return err
	}
//line funnel_edit.ego:32
	if _, err := fmt.Fprintf(w, "</div>\n\n          "); err != nil {
		return err
	}
//line funnel_edit.ego:34
	if _, err := fmt.Fprintf(w, "<div class=\"steps\">"); err != nil {
		return err
	}
//line funnel_edit.ego:34
	if _, err := fmt.Fprintf(w, "</div>\n\n          "); err != nil {
		return err
	}
//line funnel_edit.ego:36
	if _, err := fmt.Fprintf(w, "<div class=\"form-group\">\n            "); err != nil {
		return err
	}
//line funnel_edit.ego:37
	if _, err := fmt.Fprintf(w, "<a class=\"add-step\" style=\"cursor: pointer;\" class=\"btn btn-link\">Add another step"); err != nil {
		return err
	}
//line funnel_edit.ego:37
	if _, err := fmt.Fprintf(w, "</a>\n          "); err != nil {
		return err
	}
//line funnel_edit.ego:38
	if _, err := fmt.Fprintf(w, "</div>\n\n          "); err != nil {
		return err
	}
//line funnel_edit.ego:40
	if t.Funnel.ID() == 0 {
//line funnel_edit.ego:41
		if _, err := fmt.Fprintf(w, "\n            "); err != nil {
			return err
		}
//line funnel_edit.ego:41
		if _, err := fmt.Fprintf(w, "<button type=\"submit\" class=\"btn btn-primary\">Create Funnel"); err != nil {
			return err
		}
//line funnel_edit.ego:41
		if _, err := fmt.Fprintf(w, "</button>\n          "); err != nil {
			return err
		}
//line funnel_edit.ego:42
	} else {
//line funnel_edit.ego:43
		if _, err := fmt.Fprintf(w, "\n            "); err != nil {
			return err
		}
//line funnel_edit.ego:43
		if _, err := fmt.Fprintf(w, "<button type=\"submit\" class=\"btn btn-primary\">Save Funnel"); err != nil {
			return err
		}
//line funnel_edit.ego:43
		if _, err := fmt.Fprintf(w, "</button>\n          "); err != nil {
			return err
		}
//line funnel_edit.ego:44
	}
//line funnel_edit.ego:45
	if _, err := fmt.Fprintf(w, "\n          "); err != nil {
		return err
	}
//line funnel_edit.ego:45
	if _, err := fmt.Fprintf(w, "<a class=\"btn btn-link\" href=\"/funnels\">Cancel"); err != nil {
		return err
	}
//line funnel_edit.ego:45
	if _, err := fmt.Fprintf(w, "</a>\n        "); err != nil {
		return err
	}
//line funnel_edit.ego:46
	if _, err := fmt.Fprintf(w, "</form>\n      "); err != nil {
		return err
	}
//line funnel_edit.ego:47
	if _, err := fmt.Fprintf(w, "</div>\n    "); err != nil {
		return err
	}
//line funnel_edit.ego:48
	if _, err := fmt.Fprintf(w, "</div> "); err != nil {
		return err
	}
//line funnel_edit.ego:48
	if _, err := fmt.Fprintf(w, "<!-- /container -->\n  "); err != nil {
		return err
	}
//line funnel_edit.ego:49
	if _, err := fmt.Fprintf(w, "</body>\n\n  "); err != nil {
		return err
	}
//line funnel_edit.ego:51
	if _, err := fmt.Fprintf(w, "<script src=\"/assets/funnel_edit.js\">"); err != nil {
		return err
	}
//line funnel_edit.ego:51
	if _, err := fmt.Fprintf(w, "</script>\n  "); err != nil {
		return err
	}
//line funnel_edit.ego:52
	if _, err := fmt.Fprintf(w, "<script>\n    var funnel = "); err != nil {
		return err
	}
//line funnel_edit.ego:53
	json.NewEncoder(w).Encode(t.Funnel)
//line funnel_edit.ego:53
	if _, err := fmt.Fprintf(w, ";\n    var resources = "); err != nil {
		return err
	}
//line funnel_edit.ego:54
	json.NewEncoder(w).Encode(t.Resources)
//line funnel_edit.ego:54
	if _, err := fmt.Fprintf(w, ";\n    render(funnel);\n  "); err != nil {
		return err
	}
//line funnel_edit.ego:56
	if _, err := fmt.Fprintf(w, "</script>\n"); err != nil {
		return err
	}
//line funnel_edit.ego:57
	if _, err := fmt.Fprintf(w, "</html>\n\n"); err != nil {
		return err
	}
	return nil
}
