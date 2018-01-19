const electron = require('electron')
// Module to control application life.
const app = electron.app
// Module to create native browser window.
const BrowserWindow = electron.BrowserWindow

const path = require('path')
const url = require('url')

// Keep a global reference of the window object, if you don't, the window will
// be closed automatically when the JavaScript object is garbage collected.
let mainWindow

function createWindow () {
  // Create the browser window.
  mainWindow = new BrowserWindow({width: 800, height: 600})

  // and load the index.html of the app.  
  mainWindow.loadURL('http://localhost:43401/')

  // Open the DevTools.
//  mainWindow.webContents.openDevTools()

  // Emitted when the window is closed.
  mainWindow.on('closed', function () {
    // Dereference the window object, usually you would store windows
    // in an array if your app supports multi windows, this is the time
    // when you should delete the corresponding element.
    mainWindow = null
  })
}

function launchBackend() {
	
	var backendExe = "/Users/sroehling/Development/go/src/resultra/datasheet/build/dest/bin/datasheetServer";
	var backendArgs = ["--config","/Users/sroehling/Development/devTrackerDatabases/steveTrackerConfig.json"]
	var backendOpts = {
		detached: false,
		// The backend looks uses the CWD as a base path to look for static assets such as Javascript files and 
		// images.
		cwd: "/Users/sroehling/Development/go/src/resultra/datasheet/build/dest"
	}

	const spawn = require('child_process').spawn;
	const backendChildProc = spawn(backendExe, backendArgs,backendOpts);


	// Handle normal output
	backendChildProc.stdout.on('data', (data) => {
	    // As said before, convert the Uint8Array to a readable string.
	    var str = String.fromCharCode.apply(null, data);
	    console.info(str);
	});

	// Handle error output
	backendChildProc.stderr.on('data', (data) => {
	    // As said before, convert the Uint8Array to a readable string.
	    var str = String.fromCharCode.apply(null, data);
	    console.error(str);
	});

	// Handle on exit event
	backendChildProc.on('exit', (code) => {
	    var preText = `Child exited with code ${code} : `;

	    switch(code){
	        case 0:
	            console.info(preText+"Something unknown happened executing the batch.");
	            break;
	        case 1:
	            console.info(preText+"The file already exists");
	            break;
	        case 2:
	            console.info(preText+"The file doesn't exists and now is created");
	            break;
	        case 3:
	            console.info(preText+"An error ocurred while creating the file");
	            break;
	    }
	});
	
	backendChildProc.on('error', (err) => {
	  console.log('Failed to start subprocess.');
	});
	
	return backendChildProc
}

function pingToConfirmBackendStartup(pingCompleteCallback) {
	
   	var request = require("request")
	
	var numRetriesRemaining = 30
	
	function sendOnePingRequest() {
		
		function handlePingResponse(err,response,body) {
						
			if (response === undefined || response.statusCode !== 200) {
				console.log("handlePingResponse: error: " + err + " response=" + JSON.stringify(response))
				
				numRetriesRemaining--
				if(numRetriesRemaining <= 0) {
					pingCompleteCallback(false)
				} else {
					setTimeout(function() {
						sendOnePingRequest()
					},500)
				}
			} else {
					console.log("handlePingResponse: SUCCESS: body: " + JSON.stringify(body))
					pingCompleteCallback(true)
			}			
		}
		
		var pingArgs = {}
		request.post({ url:'http://localhost:43401/api/admin/ping', json: pingArgs }, handlePingResponse)
		
		
	}
	
	setTimeout(sendOnePingRequest,500)
		
}

function launchBackendThenCreateWindow() {
	
	var backendChildProc = launchBackend()
	
	pingToConfirmBackendStartup(function(success) {
		if(success) {
			createWindow()
		} else {
			// TODO - Show some kinf of startup error and quit
		}
	})
	
	app.on('quit',function() {
	  // Send SIGINT, which is equivalent to Cntrl-C and will always terminate
	  // the child process, as opposed to the default SIGTERM
	  backendChildProc.kill('SIGINT')
	})
	
	// Quit when all windows are closed.
	app.on('window-all-closed', function () {
	  // On OS X it is common for applications and their menu bar
	  // to stay active until the user quits explicitly with Cmd + Q
		
	  if (process.platform !== 'darwin') {
	    app.quit()
	  }
	})

	app.on('activate', function () {
	  // On OS X it's common to re-create a window in the app when the
	  // dock icon is clicked and there are no other windows open.
	  if (mainWindow === null) {
	    createWindow()
	  }
	})

}

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.on('ready', launchBackendThenCreateWindow )


// In this file you can include the rest of your app's specific main process
// code. You can also put them in separate files and require them here.
