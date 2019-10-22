console.info({ d: data })

// Set margins for svg 
const margin = {top: 50, right: 50, bottom: 50, left: 50};
const width = window.innerWidth - margin.left - margin.right;
const height = window.innerHeight - margin.top - margin.bottom;

const datasetIn = data.in.line;
const datasetDP = data.outDP.line;
const datasetVis = data.outVis.line;

const colors = [
    "red",
    "green",
    "blue"
]

// Scales
const xScale = d3.scaleLinear()
    .domain([d3.min(datasetIn, (d) => d[0]), d3.max(datasetIn, (d) => d[0])])
    .range([0, width]);

const yScale = d3.scaleLinear()
    .domain([d3.min(datasetIn, (d) => d[1]), d3.max(datasetIn, (d) => d[1])])
    .range([height, 0]); 

// Line generator
const inLine = d3.line()
    .x((d, i) => xScale(d[0])) 
    .y((d) => yScale(d[1]))

const outLineDP = d3.line()
    .x((d, i) => xScale(d[0]))
    .y((d) => yScale(d[1]))

const outLineVis = d3.line()
    .x((d, i) => xScale(d[0]))
    .y((d) => yScale(d[1]))

// Add SVG to page
const svg = d3.select("body").append("svg")
    .attr("width", width + margin.left + margin.right)
    .attr("height", height + margin.top + margin.bottom)
  .append("g")
    .attr("transform", "translate(" + margin.left + "," + margin.top + ")");

// Call x axis in a group tag
svg.append("g")
    .attr("class", "x axis")
    .attr("transform", "translate(0," + height + ")")
    .call(d3.axisBottom(xScale)); // Create an axis component with d3.axisBottom

// Call y axis in a group tag
svg.append("g")
    .attr("class", "y axis")
    .call(d3.axisLeft(yScale)); // Create an axis component with d3.axisLeft

// Append path, bind data, call line generator 
svg.append("path")
    .datum(datasetIn)
    .attr("fill", "none")
    .attr("stroke", colors[0])
    .attr("d", inLine);

svg.append("path")
    .datum(datasetDP)
    .attr("fill", "none")
    .attr("stroke", colors[1])
    .attr("d", outLineDP);

svg.append("path")
    .datum(datasetVis)
    .attr("fill", "none")
    .attr("stroke", colors[2])
    .attr("d", outLineVis);

// Append a circle for each datapoint
svg.selectAll(".dotDP")
    .data(datasetDP)
    .enter().append("circle") // Uses the enter().append() method
    .attr("class", "dotDP") // Assign a class for styling
    .attr("cx", (d, i) =>  (xScale(d[0])))
    .attr("cy", (d) => (yScale(d[1])))
    .attr("r", 10)
    .attr("fill", "none")
    .attr("strokeWidth", "1")
    .attr("stroke", colors[1])
    .on("mouseover", (a, b, c) => { 
  	    console.log(a) 
        //this.attr('class', 'focus')
    });

svg.selectAll(".dotVis")
    .data(datasetVis)
    .enter().append("circle") // Uses the enter().append() method
    .attr("class", "dotVis") // Assign a class for styling
    .attr("cx", (d, i) =>  (xScale(d[0])))
    .attr("cy", (d) => (yScale(d[1])))
    .attr("r", 10)
    .attr("fill", "none")
    .attr("strokeWidth", "1")
    .attr("stroke", colors[2])
    .on("mouseover", (a, b, c) => { 
  	    console.log(a) 
    });

// append some legends
const addLegends = () => {
    Object.keys(data).forEach((dataSet, i) => {
        svg.append("text")
            .attr("x", 10)
            .attr("y", 10 + 20*i)
            .text(data[dataSet].title)
            .attr("fill", colors[i])
            .attr("font-size", "15px")//.attr("alignment-baseline","middle")
            .attr("font-family", "monospace");
        svg.append("text")
            .attr("x", 160)
            .attr("y", 10 + 20*i)
            .text(`${data[dataSet].dist.toFixed(1)} km`)
            .style("font-size", "15px")
            .attr("font-family", "monospace");
    })    
}
addLegends();
