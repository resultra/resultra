function colorClassByColorScheme(colorScheme) {
	var classLookup = {
		default:"",
		info:"condFormat-info",
		primary: "condFormat-primary",
		success: "condFormat-success",
		warning: "condFormat-warning",
		danger:"condFormat-danger"
	}

	var bgClass = classLookup[colorScheme]
	if (bgClass === undefined) {
		bgClass = null
	}
	return bgClass
}


function removeConditionalFormatClasses($container) {
	$container.removeClass("condFormat-info condFormat-primary condFormat-success condFormat-warning condFormat-danger")
}
