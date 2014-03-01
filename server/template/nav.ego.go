package template

//line nav.ego:2
import (
	"fmt"
	"io"
)

//line nav.ego:1
func (t *Template) Nav(w io.Writer) error {
//line nav.ego:2
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
//line nav.ego:8
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line nav.ego:9
	if _, err := fmt.Fprintf(w, "<!-- Navbar -->\n"); err != nil {
		return err
	}
//line nav.ego:10
	if _, err := fmt.Fprintf(w, "<div class=\"navbar\" role=\"navigation\">\n  "); err != nil {
		return err
	}
//line nav.ego:11
	if _, err := fmt.Fprintf(w, "<div class=\"container-fluid\">\n    "); err != nil {
		return err
	}
//line nav.ego:12
	if _, err := fmt.Fprintf(w, "<div class=\"navbar-header\">\n      "); err != nil {
		return err
	}
//line nav.ego:13
	if _, err := fmt.Fprintf(w, "<button type=\"button\" class=\"navbar-toggle\" data-toggle=\"collapse\" data-target=\".navbar-collapse\">\n        "); err != nil {
		return err
	}
//line nav.ego:14
	if _, err := fmt.Fprintf(w, "<span class=\"sr-only\">Toggle navigation"); err != nil {
		return err
	}
//line nav.ego:14
	if _, err := fmt.Fprintf(w, "</span>\n        "); err != nil {
		return err
	}
//line nav.ego:15
	if _, err := fmt.Fprintf(w, "<span class=\"icon-bar\">"); err != nil {
		return err
	}
//line nav.ego:15
	if _, err := fmt.Fprintf(w, "</span>\n        "); err != nil {
		return err
	}
//line nav.ego:16
	if _, err := fmt.Fprintf(w, "<span class=\"icon-bar\">"); err != nil {
		return err
	}
//line nav.ego:16
	if _, err := fmt.Fprintf(w, "</span>\n        "); err != nil {
		return err
	}
//line nav.ego:17
	if _, err := fmt.Fprintf(w, "<span class=\"icon-bar\">"); err != nil {
		return err
	}
//line nav.ego:17
	if _, err := fmt.Fprintf(w, "</span>\n      "); err != nil {
		return err
	}
//line nav.ego:18
	if _, err := fmt.Fprintf(w, "</button>\n      "); err != nil {
		return err
	}
//line nav.ego:19
	if _, err := fmt.Fprintf(w, "<a class=\"navbar-brand\" href=\"/\">Skybox"); err != nil {
		return err
	}
//line nav.ego:19
	if _, err := fmt.Fprintf(w, "</a>\n    "); err != nil {
		return err
	}
//line nav.ego:20
	if _, err := fmt.Fprintf(w, "</div>\n    "); err != nil {
		return err
	}
//line nav.ego:21
	if _, err := fmt.Fprintf(w, "<div class=\"navbar-collapse collapse\">\n      "); err != nil {
		return err
	}
//line nav.ego:22
	if t.User != nil {
//line nav.ego:23
		if _, err := fmt.Fprintf(w, "\n        "); err != nil {
			return err
		}
//line nav.ego:23
		if _, err := fmt.Fprintf(w, "<ul class=\"nav navbar-nav\">\n          "); err != nil {
			return err
		}
//line nav.ego:24
		if _, err := fmt.Fprintf(w, "<li class=\"active\">"); err != nil {
			return err
		}
//line nav.ego:24
		if _, err := fmt.Fprintf(w, "<a href=\"#\">Dashboard"); err != nil {
			return err
		}
//line nav.ego:24
		if _, err := fmt.Fprintf(w, "</a>"); err != nil {
			return err
		}
//line nav.ego:24
		if _, err := fmt.Fprintf(w, "</li>\n          "); err != nil {
			return err
		}
//line nav.ego:25
		if _, err := fmt.Fprintf(w, "<li>"); err != nil {
			return err
		}
//line nav.ego:25
		if _, err := fmt.Fprintf(w, "<a href=\"#\">Projects"); err != nil {
			return err
		}
//line nav.ego:25
		if _, err := fmt.Fprintf(w, "</a>"); err != nil {
			return err
		}
//line nav.ego:25
		if _, err := fmt.Fprintf(w, "</li>\n          "); err != nil {
			return err
		}
//line nav.ego:26
		if _, err := fmt.Fprintf(w, "<li>"); err != nil {
			return err
		}
//line nav.ego:26
		if _, err := fmt.Fprintf(w, "<a href=\"#\">Users"); err != nil {
			return err
		}
//line nav.ego:26
		if _, err := fmt.Fprintf(w, "</a>"); err != nil {
			return err
		}
//line nav.ego:26
		if _, err := fmt.Fprintf(w, "</li>\n          "); err != nil {
			return err
		}
//line nav.ego:27
		if _, err := fmt.Fprintf(w, "<li>"); err != nil {
			return err
		}
//line nav.ego:27
		if _, err := fmt.Fprintf(w, "<a href=\"#\">Account"); err != nil {
			return err
		}
//line nav.ego:27
		if _, err := fmt.Fprintf(w, "</a>"); err != nil {
			return err
		}
//line nav.ego:27
		if _, err := fmt.Fprintf(w, "</li>\n        "); err != nil {
			return err
		}
//line nav.ego:28
		if _, err := fmt.Fprintf(w, "</ul>\n\n        "); err != nil {
			return err
		}
//line nav.ego:30
		if _, err := fmt.Fprintf(w, "<ul class=\"nav navbar-nav navbar-right\">\n          "); err != nil {
			return err
		}
//line nav.ego:31
		if _, err := fmt.Fprintf(w, "<li>"); err != nil {
			return err
		}
//line nav.ego:31
		if _, err := fmt.Fprintf(w, "<a href=\"/logout\">Log out"); err != nil {
			return err
		}
//line nav.ego:31
		if _, err := fmt.Fprintf(w, "</a>"); err != nil {
			return err
		}
//line nav.ego:31
		if _, err := fmt.Fprintf(w, "</li>\n        "); err != nil {
			return err
		}
//line nav.ego:32
		if _, err := fmt.Fprintf(w, "</ul>\n      "); err != nil {
			return err
		}
//line nav.ego:33
	} else {
//line nav.ego:34
		if _, err := fmt.Fprintf(w, "\n        "); err != nil {
			return err
		}
//line nav.ego:34
		if _, err := fmt.Fprintf(w, "<ul class=\"nav navbar-nav navbar-right\">\n          "); err != nil {
			return err
		}
//line nav.ego:35
		if _, err := fmt.Fprintf(w, "<li>"); err != nil {
			return err
		}
//line nav.ego:35
		if _, err := fmt.Fprintf(w, "<a href=\"/login\">Sign in"); err != nil {
			return err
		}
//line nav.ego:35
		if _, err := fmt.Fprintf(w, "</a>"); err != nil {
			return err
		}
//line nav.ego:35
		if _, err := fmt.Fprintf(w, "</li>\n          "); err != nil {
			return err
		}
//line nav.ego:36
		if _, err := fmt.Fprintf(w, "<li>"); err != nil {
			return err
		}
//line nav.ego:36
		if _, err := fmt.Fprintf(w, "<a href=\"/signup\">Sign up"); err != nil {
			return err
		}
//line nav.ego:36
		if _, err := fmt.Fprintf(w, "</a>"); err != nil {
			return err
		}
//line nav.ego:36
		if _, err := fmt.Fprintf(w, "</li>\n        "); err != nil {
			return err
		}
//line nav.ego:37
		if _, err := fmt.Fprintf(w, "</ul>\n      "); err != nil {
			return err
		}
//line nav.ego:38
	}
//line nav.ego:39
	if _, err := fmt.Fprintf(w, "\n    "); err != nil {
		return err
	}
//line nav.ego:39
	if _, err := fmt.Fprintf(w, "</div>"); err != nil {
		return err
	}
//line nav.ego:39
	if _, err := fmt.Fprintf(w, "<!--/.nav-collapse -->\n  "); err != nil {
		return err
	}
//line nav.ego:40
	if _, err := fmt.Fprintf(w, "</div>"); err != nil {
		return err
	}
//line nav.ego:40
	if _, err := fmt.Fprintf(w, "<!--/.container-fluid -->\n"); err != nil {
		return err
	}
//line nav.ego:41
	if _, err := fmt.Fprintf(w, "</div>\n"); err != nil {
		return err
	}
	return nil
}
