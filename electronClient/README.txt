This package is for building the Electron client for single-user use of the application.

In development mode, the app is build using the following:

(1) npm install
(2) npm start

This will install all the modules needed to build the app and then launch it.

This client will connect to a tracking database server running locally. More development
is needed to launch the tracking database server from the Electron app itself.


The app includes dependencies, including the request node module. To include this in the final build, I found it was needed as an explicit dependency in the package.json file: i.e.:

  "dependencies": {
      "request": "^2.83.0" 	
  },

Just as a note, if other external dependencies are included, they should also be added to the package.json file.

The app is setup to use electron-builder to package the binary (first installed via `npm install electron-builder --save-dev`). To package the distrubition, run the following:

	$ npm run dist
	
	
The application icon for Mac OS X was generated from a 512x512 .png file of Resultra's logo. The Mac Program called Image2icon is used to generate the icon.icns file in the build directory.
	
The factory templates inside the factoryTemplates/ directory are copied to the build directory and included in the packaged application. Update the file trackers.db to include any additional templates.

The license file is an RTF export from ResultraDesktopLicense.docx. So, to make any changes to the license, open ResultraDesktopLicense.docx in MS Word, then save it into the build folder as an RTF name license_en.rtf.
	