package template

//line index.ego:3
import "fmt"

//line index.ego:4
import "io"

//line index.ego:1
func (t *Template) Index(w io.Writer) error {
//line index.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line index.ego:4
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
//line index.ego:5
	if _, err := fmt.Fprintf(w, "\n\n\n"); err != nil {
		return err
	}
//line index.ego:7
	if _, err := fmt.Fprintf(w, "<!DOCTYPE html>\n"); err != nil {
		return err
	}
//line index.ego:8
	if _, err := fmt.Fprintf(w, "<html lang=\"en\">\n  "); err != nil {
		return err
	}
//line index.ego:9
	if _, err := fmt.Fprintf(w, "<head>\n    "); err != nil {
		return err
	}
//line index.ego:10
	if _, err := fmt.Fprintf(w, "<meta charset=\"utf-8\">\n    "); err != nil {
		return err
	}
//line index.ego:11
	if _, err := fmt.Fprintf(w, "<meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n    "); err != nil {
		return err
	}
//line index.ego:12
	if _, err := fmt.Fprintf(w, "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    "); err != nil {
		return err
	}
//line index.ego:13
	if _, err := fmt.Fprintf(w, "<meta name=\"description\" content=\"Open source funnel analysis\">\n\n    "); err != nil {
		return err
	}
//line index.ego:15
	if _, err := fmt.Fprintf(w, "<title>Skybox Analytics"); err != nil {
		return err
	}
//line index.ego:15
	if _, err := fmt.Fprintf(w, "</title>\n    "); err != nil {
		return err
	}
//line index.ego:16
	if _, err := fmt.Fprintf(w, "<link href=\"/assets/bootstrap.min.css\" rel=\"stylesheet\">\n    "); err != nil {
		return err
	}
//line index.ego:17
	if _, err := fmt.Fprintf(w, "<style>\n        /* Space out content a bit */\n        body {\n          padding-top: 20px;\n          padding-bottom: 20px;\n        }\n\n        .header,\n        .marketing,\n        .footer {\n          padding-right: 15px;\n          padding-left: 15px;\n        }\n\n        .header {\n          border-bottom: 1px solid #e5e5e5;\n        }\n        .header h3 {\n          padding-bottom: 19px;\n          margin-top: 0;\n          margin-bottom: 0;\n          line-height: 40px;\n        }\n\n        .footer {\n          padding-top: 19px;\n          color: #777;\n          border-top: 1px solid #e5e5e5;\n        }\n\n        @media (min-width: 768px) {\n          .container {\n            max-width: 730px;\n          }\n        }\n        .container-narrow > hr {\n          margin: 30px 0;\n        }\n\n        .jumbotron {\n          text-align: center;\n          border-bottom: 1px solid #e5e5e5;\n        }\n        .jumbotron .btn {\n          padding: 14px 24px;\n          font-size: 21px;\n        }\n\n        .marketing {\n          margin: 40px 0;\n        }\n        .marketing p + h4 {\n          margin-top: 28px;\n        }\n\n        @media screen and (min-width: 768px) {\n          .header,\n          .marketing,\n          .footer {\n            padding-right: 0;\n            padding-left: 0;\n          }\n          .header {\n            margin-bottom: 30px;\n          }\n          .jumbotron {\n            border-bottom: 0;\n          }\n        }\n    "); err != nil {
		return err
	}
//line index.ego:86
	if _, err := fmt.Fprintf(w, "</style>\n  "); err != nil {
		return err
	}
//line index.ego:87
	if _, err := fmt.Fprintf(w, "</head>\n\n  "); err != nil {
		return err
	}
//line index.ego:89
	if _, err := fmt.Fprintf(w, "<body class=\"index\">\n    "); err != nil {
		return err
	}
//line index.ego:90
	if _, err := fmt.Fprintf(w, "<div class=\"container\">\n      "); err != nil {
		return err
	}
//line index.ego:91
	if _, err := fmt.Fprintf(w, "<div class=\"header\">\n        "); err != nil {
		return err
	}
//line index.ego:92
	if _, err := fmt.Fprintf(w, "<ul class=\"nav nav-pills pull-right\">\n          "); err != nil {
		return err
	}
//line index.ego:93
	if _, err := fmt.Fprintf(w, "<li>"); err != nil {
		return err
	}
//line index.ego:93
	if _, err := fmt.Fprintf(w, "<a href=\"/login\">Log in"); err != nil {
		return err
	}
//line index.ego:93
	if _, err := fmt.Fprintf(w, "</a>"); err != nil {
		return err
	}
//line index.ego:93
	if _, err := fmt.Fprintf(w, "</li>\n          "); err != nil {
		return err
	}
//line index.ego:94
	if _, err := fmt.Fprintf(w, "<li>"); err != nil {
		return err
	}
//line index.ego:94
	if _, err := fmt.Fprintf(w, "<a href=\"/signup\">Sign up"); err != nil {
		return err
	}
//line index.ego:94
	if _, err := fmt.Fprintf(w, "</a>"); err != nil {
		return err
	}
//line index.ego:94
	if _, err := fmt.Fprintf(w, "</li>\n        "); err != nil {
		return err
	}
//line index.ego:95
	if _, err := fmt.Fprintf(w, "</ul>\n        "); err != nil {
		return err
	}
//line index.ego:96
	if _, err := fmt.Fprintf(w, "<h3 class=\"text-muted\">Skybox Analytics"); err != nil {
		return err
	}
//line index.ego:96
	if _, err := fmt.Fprintf(w, "</h3>\n      "); err != nil {
		return err
	}
//line index.ego:97
	if _, err := fmt.Fprintf(w, "</div>\n\n      "); err != nil {
		return err
	}
//line index.ego:99
	if _, err := fmt.Fprintf(w, "<div class=\"jumbotron\">\n        "); err != nil {
		return err
	}
//line index.ego:100
	if _, err := fmt.Fprintf(w, "<h1>Open Source Funnel Analysis"); err != nil {
		return err
	}
//line index.ego:100
	if _, err := fmt.Fprintf(w, "</h1>\n        "); err != nil {
		return err
	}
//line index.ego:101
	if _, err := fmt.Fprintf(w, "<p class=\"lead\">\n            Skybox is an open source funnel analysis and behavioral analytics tool.\n            Simply add a simple JavaScript tracking snippet to your site and you'll be building funnels in no time.\n        "); err != nil {
		return err
	}
//line index.ego:104
	if _, err := fmt.Fprintf(w, "</p>\n        "); err != nil {
		return err
	}
//line index.ego:105
	if _, err := fmt.Fprintf(w, "<p>\n            "); err != nil {
		return err
	}
//line index.ego:106
	if _, err := fmt.Fprintf(w, "<a class=\"btn btn-lg btn-success\" href=\"/signup\" role=\"button\">Sign up"); err != nil {
		return err
	}
//line index.ego:106
	if _, err := fmt.Fprintf(w, "</a>\n            "); err != nil {
		return err
	}
//line index.ego:107
	if _, err := fmt.Fprintf(w, "<!-- "); err != nil {
		return err
	}
//line index.ego:107
	if _, err := fmt.Fprintf(w, "<a class=\"btn btn-lg btn-primary\" href=\"/demo\" role=\"button\">Try it out"); err != nil {
		return err
	}
//line index.ego:107
	if _, err := fmt.Fprintf(w, "</a> -->\n        "); err != nil {
		return err
	}
//line index.ego:108
	if _, err := fmt.Fprintf(w, "</p>\n      "); err != nil {
		return err
	}
//line index.ego:109
	if _, err := fmt.Fprintf(w, "</div>\n\n      "); err != nil {
		return err
	}
//line index.ego:111
	if _, err := fmt.Fprintf(w, "<div class=\"row marketing\">\n        "); err != nil {
		return err
	}
//line index.ego:112
	if _, err := fmt.Fprintf(w, "<div class=\"col-lg-6\">\n          "); err != nil {
		return err
	}
//line index.ego:113
	if _, err := fmt.Fprintf(w, "<h4>Web Site Tracking"); err != nil {
		return err
	}
//line index.ego:113
	if _, err := fmt.Fprintf(w, "</h4>\n          "); err != nil {
		return err
	}
//line index.ego:114
	if _, err := fmt.Fprintf(w, "<p>\n            After you sign up you'll get a JavaScript snippet to paste on your site.\n            That's all the installation required.\n          "); err != nil {
		return err
	}
//line index.ego:117
	if _, err := fmt.Fprintf(w, "</p>\n\n          "); err != nil {
		return err
	}
//line index.ego:119
	if _, err := fmt.Fprintf(w, "<h4>Open Source"); err != nil {
		return err
	}
//line index.ego:119
	if _, err := fmt.Fprintf(w, "</h4>\n          "); err != nil {
		return err
	}
//line index.ego:120
	if _, err := fmt.Fprintf(w, "<p>\n            Have questions? Want a new feature? Come find us on our "); err != nil {
		return err
	}
//line index.ego:121
	if _, err := fmt.Fprintf(w, "<a href=\"https://github.com/skybox/skybox\" target=\"_blank\">Github"); err != nil {
		return err
	}
//line index.ego:121
	if _, err := fmt.Fprintf(w, "</a> project page!\n            Please add a Github issue if you experience any problems.\n          "); err != nil {
		return err
	}
//line index.ego:123
	if _, err := fmt.Fprintf(w, "</p>\n        "); err != nil {
		return err
	}
//line index.ego:124
	if _, err := fmt.Fprintf(w, "</div>\n\n        "); err != nil {
		return err
	}
//line index.ego:126
	if _, err := fmt.Fprintf(w, "<div class=\"col-lg-6\">\n          "); err != nil {
		return err
	}
//line index.ego:127
	if _, err := fmt.Fprintf(w, "<h4>Real-time Funnels"); err != nil {
		return err
	}
//line index.ego:127
	if _, err := fmt.Fprintf(w, "</h4>\n          "); err != nil {
		return err
	}
//line index.ego:128
	if _, err := fmt.Fprintf(w, "<p>\n            Skybox is backed by the powerful "); err != nil {
		return err
	}
//line index.ego:129
	if _, err := fmt.Fprintf(w, "<a href=\"https://github.com/skydb/sky\" target=\"_blank\">SkyDB"); err != nil {
		return err
	}
//line index.ego:129
	if _, err := fmt.Fprintf(w, "</a> behavioral analytics database.\n            That means that all your data is available immediately.\n          "); err != nil {
		return err
	}
//line index.ego:131
	if _, err := fmt.Fprintf(w, "</p>\n\n          "); err != nil {
		return err
	}
//line index.ego:133
	if _, err := fmt.Fprintf(w, "<h4>APIs and More"); err != nil {
		return err
	}
//line index.ego:133
	if _, err := fmt.Fprintf(w, "</h4>\n          "); err != nil {
		return err
	}
//line index.ego:134
	if _, err := fmt.Fprintf(w, "<p>\n            It's not quite ready yet but we'll be releasing an API so you can track user events from anywhere and pull Skybox data into your own application.\n          "); err != nil {
		return err
	}
//line index.ego:136
	if _, err := fmt.Fprintf(w, "</p>\n        "); err != nil {
		return err
	}
//line index.ego:137
	if _, err := fmt.Fprintf(w, "</div>\n      "); err != nil {
		return err
	}
//line index.ego:138
	if _, err := fmt.Fprintf(w, "</div>\n    "); err != nil {
		return err
	}
//line index.ego:139
	if _, err := fmt.Fprintf(w, "</div> "); err != nil {
		return err
	}
//line index.ego:139
	if _, err := fmt.Fprintf(w, "<!-- /container -->\n  "); err != nil {
		return err
	}
//line index.ego:140
	if _, err := fmt.Fprintf(w, "</body>\n"); err != nil {
		return err
	}
//line index.ego:141
	if _, err := fmt.Fprintf(w, "</html>\n"); err != nil {
		return err
	}
	return nil
}
