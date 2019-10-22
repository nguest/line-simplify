console.log({ d: data })

// 2. Use the margin convention practice 
const margin = {top: 50, right: 50, bottom: 50, left: 50};
const width = window.innerWidth - margin.left - margin.right; // Use the window's width 
const height = window.innerHeight - margin.top - margin.bottom; // Use the window's height

const datasetIn = data.in;// d3.range(n).map((d) => ([d3.randomUniform(1)(), d3.randomUniform(1)()]));
const datasetDP = data.outDP;
const datasetVis = data.outVis;

const colorIn = "red";
const colorDP = "green";
const colorVis = "blue";

// scales
const xScale = d3.scaleLinear()
    .domain([d3.min(datasetIn, (d) => d[0]), d3.max(datasetIn, (d) => d[0])]) // input
    .range([0, width]); // output

const yScale = d3.scaleLinear()
    .domain([d3.min(datasetIn, (d) => d[1]), d3.max(datasetIn, (d) => d[1])]) // input 
    .range([height, 0]); // output 

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

// Call the x axis in a group tag
svg.append("g")
    .attr("class", "x axis")
    .attr("transform", "translate(0," + height + ")")
    .call(d3.axisBottom(xScale)); // Create an axis component with d3.axisBottom

// Call the y axis in a group tag
svg.append("g")
    .attr("class", "y axis")
    .call(d3.axisLeft(yScale)); // Create an axis component with d3.axisLeft

// Append path, bind data, call line generator 
svg.append("path")
    .datum(datasetIn)
    .attr("fill", "none")
    .attr("stroke", colorIn)
    .attr("d", inLine)

svg.append("path")
    .datum(datasetDP)
    .attr("fill", "none")
    .attr("stroke", colorDP)
    .attr("d", outLineDP)

svg.append("path")
    .datum(datasetVis)
    .attr("fill", "none")
    .attr("stroke", colorVis)
    .attr("d", outLineVis)

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
    .attr("stroke", colorDP)
    .on("mouseover", (a, b, c) => { 
  	    console.log(a) 
        //this.attr('class', 'focus')
    })

svg.selectAll(".dotVis")
    .data(datasetVis)
    .enter().append("circle") // Uses the enter().append() method
    .attr("class", "dotVis") // Assign a class for styling
    .attr("cx", (d, i) =>  (xScale(d[0])))
    .attr("cy", (d) => (yScale(d[1])))
    .attr("r", 10)
    .attr("fill", "none")
    .attr("strokeWidth", "1")
    .attr("stroke", colorVis)
    .on("mouseover", (a, b, c) => { 
  	    console.log(a) 
        //this.attr('class', 'focus')
    })
