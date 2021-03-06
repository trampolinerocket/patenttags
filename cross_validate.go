package main

import (
	"bufio"
	"fmt"
	"github.com/gtfierro/patentcluster"
	"os"
	"strconv"
)

func main() {
	outfile, err := os.Create("cross_validation.csv")
	if err != nil {
		fmt.Println("Could not output to file cross_validation.csv", err)
		return
	}
	defer outfile.Close()
	writer := bufio.NewWriter(outfile)
	filename := os.Args[1]
	fmt.Println("reading from file", filename)
	data := patentcluster.Extract_file_contents(filename, true)
	fmt.Println(len(patentcluster.Tagset), "unique tags")
	patents := patentcluster.Make_patents(data)
	fmt.Println(len(patents), "unique patents")
	writer.WriteString("Epsilon, MinPts, Number of Clusters, Mean Cluster Size, Median Cluster Size, Largest Cluster Size\n")
	for epsilon := .5; epsilon <= .95; epsilon += .05 {
		for minpts := 3; minpts < 6; minpts += 1 {
			fmt.Println("Using epsilon =", epsilon, "and minpts =", minpts)
			fmt.Println("Initializing DBSCAN...")
			db := patentcluster.Init_DBSCAN(patents, epsilon, minpts)
			fmt.Println("Running DBSCAN...")
			db.Run()
			fmt.Println("Finished running DBSCAN. Computing Stats...")
			num_clusters, mean_size, median_size, largest, _ := db.Compute_Stats()
			fmt.Println("Number of clusters found:", num_clusters)
			fmt.Println("Mean cluster size:", mean_size)
			fmt.Println("Median cluster size:", median_size)
			fmt.Println("Largest cluster size:", largest)
			line := strconv.FormatFloat(epsilon, 'g', -1, 64) + ", " + strconv.Itoa(minpts) + ", " + strconv.Itoa(num_clusters) + ", " + strconv.FormatFloat(mean_size, 'g', -1, 64) + ", " + strconv.Itoa(median_size) + ", " + strconv.Itoa(largest) + "\n"
			writer.WriteString(line)
		}
	}
	writer.Flush()
}
