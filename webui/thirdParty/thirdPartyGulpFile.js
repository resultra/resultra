var gulp = require('gulp');
var concat = require('gulp-concat')
var inject = require("gulp-inject")
var rename = require("gulp-rename")
var gutil = require("gulp-util")
var args   = require('yargs').argv;

// The current working directory (cwd) will be the directory with the 
// gulp build-related files, so the distribution/export path is based upon this directory.
//var distDir = '../../build/dest/static'
 var distDir = './static'

gulp.task('exportThirdPartyAssets',function() {
	
	var jsFiles = {
		'node_modules/bootstrap/js/bootstrap.min.css'
	}
	
	gulp.src(jsFiles)
	  .pipe(rename(function(path) {
		  path.dirname = pathName.replace(/node_modules\/dist\//,"")
		  return path
	  }))
      .pipe(gulp.dest(distDir))

/*	  
	gulp.src(assets.cssFiles,{base:assets.basePath})
	  .pipe(rename(function(path) {
		  path.dirname = flattenPackagePathName(path.dirname)
		  return path
	  }))
      .pipe(gulp.dest(distDir))
*/

})
