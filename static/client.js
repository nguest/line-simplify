console.info({ d: data })

// Set margins for svg 
const margin = {top: 50, right: 50, bottom: 50, left: 50};
const width = window.innerWidth - margin.left - margin.right;
const height = window.innerHeight - margin.top - margin.bottom;

const datasetIn = data.in.line;
const datasetDP = data.outDP.line;
const datasetVis = data.outVis.line;
const datasetBrute = data.outBrute.line;

const colors = [
    "red",
    "green",
    "blue",
    "purple"
];

const sets = [
    {
        title: "in",
        data: data.in.line,
        circles: false,
        color: colors[0]
    },
    {
        title: "douglas-peucker",
        data: data.outDP.line,
        circles: true,
        color: colors[1]
    },
    {
        title: "visvalingam",
        data: data.outVis.line,
        circles: true,
        color: colors[2]
    },
    {
        title: "brute-force",
        data: data.outBrute.line,
        circles: true,
        color: colors[3]
    }
];

const lines = {}

// Scales
const xScale = d3.scaleLinear()
    .domain([d3.min(sets[0].data, (d) => d[0]), d3.max(sets[0].data, (d) => d[0])])
    .range([0, width]);

const yScale = d3.scaleLinear()
    .domain([d3.min(sets[0].data, (d) => d[1]), d3.max(sets[0].data, (d) => d[1])])
    .range([height, 0]); 

// Line generator
sets.forEach(set => (
    lines[set] = d3.line()
    .x((d, i) => xScale(d[0])) 
    .y((d) => yScale(d[1]))
))

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
sets.forEach(set => {
    svg.append("path")
        .datum(set.data)
        .attr("fill", "none")
        .attr("stroke", set.color)
        .attr("d", lines[set])
    // Append a circle for each datapoint
    if (set.title !== 'in') {
        svg.selectAll(`.dot-${set.title}`)
        .data(set.data)
        .enter().append("circle") // Uses the enter().append() method
        .attr("class", "dotDP") // Assign a class for styling
        .attr("cx", (d, i) =>  (xScale(d[0])))
        .attr("cy", (d) => (yScale(d[1])))
        .attr("r", 10)
        .attr("fill", "none")
        .attr("strokeWidth", "1")
        .attr("stroke", set.color)
        .on("mouseover", (a, b, c) => { 
            console.log(a) 
            //this.attr('class', 'focus')
        });
    }
})

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
