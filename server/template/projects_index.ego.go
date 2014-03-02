package template

//line projects_index.ego:3
import (
	"fmt"
	"io"
)

//line projects_index.ego:1
func (t *ProjectsTemplate) Index(w io.Writer) error {
//line projects_index.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line projects_index.ego:9
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line projects_index.ego:10
	if _, err := fmt.Fprintf(w, "<!DOCTYPE html>\n"); err != nil {
		return err
	}
//line projects_index.ego:11
	if _, err := fmt.Fprintf(w, "<html lang=\"en\">\n  "); err != nil {
		return err
	}
//line projects_index.ego:12
	t.Head(w, "")
//line projects_index.ego:13
	if _, err := fmt.Fprintf(w, "\n\n  "); err != nil {
		return err
	}
//line projects_index.ego:14
	if _, err := fmt.Fprintf(w, "<body id=\"index\">\n    "); err != nil {
		return err
	}
//line projects_index.ego:15
	if _, err := fmt.Fprintf(w, "<div class=\"container\">\n      "); err != nil {
		return err
	}
//line projects_index.ego:16
	t.Nav(w)
//line projects_index.ego:17
	if _, err := fmt.Fprintf(w, "\n\n      "); err != nil {
		return err
	}
//line projects_index.ego:18
	if _, err := fmt.Fprintf(w, "<div class=\"page-header\">\n        "); err != nil {
		return err
	}
//line projects_index.ego:19
	if _, err := fmt.Fprintf(w, "<h3>\n          Projects\n          "); err != nil {
		return err
	}
//line projects_index.ego:21
	if _, err := fmt.Fprintf(w, "<div class=\"pull-right\">\n            "); err != nil {
		return err
	}
//line projects_index.ego:22
	if _, err := fmt.Fprintf(w, "<a href=\"/projects/new\" class=\"btn btn-success\">New Project"); err != nil {
		return err
	}
//line projects_index.ego:22
	if _, err := fmt.Fprintf(w, "</a>\n          "); err != nil {
		return err
	}
//line projects_index.ego:23
	if _, err := fmt.Fprintf(w, "</div>\n        "); err != nil {
		return err
	}
//line projects_index.ego:24
	if _, err := fmt.Fprintf(w, "</h3>\n      "); err != nil {
		return err
	}
//line projects_index.ego:25
	if _, err := fmt.Fprintf(w, "</div>\n\n      "); err != nil {
		return err
	}
//line projects_index.ego:27
	if len(t.Projects) == 0 {
//line projects_index.ego:28
		if _, err := fmt.Fprintf(w, "\n        "); err != nil {
			return err
		}
//line projects_index.ego:28
		if _, err := fmt.Fprintf(w, "<div class=\"row\">\n          "); err != nil {
			return err
		}
//line projects_index.ego:29
		if _, err := fmt.Fprintf(w, "<div class=\"col-lg-12\">\n            "); err != nil {
			return err
		}
//line projects_index.ego:30
		if _, err := fmt.Fprintf(w, "<p>You do not have any projects on your account."); err != nil {
			return err
		}
//line projects_index.ego:30
		if _, err := fmt.Fprintf(w, "</p>\n          "); err != nil {
			return err
		}
//line projects_index.ego:31
		if _, err := fmt.Fprintf(w, "</div>\n        "); err != nil {
			return err
		}
//line projects_index.ego:32
		if _, err := fmt.Fprintf(w, "</div>\n      "); err != nil {
			return err
		}
//line projects_index.ego:33
	} else {
//line projects_index.ego:34
		if _, err := fmt.Fprintf(w, "\n        "); err != nil {
			return err
		}
//line projects_index.ego:34
		if _, err := fmt.Fprintf(w, "<table class=\"table\">\n          "); err != nil {
			return err
		}
//line projects_index.ego:35
		if _, err := fmt.Fprintf(w, "<thead>\n            "); err != nil {
			return err
		}
//line projects_index.ego:36
		if _, err := fmt.Fprintf(w, "<tr>\n              "); err != nil {
			return err
		}
//line projects_index.ego:37
		if _, err := fmt.Fprintf(w, "<th>Project Name"); err != nil {
			return err
		}
//line projects_index.ego:37
		if _, err := fmt.Fprintf(w, "</th>\n              "); err != nil {
			return err
		}
//line projects_index.ego:38
		if _, err := fmt.Fprintf(w, "<th>API Key"); err != nil {
			return err
		}
//line projects_index.ego:38
		if _, err := fmt.Fprintf(w, "</th>\n            "); err != nil {
			return err
		}
//line projects_index.ego:39
		if _, err := fmt.Fprintf(w, "</tr>\n          "); err != nil {
			return err
		}
//line projects_index.ego:40
		if _, err := fmt.Fprintf(w, "</thead>\n          "); err != nil {
			return err
		}
//line projects_index.ego:41
		if _, err := fmt.Fprintf(w, "<tbody>\n            "); err != nil {
			return err
		}
//line projects_index.ego:42
		for _, p := range t.Projects {
//line projects_index.ego:43
			if _, err := fmt.Fprintf(w, "\n              "); err != nil {
				return err
			}
//line projects_index.ego:43
			if _, err := fmt.Fprintf(w, "<tr>\n                "); err != nil {
				return err
			}
//line projects_index.ego:44
			if _, err := fmt.Fprintf(w, "<td>\n                  "); err != nil {
				return err
			}
//line projects_index.ego:45
			if _, err := fmt.Fprintf(w, "<a href=\"/projects/"); err != nil {
				return err
			}
//line projects_index.ego:45
			if _, err := fmt.Fprintf(w, "%v", p.ID()); err != nil {
				return err
			}
//line projects_index.ego:45
			if _, err := fmt.Fprintf(w, "\">"); err != nil {
				return err
			}
//line projects_index.ego:45
			if _, err := fmt.Fprintf(w, "%v", p.Name); err != nil {
				return err
			}
//line projects_index.ego:45
			if _, err := fmt.Fprintf(w, "</a>\n                "); err != nil {
				return err
			}
//line projects_index.ego:46
			if _, err := fmt.Fprintf(w, "</td>\n                "); err != nil {
				return err
			}
//line projects_index.ego:47
			if _, err := fmt.Fprintf(w, "<td>"); err != nil {
				return err
			}
//line projects_index.ego:47
			if _, err := fmt.Fprintf(w, "%v", p.APIKey); err != nil {
				return err
			}
//line projects_index.ego:47
			if _, err := fmt.Fprintf(w, "</td>\n              "); err != nil {
				return err
			}
//line projects_index.ego:48
			if _, err := fmt.Fprintf(w, "</tr>\n            "); err != nil {
				return err
			}
//line projects_index.ego:49
		}
//line projects_index.ego:50
		if _, err := fmt.Fprintf(w, "\n          "); err != nil {
			return err
		}
//line projects_index.ego:50
		if _, err := fmt.Fprintf(w, "</tbody>\n        "); err != nil {
			return err
		}
//line projects_index.ego:51
		if _, err := fmt.Fprintf(w, "</table>\n      "); err != nil {
			return err
		}
//line projects_index.ego:52
	}
//line projects_index.ego:53
	if _, err := fmt.Fprintf(w, "\n    "); err != nil {
		return err
	}
//line projects_index.ego:53
	if _, err := fmt.Fprintf(w, "</div> "); err != nil {
		return err
	}
//line projects_index.ego:53
	if _, err := fmt.Fprintf(w, "<!-- /container -->\n  "); err != nil {
		return err
	}
//line projects_index.ego:54
	if _, err := fmt.Fprintf(w, "</body>\n"); err != nil {
		return err
	}
//line projects_index.ego:55
	if _, err := fmt.Fprintf(w, "</html>\n\n"); err != nil {
		return err
	}
	return nil
}
