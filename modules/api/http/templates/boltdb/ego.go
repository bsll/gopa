// Generated by ego.
// DO NOT EDIT

package templates

import (
	"fmt"
	"github.com/boltdb/bolt"
	"html"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"unsafe"
)

var _ = fmt.Sprint("") // just so that we can keep the fmt import for now
//line error.ego:1
func Error(w io.Writer, err error) error {
//line error.ego:2
	_, _ = io.WriteString(w, "\n\n<!DOCTYPE html>\n<html lang=\"en\">\n  <head>\n    <meta charset=\"utf-8\">\n    <title>boltd</title>\n  </head>\n\n  <body class=\"error\">\n    <div class=\"container\">\n      <div class=\"header\">\n        <h3 class=\"text-muted\">Error</h3>\n      </div>\n\n      An error has occurred: ")
//line error.ego:16
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(err)))
//line error.ego:17
	_, _ = io.WriteString(w, "\n    </div> <!-- /container -->\n  </body>\n</html>\n")
	return nil
}

//line head.ego:1
func head(w io.Writer, tx *bolt.Tx) error {
//line head.ego:2
	_, _ = io.WriteString(w, "\n\n")
//line head.ego:4
	_, _ = io.WriteString(w, "\n")
//line head.ego:5
	_, _ = io.WriteString(w, "\n\n<head>\n  <meta charset=\"utf-8\">\n  <title>")
//line head.ego:8
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(filepath.Base(tx.DB().Path()))))
//line head.ego:8
	_, _ = io.WriteString(w, " - GOPA</title>\n  <link rel=\"stylesheet\" href=\"/ui/assets/uikit-2.27.1/css/uikit.min.css\" />\n  <link rel=\"stylesheet\" href=\"/ui/assets/css/tasks.css\" />\n  <script src=\"/ui/assets/js/jquery.min.js\"></script>\n  <script src=\"/ui/assets/js/echarts.min.js\"></script>\n  <script src=\"/ui/assets/uikit-2.27.1/js/uikit.min.js\"></script>\n  <script src=\"/ui/assets/js/jquery.timeago.js\"></script>\n  <script src=\"/ui/assets/js/page/tasks.js\"></script>\n  <style>\n    table {\n      border-collapse:collapse;\n      word-break:break-all; word-wrap:break-word;\n    }\n    \n    table, th, td {\n      border: 1px solid black;\n    }\n\n    th, td { \n      min-width: 100px;\n      padding: 2px 5px;\n    }\n  </style>\n</head>\n\n")
	return nil
}

//line index.ego:1
func Index(w io.Writer) error {
//line index.ego:2
	_, _ = io.WriteString(w, "\n\n<!DOCTYPE html>\n<html lang=\"en\">\n  <head>\n    <meta http-equiv=\"refresh\" content=\"0; url=page\">\n  </head>\n\n  <body>redirecting...</body>\n</html>\n")
	return nil
}

//line nav.ego:1
func nav(w io.Writer, tx *bolt.Tx) error {
//line nav.ego:2
	_, _ = io.WriteString(w, "\n\n")
//line nav.ego:4
	_, _ = io.WriteString(w, "\n")
//line nav.ego:5
	_, _ = io.WriteString(w, "\n\n<h1>")
//line nav.ego:6
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(filepath.Base(tx.DB().Path()))))
//line nav.ego:6
	_, _ = io.WriteString(w, "</h1>")
	return nil
}

//line page.ego:1
func Page(w io.Writer, r *http.Request, tx *bolt.Tx, indexes []int, directID int, showUsage bool) error {
//line page.ego:2
	_, _ = io.WriteString(w, "\n\n")
//line page.ego:4
	_, _ = io.WriteString(w, "\n")
//line page.ego:5
	_, _ = io.WriteString(w, "\n")
//line page.ego:6
	_, _ = io.WriteString(w, "\n")
//line page.ego:7
	_, _ = io.WriteString(w, "\n\n")
//line page.ego:9
	p, ids, err := find(tx, directID, indexes)
	if err != nil {
		return err
	}

	// Generate page stats.
	pageSize := tx.DB().Info().PageSize
	stats := p.stats(pageSize)

	// Generate histogram of all nested page usage.
	var histogram map[int]int
	if showUsage {
		histogram = usage(tx, p.id)
	}

//line page.ego:25
	_, _ = io.WriteString(w, "\n\n<!DOCTYPE html>\n<html lang=\"en\">\n")
//line page.ego:28
	head(w, tx)
//line page.ego:29
	_, _ = io.WriteString(w, "\n<body>\n\n<div class=\"uk-container uk-container-center uk-margin-top uk-margin-large-bottom\">\n    <nav class=\"uk-navbar uk-margin-large-bottom\">\n        <a class=\"uk-navbar-brand uk-hidden-small\" href=\"/ui/\">GOPA</a>\n        <ul class=\"uk-navbar-nav uk-hidden-small\">\n            <li >\n                <a href=\"/ui/page/index.html\" >Home</a>\n            </li>\n            <li >\n                <a href=\"/ui/page/console.html\" >Console</a>\n            </li>\n            <li >\n                <a href=\"/ui/page/tasks.html\">Tasks</a>\n            </li>\n            <li class=\"uk-active\">\n                <a href=\"/ui/boltdb\" >Boltdb</a>\n            </li>\n        </ul>\n        <a href=\"#offcanvas\" class=\"uk-navbar-toggle uk-visible-small\" data-uk-offcanvas></a>\n        <div class=\"uk-navbar-brand uk-navbar-center uk-visible-small\">GOPA</div>\n    </nav>\n\n    <div class=\"uk-grid\" data-uk-grid-margin>\n        <div class=\"uk-width-large-1-1 uk-visible-large\">\n            <div class=\"uk-alert\" >")
//line page.ego:54
	nav(w, tx)
//line page.ego:54
	_, _ = io.WriteString(w, "</div>\n\n\n          <h2>\n                ")
//line page.ego:58
	for i, id := range ids {
//line page.ego:59
		_, _ = io.WriteString(w, "\n                  ")
//line page.ego:59
		if i > 0 {
//line page.ego:59
			_, _ = io.WriteString(w, "&raquo;")
//line page.ego:59
		}
//line page.ego:60
		_, _ = io.WriteString(w, "\n                  <a href=\"")
//line page.ego:60
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(pagelink(indexes[:i+1]))))
//line page.ego:60
		_, _ = io.WriteString(w, "\">#")
//line page.ego:60
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(id)))
//line page.ego:60
		_, _ = io.WriteString(w, "</a>\n                ")
//line page.ego:61
	}
//line page.ego:62
	_, _ = io.WriteString(w, "\n              </h2>\n\n              <h3>Page Information</h3>\n              <p>\n                <strong>ID:</strong> ")
//line page.ego:66
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(comma(int(p.id)))))
//line page.ego:66
	_, _ = io.WriteString(w, "<br/>\n                <strong>Type:</strong> ")
//line page.ego:67
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(fmt.Sprintf("%s (%x)", p.typ(), p.flags))))
//line page.ego:67
	_, _ = io.WriteString(w, "<br/>\n                <strong>Overflow:</strong> ")
//line page.ego:68
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(p.overflow)))
//line page.ego:68
	_, _ = io.WriteString(w, "<br/><br/>\n\n                <strong>Alloc:</strong> ")
//line page.ego:70
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(comma(stats.alloc))))
//line page.ego:70
	_, _ = io.WriteString(w, "<br/>\n                <strong>In Use:</strong> ")
//line page.ego:71
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(comma(stats.inuse))))
//line page.ego:71
	_, _ = io.WriteString(w, "<br/>\n                <strong>Utilization:</strong> ")
//line page.ego:72
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(fmt.Sprintf("%.2f%%", stats.utilization*100))))
//line page.ego:72
	_, _ = io.WriteString(w, "<br/>\n              </p>\n\n              ")
//line page.ego:75
	if (p.flags & branchPageFlag) != 0 {
//line page.ego:76
		_, _ = io.WriteString(w, "\n                <h3>Branch Elements (")
//line page.ego:76
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(p.count)))
//line page.ego:76
		_, _ = io.WriteString(w, ")</h3>\n                <table>\n                  <thead>\n                    <tr>\n                      <th align=\"left\">Key</th>\n                      <th align=\"left\">Page ID</th>\n                      <th align=\"left\">Size (k)</th>\n                      <th align=\"center\">%%Util</th>\n                    </tr>\n                  </thead>\n                  <tbody>\n                    ")
//line page.ego:87
		for i := uint16(0); i < p.count; i++ {
//line page.ego:88
			_, _ = io.WriteString(w, "\n                      ")
//line page.ego:88
			e := p.branchPageElement(i)
//line page.ego:89
			_, _ = io.WriteString(w, "\n                      ")
//line page.ego:89
			subpage := pageAt(tx, e.pgid)
//line page.ego:90
			_, _ = io.WriteString(w, "\n                      ")
//line page.ego:90
			substats := subpage.stats(pageSize)
//line page.ego:91
			_, _ = io.WriteString(w, "\n                      <tr>\n                        <td>")
//line page.ego:92
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(trunc(tostr(e.key()), 150))))
//line page.ego:92
			_, _ = io.WriteString(w, "</td>\n                        <td><a href=\"")
//line page.ego:93
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(subpagelink(indexes, int(i)))))
//line page.ego:93
			_, _ = io.WriteString(w, "\">")
//line page.ego:93
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(e.pgid)))
//line page.ego:93
			_, _ = io.WriteString(w, "</a></td>\n                        <td>")
//line page.ego:94
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(len(e.key()))))
//line page.ego:94
			_, _ = io.WriteString(w, "</td>\n                        <td align=\"right\">")
//line page.ego:95
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(fmt.Sprintf("%.2f%%", substats.utilization*100))))
//line page.ego:95
			_, _ = io.WriteString(w, "</td>\n                      </tr>\n                    ")
//line page.ego:97
		}
//line page.ego:98
		_, _ = io.WriteString(w, "\n                  </tbody>\n                </table>\n\n              ")
//line page.ego:101
	} else if (p.flags & leafPageFlag) != 0 {
//line page.ego:102
		_, _ = io.WriteString(w, "\n                <h3>Leaf Elements (")
//line page.ego:102
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(p.count)))
//line page.ego:102
		_, _ = io.WriteString(w, ")</h3>\n                <table>\n                  <thead>\n                    <tr>\n                      <th align=\"left\">Key</th>\n                      <th align=\"left\">Value</th>\n                      <th align=\"left\">Size (k/v)</th>\n                      <th align=\"center\">%%Util</th>\n                    </tr>\n                  </thead>\n                  <tbody>\n                    ")
//line page.ego:113
		for i := uint16(0); i < p.count; i++ {
//line page.ego:114
			_, _ = io.WriteString(w, "\n                      ")
//line page.ego:114
			e := p.leafPageElement(i)
//line page.ego:115
			_, _ = io.WriteString(w, "\n                      ")
//line page.ego:115
			if (e.flags & bucketLeafFlag) != 0 {
//line page.ego:116
				_, _ = io.WriteString(w, "\n                        ")
//line page.ego:116
				b := ((*bucket)(unsafe.Pointer(&e.value()[0])))
//line page.ego:117
				_, _ = io.WriteString(w, "\n                        ")
//line page.ego:118
				util := "-"
				if b.root != 0 {
					substats := pageAt(tx, b.root).stats(pageSize)
					util = fmt.Sprintf("%.2f%%", substats.utilization*100)
				}

//line page.ego:125
				_, _ = io.WriteString(w, "\n                        <tr>\n                          <td><strong>")
//line page.ego:126
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(trunc(tostr(e.key()), 150))))
//line page.ego:126
				_, _ = io.WriteString(w, "</strong></td>\n                          <td>\n                            &lt;bucket(root=")
//line page.ego:128
				if b.root != 0 {
//line page.ego:128
					_, _ = io.WriteString(w, "<a href=\"")
//line page.ego:128
					_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(subpagelink(indexes, int(i)))))
//line page.ego:128
					_, _ = io.WriteString(w, "\">")
//line page.ego:128
				}
//line page.ego:128
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(b.root)))
//line page.ego:128
				if b.root != 0 {
//line page.ego:128
					_, _ = io.WriteString(w, "</a>")
//line page.ego:128
				}
//line page.ego:128
				_, _ = io.WriteString(w, "; seq=")
//line page.ego:128
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(b.sequence)))
//line page.ego:128
				_, _ = io.WriteString(w, ")&gt;\n                          </td>\n                          <td>")
//line page.ego:130
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(len(e.key()))))
//line page.ego:130
				_, _ = io.WriteString(w, " / ")
//line page.ego:130
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(len(e.value()))))
//line page.ego:130
				_, _ = io.WriteString(w, "</td>\n                          <td align=\"right\">")
//line page.ego:131
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(util)))
//line page.ego:131
				_, _ = io.WriteString(w, "</td>\n                        </tr>\n                      ")
//line page.ego:133
			} else {
//line page.ego:134
				_, _ = io.WriteString(w, "\n                        <tr>\n                          <td>")
//line page.ego:135
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(trunc(tostr(e.key()), 150))))
//line page.ego:135
				_, _ = io.WriteString(w, "</td>\n                          <td>")
//line page.ego:136
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(trunc(tostr(e.value()), 150))))
//line page.ego:136
				_, _ = io.WriteString(w, "</td>\n                          <td>")
//line page.ego:137
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(len(e.key()))))
//line page.ego:137
				_, _ = io.WriteString(w, " / ")
//line page.ego:137
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(len(e.value()))))
//line page.ego:137
				_, _ = io.WriteString(w, "</td>\n                          <td>&nbsp;</td>\n                        </tr>\n                      ")
//line page.ego:140
			}
//line page.ego:141
			_, _ = io.WriteString(w, "\n                    ")
//line page.ego:141
		}
//line page.ego:142
		_, _ = io.WriteString(w, "\n                  </tbody>\n                </table>\n              ")
//line page.ego:144
	}
//line page.ego:145
	_, _ = io.WriteString(w, "\n\n              ")
//line page.ego:146
	if showUsage {
//line page.ego:147
		_, _ = io.WriteString(w, "\n                ")
//line page.ego:148
		mins, maxs, values := bucketize(histogram)
		vmax, maxlen := 0, 20
		for _, v := range values {
			if v > vmax {
				vmax = v
			}
		}

//line page.ego:157
		_, _ = io.WriteString(w, "\n\n                <h3>Page Usage Histogram</h3>\n                <table>\n                  <thead>\n                    <tr>\n                      <th align=\"left\">Usage (bytes)</th>\n                      <th align=\"left\">Count</th>\n                      <th>&nbsp;</th>\n                    </tr>\n                  </thead>\n                  <tbody>\n                    ")
//line page.ego:168
		for i := 0; i < len(values); i++ {
//line page.ego:169
			_, _ = io.WriteString(w, "\n                      <tr>\n                        <td>")
//line page.ego:170
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(mins[i])))
//line page.ego:170
			_, _ = io.WriteString(w, " - ")
//line page.ego:170
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(maxs[i])))
//line page.ego:170
			_, _ = io.WriteString(w, "</th>\n                        <td>")
//line page.ego:171
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(values[i])))
//line page.ego:171
			_, _ = io.WriteString(w, "</th>\n                        <td>")
//line page.ego:172
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(strings.Repeat("█", int((float64(values[i])/float64(vmax))*float64(maxlen))))))
//line page.ego:172
			_, _ = io.WriteString(w, "</td>\n                      </tr>\n                    ")
//line page.ego:174
		}
//line page.ego:175
		_, _ = io.WriteString(w, "\n                  </tbody>\n                </table>\n              ")
//line page.ego:177
	} else {
//line page.ego:178
		_, _ = io.WriteString(w, "\n                ")
//line page.ego:179
		u, q := r.URL, r.URL.Query()
		q.Set("usage", "true")
		u.RawQuery = q.Encode()

//line page.ego:184
		_, _ = io.WriteString(w, "\n\n                <p><a href=\"")
//line page.ego:185
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(u.String())))
//line page.ego:185
		_, _ = io.WriteString(w, "\">Show Page Usage</a></p>\n              ")
//line page.ego:186
	}
//line page.ego:187
	_, _ = io.WriteString(w, "\n\n              <br/><br/>\n              <form action=\"boltdb\" method=\"GET\">\n                Go to page: <input type=\"text\" name=\"id\"/>\n                <button type=\"submit\">Go</button>\n              </form>\n\n            </div>\n    </div>\n\n</div>\n  </body>\n</html>\n")
	return nil
}