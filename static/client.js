console.log({ d: data.out })

// 2. Use the margin convention practice 
const margin = {top: 50, right: 50, bottom: 50, left: 50};
const width = window.innerWidth - margin.left - margin.right; // Use the window's width 
const height = window.innerHeight - margin.top - margin.bottom; // Use the window's height

const dataset1 = data.in// d3.range(n).map((d) => ([d3.randomUniform(1)(), d3.randomUniform(1)()]));
const dataset2 = data.out
// The number of datapoints
const n = 21;
// 5. X scale will use the index of our data
const xScale = d3.scaleLinear()
    .domain([d3.min(dataset1, (d) => d[0]), d3.max(dataset1, (d) => d[0])]) // input
    .range([0, width]); // output

// 6. Y scale will use the randomly generate number 
const yScale = d3.scaleLinear()
    .domain([d3.min(dataset1, (d) => d[1]), d3.max(dataset1, (d) => d[1])]) // input 
    .range([height, 0]); // output 

// 7. d3's line generator
const inLine = d3.line()
    .x((d, i) => xScale(d[0])) // set the x values for the line generator
    .y((d) => yScale(d[1])) // set the y values for the line generator 
    //.curve(d3.curveMonotoneX) // apply smoothing to the line
const outLine = d3.line()
    .x((d, i) => xScale(d[0])) // set the x values for the line generator
    .y((d) => yScale(d[1])) // set the y values for the line generator 
    //.curve(d3.curveMonotoneX) // apply smoothing to the line
// 8. An array of objects of length N. Each object has key -> value pair, the key being "y" and the value is a random number

//console.log(dataset1)
// 1. Add the SVG to the page and employ #2
const svg = d3.select("body").append("svg")
    .attr("width", width + margin.left + margin.right)
    .attr("height", height + margin.top + margin.bottom)
  .append("g")
    .attr("transform", "translate(" + margin.left + "," + margin.top + ")");

// 3. Call the x axis in a group tag
svg.append("g")
    .attr("class", "x axis")
    .attr("transform", "translate(0," + height + ")")
    .call(d3.axisBottom(xScale)); // Create an axis component with d3.axisBottom

// 4. Call the y axis in a group tag
svg.append("g")
    .attr("class", "y axis")
    .call(d3.axisLeft(yScale)); // Create an axis component with d3.axisLeft

// 9. Append the path, bind the data, and call the line generator 
svg.append("path")
    .datum(dataset1) // 10. Binds data to the line 
    .attr("fill", "none")
    .attr("stroke", "red")
    .attr("d", inLine); // 11. Calls the line generator 

svg.append("path")
    .datum(dataset2) // 10. Binds data to the line 
    .attr("fill", "none")
    .attr("stroke", "green")
    .attr("d", outLine); // 11. Calls the line generator 

// 12. Appends a circle for each datapoint

svg.selectAll(".dot")
    .data(dataset2)
    .enter().append("circle") // Uses the enter().append() method
    .attr("class", "dot") // Assign a class for styling
    .attr("cx", (d, i) =>  (xScale(d[0])))
    .attr("cy", (d) => (yScale(d[1])))
    .attr("r", 5)
      .on("mouseover", (a, b, c) => { 
  			console.log(a) 
        //this.attr('class', 'focus')
    })


svg.selectAll(".dot2")
    .data(dataset2).filter((d, i) => d)
    .enter().append("circle") // Uses the enter().append() method
    .attr("class", "dot2") // Assign a class for styling
    .attr("cx", (d, i) =>  (xScale(d[0])))
    .attr("cy", (d) => (yScale(d[1])))
    .attr("r", 15)
    .attr("fill", "none")
    .attr("stroke", "black")

