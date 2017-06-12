var gulp = require('gulp');
var concat = require('gulp-concat')
var inject = require("gulp-inject")
var rename = require("gulp-rename")
var gutil = require("gulp-util")
var merge = require('merge-stream');
var args   = require('yargs').argv;

// The current working directory (cwd) will be the directory with the 
// gulp build-related files, so the distribution/export path is based upon this directory.
var distDir = '../../build/dest/static'

var cwdAbsPath = __dirname

gutil.log("Loading package asset list from file: " + args.pkgassets)
var pkgAssets = require(args.pkgassets)

var pathPrefix = args.pathprefix
gutil.log("Path prefix: " + pathPrefix)

function prefixFilesWithPathPrefix(prefix, fileList) {
	var absFiles = []
	for(var f in fileList) {
		var fName = prefix + '/' + fileList[f]
		absFiles.push(fName)
	}
	return absFiles
}

gulp.task('exportThirdPartyAssets',function() {
	
	for(var pkgName in pkgAssets.packages) {
		
		gutil.log("Exporting package assets: package name = " + pkgName)
		var pkgInfo = pkgAssets.packages[pkgName]
		gutil.log("package info: " + JSON.stringify(pkgInfo))

		function exportFileList(pkgName, pkgInfo, pkgFiles) {
			
			if (pkgFiles === undefined) { return }
			
	  		// The asset file name given in the package asset file is only the filename
	  		// relative to the pkgPrefix (also in the package asset file). To get the 
	  		// fully-qualified file name, both the package prefix and third party 
	  		// directory name must be pre-pended.
	  		var absFiles = prefixFilesWithPathPrefix(pkgInfo.pkgPrefix,pkgFiles)
	  		absFiles = prefixFilesWithPathPrefix(pathPrefix,absFiles)
		
	  	  	var basePath = "/Users/sroehling/Development/go/src/resultra/datasheet/webui"
	
	  	  	gulp.src(absFiles, {base:basePath})
	  	  	  .pipe(rename(function(path) {
				  	  			  // Replace the package prefix (package directory location relative
	  			  // to the third party directory) with just the package name.
	  	  		  path.dirname = path.dirname.replace(pkgInfo.pkgPrefix,pkgName)
	  	  		  return path
	  	  	  }))
	  	      .pipe(gulp.dest(distDir))
		}
		
		exportFileList(pkgName, pkgInfo,pkgInfo.jsFiles)
		exportFileList(pkgName, pkgInfo,pkgInfo.cssFiles)
		exportFileList(pkgName, pkgInfo,pkgInfo.fontFiles)
		exportFileList(pkgName, pkgInfo,pkgInfo.imageFiles)
		
	}


});

gulp.task('injectHTMLFilesWithIndividualPkgAssets', function() {
	
	var absHTMLFiles = prefixFilesWithPathPrefix(pathPrefix,pkgAssets.htmlFiles)
	gutil.log("Injecting package assets: html files  = " + JSON.stringify(absHTMLFiles))
	
	var basePath = "/Users/sroehling/Development/go/src/resultra/datasheet/webui"
	var htmlTarget = gulp.src(absHTMLFiles,{base:basePath})	
	
	// The files for individual packages are merged into a single stream, then
	// transformJSPkgFileForInjection is called for all the different packages.
	// absFileToLinkFile is needed to map the absolute file names from the source
	// stream to the script/link reference.
	var absFileToLinkFile = {}
	function transformJSPkgFileForInjection(filepath, file, index, length, targetFile) {
		return '<script src="' + absFileToLinkFile[file.path] + '"></script>'
	}
	
	function transformCSSPkgFileForInjection(filepath, file, index, length, targetFile) {
		return '<link rel="stylesheet" href="' + absFileToLinkFile[file.path]  + '">'
	}
	
	
	var allPkgJSFileRefSrcs = merge(); // Create an empty stream
	for(var pkgName in pkgAssets.packages) {
		
		gutil.log("Injecting package assets: package name = " + pkgName)
		var pkgInfo = pkgAssets.packages[pkgName]
			
		// The fully-qualified package file name are needed for injection. Using
		// absFileToLinkFile, this fully-qualified name is then mapped to the Javascript
		// reference name transformJSPkgFileForInjection. 
	  	var absFiles = prefixFilesWithPathPrefix(pkgInfo.pkgPrefix,pkgInfo.jsFiles)
	  	var fileRefs = prefixFilesWithPathPrefix(pathPrefix,absFiles)
		gutil.log("Injecting package assets: files  = " + JSON.stringify(fileRefs))
		
		// Populate absFileToLinkFile for each individual file reference (see above for details).
		for (var absFileIndex in fileRefs) {
			var absFileName = fileRefs[absFileIndex]
			var linkJSPath = absFileName.replace(basePath,'') // remove the base path
			linkJSPath = linkJSPath.replace(pkgInfo.pkgPrefix,pkgName) // replace package prefix with just the package name
			linkJSPath = 'static' + linkJSPath // prepend the static reference.
			absFileToLinkFile[absFileName] = linkJSPath
		}
		
		// Merge the file source streams for individual packages into a merged stream 
		// to be used for injection into the HTML file(s)
		var fileRefSrcs = gulp.src(fileRefs,{read: false}) 
		allPkgJSFileRefSrcs.add(fileRefSrcs)
						
	}
	
	gutil.log("Mapping from package file to JS link:" + JSON.stringify(absFileToLinkFile))
	
	htmlTarget.pipe(inject(allPkgJSFileRefSrcs,{name: pkgAssets.injectPlaceholderName, transform: transformJSPkgFileForInjection}))
	    .pipe(gulp.dest(distDir))
				
});
