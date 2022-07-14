# Comet Backup Skill Test

4 Inputs:

-   Road 1 (North) CPM
-   Road 2 (East) CPM
-   Road 3 (South) CPM
-   Road 4 (West) CPM

Performance table (Constants)
|Control Method|“High Throughput” (Total CPM >= 20)| “Medium Throughput” (20 > Total CPM >= 10) | “Low Throughput” (Total CPM < 10)|
|----------|:-------------:|------:|------:|
| Roundabout | 50% efficient | 75% efficient | 90% efficient
| Stop Signs | 20% efficient | 30% efficient | 40% efficient
| Traffic Lights | 90% efficient | 75% efficient | 30% efficient

## Extensions
- [x] The Council would like to be able to manipulate the CPM values for each four roads without recompiling every time. An interactive prompt would be helpful. You may choose to prompt for input values on stdin, or accept them as command-line parameters.
- [x] You might want to use visual formatting to make it more readable and visually appealing, or implement a GUI interface using Qt. If your program is a command-line program, you could add support for a "/?" or "--help" or "--usage" command-line flag.
- [x] You might want to implement a batch system that produces the result for a large number of CPM values, read from a CSV file on disk or stdin.
- [ ] You might want to extend the system to support arbitrary traffic control methods.
- [x] Roundabouts tend to be more efficient when one through-road is busy, and the other is
quiet. So, if roads 1 and 3 have a high CPM relative to 2 and 4, then a bonus efficiency score
of +10% might be earned for the roundabout control method.
- [ ] You might have noticed that stop signs never get the best score. However, they are very
cheap to make. Your program could allow the user to enter costs of each of the 3 control methods. You can then output a CPM-per-dollar value, indicating the number of cars the intersection can support, per dollar of construction cost. Roundabouts cost $100k, stop signs cost $40k and traffic lights cost $200k.
- [x] Rather than three static states of throughput (low, medium, high), could these states be a continuous curve? Modify the algorithm to smoothly transition between these states. **I got the equations from**[reference](https://www.mathcelebrity.com/3ptquad.php?p1=10%2C0.9&p2=15%2C0.75&p3=20%2C0.50&pl=Calculate+Equation)
- [ ] Add support for dual carriageway and/or pedestrian lanes. Can you allow a variable number of lanes in each direction?
- [ ] Traffic lights can allow north-to-south and south-to-north car movement at the same time. What are all the other path combinations? What is the smallest covering set of paths that can be travelled simultaneously? Explore the maximum efficiency when each path combination has a different CPM rate.
- [ ] Each of the 3 control methods has a different average latency before a new car is able to pass. At high throughput levels, traffic lights are more efficient, but a car might be able to
enter an empty roundabout sooner. Invent some realistic latency numbers for each control method, and then investigate the latency-vs-throughput curve. What is the effect of cars “piling up” on average latency? Simulate this.