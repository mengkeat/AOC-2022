package main

import "fmt"

type BP struct {
	ore               int
	clay              int
	obs_ore, obs_clay int
	geo_ore, geo_obs  int
}

func newBluePrint(s string) BP {
	bp := BP{}
	bpnum := 0
	fmt.Sscanf(s, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
		&bpnum, &bp.ore, &bp.clay, &bp.obs_ore, &bp.obs_clay, &bp.geo_ore, &bp.geo_obs)
	return bp
}

type State struct {
	ore, clay, obs, geo                 int
	rob_ore, rob_clay, rob_obs, rob_geo int
	time                                int
}

func Step(s State) State {
	next_state := s
	next_state.ore += s.rob_ore
	next_state.clay += s.rob_clay
	next_state.obs += s.rob_obs
	next_state.geo += s.rob_geo
	next_state.time--
	return next_state
}

func MaxGeode(bp BP, st State) int {
	Q := make([]State, 1)
	Q[0] = st
	visited := make(map[State]bool)
	best_geo := 0

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur.geo > best_geo {
			best_geo = cur.geo
		}

		ore_max := Max(bp.ore, bp.clay, bp.obs_ore, bp.geo_ore)
		cur.rob_ore = Min(cur.rob_ore, ore_max)
		cur.rob_clay = Min(cur.rob_clay, bp.obs_clay)
		cur.rob_obs = Min(cur.rob_obs, bp.geo_obs)

		cur.ore = Min(cur.ore, cur.time*ore_max-cur.rob_ore*(cur.time-1))
		cur.clay = Min(cur.clay, cur.time*bp.obs_clay-cur.rob_clay*(cur.time-1))
		cur.obs = Min(cur.obs, cur.time*bp.geo_obs-cur.rob_obs*(cur.time-1))

		if cur.time == 0 || InMap(cur, visited) {
			continue
		}
		visited[cur] = true

		next_step := Step(cur)
		Q = append(Q, next_step)
		if cur.ore >= bp.geo_ore && cur.obs >= bp.geo_obs {
			temp := next_step
			temp.rob_geo += 1
			temp.ore, temp.obs = temp.ore-bp.geo_ore, temp.obs-bp.geo_obs
			Q = append(Q, temp)
		}
		if cur.ore >= bp.obs_ore && cur.clay >= bp.obs_clay {
			temp := next_step
			temp.rob_obs += 1
			temp.ore, temp.clay = temp.ore-bp.obs_ore, temp.clay-bp.obs_clay
			Q = append(Q, temp)
		}
		if cur.ore >= bp.clay {
			temp := next_step
			temp.rob_clay += 1
			temp.ore -= bp.clay
			Q = append(Q, temp)
		}
		if cur.ore >= bp.ore {
			temp := next_step
			temp.rob_ore += 1
			temp.ore -= bp.ore
			Q = append(Q, temp)
		}

	}
	return best_geo
}

func main() {
	lines := LinesFromFile("input19.txt")

	start := State{rob_ore: 1, time: 24}
	quality := 0
	for i, ln := range lines {
		bp := newBluePrint(ln)
		max := MaxGeode(bp, start)
		quality += (i + 1) * max
	}
	fmt.Println("Part A:", quality)

	start2 := State{rob_ore: 1, time: 32}
	b1 := MaxGeode(newBluePrint(lines[0]), start2)
	b2 := MaxGeode(newBluePrint(lines[1]), start2)
	b3 := MaxGeode(newBluePrint(lines[2]), start2)
	fmt.Println("Part B:", b1*b2*b3)
}
