## Running a Single API Test

	% python fileName.py TestCaseName.testName

Or, assuming the python script itself is executable:

	% ./fileName.py TestCaseName.testName
	
	
## Run all Tests

	% ./runAllTests.py
	
Note files with unit tests need to start with "test" in order to be discovered by the runAllTests.py script.


## Testing TODO List

* Create form components, then retrieve them back.
* Set properties on a form component.
* Retrieve the list of filter rules associated with a filter.
* Calculated field formulas
	* Validate formula
	* Get raw text from formula
	* Set formula text.