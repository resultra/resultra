# GAE Project File Organization for Development vs. Testing vs. Deployment

When go language projects are run under Google App Engine (GAE), the process which builds the GAE project
looks at all subdirectories underneath app.yaml. This can cause conflicts during the build of go language files.

The recommendation is then to place the app.yaml outside the root of the golang source tree (see https://goo.gl/xjCbpz
)

As a convention for this project, all static files (html, css, javascript, png, jpg, etc) are maintained
alongside the golang source, then copied over to the "static" subfolder of this directory. This is in
the interest of the package principle that "what changes together stays together". Specifically, in the 
case of golang source files which render html pages from templates:

* There is a close coupling/dependency on the golang source file and the html templates.
	* The golang "page renderer" files refer to the html templates by name.
	* The golang page renderer's pass parameters which are referred to by name from the 
	  templates themselves. 
* The html templates in turn have a close coupling with the javascript, css and image files. 
  In particular, the javascript files implement dynamic behavior for the DOM's elements defined
  within the html templates. 

This project file organization emphasizes packages and their dependencies, for purposes of 
maintaining the source code, over the directory structure for the running web app. 

The simple script mirrorStaticFilesToGoogleAppEngProjFolder (sorry for the long name) copies
all the static files under the GAE project folder in one step. The only downside is that 
this script must be re-run every time a static file changes. 