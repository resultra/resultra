// This code originated from https://gist.github.com/tomerd/1499279, which the author released into public
// domain for commercial/personal use ("no special license. feel free to use it however you like."). This
// code served as a starting point, which was subsequently heavily modified for the project.

function GaugeUIControl($gaugeContainer, configuration)
{
	
	var valueBandClass = "valueBand"
	
//	this.placeholderName = placeholderName;
	this.gaugeContainerElem = $gaugeContainer.get(0)
	
	var self = this; // for internal d3 functions
	
	this.configure = function(configuration)
	{
		this.config = configuration;
		
		this.config.size = this.config.size * 0.9;
		
		this.config.raduis = this.config.size * 0.97 / 2;
		this.config.cx = this.config.size / 2;
		this.config.cy = this.config.size / 2;
		
		this.config.min = undefined != configuration.min ? configuration.min : 0; 
		this.config.max = undefined != configuration.max ? configuration.max : 100; 
		this.config.range = this.config.max - this.config.min;
		
		this.config.majorTicks = configuration.majorTicks || 5;
		this.config.minorTicks = configuration.minorTicks || 2;
		
		this.config.greenColor 	= configuration.greenColor || "#109618";
		this.config.yellowColor = configuration.yellowColor || "#FF9900";
		this.config.redColor 	= configuration.redColor || "#DC3912";
		
		this.config.transitionDuration = configuration.transitionDuration || 500;
	}
	
	this.valueToDegrees = function(value)
	{
		// thanks @closealert
		//return value / this.config.range * 270 - 45;
		return value / this.config.range * 270 - (this.config.min / this.config.range * 270 + 45);
	}
	
	this.valueToRadians = function(value)
	{
		return this.valueToDegrees(value) * Math.PI / 180;
	}
	
	this.valueToPoint = function(value, factor)
	{
		return { 	x: this.config.cx - this.config.raduis * factor * Math.cos(this.valueToRadians(value)),
					y: this.config.cy - this.config.raduis * factor * Math.sin(this.valueToRadians(value)) 		};
	}
	
	this.valueToThresholdColor = function(value) {
		
		for (var index in this.config.greenZones)
		{
			var range = this.config.greenZones[index]		
			if (value >= range.from && value <= range.to) {
				return self.config.greenColor
			}
		}
		
		for (var index in this.config.yellowZones)
		{
			var range = this.config.yellowZones[index]		
			if (value >= range.from && value <= range.to) {
				return self.config.yellowColor
			}
		}
		
		for (var index in this.config.redZones)
		{
			var range = this.config.redZones[index]		
			if (value >= range.from && value <= range.to) {
				return self.config.redColor
			}
		}
		return self.config.greenColor
		
	}
	
	this.createArc = function(startVal,endVal,innerBandRadiusPerc,outerBandRadiusPerc) {
		var arc = d3.arc()
			.startAngle(this.valueToRadians(startVal))
			.endAngle(this.valueToRadians(endVal))
			.innerRadius(innerBandRadiusPerc * this.config.raduis)
			.outerRadius(outerBandRadiusPerc * this.config.raduis)
		return arc
		
	}
	
	this.renderBand = function(startVal,endVal,color,innerBandRadiusPerc,outerBandRadiusPerc,arcOpacity,bandClass) {
		
		var arc = this.createArc(startVal,endVal,innerBandRadiusPerc,outerBandRadiusPerc)

		this.body.append("svg:path")
					.style("fill", color)
					.style("opacity",arcOpacity)
					.attr("d", arc)
					.attr("class",bandClass)
					.attr("transform", function() { return "translate(" + self.config.cx + ", " + self.config.cy + ") rotate(270)" });
	}
	
	this.renderThresholdBand = function(startVal, endVal, color)
	{
		if (0 >= endVal - startVal) return;
		
		// Define the thickness of the bands.
		var innerBandRadiusPerc = 0.50
		var outerBandRadiusPerc = 0.95
		
		var arcOpacity = 0.1 // range is 0-1, with 0 being completely transparent, 1 being opaque
		
		var thresholdBandClass = "thresholdBand"
		
		this.renderBand(startVal,endVal,color,innerBandRadiusPerc,outerBandRadiusPerc,arcOpacity,thresholdBandClass)
		
	}
	
	this.renderThresholdBands = function() {
		for (var index in this.config.greenZones)
		{
			this.renderThresholdBand(this.config.greenZones[index].from, this.config.greenZones[index].to, self.config.greenColor);
		}
		
		for (var index in this.config.yellowZones)
		{
			this.renderThresholdBand(this.config.yellowZones[index].from, this.config.yellowZones[index].to, self.config.yellowColor);
		}
		
		for (var index in this.config.redZones)
		{
			this.renderThresholdBand(this.config.redZones[index].from, this.config.redZones[index].to, self.config.redColor);
		}		
	}

	this.renderValueBand = function(startVal, endVal, color)
	{
		if ((endVal-startVal)<=0) return;
		
		// Define the thickness of the bands.
		var innerBandRadiusPerc = 0.60
		var outerBandRadiusPerc = 0.85
		
		// The actual value is fully opaque, but the inner and outer radius of the value band is smaller
		// than the thresholds so the user can still see where the value is at in relation to thresholds.
		var arcOpacity = 1.0
		
		var valueBandClass = "valueBand"
		
		this.renderBand(startVal,endVal,color,innerBandRadiusPerc,outerBandRadiusPerc,arcOpacity,valueBandClass)
		
	}
	
	this.redrawValueBand = function(newVal) {
		
		// Define the thickness of the bands.
		var innerBandRadiusPerc = 0.60
		var outerBandRadiusPerc = 0.85

		var startVal = 0 // TODO - This can be specified by the user
		
		var newArc = this.createArc(startVal,newVal,innerBandRadiusPerc,outerBandRadiusPerc)
		var newColor = this.valueToThresholdColor(newVal)


		var valueBand = this.body.selectAll("."+valueBandClass)
		// TODO - Transition the arc using attrTween
		valueBand.attr("d",newArc)
			.style("fill", newColor)
			
		
	}
	
	this.renderMajorMinorTicks = function() {
		var fontSize = Math.round(this.config.size / 16);
		var majorDelta = this.config.range / (this.config.majorTicks - 1);
		for (var major = this.config.min; major <= this.config.max; major += majorDelta)
		{
			var minorDelta = majorDelta / this.config.minorTicks;
			for (var minor = major + minorDelta; minor < Math.min(major + majorDelta, this.config.max); minor += minorDelta)
			{
				
				var minorTicksInnerRadiusPerc = 0.8
				var minorTicksOuterRadiusPerc = 0.95
				
				var point1 = this.valueToPoint(minor, minorTicksInnerRadiusPerc);
				var point2 = this.valueToPoint(minor, minorTicksOuterRadiusPerc);
				
				this.body.append("svg:line")
							.attr("x1", point1.x)
							.attr("y1", point1.y)
							.attr("x2", point2.x)
							.attr("y2", point2.y)
							.style("stroke", "#666")
							.style("stroke-width", "1px");
			}
			
			var majorTicksInnerRadiusPerc = 0.7
			var majorTicksOuterRadiusPerc = 0.95
			
			var point1 = this.valueToPoint(major, majorTicksInnerRadiusPerc);
			var point2 = this.valueToPoint(major, majorTicksOuterRadiusPerc);	
			
			this.body.append("svg:line")
						.attr("x1", point1.x)
						.attr("y1", point1.y)
						.attr("x2", point2.x)
						.attr("y2", point2.y)
						.style("stroke", "#333")
						.style("stroke-width", "2px");
			
			if (major == this.config.min || major == this.config.max)
			{
				var point = this.valueToPoint(major, 0.63);
				
				this.body.append("svg:text")
				 			.attr("x", point.x)
				 			.attr("y", point.y)
				 			.attr("dy", fontSize / 3)
				 			.attr("text-anchor", major == this.config.min ? "start" : "end")
				 			.text(major)
				 			.style("font-size", fontSize + "px")
							.style("fill", "#333")
							.style("stroke-width", "0px");
			}
		}
		
	}
	
	this.renderLabel = function() {
		if (undefined != this.config.label)
		{
			var fontSize = Math.round(this.config.size / 9);
			this.body.append("svg:text")
						.attr("x", this.config.cx)
						.attr("y",this.config.cy + fontSize/3)
						.attr("text-anchor", "middle")
						.text(this.config.label)
						.style("font-size", fontSize + "px")
						.style("fill", "#333")
						.style("stroke-width", "0px");
		}
	}
	
	this.buildPointerPath = function(value)
	{
		var delta = this.config.range / 25;
		
		var head = valueToPoint(value, 0.55);
		var head1 = valueToPoint(value - delta, 0.45);
		var head2 = valueToPoint(value + delta, 0.45);
		
		return [head, head1, head2]
		
		function valueToPoint(value, factor)
		{
			var point = self.valueToPoint(value, factor);
			point.x -= self.config.cx;
			point.y -= self.config.cy;
			return point;
		}
	}
	
	
	this.renderPointer = function() {
		var pointerContainer = this.body.append("svg:g").attr("class", "pointerContainer");
		
		var midValue = (this.config.min + this.config.max) / 2;
		
		var pointerPath = this.buildPointerPath(midValue);
		
		var pointerLine = d3.line()
			.x(function(d) { return d.x })
			.y(function(d) { return d.y })
			
		var triangle = d3.symbol()
				.type(d3.symbolTriangle)
            	.size(30)
		
		pointerContainer.selectAll("path")
							.data([pointerPath])
							.enter()
								.append("svg:path")
									.attr("d", pointerLine)
									.style("fill", "#000")
									.style("stroke", "#000")
									.style("fill-opacity", 0.7)					
		
		var fontSize = Math.round(this.config.size / 10);
		pointerContainer.selectAll("text")
							.data([midValue])
							.enter()
								.append("svg:text")
									.attr("x", this.config.cx)
									.attr("y", this.config.size - this.config.cy / 4 - fontSize)
									.attr("dy", fontSize / 2)
									.attr("text-anchor", "middle")
									.style("font-size", fontSize + "px")
									.style("fill", "#000")
									.style("stroke-width", "0px");
		
	}
	
	this.redrawPointer = function(value) {
		
		var pointerContainer = this.body.select(".pointerContainer");
		
		
		
		var pointer = pointerContainer.selectAll("path");
		pointer.transition()
					.duration(this.config.transitionDuration)
					.attrTween("transform", function()
					{
						var pointerValue = value;
						if (value > self.config.max) pointerValue = self.config.max + 0.02*self.config.range;
						else if (value < self.config.min) pointerValue = self.config.min - 0.02*self.config.range;
						var targetRotation = (self.valueToDegrees(pointerValue) - 90);
						var currentRotation = self._currentRotation || targetRotation;
						self._currentRotation = targetRotation;
						
						return function(step) 
						{
							var rotation = currentRotation + (targetRotation-currentRotation)*step;
							return "translate(" + self.config.cx + ", " + self.config.cy + ") rotate(" + rotation + ")"; 
						}
					});
		
	}

	
	this.render = function()
	{
		// Main container for the overall gauge
		this.body = d3.select(this.gaugeContainerElem)
							.append("svg:svg")
							.attr("class", "gauge")
							.attr("width", this.config.size)
							.attr("height", this.config.size);
							
		
		this.renderThresholdBands()
		this.renderValueBand(0,40,self.config.yellowColor)
		this.renderLabel()
		this.renderMajorMinorTicks()
		this.renderPointer()
		
		this.redraw(this.config.min, 0);
	}
	
	
	
	this.redraw = function(value, transitionDuration)
	{
		var pointerContainer = this.body.select(".pointerContainer");
		pointerContainer.selectAll("text").text(Math.round(value)); // update the value text
		
		this.redrawValueBand(value)
		this.redrawPointer(value)
		
	}
	
	
	// initialization
	this.configure(configuration);	
}