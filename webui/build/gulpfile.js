var gulp = require('gulp');
var concat = require('gulp-concat')
var uglify = require('gulp-uglify')
var minifyCSS = require("gulp-minify-css")
var inject = require("gulp-inject")
var gutil = require("gulp-util")
var args   = require('yargs').argv;

// The current working directory (cwd) will be the directory with the 
// gulp build-related files, so the distribution/export path is based upon this directory.
var distDir = '../../build/dest/static'

var cwdAbsPath = __dirname

gutil.log("Loading asset list from file: " + args.assets)
var assets = require(args.assets);

gulp.task('debugExportAssets', function() {

	gutil.log("Exporting individual javascript files from asset list: # files = " + assets.jsFiles.length)
	gulp.src(assets.jsFiles,{base:assets.basePath})
      .pipe(gulp.dest(distDir))
	
	gutil.log("Exporting individual css files from asset list: # files = " + assets.cssFiles.length)
	gulp.src(assets.cssFiles,{base:assets.basePath})
      .pipe(gulp.dest(distDir))
})

gulp.task('exportMinifiedAssets', function() {
	
	gutil.log("Building minified javascript file from asset list: files = " + assets.jsFiles.length 
		+ ", target file = " + assets.minJSFile)
	
  gulp.src(assets.jsFiles,{base:assets.basePath})
      .pipe(concat(assets.minJSFile))
      .pipe(uglify())
      .pipe(gulp.dest(distDir))

  // A single CSS file for release builds
	gutil.log("Building minified CSS file from asset list: files = " + assets.cssFiles.length 
		+ ", target file = " + assets.minCSSFile)
  gulp.src(assets.cssFiles,{base:assets.basePath})
	  .pipe(concat(assets.minCSSFile))
	  .pipe(minifyCSS({keepBreaks:true}))
	  .pipe(gulp.dest(distDir))

})

gulp.task('injectHTMLFiles', function() {
		
  	// Inject the list of javascript sources into the HTML template(s).
	// Injection is based upon the absolute path of the JS and HTML file, and uses
	// a transformation function to do the actual mapping.
	var htmlTarget = gulp.src(assets.htmlFiles)
		
	var jsSources = gulp.src(assets.jsFiles, {read: false})
	var cssSources = gulp.src(assets.cssFiles,{read:false})

	gutil.log("Injecting HTML files with JS and CSS references: html files = " + assets.htmlFiles.length 
		+ ", inject name = " + assets.injectJSName)
	
	// JS file inject is based upon the fully qualified (absolute) file names given 
	// to this task. The transformation simply strips off the basePath of the JS file
	// before returning the HTML reference to be inserted into to the target file.
	// The available gulp-inject options couldn't be made to work right correct prefix
	// to the JS file name.	
	function transformJSFileForInjection(filepath, file, index, length, targetFile) {
		var newJSPath = '/static' + file.path.replace(assets.basePath,"")
		gutil.log("transformJSFileForInjection: JS filepath: " + JSON.stringify(file.path) 
				+ " target HTML file: " + JSON.stringify(targetFile.path))
		gutil.log("transformJSFileForInjection: transformed path: " + newJSPath)
		return '<script src="' + newJSPath  + '"></script>'
	}
	
	function transformCSSFileForInjection(filepath, file, index, length, targetFile) {
		var newCSSPath = '/static' + file.path.replace(assets.basePath,"")
		gutil.log("transformCSSFileForInjection: CSS filepath: " + JSON.stringify(file.path) 
				+ " target HTML file: " + JSON.stringify(targetFile.path))
		gutil.log("transformCSSFileForInjection: transformed path: " + newCSSPath)
		return '<link rel="stylesheet" href="' + newCSSPath  + '">'
	}
	
	htmlTarget.pipe(inject(jsSources,{name: assets.injectPlaceholderName, transform: transformJSFileForInjection}))
		.pipe(inject(cssSources,{name: assets.injectPlaceholderName, transform: transformCSSFileForInjection}))
		.pipe(gulp.dest(distDir))
});

/*

gulp.task('designFormRelease', function() {
  // A single JS files for release builds
  var minifiedJS = gulp.src(designFormAssets.jsFiles,{base:designFormAssets.basePath})
      .pipe(concat('form/designForm.min.js'))
      .pipe(uglify())
      .pipe(gulp.dest(distDir))
	
  var minifiedJSSource = gulp.src(designFormAssets.jsFiles,{base:designFormAssets.basePath})
      .pipe(concat('form/designForm.min.js'))
	
  // A single CSS file for release builds
  gulp.src(designFormAssets.cssFiles,{base:designFormAssets.basePath})
	  .pipe(concat('designForm.min.css'))
	  .pipe(minifyCSS({keepBreaks:true}))
	  .pipe(gulp.dest(distDir))
	
	var target = gulp.src('./designForm.html')

	target.pipe(inject(minifiedJSSource,{name: 'designFormJS',relative:true,addPrefix:'/static/form'}))
		.pipe(gulp.dest(distDir))

});

*/
