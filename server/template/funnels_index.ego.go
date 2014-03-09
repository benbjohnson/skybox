package template

//line funnels_index.ego:3
import "fmt"

//line funnels_index.ego:4
import "io"

//line funnels_index.ego:1
func (t *FunnelsTemplate) Index(w io.Writer) error {
//line funnels_index.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line funnels_index.ego:4
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
//line funnels_index.ego:5
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line funnels_index.ego:6
	if _, err := fmt.Fprintf(w, "<!DOCTYPE html>\n"); err != nil {
		return err
	}
//line funnels_index.ego:7
	if _, err := fmt.Fprintf(w, "<html lang=\"en\">\n  "); err != nil {
		return err
	}
//line funnels_index.ego:8
	t.Head(w, "")
//line funnels_index.ego:9
	if _, err := fmt.Fprintf(w, "\n\n  "); err != nil {
		return err
	}
//line funnels_index.ego:10
	if _, err := fmt.Fprintf(w, "<body id=\"index\">\n    "); err != nil {
		return err
	}
//line funnels_index.ego:11
	if _, err := fmt.Fprintf(w, "<div class=\"container\">\n      "); err != nil {
		return err
	}
//line funnels_index.ego:12
	t.Nav(w)
//line funnels_index.ego:13
	if _, err := fmt.Fprintf(w, "\n\n      "); err != nil {
		return err
	}
//line funnels_index.ego:14
	if _, err := fmt.Fprintf(w, "<div class=\"page-header\">\n        "); err != nil {
		return err
	}
//line funnels_index.ego:15
	if _, err := fmt.Fprintf(w, "<h3>\n          Funnels\n          "); err != nil {
		return err
	}
//line funnels_index.ego:17
	if _, err := fmt.Fprintf(w, "<div class=\"pull-right\">\n            "); err != nil {
		return err
	}
//line funnels_index.ego:18
	if _, err := fmt.Fprintf(w, "<a href=\"/funnels/0/edit\" class=\"btn btn-success\">New Funnel"); err != nil {
		return err
	}
//line funnels_index.ego:18
	if _, err := fmt.Fprintf(w, "</a>\n          "); err != nil {
		return err
	}
//line funnels_index.ego:19
	if _, err := fmt.Fprintf(w, "</div>\n        "); err != nil {
		return err
	}
//line funnels_index.ego:20
	if _, err := fmt.Fprintf(w, "</h3>\n      "); err != nil {
		return err
	}
//line funnels_index.ego:21
	if _, err := fmt.Fprintf(w, "</div>\n\n      "); err != nil {
		return err
	}
//line funnels_index.ego:23
	if len(t.Funnels) == 0 {
//line funnels_index.ego:24
		if _, err := fmt.Fprintf(w, "\n        "); err != nil {
			return err
		}
//line funnels_index.ego:24
		if _, err := fmt.Fprintf(w, "<div class=\"row\">\n          "); err != nil {
			return err
		}
//line funnels_index.ego:25
		if _, err := fmt.Fprintf(w, "<div class=\"col-lg-12\">\n            "); err != nil {
			return err
		}
//line funnels_index.ego:26
		if _, err := fmt.Fprintf(w, "<p>You do not have any funnels on your account."); err != nil {
			return err
		}
//line funnels_index.ego:26
		if _, err := fmt.Fprintf(w, "</p>\n          "); err != nil {
			return err
		}
//line funnels_index.ego:27
		if _, err := fmt.Fprintf(w, "</div>\n        "); err != nil {
			return err
		}
//line funnels_index.ego:28
		if _, err := fmt.Fprintf(w, "</div>\n      "); err != nil {
			return err
		}
//line funnels_index.ego:29
	} else {
//line funnels_index.ego:30
		if _, err := fmt.Fprintf(w, "\n        "); err != nil {
			return err
		}
//line funnels_index.ego:30
		if _, err := fmt.Fprintf(w, "<table class=\"table\">\n          "); err != nil {
			return err
		}
//line funnels_index.ego:31
		if _, err := fmt.Fprintf(w, "<thead>\n            "); err != nil {
			return err
		}
//line funnels_index.ego:32
		if _, err := fmt.Fprintf(w, "<tr>\n              "); err != nil {
			return err
		}
//line funnels_index.ego:33
		if _, err := fmt.Fprintf(w, "<th>Funnel Name"); err != nil {
			return err
		}
//line funnels_index.ego:33
		if _, err := fmt.Fprintf(w, "</th>\n            "); err != nil {
			return err
		}
//line funnels_index.ego:34
		if _, err := fmt.Fprintf(w, "</tr>\n          "); err != nil {
			return err
		}
//line funnels_index.ego:35
		if _, err := fmt.Fprintf(w, "</thead>\n          "); err != nil {
			return err
		}
//line funnels_index.ego:36
		if _, err := fmt.Fprintf(w, "<tbody>\n            "); err != nil {
			return err
		}
//line funnels_index.ego:37
		for _, f := range t.Funnels {
//line funnels_index.ego:38
			if _, err := fmt.Fprintf(w, "\n              "); err != nil {
				return err
			}
//line funnels_index.ego:38
			if _, err := fmt.Fprintf(w, "<tr>\n                "); err != nil {
				return err
			}
//line funnels_index.ego:39
			if _, err := fmt.Fprintf(w, "<td>\n                  "); err != nil {
				return err
			}
//line funnels_index.ego:40
			if _, err := fmt.Fprintf(w, "<a href=\"/funnels/"); err != nil {
				return err
			}
//line funnels_index.ego:40
			if _, err := fmt.Fprintf(w, "%v", f.ID()); err != nil {
				return err
			}
//line funnels_index.ego:40
			if _, err := fmt.Fprintf(w, "\">"); err != nil {
				return err
			}
//line funnels_index.ego:40
			if _, err := fmt.Fprintf(w, "%v", f.Name); err != nil {
				return err
			}
//line funnels_index.ego:40
			if _, err := fmt.Fprintf(w, "</a>\n                "); err != nil {
				return err
			}
//line funnels_index.ego:41
			if _, err := fmt.Fprintf(w, "</td>\n              "); err != nil {
				return err
			}
//line funnels_index.ego:42
			if _, err := fmt.Fprintf(w, "</tr>\n            "); err != nil {
				return err
			}
//line funnels_index.ego:43
		}
//line funnels_index.ego:44
		if _, err := fmt.Fprintf(w, "\n          "); err != nil {
			return err
		}
//line funnels_index.ego:44
		if _, err := fmt.Fprintf(w, "</tbody>\n        "); err != nil {
			return err
		}
//line funnels_index.ego:45
		if _, err := fmt.Fprintf(w, "</table>\n      "); err != nil {
			return err
		}
//line funnels_index.ego:46
	}
//line funnels_index.ego:47
	if _, err := fmt.Fprintf(w, "\n    "); err != nil {
		return err
	}
//line funnels_index.ego:47
	if _, err := fmt.Fprintf(w, "</div> "); err != nil {
		return err
	}
//line funnels_index.ego:47
	if _, err := fmt.Fprintf(w, "<!-- /container -->\n  "); err != nil {
		return err
	}
//line funnels_index.ego:48
	if _, err := fmt.Fprintf(w, "</body>\n"); err != nil {
		return err
	}
//line funnels_index.ego:49
	if _, err := fmt.Fprintf(w, "</html>\n"); err != nil {
		return err
	}
	return nil
}
