package template

//line funnel_show.ego:3
import "encoding/json"

//line funnel_show.ego:4
import "fmt"

//line funnel_show.ego:5
import "io"

//line funnel_show.ego:6
import "regexp"

//line funnel_show.ego:1
func (t *FunnelTemplate) Show(w io.Writer) error {
//line funnel_show.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line funnel_show.ego:4
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
//line funnel_show.ego:5
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
//line funnel_show.ego:6
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
//line funnel_show.ego:7
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line funnel_show.ego:8
	if _, err := fmt.Fprintf(w, "<!DOCTYPE html>\n"); err != nil {
		return err
	}
//line funnel_show.ego:9
	if _, err := fmt.Fprintf(w, "<html lang=\"en\">\n  "); err != nil {
		return err
	}
//line funnel_show.ego:10
	t.Head(w, "")
//line funnel_show.ego:11
	if _, err := fmt.Fprintf(w, "\n\n  "); err != nil {
		return err
	}
//line funnel_show.ego:12
	if _, err := fmt.Fprintf(w, "<body id=\"index\">\n    "); err != nil {
		return err
	}
//line funnel_show.ego:13
	if _, err := fmt.Fprintf(w, "<div class=\"container\">\n      "); err != nil {
		return err
	}
//line funnel_show.ego:14
	t.Nav(w)
//line funnel_show.ego:15
	if _, err := fmt.Fprintf(w, "\n\n      "); err != nil {
		return err
	}
//line funnel_show.ego:16
	if _, err := fmt.Fprintf(w, "<div class=\"page-header\">\n        "); err != nil {
		return err
	}
//line funnel_show.ego:17
	if _, err := fmt.Fprintf(w, "<h3>\n          "); err != nil {
		return err
	}
//line funnel_show.ego:18
	if _, err := fmt.Fprintf(w, "%v", t.Funnel.Name); err != nil {
		return err
	}
//line funnel_show.ego:19
	if _, err := fmt.Fprintf(w, "\n          "); err != nil {
		return err
	}
//line funnel_show.ego:19
	if _, err := fmt.Fprintf(w, "<div class=\"pull-right\">\n            "); err != nil {
		return err
	}
//line funnel_show.ego:20
	if _, err := fmt.Fprintf(w, "<a href=\"/funnels\" class=\"btn btn-default\">Back"); err != nil {
		return err
	}
//line funnel_show.ego:20
	if _, err := fmt.Fprintf(w, "</a>\n          "); err != nil {
		return err
	}
//line funnel_show.ego:21
	if _, err := fmt.Fprintf(w, "</div>\n        "); err != nil {
		return err
	}
//line funnel_show.ego:22
	if _, err := fmt.Fprintf(w, "</h3>\n      "); err != nil {
		return err
	}
//line funnel_show.ego:23
	if _, err := fmt.Fprintf(w, "</div>\n\n      "); err != nil {
		return err
	}
//line funnel_show.ego:25
	if _, err := fmt.Fprintf(w, "<div class=\"row\">\n        "); err != nil {
		return err
	}
//line funnel_show.ego:26
	if _, err := fmt.Fprintf(w, "<div class=\"chart\">"); err != nil {
		return err
	}
//line funnel_show.ego:26
	if _, err := fmt.Fprintf(w, "</div>\n      "); err != nil {
		return err
	}
//line funnel_show.ego:27
	if _, err := fmt.Fprintf(w, "</div>\n\n      "); err != nil {
		return err
	}
//line funnel_show.ego:29
	if _, err := fmt.Fprintf(w, "<table class=\"table\">\n        "); err != nil {
		return err
	}
//line funnel_show.ego:30
	if _, err := fmt.Fprintf(w, "<thead>\n          "); err != nil {
		return err
	}
//line funnel_show.ego:31
	if _, err := fmt.Fprintf(w, "<tr>\n            "); err != nil {
		return err
	}
//line funnel_show.ego:32
	if _, err := fmt.Fprintf(w, "<th>Step"); err != nil {
		return err
	}
//line funnel_show.ego:32
	if _, err := fmt.Fprintf(w, "</th>\n            "); err != nil {
		return err
	}
//line funnel_show.ego:33
	if _, err := fmt.Fprintf(w, "<th>Count"); err != nil {
		return err
	}
//line funnel_show.ego:33
	if _, err := fmt.Fprintf(w, "</th>\n          "); err != nil {
		return err
	}
//line funnel_show.ego:34
	if _, err := fmt.Fprintf(w, "</tr>\n        "); err != nil {
		return err
	}
//line funnel_show.ego:35
	if _, err := fmt.Fprintf(w, "</thead>\n        "); err != nil {
		return err
	}
//line funnel_show.ego:36
	if _, err := fmt.Fprintf(w, "<tbody>\n          "); err != nil {
		return err
	}
//line funnel_show.ego:37
	for _, f := range t.FunnelResult.Steps {
//line funnel_show.ego:38
		if _, err := fmt.Fprintf(w, "\n            "); err != nil {
			return err
		}
//line funnel_show.ego:38
		if _, err := fmt.Fprintf(w, "<tr>\n              "); err != nil {
			return err
		}
//line funnel_show.ego:39
		if _, err := fmt.Fprintf(w, "<td>\n                "); err != nil {
			return err
		}
//line funnel_show.ego:40
		if _, err := fmt.Fprintf(w, "%v", regexp.MustCompile(`^resource == "(.+)"$`).ReplaceAllString(f.Condition, "$1")); err != nil {
			return err
		}
//line funnel_show.ego:41
		if _, err := fmt.Fprintf(w, "\n              "); err != nil {
			return err
		}
//line funnel_show.ego:41
		if _, err := fmt.Fprintf(w, "</td>\n              "); err != nil {
			return err
		}
//line funnel_show.ego:42
		if _, err := fmt.Fprintf(w, "<td>\n                "); err != nil {
			return err
		}
//line funnel_show.ego:43
		if _, err := fmt.Fprintf(w, "%v", f.Count); err != nil {
			return err
		}
//line funnel_show.ego:44
		if _, err := fmt.Fprintf(w, "\n              "); err != nil {
			return err
		}
//line funnel_show.ego:44
		if _, err := fmt.Fprintf(w, "</td>\n            "); err != nil {
			return err
		}
//line funnel_show.ego:45
		if _, err := fmt.Fprintf(w, "</tr>\n          "); err != nil {
			return err
		}
//line funnel_show.ego:46
	}
//line funnel_show.ego:47
	if _, err := fmt.Fprintf(w, "\n        "); err != nil {
		return err
	}
//line funnel_show.ego:47
	if _, err := fmt.Fprintf(w, "</tbody>\n      "); err != nil {
		return err
	}
//line funnel_show.ego:48
	if _, err := fmt.Fprintf(w, "</div>\n    "); err != nil {
		return err
	}
//line funnel_show.ego:49
	if _, err := fmt.Fprintf(w, "</div> "); err != nil {
		return err
	}
//line funnel_show.ego:49
	if _, err := fmt.Fprintf(w, "<!-- /container -->\n  "); err != nil {
		return err
	}
//line funnel_show.ego:50
	if _, err := fmt.Fprintf(w, "</body>\n\n  "); err != nil {
		return err
	}
//line funnel_show.ego:52
	if _, err := fmt.Fprintf(w, "<script src=\"/assets/funnel_show.js\">"); err != nil {
		return err
	}
//line funnel_show.ego:52
	if _, err := fmt.Fprintf(w, "</script>\n  "); err != nil {
		return err
	}
//line funnel_show.ego:53
	if _, err := fmt.Fprintf(w, "<script>\n    var result = "); err != nil {
		return err
	}
//line funnel_show.ego:54
	json.NewEncoder(w).Encode(t.FunnelResult)
//line funnel_show.ego:54
	if _, err := fmt.Fprintf(w, ";\n    update(result);\n  "); err != nil {
		return err
	}
//line funnel_show.ego:56
	if _, err := fmt.Fprintf(w, "</script>\n"); err != nil {
		return err
	}
//line funnel_show.ego:57
	if _, err := fmt.Fprintf(w, "</html>\n\n"); err != nil {
		return err
	}
	return nil
}
