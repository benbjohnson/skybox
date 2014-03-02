package template

//line project_edit.ego:3
import (
	"fmt"
	"io"
)

//line project_edit.ego:1
func (t *ProjectTemplate) Edit(w io.Writer) error {
//line project_edit.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line project_edit.ego:9
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line project_edit.ego:10
	if _, err := fmt.Fprintf(w, "<!DOCTYPE html>\n"); err != nil {
		return err
	}
//line project_edit.ego:11
	if _, err := fmt.Fprintf(w, "<html lang=\"en\">\n  "); err != nil {
		return err
	}
//line project_edit.ego:12
	t.Head(w, "")
//line project_edit.ego:13
	if _, err := fmt.Fprintf(w, "\n\n  "); err != nil {
		return err
	}
//line project_edit.ego:14
	if _, err := fmt.Fprintf(w, "<body id=\"index\">\n    "); err != nil {
		return err
	}
//line project_edit.ego:15
	if _, err := fmt.Fprintf(w, "<div class=\"container\">\n      "); err != nil {
		return err
	}
//line project_edit.ego:16
	t.Nav(w)
//line project_edit.ego:17
	if _, err := fmt.Fprintf(w, "\n\n      "); err != nil {
		return err
	}
//line project_edit.ego:18
	if _, err := fmt.Fprintf(w, "<div class=\"page-header\">\n        "); err != nil {
		return err
	}
//line project_edit.ego:19
	if _, err := fmt.Fprintf(w, "<h3>\n          "); err != nil {
		return err
	}
//line project_edit.ego:20
	if t.Project.ID() == 0 {
//line project_edit.ego:21
		if _, err := fmt.Fprintf(w, "\n            New Project\n          "); err != nil {
			return err
		}
//line project_edit.ego:22
	} else {
//line project_edit.ego:23
		if _, err := fmt.Fprintf(w, "\n            Edit Project\n          "); err != nil {
			return err
		}
//line project_edit.ego:24
	}
//line project_edit.ego:25
	if _, err := fmt.Fprintf(w, "\n        "); err != nil {
		return err
	}
//line project_edit.ego:25
	if _, err := fmt.Fprintf(w, "</h3>\n      "); err != nil {
		return err
	}
//line project_edit.ego:26
	if _, err := fmt.Fprintf(w, "</div>\n\n      "); err != nil {
		return err
	}
//line project_edit.ego:28
	if _, err := fmt.Fprintf(w, "<div class=\"row\">\n        "); err != nil {
		return err
	}
//line project_edit.ego:29
	if _, err := fmt.Fprintf(w, "<form action=\"/projects/"); err != nil {
		return err
	}
//line project_edit.ego:29
	if _, err := fmt.Fprintf(w, "%v", t.Project.ID()); err != nil {
		return err
	}
//line project_edit.ego:29
	if _, err := fmt.Fprintf(w, "\" method=\"POST\" role=\"form\" class=\"col-sm-6 col-md-5 col-lg-5\">\n          "); err != nil {
		return err
	}
//line project_edit.ego:30
	if _, err := fmt.Fprintf(w, "<input type=\"hidden\" name=\"id\" value=\""); err != nil {
		return err
	}
//line project_edit.ego:30
	if _, err := fmt.Fprintf(w, "%v", t.Project.ID()); err != nil {
		return err
	}
//line project_edit.ego:30
	if _, err := fmt.Fprintf(w, "\"/>\n\n          "); err != nil {
		return err
	}
//line project_edit.ego:32
	if _, err := fmt.Fprintf(w, "<div class=\"form-group\">\n            "); err != nil {
		return err
	}
//line project_edit.ego:33
	if _, err := fmt.Fprintf(w, "<label for=\"name\">Project Name"); err != nil {
		return err
	}
//line project_edit.ego:33
	if _, err := fmt.Fprintf(w, "</label>\n            "); err != nil {
		return err
	}
//line project_edit.ego:34
	if _, err := fmt.Fprintf(w, "<input type=\"text\" class=\"form-control\" id=\"name\" name=\"name\" value=\""); err != nil {
		return err
	}
//line project_edit.ego:34
	if _, err := fmt.Fprintf(w, "%v", t.Project.Name); err != nil {
		return err
	}
//line project_edit.ego:34
	if _, err := fmt.Fprintf(w, "\"/>\n          "); err != nil {
		return err
	}
//line project_edit.ego:35
	if _, err := fmt.Fprintf(w, "</div>\n\n          "); err != nil {
		return err
	}
//line project_edit.ego:37
	if t.Project.ID() != 0 {
//line project_edit.ego:38
		if _, err := fmt.Fprintf(w, "\n            "); err != nil {
			return err
		}
//line project_edit.ego:38
		if _, err := fmt.Fprintf(w, "<div class=\"form-group\">\n              "); err != nil {
			return err
		}
//line project_edit.ego:39
		if _, err := fmt.Fprintf(w, "<label>API Key"); err != nil {
			return err
		}
//line project_edit.ego:39
		if _, err := fmt.Fprintf(w, "</label>\n              "); err != nil {
			return err
		}
//line project_edit.ego:40
		if _, err := fmt.Fprintf(w, "<input type=\"text\" class=\"form-control\" value=\""); err != nil {
			return err
		}
//line project_edit.ego:40
		if _, err := fmt.Fprintf(w, "%v", t.Project.APIKey); err != nil {
			return err
		}
//line project_edit.ego:40
		if _, err := fmt.Fprintf(w, "\" disabled/>\n            "); err != nil {
			return err
		}
//line project_edit.ego:41
		if _, err := fmt.Fprintf(w, "</div>\n          "); err != nil {
			return err
		}
//line project_edit.ego:42
	}
//line project_edit.ego:43
	if _, err := fmt.Fprintf(w, "\n\n          "); err != nil {
		return err
	}
//line project_edit.ego:44
	if t.Project.ID() == 0 {
//line project_edit.ego:45
		if _, err := fmt.Fprintf(w, "\n            "); err != nil {
			return err
		}
//line project_edit.ego:45
		if _, err := fmt.Fprintf(w, "<button type=\"submit\" class=\"btn btn-primary\">Create Project"); err != nil {
			return err
		}
//line project_edit.ego:45
		if _, err := fmt.Fprintf(w, "</button>\n          "); err != nil {
			return err
		}
//line project_edit.ego:46
	} else {
//line project_edit.ego:47
		if _, err := fmt.Fprintf(w, "\n            "); err != nil {
			return err
		}
//line project_edit.ego:47
		if _, err := fmt.Fprintf(w, "<button type=\"submit\" class=\"btn btn-primary\">Save Project"); err != nil {
			return err
		}
//line project_edit.ego:47
		if _, err := fmt.Fprintf(w, "</button>\n          "); err != nil {
			return err
		}
//line project_edit.ego:48
	}
//line project_edit.ego:49
	if _, err := fmt.Fprintf(w, "\n          "); err != nil {
		return err
	}
//line project_edit.ego:49
	if _, err := fmt.Fprintf(w, "<button class=\"btn btn-link\" onclick=\"window.history.back(); return false\">Cancel"); err != nil {
		return err
	}
//line project_edit.ego:49
	if _, err := fmt.Fprintf(w, "</button>\n        "); err != nil {
		return err
	}
//line project_edit.ego:50
	if _, err := fmt.Fprintf(w, "</form>\n      "); err != nil {
		return err
	}
//line project_edit.ego:51
	if _, err := fmt.Fprintf(w, "</div>\n    "); err != nil {
		return err
	}
//line project_edit.ego:52
	if _, err := fmt.Fprintf(w, "</div> "); err != nil {
		return err
	}
//line project_edit.ego:52
	if _, err := fmt.Fprintf(w, "<!-- /container -->\n  "); err != nil {
		return err
	}
//line project_edit.ego:53
	if _, err := fmt.Fprintf(w, "</body>\n"); err != nil {
		return err
	}
//line project_edit.ego:54
	if _, err := fmt.Fprintf(w, "</html>\n\n"); err != nil {
		return err
	}
	return nil
}
