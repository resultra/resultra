var gulp = require('gulp');
var concat = require('gulp-concat')
var uglify = require('gulp-uglify')
var minifyCSS = require("gulp-minify-css")
var inject = require("gulp-inject")
var gutil = require("gulp-util")
var stripDebug = require("gulp-strip-debug")
var args   = require('yargs').argv;

// The current working directory (cwd) will be the directory with the 
// gulp build-related files, so the distribution/export path is based upon this directory.
var distDir = '../../build/dest/static'

var cwdAbsPath = __dirname

gutil.log("Loading asset list from file: " + args.assets)
var assets = require(args.assets);

gulp.task('exportIndividualAssets', function() {

	gutil.log("Exporting individual javascript files from asset list: # files = " + assets.jsFiles.length)
	gulp.src(assets.jsFiles,{base:assets.basePath})
      .pipe(gulp.dest(distDir))
	
	gutil.log("Exporting individual css files from asset list: # files = " + assets.cssFiles.length)
	gulp.src(assets.cssFiles,{base:assets.basePath})
      .pipe(gulp.dest(distDir))

	gutil.log("Exporting individual html template files from asset list: # files = " + assets.htmlFiles.length)
	gulp.src(assets.htmlFiles,{base:assets.basePath})
      .pipe(gulp.dest(distDir))

	gutil.log("Exporting individual image files from asset list: # files = " + assets.imageFiles.length)
	gulp.src(assets.imageFiles,{base:assets.basePath})
      .pipe(gulp.dest(distDir))

	
})

gulp.task('exportMinifiedAssets', function() {
	
	gutil.log("Building minified javascript file from asset list: files = " + assets.jsFiles.length 
		+ ", target file = " + assets.minJSFile)
	
	gulp.src(assets.jsFiles,{base:assets.basePath})
      .pipe(concat(assets.minJSFile))
	  .pipe(stripDebug()) // strip out console.log() debug messages.
      .pipe(uglify(
		  {mangle: {toplevel: true}} // mange top-level names as well as names within functions
      ))
      .pipe(gulp.dest(distDir))

  // A single CSS file for release builds
	gutil.log("Building minified CSS file from asset list: files = " + assets.cssFiles.length 
		+ ", target file = " + assets.minCSSFile)
	gulp.src(assets.cssFiles,{base:assets.basePath})
	  .pipe(concat(assets.minCSSFile))
	  .pipe(gulp.dest(distDir))

	// TODO - Export concatenated html templates when in release build.
	gutil.log("Exporting individual html template files from asset list: # files = " + assets.htmlFiles.length)
	gulp.src(assets.htmlFiles,{base:assets.basePath})
      .pipe(gulp.dest(distDir))

	gutil.log("Exporting individual image files from asset list: # files = " + assets.imageFiles.length)
	gulp.src(assets.imageFiles,{base:assets.basePath})
      .pipe(gulp.dest(distDir))

})

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

gulp.task('injectHTMLFilesWithIndividualAssets', function() {
		
  	// Inject the list of javascript sources into the HTML template(s).
	// Injection is based upon the absolute path of the JS and HTML file, and uses
	// a transformation function to do the actual mapping.
	var htmlTarget = gulp.src(assets.htmlFiles,{base:assets.basePath})
		
	var jsSources = gulp.src(assets.jsFiles, {read: false})
	var cssSources = gulp.src(assets.cssFiles,{read:false})

	gutil.log("Injecting HTML files with JS and CSS references: html files = " + assets.htmlFiles.length 
		+ ", inject name = " + assets.injectPlaceholderName)
	
	htmlTarget.pipe(inject(jsSources,{name: assets.injectPlaceholderName, transform: transformJSFileForInjection}))
		.pipe(inject(cssSources,{name: assets.injectPlaceholderName, transform: transformCSSFileForInjection}))
		.pipe(gulp.dest(distDir))
});

gulp.task('injectHTMLFilesWithMinifiedAssets', function() {

	var htmlTarget = gulp.src(assets.htmlFiles,{base:assets.basePath})

	gutil.log("Injecting HTML files with minified JS and CSS references: html files = " + assets.htmlFiles.length 
		+ ", inject placholder name = " + assets.injectPlaceholderName)

	var jsSource = gulp.src(assets.jsFiles,{base:assets.basePath})
      .pipe(concat(assets.minJSFile))

	var cssSource = gulp.src(assets.cssFiles,{base:assets.basePath})
      .pipe(concat(assets.minCSSFile))
	
	htmlTarget.pipe(inject(jsSource,{name: assets.injectPlaceholderName,transform: transformJSFileForInjection}))
		.pipe(inject(cssSource,{name: assets.injectPlaceholderName, transform: transformCSSFileForInjection}))
		.pipe(gulp.dest(distDir))
		
});