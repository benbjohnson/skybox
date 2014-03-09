var chart = d3.select(".chart");
var svg = chart.append("svg");
var g = {
    bars: svg.append("g"),
    axis: {
        x: svg.append("g"),
        y: svg.append("g"),
    },
};
var margin = {
    top: 20,
    bottom: 20,
    left: 50,
    right: 20,
};

// Renders a funnel result as a chart.
function update(result) {
    var width = $(".chart").width();
    var height = 200;

    svg.attr("width", width).attr("height", height);

    // Setup scales.
    var scales = {
        x: d3.scale.ordinal(),
        y: d3.scale.linear(),
    };
    scales.x
        .rangeRoundBands([0, width - margin.left - margin.right], 0.1)
        .domain(result.steps.map(function(d) { return d.condition; }));
    scales.y
        .range([height - margin.top - margin.bottom, 0])
        .domain([0, d3.max(result.steps, function(d) { return d.count; })]);

    // Setup g elements.
    g.axis.x.attr("width", width).attr("height", height);
    g.axis.y.attr("width", width).attr("height", height);
    g.bars.attr("width", width).attr("height", height);

    // Draw axes.
    var axes = {
        x: d3.svg.axis().scale(scales.x).orient("bottom").tickFormat(function(d) { return d.replace(/^resource == "(.+)"$/, "$1"); }),
        y: d3.svg.axis().scale(scales.y).orient("left").ticks(3, "0s"),
    };
    g.axis.x
        .attr("class", "x axis")
        .attr("transform", "translate(" + margin.left + "," + (height - margin.bottom) + ")")
        .call(axes.x);
    g.axis.y
        .attr("class", "y axis")
        .attr("transform", "translate(" + margin.left + "," + margin.top + ")")
        .call(axes.y)

    // Draw bars.
    g.bars.attr("transform", "translate(" + margin.left + "," + margin.top + ")");
    g.bars.selectAll(".bar").data(result.steps)
      .call(function() {
        var bar = this.enter().append("rect")
            .attr("class", "bar");
        bar = this
            .attr("x", function(d) { return scales.x(d.condition); })
            .attr("width", scales.x.rangeBand())
            .attr("y", function(d) { return scales.y(d.count); })
            .attr("height", function(d) { return height - scales.y(d.count) - margin.top - margin.bottom; });
      })
}
