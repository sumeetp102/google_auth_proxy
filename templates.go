package main

import (
	"html/template"
	"log"
)

func getTemplates() *template.Template {
	t, err := template.New("foo").Parse(`{{define "sign_in.html"}}
<!DOCTYPE html>
<html lang="en" charset="utf-8">
<head><title>Sign In</title></head>
<style>body {
  background-color: #25272b;
  color: #fff;
}

#container {
  max-width: 920px;
  margin: 40px auto;
}

#title {
  float: left;
  width: 470px;
}

#logo {
  width: 443px;
  height: 467px;
  text-indent: -10000px;
  background: url(logo.jpg) top left no-repeat;
}

#small-logo { display: none; padding-top: 30px; }

#main {
  padding-top: 100px;
  margin-left: 470px;
  max-width: 450px;
}

h1 {
  font-weight: normal;
  font-style: italic;
  font-size: 1.6em;
  border-bottom: 1px solid #eee;
  margin-bottom: 1.5em;
  font-family: Palatino, "Times New Roman", serif;
}

h2 {
  font-weight: normal;
  font-size: 1.2em;
  color: #CCC;
}

a.button {
  display: block;
  color: #fff;
  text-decoration: none;
  text-shadow: 0 -1px 0 rgba(0,0,0,.5);
  background-color: #3072b3; /* Old browsers */
  background-repeat: repeat-x; /* Repeat the gradient */
  background-image: -moz-linear-gradient(top, #599bdc 0%, #3072b3 100%); /* FF3.6+ */
  background-image: -webkit-gradient(linear, left top, left bottom, color-stop(0%,#599bdc), color-stop(100%,#3072b3)); /* Chrome,Safari4+ */
  background-image: -webkit-linear-gradient(top, #599bdc 0%,#3072b3 100%); /* Chrome 10+,Safari 5.1+ */
  background-image: -ms-linear-gradient(top, #599bdc 0%,#3072b3 100%); /* IE10+ */
  background-image: -o-linear-gradient(top, #599bdc 0%,#3072b3 100%); /* Opera 11.10+ */
  filter: progid:DXImageTransform.Microsoft.gradient( startColorstr='#599bdc', endColorstr='#3072b3',GradientType=0 ); /* IE6-9 */
  background-image: linear-gradient(top, #599bdc 0%,#3072b3 100%); /* W3C */
  border: 1px solid #2967a4;
  -webkit-transition: none;
  -moz-transition: none;
  transition: none;
  -webkit-box-shadow: inset 0 1px 0 rgba(255,255,255,.1), 0 1px 2px rgba(0,0,0,.2);
  -moz-box-shadow: inset 0 1px 0 rgba(255,255,255,.1), 0 1px 2px rgba(0,0,0,.2);
  box-shadow: inset 0 1px 0 rgba(255,255,255,.1), 0 1px 2px rgba(0,0,0,.2);

  font-size: 20px;
  font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
  padding: 12px 24px;
  margin-bottom: 20px;
  -webkit-border-radius: 6px;
  -moz-border-radius: 6px;
  border-radius: 6px;
}

a.button:hover {
  background-position: 0 -15px;
}

a.button:active {
  background-color: #3072b3; /* Old browsers */
  background-repeat: repeat-x; /* Repeat the gradient */
  background-image: -moz-linear-gradient(top, #285e93 0%, #3072b3 100%); /* FF3.6+ */
  background-image: -webkit-gradient(linear, left top, left bottom, color-stop(0%,#285e93), color-stop(100%,#3072b3)); /* Chrome,Safari4+ */
  background-image: -webkit-linear-gradient(top, #285e93 0%,#3072b3 100%); /* Chrome 10+,Safari 5.1+ */
  background-image: -ms-linear-gradient(top, #285e93 0%,#3072b3 100%); /* IE10+ */
  background-image: -o-linear-gradient(top, #285e93 0%,#3072b3 100%); /* Opera 11.10+ */
  filter: progid:DXImageTransform.Microsoft.gradient( startColorstr='#285e93', endColorstr='#3072b3',GradientType=0 ); /* IE6-9 */
  background-image: linear-gradient(top, #285e93 0%,#3072b3 100%); /* W3C */
  -webkit-box-shadow: inset 0 10px 10px rgba(0,0,0,.15), 0 1px 2px rgba(0,0,0,.2);
  -moz-box-shadow: inset 0 10px 10px rgba(0,0,0,.15), 0 1px 2px rgba(0,0,0,.2);
  box-shadow: inset 0 10px 10px rgba(0,0,0,.15), 0 1px 2px rgba(0,0,0,.2);
}

a.button img {
  background-color: #fff;
  margin-right: 5px;
  padding: 2px;
  border-radius: 2px;
  margin-bottom: -3px;
  width: 16px;
  height: 16px;
}

@media screen and (max-width: 776px) {
  #title { display: none; }
  #main { margin: 0 auto; padding-top: 0px; }
  #small-logo { display: block; }
  #small-logo img { width: 100%; }
}


</style></head><body><div id="container"><div id="title"><div id="logo">Doorman</div></div><div id="main"><h1>How may I help you?</h1><div id="provider-list"><a href="/oauth2/start" class="button"><img src="google.ico.png"/>Sign in with Google
</a></div><div id="small-logo"><img src="logo.jpg"></div></div></div></body></html>
</html>
{{end}}`)
	if err != nil {
		log.Fatalf("failed parsing template %s", err.Error())
	}

	t, err = t.Parse(`{{define "error.html"}}
<!DOCTYPE html>
<html lang="en" charset="utf-8">
<head><title>{{.Title}}</title></head>
<body>
	<h2>{{.Title}}</h2>
	<p>{{.Message}}</p>
	<hr>
	<p><a href="/oauth2/sign_in">Sign In</a></p>
</body>
</html>{{end}}`)
	if err != nil {
		log.Fatalf("failed parsing template %s", err.Error())
	}
	return t
}
