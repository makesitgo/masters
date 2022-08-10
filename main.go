package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("data_new.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	golfers := make([]golfer, len(lines))
	for i, line := range lines {
		cost, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		chance, err := strconv.ParseFloat(line[6], 64)
		if err != nil {
			log.Fatal(err)
		}
		golfers[i] = golfer{line[0], cost, chance}
	}

	var entries entries

	entries = append(entries, findEntries3(golfers)...)
	entries = append(entries, findEntries4(golfers)...)
	entries = append(entries, findEntries5(golfers)...)
	entries = append(entries, findEntries6(golfers)...)
	entries = append(entries, findEntries7(golfers)...)
	entries = append(entries, findEntries8(golfers)...)
	for _, entry := range entries {
		fmt.Println(entry)
	}
	fmt.Println("entries computed:", len(entries))
}

func findEntries3(golfers []golfer) entries {
	var out entries
	for i := 0; i < len(golfers)-2; i++ {
		for j := i + 1; j < len(golfers)-1; j++ {
			for k := j + 1; k < len(golfers); k++ {
				g1, g2, g3 := golfers[i], golfers[j], golfers[k]

				cost := g1.cost + g2.cost + g3.cost
				if cost > 20 {
					continue
				}

				chance := g1.chance + g2.chance + g3.chance

				out = append(out, entry{[]string{g1.name, g2.name, g3.name}, cost, chance})
			}
		}
	}
	return out
}

func findEntries4(golfers []golfer) entries {
	var out entries
	for i := 0; i < len(golfers)-3; i++ {
		for j := i + 1; j < len(golfers)-2; j++ {
			for k := j + 1; k < len(golfers)-1; k++ {
				for x := k + 1; k < len(golfers); k++ {
					g1, g2, g3, g4 := golfers[i], golfers[j], golfers[k], golfers[x]

					cost := g1.cost + g2.cost + g3.cost + g4.cost
					if cost > 20 {
						continue
					}

					chance := g1.chance + g2.chance + g3.chance + g4.chance

					out = append(out, entry{[]string{g1.name, g2.name, g3.name, g4.name}, cost, chance})
				}
			}
		}
	}
	return out
}

func findEntries5(golfers []golfer) entries {
	var out entries
	for i := 0; i < len(golfers)-4; i++ {
		for j := i + 1; j < len(golfers)-3; j++ {
			for k := j + 1; k < len(golfers)-2; k++ {
				for x := k + 1; k < len(golfers)-1; k++ {
					for y := x + 1; k < len(golfers); k++ {
						g1, g2, g3, g4, g5 := golfers[i], golfers[j], golfers[k], golfers[x], golfers[y]

						cost := g1.cost + g2.cost + g3.cost + g4.cost + g5.cost
						if cost > 20 {
							continue
						}

						chance := g1.chance + g2.chance + g3.chance + g4.chance + g5.chance

						out = append(out, entry{[]string{g1.name, g2.name, g3.name, g4.name, g5.name}, cost, chance})
					}
				}
			}
		}
	}
	return out
}

func findEntries6(golfers []golfer) entries {
	var out entries
	for i := 0; i < len(golfers)-5; i++ {
		for j := i + 1; j < len(golfers)-4; j++ {
			for k := j + 1; k < len(golfers)-3; k++ {
				for x := k + 1; k < len(golfers)-2; k++ {
					for y := x + 1; k < len(golfers)-1; k++ {
						for z := y + 1; k < len(golfers); k++ {
							g1, g2, g3, g4, g5, g6 := golfers[i], golfers[j], golfers[k], golfers[x], golfers[y], golfers[z]

							cost := g1.cost + g2.cost + g3.cost + g4.cost + g5.cost + g6.cost
							if cost > 20 {
								continue
							}

							chance := g1.chance + g2.chance + g3.chance + g4.chance + g5.chance + g6.chance

							out = append(out, entry{[]string{g1.name, g2.name, g3.name, g4.name, g5.name, g6.name}, cost, chance})
						}
					}
				}
			}
		}
	}
	return out
}

func findEntries7(golfers []golfer) entries {
	var out entries
	for i := 0; i < len(golfers)-6; i++ {
		for j := i + 1; j < len(golfers)-5; j++ {
			for k := j + 1; k < len(golfers)-4; k++ {
				for x := k + 1; k < len(golfers)-3; k++ {
					for y := x + 1; k < len(golfers)-2; k++ {
						for z := y + 1; k < len(golfers)-1; k++ {
							for m := z + 1; k < len(golfers); k++ {
								g1, g2, g3, g4, g5, g6, g7 := golfers[i], golfers[j], golfers[k], golfers[x], golfers[y], golfers[z], golfers[m]

								cost := g1.cost + g2.cost + g3.cost + g4.cost + g5.cost + g6.cost + g7.cost
								if cost > 20 {
									continue
								}

								chance := g1.chance + g2.chance + g3.chance + g4.chance + g5.chance + g6.chance + g7.chance

								out = append(out, entry{[]string{g1.name, g2.name, g3.name, g4.name, g5.name, g6.name, g7.name}, cost, chance})
							}
						}
					}
				}
			}
		}
	}
	return out
}

func findEntries8(golfers []golfer) entries {
	var out entries
	for i := 0; i < len(golfers)-7; i++ {
		for j := i + 1; j < len(golfers)-6; j++ {
			for k := j + 1; k < len(golfers)-5; k++ {
				for x := k + 1; k < len(golfers)-4; k++ {
					for y := x + 1; k < len(golfers)-3; k++ {
						for z := y + 1; k < len(golfers)-2; k++ {
							for m := z + 1; k < len(golfers)-1; k++ {
								for n := m + 1; k < len(golfers); k++ {
									g1, g2, g3, g4, g5, g6, g7, g8 := golfers[i], golfers[j], golfers[k], golfers[x], golfers[y], golfers[z], golfers[m], golfers[n]

									cost := g1.cost + g2.cost + g3.cost + g4.cost + g5.cost + g6.cost + g7.cost + g8.cost
									if cost > 20 {
										continue
									}

									chance := g1.chance + g2.chance + g3.chance + g4.chance + g5.chance + g6.chance + g7.chance + g8.chance

									out = append(out, entry{[]string{g1.name, g2.name, g3.name, g4.name, g5.name, g6.name, g7.name, g8.name}, cost, chance})
								}
							}
						}
					}
				}
			}
		}
	}
	return out
}

type golfer struct {
	name   string
	cost   float64
	chance float64
}

type entry struct {
	golfers []string
	cost    float64
	chance  float64
}

type entries []entry

func (e entries) Len() int {
	return len(e)
}

func (e entries) Less(i, j int) bool {
	return e[i].chance < e[j].chance
}

func (e entries) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
